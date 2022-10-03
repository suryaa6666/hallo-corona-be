package routes

import (
	"hallocorona/handlers"
	"hallocorona/pkg/middleware"
	"hallocorona/pkg/mysql"
	"hallocorona/repositories"

	"github.com/gorilla/mux"
)

func ArticleRoutes(r *mux.Router) {
	articleRepository := repositories.RepositoryArticle(mysql.DB)
	h := handlers.HandlerArticle(articleRepository)

	r.HandleFunc("/articles", h.FindArticles).Methods("GET")
	r.HandleFunc("/article/{id}", h.GetArticle).Methods("GET")
	r.HandleFunc("/article", middleware.Auth(middleware.UploadFile(h.CreateArticle, ""), "member")).Methods("POST")
	r.HandleFunc("/article/{id}", middleware.Auth(middleware.UploadFile(h.UpdateArticle, "edit"), "member")).Methods("PATCH")
	r.HandleFunc("/article/{id}", middleware.Auth(h.DeleteArticle, "member")).Methods("DELETE")
}
