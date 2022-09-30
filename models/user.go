package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	FullName  string    `json:"fullName" gorm:"type: varchar(255)"`
	Email     string    `json:"email" gorm:"type: varchar(255)"`
	Username  string    `json:"username" gorm:"type: varchar(255)"`
	Password  string    `json:"-" gorm:"type: varchar(255)"`
	ListAs    string    `json:"listAs" gorm:"type: varchar(50)"`
	Gender    string    `json:"gender" gorm:"type: varchar(100)"`
	Phone     string    `json:"phone" gorm:"type: varchar(50)"`
	Address   string    `json:"address" gorm:"type: varchar(255)"`
	Image     string    `json:"image" form:"image" gorm:"type: varchar(255)"`
	Role      string    `json:"role" gorm:"type: varchar(50)"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UsersResponse struct {
	ID        int       `json:"id"`
	FullName  string    `json:"fullName"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	ListAs    string    `json:"listAs"`
	Gender    string    `json:"gender"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Image     string    `json:"image"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (UsersResponse) TableName() string {
	return "users"
}
