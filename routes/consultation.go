package routes

import (
	"hallocorona/handlers"
	"hallocorona/pkg/middleware"
	"hallocorona/pkg/mysql"
	"hallocorona/repositories"

	"github.com/gorilla/mux"
)

func ConsultationRoutes(r *mux.Router) {
	consultationRepository := repositories.RepositoryConsultation(mysql.DB)
	h := handlers.HandlerConsultation(consultationRepository)

	r.HandleFunc("/consultations", middleware.Auth(h.FindConsultations, "member")).Methods("GET")
	r.HandleFunc("/consultation/{id}", middleware.Auth(h.GetConsultation, "member")).Methods("GET")
	r.HandleFunc("/consultation", middleware.Auth(h.CreateConsultation, "member")).Methods("POST")
	r.HandleFunc("/consultation/{id}", middleware.Auth(h.UpdateConsultation, "member")).Methods("PATCH")
	r.HandleFunc("/consultationstatus/{id}", middleware.Auth(h.UpdateConsultationStatus, "member")).Methods("PATCH")
	r.HandleFunc("/consultation/{id}", middleware.Auth(h.DeleteConsultation, "member")).Methods("DELETE")
}
