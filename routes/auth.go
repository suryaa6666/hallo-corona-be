package routes

import (
	"hallocorona/handlers"
	"hallocorona/pkg/middleware"
	"hallocorona/pkg/mysql"
	"hallocorona/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	authRepository := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(authRepository)

	r.HandleFunc("/login", h.LoginAuth).Methods("POST")
	r.HandleFunc("/register", h.RegisterAuth).Methods("POST")
	r.HandleFunc("/check-auth", middleware.Auth(h.CheckAuth, "member")).Methods("GET")
}
