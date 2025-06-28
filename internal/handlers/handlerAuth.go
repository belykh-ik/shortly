package handlers

import (
	"api/shorturl/internal/models"
	"api/shorturl/internal/service"
	"api/shorturl/middleware"
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
	router.Handle("POST /auth/login", middleware.IsAuth(handler.login()))
	router.Handle("POST /auth/register", middleware.IsAuth(handler.register()))
}

func (h *AuthHandler) login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Login")
		s, err := service.RequestJson[models.LoginRequest](req)
		if err != nil {
			service.ResponseJson(w, err, http.StatusBadRequest)
		}
		service.ResponseJson(w, s, http.StatusOK)
	})
}

func (h *AuthHandler) register() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Register")
		s, err := service.RequestJson[models.RegisterRequest](req)
		if err != nil {
			service.ResponseJson(w, err, http.StatusBadRequest)
		}
		service.ResponseJson(w, s, http.StatusOK)
	})
}
