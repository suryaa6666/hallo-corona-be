package replydto

type CreateReplyRequest struct {
	Response string `json:"response" validate:"required"`
	MeetLink string `json:"meetLink" validate:"required"`
	MeetType string `json:"meetType" validate:"required"`
}

type UpdateReplyRequest struct {
	Response string `json:"response" validate:"required"`
	MeetLink string `json:"meetLink" validate:"required"`
	MeetType string `json:"meetType" validate:"required"`
	UserID   int    `json:"userId" validate:"required"`
}
