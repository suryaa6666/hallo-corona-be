package consultationdto

type CreateConsultationRequest struct {
	FullName    string `json:"fullName" validate:"required"`
	Phone       string `json:"phone" validate:"required"`
	BornDate    int    `json:"bornDate" validate:"required"`
	Age         int    `json:"age" validate:"required"`
	Height      int    `json:"height" validate:"required"`
	Weight      int    `json:"weight" validate:"required"`
	Gender      string `json:"gender" validate:"required"`
	Subject     string `json:"subject" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateConsultationRequest struct {
	FullName    string `json:"fullName" validate:"required"`
	Phone       string `json:"phone" validate:"required"`
	BornDate    int    `json:"bornDate" validate:"required"`
	Age         int    `json:"age" validate:"required"`
	Height      int    `json:"height" validate:"required"`
	Weight      int    `json:"weight" validate:"required"`
	Gender      string `json:"gender" validate:"required"`
	Subject     string `json:"subject" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateConsultationStatus struct {
	LiveConsultation int    `json:"liveConsultation" validate:"required"`
	Status           string `json:"status" validate:"required"`
	ReplyID          int    `json:"replyId" validate:"required"`
}
