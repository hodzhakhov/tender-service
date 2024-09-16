package tenders

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tender-service/internal/employee"
	"tender-service/internal/orgresponsible"

	"github.com/gorilla/mux"
)

type TenderHandler struct {
	tenderService   *TenderService
	employeeService *employee.EmployeeService
	orgRespService  *orgresponsible.OrgRespService
}

func NewTenderHandler(
	tenderSerive *TenderService,
	employeeService *employee.EmployeeService,
	orgRespService *orgresponsible.OrgRespService,
) *TenderHandler {
	return &TenderHandler{
		tenderService:   tenderSerive,
		employeeService: employeeService,
		orgRespService:  orgRespService,
	}
}

// получение тендеров
func (h *TenderHandler) GetTenders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	serviceType := r.URL.Query().Get("serviceType")

	tenders, err := h.tenderService.GetTenders(ctx, serviceType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tenders)
}

// создание тендера
func (h *TenderHandler) CreateTender(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var tender Tender

	if err := json.NewDecoder(r.Body).Decode(&tender); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	employee, err := h.employeeService.GetEmployeeByUsername(ctx, tender.CreatorUsername)
	if err != nil {
		http.Error(w, "Invalid user", http.StatusBadRequest)
		return
	}

	responsible, err := h.orgRespService.GetResponsible(ctx, employee.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if responsible.OrganizationId != tender.OrganizationId {
		http.Error(w, "Not responsible for organization", http.StatusForbidden)
		return
	}

	createdTender, err := h.tenderService.CreateTender(ctx, tender)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createdTender)
}

// получение тендеров пользователя
func (h *TenderHandler) GetMyTenders(w http.ResponseWriter, r *http.Request) {
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

	tenders, err := h.tenderService.GetMyTenders(ctx, responsible.OrganizationId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tenders)
}

// редактирование тендера
func (h *TenderHandler) EditTender(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)
	tenderID := params["tenderId"]

	var tender Tender

	err := json.NewDecoder(r.Body).Decode(&tender)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(tenderID)
	if err != nil {
		http.Error(w, "Invalid query params", http.StatusBadRequest)
		return
	}

	tender.ID = int32(id)

	updatedTender, err := h.tenderService.EditTender(ctx, tender)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTender)
}

// изменение статуса тендера
func (h *TenderHandler) SetTenderStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)
	tenderID := params["tenderId"]

	intTenderId, err := strconv.Atoi(tenderID)
	if err != nil {
		http.Error(w, "Invalid query params", http.StatusBadRequest)
		return
	}

	var tender Tender

	err = json.NewDecoder(r.Body).Decode(&tender)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.tenderService.SetTenderStatus(ctx, int32(intTenderId), tender.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// откат версии тендера
func (h *TenderHandler) RollbackTender(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)
	tenderID := params["tenderId"]
	version := params["version"]

	intTenderId, err := strconv.Atoi(tenderID)
	if err != nil {
		http.Error(w, "Invalid query params", http.StatusBadRequest)
		return
	}

	intVersion, err := strconv.Atoi(version)
	if err != nil {
		http.Error(w, "Invalid query params", http.StatusBadRequest)
		return
	}

	updatedTender, err := h.tenderService.RollbackTender(ctx, int32(intTenderId), int32(intVersion))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTender)
}
