package reviews

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tender-service/internal/bids"
	"tender-service/internal/employee"
	"tender-service/internal/orgresponsible"

	"github.com/gorilla/mux"
)

type ReviewHandler struct {
	reviewService   *ReviewService
	bidService      *bids.BidService
	employeeService *employee.EmployeeService
	orgRespService  *orgresponsible.OrgRespService
}

func NewTenderHandler(
	reviewSerive *ReviewService,
	bidService *bids.BidService,
	employeeService *employee.EmployeeService,
	orgRespService *orgresponsible.OrgRespService,
) *ReviewHandler {
	return &ReviewHandler{
		reviewService:   reviewSerive,
		bidService:      bidService,
		employeeService: employeeService,
		orgRespService:  orgRespService,
	}
}

func (h *ReviewHandler) CreateReview(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var review Review

	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	employee, err := h.employeeService.GetEmployeeByUsername(ctx, review.AuthorUsername)
	if err != nil {
		http.Error(w, "Invalid user", http.StatusBadRequest)
		return
	}

	_, err = h.orgRespService.GetResponsible(ctx, employee.ID)
	if err != nil {
		http.Error(w, "Not responsible for organization", http.StatusForbidden)
		return
	}

	createdReview, err := h.reviewService.CreateReview(ctx, review)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createdReview)
}

func (h *ReviewHandler) GetReviewsByBidId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)
	tenderId := params["tenderId"]

	intTenderId, err := strconv.Atoi(tenderId)
	if err != nil {
		http.Error(w, "Invalid query params", http.StatusBadRequest)
		return
	}

	authorUsername := r.URL.Query().Get("authorUsername")

	bids, err := h.bidService.GetBidsByCreatorUsername(ctx, authorUsername)
	if err != nil {
		http.Error(w, "No reviews for this organization", http.StatusForbidden)
		return
	}

	isBidFit := false
	for _, bid := range bids {
		if bid.TenderId == int32(intTenderId) {
			isBidFit = true
			break
		}
	}

	if !isBidFit {
		http.Error(w, "No bids from this user", http.StatusForbidden)
		return
	}

	outReviews := make([]Review, 0)

	for _, bid := range bids {
		reviews, err := h.reviewService.GetReviews(ctx, int32(bid.ID))
		if err == nil {
			outReviews = append(outReviews, reviews...)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(outReviews)
}
