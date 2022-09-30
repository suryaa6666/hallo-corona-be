package consultationdto

import (
	"hallocorona/models"
)

type ConsultationResponse struct {
	ID               int          `json:"id"`
	FullName         string       `json:"fullName"`
	Phone            string       `json:"phone"`
	BornDate         int          `json:"bornDate"`
	Age              int          `json:"age"`
	Height           int          `json:"height"`
	Weight           int          `json:"weight"`
	Gender           string       `json:"gender"`
	Subject          string       `json:"subject"`
	LiveConsultation int          `json:"liveConsultation"`
	Description      string       `json:"description"`
	Status           string       `json:"status"`
	ReplyID          int          `json:"-"`
	Reply            models.Reply `json:"reply"`
	UserID           int          `json:"-"`
	User             models.User  `json:"user"`
}
