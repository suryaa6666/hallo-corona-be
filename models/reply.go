package models

import (
	"time"
)

type Reply struct {
	ID        int           `json:"id" gorm:"primary_key:auto_increment"`
	Response  string        `json:"response" gorm:"type: text"`
	MeetLink  string        `json:"meetLink" gorm:"type: varchar(255)"`
	MeetType  string        `json:"meetType" gorm:"type: varchar(255)"`
	UserID    int           `json:"-"`
	User      UsersResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
}

type ReplyResponse struct {
	ID        int           `json:"id"`
	Response  string        `json:"response"`
	MeetLink  string        `json:"meetLink"`
	MeetType  string        `json:"meetType"`
	UserID    int           `json:"-"`
	User      UsersResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
}

func (ReplyResponse) TableName() string {
	return "replies"
}
