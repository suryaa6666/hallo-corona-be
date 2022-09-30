package middleware

import (
	"context"
	"encoding/json"
	dto "hallocorona/dto/result"
	jwtToken "hallocorona/pkg/jwt"
	"net/http"
	"strings"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Auth(next http.HandlerFunc, status string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		token := r.Header.Get("Authorization")

		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "unauthorized"}
			json.NewEncoder(w).Encode(response)
			return
		}

		token = strings.Split(token, " ")[1]
		claims, err := jwtToken.DecodeToken(token)

		if claims["role"] == "member" && status == "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			response := Result{Code: http.StatusUnauthorized, Message: "role unaccepted"}
			json.NewEncoder(w).Encode(response)
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			response := Result{Code: http.StatusUnauthorized, Message: "unauthorized"}
			json.NewEncoder(w).Encode(response)
			return
		}

		userInfo := "userInfo"
		ctx := context.WithValue(r.Context(), userInfo, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
