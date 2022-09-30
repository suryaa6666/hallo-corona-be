package repositories

import (
	"hallocorona/models"

	"gorm.io/gorm"
)

type ReplyRepository interface {
	FindReplies() ([]models.Reply, error)
	GetReply(ID int) (models.Reply, error)
	CreateReply(reply models.Reply) (models.Reply, error)
	UpdateReply(reply models.Reply) (models.Reply, error)
	DeleteReply(reply models.Reply) (models.Reply, error)
	GetReplyAuthor(UserID int) (models.User, error)
}

func RepositoryReply(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetReplyAuthor(UserID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, UserID).Error

	return user, err
}

func (r *repository) FindReplies() ([]models.Reply, error) {
	var replies []models.Reply
	err := r.db.Find(&replies).Error

	return replies, err
}

func (r *repository) GetReply(ID int) (models.Reply, error) {
	var reply models.Reply
	err := r.db.First(&reply, ID).Error

	return reply, err
}

func (r *repository) CreateReply(reply models.Reply) (models.Reply, error) {
	err := r.db.Create(&reply).Error

	return reply, err
}

func (r *repository) UpdateReply(reply models.Reply) (models.Reply, error) {
	err := r.db.Save(&reply).Error

	return reply, err
}

func (r *repository) DeleteReply(reply models.Reply) (models.Reply, error) {
	err := r.db.Delete(&reply).Error

	return reply, err
}
