package routes

import (
	"hallocorona/handlers"
	"hallocorona/pkg/middleware"
	"hallocorona/pkg/mysql"
	"hallocorona/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/users", middleware.Auth(h.FindUsers, "admin")).Methods("GET")
	r.HandleFunc("/user/{id}", middleware.Auth(h.GetUser, "admin")).Methods("GET")
	r.HandleFunc("/user", middleware.Auth(h.CreateUser, "admin")).Methods("POST")
	r.HandleFunc("/user/{id}", middleware.Auth(middleware.UploadFile(h.UpdateUser), "member")).Methods("PATCH")
	r.HandleFunc("/user/{id}", middleware.Auth(h.DeleteUser, "admin")).Methods("DELETE")
}
