package articledto

import (
	"hallocorona/models"
)

type ArticleResponse struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Image       string      `json:"image"`
	Description string      `json:"description"`
	UserID      int         `json:"-"`
	User        models.User `json:"user"`
}
