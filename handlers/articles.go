package handlers

import (
	"context"
	"encoding/json"
	articledto "hallocorona/dto/article"
	dto "hallocorona/dto/result"
	"hallocorona/models"
	"hallocorona/repositories"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerArticle struct {
	ArticleRepositories repositories.ArticleRepository
}

func HandlerArticle(ArticleRepositories repositories.ArticleRepository) *handlerArticle {
	return &handlerArticle{ArticleRepositories}
}

func (h *handlerArticle) FindArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	articles, err := h.ArticleRepositories.FindArticles()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var result []articledto.ArticleResponse

	for _, article := range articles {
		user, _ := h.ArticleRepositories.GetArticleAuthor(article.UserID)
		result = append(result, articledto.ArticleResponse{
			ID:          article.ID,
			Title:       article.Title,
			Description: article.Description,
			Image:       article.Image,
			User:        user,
		})
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: result}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerArticle) GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	article, err := h.ArticleRepositories.GetArticle(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := h.ArticleRepositories.GetArticleAuthor(article.UserID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	result := articledto.ArticleResponse{
		ID:          article.ID,
		Title:       article.Title,
		Image:       article.Image,
		Description: article.Description,
		User:        user,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: result}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerArticle) CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContext := r.Context().Value(string("dataFile"))
	filepath := dataContext.(string)

	userInfo := r.Context().Value(string("userInfo")).(jwt.MapClaims)
	userInfoId := userInfo["id"].(float64)

	request := articledto.CreateArticleRequest{
		Title:       r.FormValue("title"),
		Image:       filepath,
		Description: r.FormValue("description"),
		UserID:      int(userInfoId),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "hallocorona"})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	article := models.Article{
		Title:       request.Title,
		Image:       resp.SecureURL,
		Description: request.Description,
		UserID:      request.UserID,
	}

	data, err := h.ArticleRepositories.CreateArticle(article)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := h.ArticleRepositories.GetArticleAuthor(data.UserID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	result := articledto.ArticleResponse{
		ID:          data.ID,
		Title:       data.Title,
		Image:       data.Image,
		Description: data.Description,
		User:        user,
	}

	w.WriteHeader(http.StatusCreated)
	response := dto.SuccessResult{Code: http.StatusCreated, Data: result}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerArticle) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContext := r.Context().Value(string("dataFile"))
	filepath := dataContext.(string)

	userInfo := r.Context().Value(string("userInfo")).(jwt.MapClaims)
	userInfoId := userInfo["id"].(float64)

	request := articledto.UpdateArticleRequest{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		UserID:      int(userInfoId),
		Image:       filepath,
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	article, err := h.ArticleRepositories.GetArticle(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "hallocorona"})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Title != "" {
		article.Title = request.Title
	}

	if request.UserID != 0 {
		article.UserID = request.UserID
	}

	if request.Image != "" {
		article.Image = resp.SecureURL
	}

	if request.Description != "" {
		article.Description = request.Description
	}

	if article.Title == "" || article.Description == "" || article.Image == "" || article.UserID == 0 {
		validation := validator.New()
		err := validation.Struct(request)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	data, err := h.ArticleRepositories.UpdateArticle(article)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := h.ArticleRepositories.GetArticleAuthor(data.UserID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	result := articledto.ArticleResponse{
		ID:          data.ID,
		Title:       data.Title,
		Image:       data.Image,
		Description: data.Description,
		User:        user,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: result}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerArticle) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	article, err := h.ArticleRepositories.GetArticle(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.ArticleRepositories.DeleteArticle(article)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := h.ArticleRepositories.GetArticleAuthor(data.UserID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	result := articledto.ArticleResponse{
		ID:          data.ID,
		Title:       data.Title,
		Image:       data.Image,
		Description: data.Description,
		User:        user,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: result}
	json.NewEncoder(w).Encode(response)
}
