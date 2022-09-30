package handlers

import (
	"encoding/json"
	consultationdto "hallocorona/dto/consultation"
	dto "hallocorona/dto/result"
	"hallocorona/models"
	"hallocorona/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerConsultation struct {
	consultationRepositories repositories.ConsultationRepository
}

func HandlerConsultation(consultationRepositories repositories.ConsultationRepository) *handlerConsultation {
	return &handlerConsultation{consultationRepositories}
}

func (h *handlerConsultation) FindConsultations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value(string("userInfo")).(jwt.MapClaims)
	userInfoId := userInfo["id"].(float64)

	consultations, err := h.consultationRepositories.FindConsultations(int(userInfoId))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var result []consultationdto.ConsultationResponse

	for _, consultation := range consultations {
		user, _ := h.consultationRepositories.GetConsultationAuthor(consultation.UserID)
		reply, _ := h.consultationRepositories.GetConsultationReply(consultation.ReplyID)
		result = append(result, consultationdto.ConsultationResponse{
			ID:               consultation.ID,
			FullName:         consultation.FullName,
			Phone:            consultation.Phone,
			BornDate:         consultation.BornDate,
			Age:              consultation.Age,
			Height:           consultation.Height,
			Weight:           consultation.Weight,
			Gender:           consultation.Gender,
			Subject:          consultation.Subject,
			LiveConsultation: consultation.LiveConsultation,
			Description:      consultation.Description,
			Status:           consultation.Status,
			Reply:            reply,
			User:             user,
			CreatedAt:        consultation.CreatedAt,
			UpdatedAt:        consultation.UpdatedAt,
		})
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: result}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerConsultation) GetConsultation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	consultation, err := h.consultationRepositories.GetConsultation(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := h.consultationRepositories.GetConsultationAuthor(consultation.UserID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var reply models.Reply

	if consultation.ReplyID != 0 {
		reply, err = h.consultationRepositories.GetConsultationReply(consultation.ReplyID)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	result := consultationdto.ConsultationResponse{
		ID:               consultation.ID,
		FullName:         consultation.FullName,
		Phone:            consultation.Phone,
		BornDate:         consultation.BornDate,
		Age:              consultation.Age,
		Height:           consultation.Height,
		Weight:           consultation.Weight,
		Gender:           consultation.Gender,
		Subject:          consultation.Subject,
		LiveConsultation: consultation.LiveConsultation,
		Description:      consultation.Description,
		Status:           consultation.Status,
		Reply:            reply,
		User:             user,
		CreatedAt:        consultation.CreatedAt,
		UpdatedAt:        consultation.UpdatedAt,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: result}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerConsultation) CreateConsultation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value(string("userInfo")).(jwt.MapClaims)
	userInfoId := userInfo["id"].(float64)

	request := new(consultationdto.CreateConsultationRequest)

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	consultation := models.Consultation{
		FullName:         request.FullName,
		Phone:            request.Phone,
		BornDate:         request.BornDate,
		Age:              request.Age,
		Height:           request.Height,
		Weight:           request.Weight,
		LiveConsultation: request.LiveConsultation,
		Gender:           request.Gender,
		Subject:          request.Subject,
		Description:      request.Description,
		Status:           "pending",
		UserID:           int(userInfoId),
	}

	data, err := h.consultationRepositories.CreateConsultation(consultation)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := h.consultationRepositories.GetConsultationAuthor(data.UserID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	result := consultationdto.ConsultationResponse{
		ID:               data.ID,
		FullName:         data.FullName,
		Phone:            data.Phone,
		BornDate:         data.BornDate,
		Age:              data.Age,
		Height:           data.Height,
		Weight:           data.Weight,
		Gender:           data.Gender,
		Subject:          data.Subject,
		LiveConsultation: data.LiveConsultation,
		Description:      data.Description,
		Status:           data.Status,
		User:             user,
	}

	w.WriteHeader(http.StatusCreated)
	response := dto.SuccessResult{Code: http.StatusCreated, Data: result}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerConsultation) UpdateConsultationStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(consultationdto.UpdateConsultationStatus)

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	consultation, err := h.consultationRepositories.GetConsultation(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Status != "" {
		consultation.Status = request.Status
	}

	if request.ReplyID != 0 {
		consultation.ReplyID = request.ReplyID
	}

	if consultation.Status == "" || consultation.ReplyID == 0 || consultation.UserID == 0 {
		validation := validator.New()
		err := validation.Struct(request)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	data, err := h.consultationRepositories.UpdateConsultationStatus(consultation)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := h.consultationRepositories.GetConsultationAuthor(data.UserID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var reply models.Reply

	if data.ReplyID != 0 {
		reply, err = h.consultationRepositories.GetConsultationReply(data.ReplyID)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	result := consultationdto.ConsultationResponse{
		ID:               consultation.ID,
		FullName:         consultation.FullName,
		Phone:            consultation.Phone,
		BornDate:         consultation.BornDate,
		Age:              consultation.Age,
		Height:           consultation.Height,
		Weight:           consultation.Weight,
		Gender:           consultation.Gender,
		Subject:          consultation.Subject,
		LiveConsultation: consultation.LiveConsultation,
		Description:      consultation.Description,
		Status:           consultation.Status,
		Reply:            reply,
		User:             user,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: result}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerConsultation) UpdateConsultation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(consultationdto.UpdateConsultationRequest)

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	consultation, err := h.consultationRepositories.GetConsultation(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.FullName != "" {
		consultation.FullName = request.FullName
	}

	if request.Phone != "" {
		consultation.Phone = request.Phone
	}

	if request.BornDate != 0 {
		consultation.BornDate = request.BornDate
	}

	if request.Age != 0 {
		consultation.Age = request.Age
	}

	if request.Weight != 0 {
		consultation.Weight = request.Weight
	}

	if request.Height != 0 {
		consultation.Height = request.Height
	}

	if request.Gender != "" {
		consultation.Gender = request.Gender
	}

	if request.Subject != "" {
		consultation.Subject = request.Subject
	}

	if request.Description != "" {
		consultation.Description = request.Description
	}

	if request.LiveConsultation != 0 {
		consultation.LiveConsultation = request.LiveConsultation
	}

	if consultation.FullName == "" || consultation.Phone == "" || consultation.BornDate == 0 || consultation.Age == 0 || consultation.Weight == 0 || consultation.Height == 0 || consultation.Gender == "" || consultation.Subject == "" || consultation.Description == "" || consultation.LiveConsultation == 0 {
		validation := validator.New()
		err := validation.Struct(request)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	data, err := h.consultationRepositories.UpdateConsultation(consultation, consultation.ID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := h.consultationRepositories.GetConsultationAuthor(data.UserID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var reply models.Reply

	if data.ReplyID != 0 {
		reply, err = h.consultationRepositories.GetConsultationReply(data.ReplyID)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	result := consultationdto.ConsultationResponse{
		ID:               consultation.ID,
		FullName:         consultation.FullName,
		Phone:            consultation.Phone,
		BornDate:         consultation.BornDate,
		Age:              consultation.Age,
		Height:           consultation.Height,
		Weight:           consultation.Weight,
		Gender:           consultation.Gender,
		Subject:          consultation.Subject,
		LiveConsultation: consultation.LiveConsultation,
		Description:      consultation.Description,
		Status:           consultation.Status,
		Reply:            reply,
		User:             user,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: result}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerConsultation) DeleteConsultation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	consultation, err := h.consultationRepositories.GetConsultation(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.consultationRepositories.DeleteConsultation(consultation)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := h.consultationRepositories.GetConsultationAuthor(data.UserID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var reply models.Reply

	if data.ReplyID != 0 {
		reply, err = h.consultationRepositories.GetConsultationReply(data.ReplyID)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	result := consultationdto.ConsultationResponse{
		ID:               consultation.ID,
		FullName:         consultation.FullName,
		Phone:            consultation.Phone,
		BornDate:         consultation.BornDate,
		Age:              consultation.Age,
		Height:           consultation.Height,
		Weight:           consultation.Weight,
		Gender:           consultation.Gender,
		Subject:          consultation.Subject,
		LiveConsultation: consultation.LiveConsultation,
		Description:      consultation.Description,
		Status:           consultation.Status,
		Reply:            reply,
		User:             user,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: result}
	json.NewEncoder(w).Encode(response)
}
