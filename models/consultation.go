package models

import (
	"time"
)

type Consultation struct {
	ID               int           `json:"id" gorm:"primary_key:auto_increment"`
	FullName         string        `json:"fullName" gorm:"type: varchar(255)"`
	Phone            string        `json:"phone" gorm:"type: varchar(50)"`
	BornDate         int           `json:"bornDate"`
	Age              int           `json:"age"`
	Height           int           `json:"height"`
	Weight           int           `json:"weight"`
	Gender           string        `json:"gender" gorm:"type: varchar(50)"`
	Subject          string        `json:"subject" gorm:"type: varchar(255)"`
	LiveConsultation int           `json:"liveConsultation"`
	Description      string        `json:"description" gorm:"type: text"`
	Status           string        `json:"status" gorm:"type: varchar(100)"`
	ReplyID          int           `json:"-" gorm:"default: null"`
	Reply            ReplyResponse `json:"reply" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID           int           `json:"-"`
	User             UsersResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt        time.Time     `json:"createdAt"`
	UpdatedAt        time.Time     `json:"updatedAt"`
}
