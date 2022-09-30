package routes

import (
	"hallocorona/handlers"
	"hallocorona/pkg/middleware"
	"hallocorona/pkg/mysql"
	"hallocorona/repositories"

	"github.com/gorilla/mux"
)

func ReplyRoutes(r *mux.Router) {
	replyRepository := repositories.RepositoryReply(mysql.DB)
	h := handlers.HandlerReply(replyRepository)

	r.HandleFunc("/replies", middleware.Auth(h.FindReplies, "member")).Methods("GET")
	r.HandleFunc("/reply/{id}", middleware.Auth(h.GetReply, "member")).Methods("GET")
	r.HandleFunc("/reply", middleware.Auth(h.CreateReply, "member")).Methods("POST")
	r.HandleFunc("/reply/{id}", middleware.Auth(h.UpdateReply, "member")).Methods("PATCH")
	r.HandleFunc("/reply/{id}", middleware.Auth(h.DeleteReply, "member")).Methods("DELETE")
}
