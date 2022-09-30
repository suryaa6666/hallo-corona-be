package authdto

import (
	"hallocorona/models"
)

type LoginResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"-"`
	ListAs   string `json:"listAs"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Image    string `json:"image"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

type CheckAuthResponse struct {
	User   models.User `json:"user"`
	Status string      `json:"status"`
}
