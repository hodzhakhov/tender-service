package bids

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tender-service/internal/employee"
	"tender-service/internal/orgresponsible"
	"tender-service/internal/tenders"

	"github.com/gorilla/mux"
)

type BidHandler struct {
	bidService      *BidService
	tenderService   *tenders.TenderService
	employeeService *employee.EmployeeService
	orgRespService  *orgresponsible.OrgRespService
}

func NewBidHandler(
	bidSerive *BidService,
	tenderSerive *tenders.TenderService,
	employeeService *employee.EmployeeService,
	orgRespService *orgresponsible.OrgRespService,
) *BidHandler {
	return &BidHandler{
		bidService:      bidSerive,
		tenderService:   tenderSerive,
		employeeService: employeeService,
		orgRespService:  orgRespService,
	}
}

// создание предложения
func (h *BidHandler) CreateBid(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var bid Bid

	if err := json.NewDecoder(r.Body).Decode(&bid); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	employee, err := h.employeeService.GetEmployeeByUsername(ctx, bid.CreatorUsername)
	if err != nil {
		http.Error(w, "Invalid user", http.StatusBadRequest)
		return
	}

	responsible, err := h.orgRespService.GetResponsible(ctx, employee.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if responsible.OrganizationId != bid.OrganizationId {
		http.Error(w, "Not responsible for organization", http.StatusForbidden)
		return
	}

	_, err = h.tenderService.GetTenderById(ctx, bid.TenderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	createdBid, err := h.bidService.CreateBid(ctx, bid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createdBid)
}

// получение предложений пользователя
func (h *BidHandler) GetMyBids(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Invalid request params", http.StatusBadRequest)
		return
	}

	employee, err := h.employeeService.GetEmployeeByUsername(ctx, username)
	if err != nil {
		http.Error(w, "Invalid user", http.StatusBadRequest)
		return
	}

	responsible, err := h.orgRespService.GetResponsible(ctx, employee.ID)
	if err != nil {
		http.Error(w, "Not responsible for organization", http.StatusForbidden)
		return
	}

	bids, err := h.bidService.GetMyBids(ctx, responsible.OrganizationId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bids)
}

func (h *BidHandler) GetBidsByTenderId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)
	tenderId := params["tenderId"]

	id, err := strconv.Atoi(tenderId)
	if err != nil {
		http.Error(w, "Invalid query params", http.StatusBadRequest)
		return
	}

	bids, err := h.bidService.GetBidsByTenderId(ctx, int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bids)
}

func (h *BidHandler) EditBid(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)
	bidID := params["bidId"]

	var bid Bid

	err := json.NewDecoder(r.Body).Decode(&bid)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(bidID)
	if err != nil {
		http.Error(w, "Invalid bid id", http.StatusBadRequest)
		return
	}

	bid.ID = int32(id)

	updatedBid, err := h.bidService.EditBid(ctx, bid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedBid)
}

func (h *BidHandler) SetBidStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)
	bidID := params["bidId"]

	intBidId, err := strconv.Atoi(bidID)
	if err != nil {
		http.Error(w, "Invalid query params", http.StatusBadRequest)
		return
	}

	var bid Bid

	err = json.NewDecoder(r.Body).Decode(&bid)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.bidService.SetBidStatus(ctx, int32(intBidId), bid.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BidHandler) RollbackBid(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)
	bidID := params["bidId"]
	version := params["version"]

	intBidId, err := strconv.Atoi(bidID)
	if err != nil {
		http.Error(w, "Invalid query params", http.StatusBadRequest)
		return
	}

	intVersion, err := strconv.Atoi(version)
	if err != nil {
		http.Error(w, "Invalid query params", http.StatusBadRequest)
		return
	}

	updatedBid, err := h.bidService.RollbackBid(ctx, int32(intBidId), int32(intVersion))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedBid)
}

func (h *BidHandler) VoteOnBid(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)
	bidID := params["bidId"]

	intBidId, err := strconv.Atoi(bidID)
	if err != nil {
		http.Error(w, "Invalid query params", http.StatusBadRequest)
		return
	}

	var vote BidVote

	err = json.NewDecoder(r.Body).Decode(&vote)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	vote.BidId = int32(intBidId)

	employee, err := h.employeeService.GetEmployeeByUsername(ctx, vote.CreatorUsername)
	if err != nil {
		http.Error(w, "Invalid user", http.StatusBadRequest)
		return
	}

	responsible, err := h.orgRespService.GetResponsible(ctx, employee.ID)
	if err != nil {
		http.Error(w, "Not responsible for organization", http.StatusForbidden)
		return
	}

	bid, err := h.bidService.GetBidById(ctx, vote.BidId)
	if err != nil {
		http.Error(w, "No bid", http.StatusBadRequest)
		return
	}

	tender, err := h.tenderService.GetTenderById(ctx, bid.TenderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if tender.OrganizationId != responsible.OrganizationId {
		http.Error(w, "Bid is not available", http.StatusForbidden)
		return
	}

	bidVote, err := h.bidService.VoteOnBid(ctx, vote)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if bidVote.ID == 0 {
		http.Error(w, "Bid is not available", http.StatusForbidden)
		return
	}

	bidVote.CreatorUsername = vote.CreatorUsername

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bidVote)
}

func (h *BidHandler) CheckQuorum(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)
	tenderId := params["tenderId"]
	bidId := params["bidId"]

	intTenderId, err := strconv.Atoi(tenderId)
	if err != nil {
		http.Error(w, "Invalid query params", http.StatusBadRequest)
		return
	}

	intBidId, err := strconv.Atoi(bidId)
	if err != nil {
		http.Error(w, "Invalid query params", http.StatusBadRequest)
		return
	}

	tender, err := h.tenderService.GetTenderById(ctx, int32(intTenderId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responsibles, err := h.orgRespService.GetResponsiblesByOrgId(ctx, tender.OrganizationId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	approveCount, err := h.tenderService.CheckDecisionCount(ctx, int32(intBidId), true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rejectCount, err := h.tenderService.CheckDecisionCount(ctx, int32(intBidId), false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rejectCount > 0 {
		err := h.bidService.SetBidStatus(ctx, int32(intBidId), "CANCELED")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}

	if approveCount >= min(3, int32(len(responsibles))) {
		err := h.bidService.SetBidStatus(ctx, int32(intBidId), "PUBLISHED")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = h.tenderService.SetTenderStatus(ctx, int32(intTenderId), "CANCELED")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	}
}
