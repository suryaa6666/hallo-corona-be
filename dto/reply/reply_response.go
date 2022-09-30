package replydto

import (
	"hallocorona/models"
	"time"
)

type ReplyResponse struct {
	ID        int         `json:"id"`
	Response  string      `json:"response"`
	MeetLink  string      `json:"meetLink"`
	MeetType  string      `json:"meetType"`
	User      models.User `json:"user"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}
