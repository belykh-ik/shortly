package handlers

import (
	"api/shorturl/internal/models"
	"api/shorturl/internal/service"
	"api/shorturl/internal/service/jwt"
	"api/shorturl/middleware"
	"net/http"
)

type AuthHandler struct {
	Db             *models.Config
	userRepository *service.UserRepository
}

func RegisterAuthRoutes(router *http.ServeMux, db *models.Config, userRepository *service.UserRepository) {
	handler := &AuthHandler{
		Db:             db,
		userRepository: userRepository,
	}
	router.Handle("POST /auth/login", middleware.IsAuth(handler.login()))
	router.Handle("POST /auth/register", middleware.IsAuth(handler.register()))
}

func (h *AuthHandler) login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		data, err := service.RequestJson[models.LoginRequest](req)
		if err != nil {
			service.ResponseJson(w, err, http.StatusBadRequest)
		}
		userEmail, err := h.userRepository.LoginUser(data)
		if err != nil {
			service.ResponseJson(w, err, http.StatusUnauthorized)
			return
		}
		token, err := jwt.NewJWT(h.Db.Secret).Create(userEmail)
		if err != nil {
			service.ResponseJson(w, err, http.StatusBadGateway)
			return
		}
		service.ResponseJson(w, token, http.StatusOK)
	})
}

func (h *AuthHandler) register() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		data, err := service.RequestJson[models.RegisterRequest](req)
		if err != nil {
			service.ResponseJson(w, err, http.StatusBadRequest)
		}
		userEmail, err := h.userRepository.CreateUser(data)
		if err != nil {
			service.ResponseJson(w, err, http.StatusBadRequest)
			return
		}
		token, err := jwt.NewJWT(h.Db.Secret).Create(userEmail)
		if err != nil {
			service.ResponseJson(w, err, http.StatusBadGateway)
			return
		}
		service.ResponseJson(w, token, http.StatusCreated)
	})
}
