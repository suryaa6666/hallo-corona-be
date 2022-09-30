package repositories

import (
	"hallocorona/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	RegisterAuth(auth models.User) (models.User, error)
	LoginAuth(email string) (models.User, error)
	GetUserAuth(ID int) (models.User, error)
	RegisterUpdateAuth(user models.User) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) RegisterUpdateAuth(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) LoginAuth(email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}

func (r *repository) RegisterAuth(auth models.User) (models.User, error) {
	err := r.db.Create(&auth).Error

	return auth, err
}

func (r *repository) GetUserAuth(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error

	return user, err
}
