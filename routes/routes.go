package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	AuthRoutes(r)
	ArticleRoutes(r)
	ReplyRoutes(r)
	ConsultationRoutes(r)
}
