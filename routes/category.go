package routes

import (
	"hallocorona/handlers"
	"hallocorona/pkg/middleware"
	"hallocorona/pkg/mysql"
	"hallocorona/repositories"

	"github.com/gorilla/mux"
)

func categoryRoutes(r *mux.Router) {
	categoryRepository := repositories.RepositoryCategory(mysql.DB)
	h := handlers.HandlerCategory(categoryRepository)

	r.HandleFunc("/categories", middleware.Auth(h.FindCategories, "member")).Methods("GET")
	r.HandleFunc("/category/{id}", middleware.Auth(h.GetCategory, "member")).Methods("GET")
	r.HandleFunc("/category", middleware.Auth(h.CreateCategory, "member")).Methods("POST")
	r.HandleFunc("/category/{id}", middleware.Auth(h.UpdateCategory, "member")).Methods("PATCH")
	r.HandleFunc("/category/{id}", middleware.Auth(h.DeleteCategory, "member")).Methods("DELETE")
}
