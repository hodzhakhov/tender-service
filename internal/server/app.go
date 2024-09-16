package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"tender-service/internal/bids"
	"tender-service/internal/employee"
	"tender-service/internal/orgresponsible"
	"tender-service/internal/reviews"
	"tender-service/internal/tenders"

	"github.com/gorilla/mux"
)

func Run(psqlDB *sql.DB) {
	tenderService := tenders.NewTenderService(
		tenders.NewTenderRepo(psqlDB),
	)

	bidService := bids.NewBidService(
		bids.NewBidRepo(psqlDB),
	)

	rebiewService := reviews.NewReviewService(
		reviews.NewReviewRepo(psqlDB),
	)

	employeeService := employee.NewEmployeeService(
		employee.NewEmployeeRepo(psqlDB),
	)

	orgRespService := orgresponsible.NewOrgRespService(
		orgresponsible.NewOrgRespRepo(psqlDB),
	)

	tenderHandler := tenders.NewTenderHandler(
		tenderService,
		employeeService,
		orgRespService,
	)

	bidHandler := bids.NewBidHandler(
		bidService,
		tenderService,
		employeeService,
		orgRespService,
	)

	reviewHandler := reviews.NewTenderHandler(
		rebiewService,
		bidService,
		employeeService,
		orgRespService,
	)

	r := mux.NewRouter()

	// тендеры
	r.HandleFunc("/api/ping", PingHandler).Methods("GET")
	r.HandleFunc("/api/tenders", tenderHandler.GetTenders).Methods("GET")
	r.HandleFunc("/api/tenders/new", tenderHandler.CreateTender).Methods("POST")
	r.HandleFunc("/api/tenders/my", tenderHandler.GetMyTenders).Methods("GET")
	r.HandleFunc("/api/tenders/{tenderId}/edit", tenderHandler.EditTender).Methods("PATCH")
	r.HandleFunc("/api/tenders/{tenderId}/status", tenderHandler.SetTenderStatus).Methods("PATCH")
	r.HandleFunc("/api/tenders/{tenderId}/rollback/{version}", tenderHandler.RollbackTender).Methods("PUT")

	// предложения
	r.HandleFunc("/api/bids/new", bidHandler.CreateBid).Methods("POST")
	r.HandleFunc("/api/bids/my", bidHandler.GetMyBids).Methods("GET")
	r.HandleFunc("/api/bids/{tenderId}/list", bidHandler.GetBidsByTenderId).Methods("GET")
	r.HandleFunc("/api/bids/{bidId}/edit", bidHandler.EditBid).Methods("PATCH")
	r.HandleFunc("/api/bids/{bidId}/rollback/{version}", bidHandler.RollbackBid).Methods("PUT")

	// согласование предложений
	r.HandleFunc("/api/bids/{bidId}/submit_decision", bidHandler.VoteOnBid).Methods("POST")
	r.HandleFunc("/api/tenders/{tenderId}/quorum/{bidId}", bidHandler.CheckQuorum).Methods("GET")

	// отзывыx
	r.HandleFunc("/api/bids/feedback", reviewHandler.CreateReview).Methods("POST")

	// посчитал, что organizationId излишне
	r.HandleFunc("/api/bids/{tenderId}/reviews", reviewHandler.GetReviewsByBidId).Methods("GET")

	// Запуск сервера
	address := os.Getenv("SERVER_ADDRESS")
	log.Printf("Server is running at %s", address)
	log.Fatal(http.ListenAndServe(address, r))
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
