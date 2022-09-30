package models

import (
	"time"
)

type Article struct {
	ID          int           `json:"id" gorm:"primary_key:auto_increment"`
	Title       string        `json:"title" form:"title" gorm:"type: varchar(255)"`
	Image       string        `json:"image" form:"image" gorm:"type: varchar(255)"`
	Description string        `json:"description" form:"description" gorm:"type: text"`
	UserID      int           `json:"-"`
	User        UsersResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
}
