package repositories

import (
	"hallocorona/models"

	"gorm.io/gorm"
)

type ConsultationRepository interface {
	FindConsultations(UserID int) ([]models.Consultation, error)
	GetConsultation(ID int) (models.Consultation, error)
	CreateConsultation(consultation models.Consultation) (models.Consultation, error)
	UpdateConsultation(consultation models.Consultation, consultationID int) (models.Consultation, error)
	DeleteConsultation(consultation models.Consultation) (models.Consultation, error)
	GetConsultationAuthor(UserID int) (models.User, error)
	GetConsultationReply(ID int) (models.Reply, error)
	UpdateConsultationStatus(consultation models.Consultation) (models.Consultation, error)
}

func RepositoryConsultation(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetConsultationReply(ID int) (models.Reply, error) {
	var reply models.Reply
	err := r.db.Preload("User").First(&reply, ID).Error

	return reply, err
}

func (r *repository) GetConsultationAuthor(UserID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, UserID).Error

	return user, err
}

func (r *repository) FindConsultations(UserID int) ([]models.Consultation, error) {
	var consultations []models.Consultation
	err := r.db.Raw("SELECT * FROM consultations WHERE user_id=?", UserID).Scan(&consultations).Error

	return consultations, err
}

func (r *repository) GetConsultation(ID int) (models.Consultation, error) {
	var consultation models.Consultation
	err := r.db.First(&consultation, ID).Error

	return consultation, err
}

func (r *repository) CreateConsultation(consultation models.Consultation) (models.Consultation, error) {
	err := r.db.Create(&consultation).Error

	return consultation, err
}

func (r *repository) UpdateConsultationStatus(consultation models.Consultation) (models.Consultation, error) {
	err := r.db.Save(&consultation).Error

	return consultation, err
}

func (r *repository) UpdateConsultation(consultation models.Consultation, consultationID int) (models.Consultation, error) {
	err := r.db.Raw("UPDATE consultations SET full_name=?,phone=?,born_date =?,age=?,height=?,weight=?,gender=?,subject=?,live_consultation=?,description=?,status=?,user_id=? WHERE id=?", consultation.FullName, consultation.Phone, consultation.BornDate, consultation.Age, consultation.Height, consultation.Weight, consultation.Gender, consultation.Subject, consultation.LiveConsultation, consultation.Description, consultation.Status, consultation.UserID, consultationID).Scan(&consultation).Error
	return consultation, err
}

func (r *repository) DeleteConsultation(consultation models.Consultation) (models.Consultation, error) {
	err := r.db.Delete(&consultation).Error

	return consultation, err
}
