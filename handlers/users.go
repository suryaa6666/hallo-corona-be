package handlers

import (
	"context"
	"encoding/json"
	dto "hallocorona/dto/result"
	usersdto "hallocorona/dto/users"
	"hallocorona/models"
	"hallocorona/repositories"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handler struct {
	UserRepositories repositories.UserRepository
}

func HandlerUser(UserRepositories repositories.UserRepository) *handler {
	return &handler{UserRepositories}
}

func (h *handler) FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.UserRepositories.FindUsers()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: users}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.UserRepositories.GetUser(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: user}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(usersdto.CreateUserRequest)

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

	if _, err := h.UserRepositories.UserCheckEmail(request.Email); err == nil {
		w.WriteHeader(http.StatusForbidden)
		response := dto.ErrorResult{Code: http.StatusForbidden, Message: "Email already registered!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	user := models.User{
		FullName: request.FullName,
		Email:    request.Email,
		Username: request.Username,
		Password: request.Password,
		ListAs:   request.ListAs,
		Gender:   request.Gender,
		Phone:    request.Phone,
		Address:  request.Address,
		Image:    "",
		Role:     request.Role,
	}

	data, err := h.UserRepositories.CreateUser(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := dto.SuccessResult{Code: http.StatusCreated, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContext := r.Context().Value(string("dataFile"))
	request := usersdto.UpdateUserRequest{
		FullName: r.FormValue("fullName"),
		Email:    r.FormValue("email"),
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
		ListAs:   r.FormValue("listAs"),
		Gender:   r.FormValue("gender"),
		Phone:    r.FormValue("phone"),
		Address:  r.FormValue("address"),
	}

	if dataContext != nil {
		filepath := dataContext.(string)
		request := usersdto.UpdateUserRequest{
			FullName: r.FormValue("fullName"),
			Email:    r.FormValue("email"),
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
			ListAs:   r.FormValue("listAs"),
			Gender:   r.FormValue("gender"),
			Phone:    r.FormValue("phone"),
			Address:  r.FormValue("address"),
			Image:    filepath,
		}

		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		user, err := h.UserRepositories.GetUser(id)

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

		if request.FullName != "" {
			user.FullName = request.FullName
		}

		if request.Email != "" {
			user.Email = request.Email
		}

		if request.Username != "" {
			user.Username = request.Username
		}

		if request.Password != "" {
			user.Password = request.Password
		}

		if request.ListAs != "" {
			user.ListAs = request.ListAs
		}

		if request.Gender != "" {
			user.Gender = request.Gender
		}

		if request.Phone != "" {
			user.Phone = request.Phone
		}

		if request.Address != "" {
			user.Address = request.Address
		}

		if request.Image != "" {
			user.Image = resp.SecureURL
		}

		if user.FullName == "" || user.Email == "" || user.Username == "" || user.Password == "" || user.ListAs == "" || user.Gender == "" || user.Phone == "" || user.Address == "" {
			validation := validator.New()
			err := validation.Struct(request)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
				json.NewEncoder(w).Encode(response)
				return
			}
		}

		data, err := h.UserRepositories.UpdateUser(user)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Code: http.StatusOK, Data: data}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.UserRepositories.GetUser(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.FullName != "" {
		user.FullName = request.FullName
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Username != "" {
		user.Username = request.Username
	}

	if request.Password != "" {
		user.Password = request.Password
	}

	if request.ListAs != "" {
		user.ListAs = request.ListAs
	}

	if request.Gender != "" {
		user.Gender = request.Gender
	}

	if request.Phone != "" {
		user.Phone = request.Phone
	}

	if request.Address != "" {
		user.Address = request.Address
	}

	if user.FullName == "" || user.Email == "" || user.Username == "" || user.Password == "" || user.ListAs == "" || user.Gender == "" || user.Phone == "" || user.Address == "" {
		validation := validator.New()
		err := validation.Struct(request)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	data, err := h.UserRepositories.UpdateUser(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.UserRepositories.GetUser(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.UserRepositories.DeleteUser(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
