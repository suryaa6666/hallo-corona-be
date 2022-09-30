package replydto

import "hallocorona/models"

type ReplyResponse struct {
	ID       int         `json:"id"`
	Response string      `json:"response"`
	MeetLink string      `json:"meetLink"`
	MeetType string      `json:"meetType"`
	User     models.User `json:"user"`
}
