package handlers

import (
	"encoding/json"
	replydto "hallocorona/dto/reply"
	dto "hallocorona/dto/result"
	"hallocorona/models"
	"hallocorona/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerReply struct {
	replyRepositories repositories.ReplyRepository
}

func HandlerReply(replyRepositories repositories.ReplyRepository) *handlerReply {
	return &handlerReply{replyRepositories}
}

func (h *handlerReply) FindReplies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	replies, err := h.replyRepositories.FindReplies()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var result []replydto.ReplyResponse

	for _, reply := range replies {
		user, _ := h.replyRepositories.GetReplyAuthor(reply.UserID)
		result = append(result, replydto.ReplyResponse{
			ID:        reply.ID,
			Response:  reply.Response,
			MeetLink:  reply.MeetLink,
			MeetType:  reply.MeetType,
			User:      user,
			CreatedAt: reply.CreatedAt,
			UpdatedAt: reply.UpdatedAt,
		})
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: result}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerReply) GetReply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	reply, err := h.replyRepositories.GetReply(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := h.replyRepositories.GetReplyAuthor(reply.UserID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	result := replydto.ReplyResponse{
		ID:        reply.ID,
		Response:  reply.Response,
		MeetLink:  reply.MeetLink,
		MeetType:  reply.MeetType,
		User:      user,
		CreatedAt: reply.CreatedAt,
		UpdatedAt: reply.UpdatedAt,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: result}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerReply) CreateReply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value(string("userInfo")).(jwt.MapClaims)
	userInfoId := userInfo["id"].(float64)

	request := new(replydto.CreateReplyRequest)

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

	reply := models.Reply{
		Response: request.Response,
		MeetLink: request.MeetLink,
		MeetType: request.MeetType,
		UserID:   int(userInfoId),
	}

	data, err := h.replyRepositories.CreateReply(reply)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := h.replyRepositories.GetReplyAuthor(data.UserID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	result := replydto.ReplyResponse{
		ID:       data.ID,
		Response: data.Response,
		MeetLink: data.MeetLink,
		MeetType: data.MeetType,
		User:     user,
	}

	w.WriteHeader(http.StatusCreated)
	response := dto.SuccessResult{Code: http.StatusCreated, Data: result}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerReply) UpdateReply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value(string("userInfo")).(jwt.MapClaims)
	userInfoId := userInfo["id"].(float64)

	request := new(replydto.UpdateReplyRequest)

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	reply, err := h.replyRepositories.GetReply(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Response != "" {
		reply.Response = request.Response
	}

	if request.MeetLink != "" {
		reply.MeetLink = request.MeetLink
	}

	if request.MeetType != "" {
		reply.MeetType = request.MeetType
	}

	if request.UserID != 0 {
		reply.UserID = int(userInfoId)
	}

	if reply.Response == "" || reply.MeetLink == "" || reply.MeetType == "" || reply.UserID == 0 {
		validation := validator.New()
		err := validation.Struct(request)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	data, err := h.replyRepositories.UpdateReply(reply)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := h.replyRepositories.GetReplyAuthor(data.UserID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	result := replydto.ReplyResponse{
		ID:       data.ID,
		Response: data.Response,
		MeetLink: data.MeetLink,
		MeetType: data.MeetType,
		User:     user,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: result}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerReply) DeleteReply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	reply, err := h.replyRepositories.GetReply(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.replyRepositories.DeleteReply(reply)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := h.replyRepositories.GetReplyAuthor(data.UserID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	result := replydto.ReplyResponse{
		ID:       data.ID,
		Response: data.Response,
		MeetLink: data.MeetLink,
		MeetType: data.MeetType,
		User:     user,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: result}
	json.NewEncoder(w).Encode(response)
}
