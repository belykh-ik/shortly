package handlers

import (
	"api/shorturl/internal/models"
	"api/shorturl/internal/service"
	"fmt"
	"net/http"
)

type AuthHandler struct {
	Db *models.Config
}

func RegisterAuthRoutes(router *http.ServeMux, db *models.Config) {
	handler := &AuthHandler{
		Db: db,
	}
	router.HandleFunc("POST /auth/login", handler.login())
	router.HandleFunc("POST /auth/register", handler.register())
}

func (h *AuthHandler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Login")
		s, err := service.RequestJson[models.LoginRequest](req)
		if err != nil {
			service.ResponseJson(w, err, http.StatusBadRequest)
		}
		service.ResponseJson(w, s, http.StatusOK)
	}
}

func (h *AuthHandler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Register")
		s, err := service.RequestJson[models.RegisterRequest](req)
		if err != nil {
			service.ResponseJson(w, err, http.StatusBadRequest)
		}
		service.ResponseJson(w, s, http.StatusOK)
	}
}
