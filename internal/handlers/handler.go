package handlers

import (
	"api/shorturl/internal/models"
	"api/shorturl/internal/service"
	"api/shorturl/middleware"
	"net/http"
	"strconv"
)

type handlers struct {
	link *service.LinkDeps
}

func RegisterRoutes(router *http.ServeMux, link *service.LinkDeps) {
	handler := &handlers{
		link: link,
	}
	router.HandleFunc("GET /link/{hash}", handler.getUrlByHash)
	router.Handle("POST /create", middleware.IsAuth(handler.createUrl()))
	router.Handle("PATCH /update/{id}", middleware.IsAuth(handler.updateUrl()))
	router.Handle("DELETE /delete/{id}", middleware.IsAuth(handler.deleteLink()))
}

func (h handlers) getUrlByHash(w http.ResponseWriter, req *http.Request) {
	hash := req.PathValue("hash")
	originalLink := h.link.LinkGet(hash)
	// http.Redirect(w, req, originalLink.Url, http.StatusPermanentRedirect)
	service.ResponseJson(w, originalLink, http.StatusOK)
}

func (h handlers) createUrl() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		link, err := service.RequestJson[models.Url](req)
		if err != nil {
			service.ResponseJson(w, err, http.StatusBadRequest)
		}
		h.link.LinkCreate(link)
		service.ResponseJson(w, link, http.StatusCreated)
	})
}

func (h handlers) updateUrl() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		link, err := service.RequestJson[models.Url](req)
		if err != nil {
			service.ResponseJson(w, err, http.StatusBadRequest)
		}
		idString := req.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			service.ResponseJson(w, err, http.StatusBadRequest)
		}
		h.link.LinkUpdate(link, &id)
		service.ResponseJson(w, link, http.StatusOK)
	})
}

func (h handlers) deleteLink() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		idString := req.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			service.ResponseJson(w, err, http.StatusBadRequest)
		}
		err = h.link.LinkDelete(&id)
		if err != nil {
			service.ResponseJson(w, err, http.StatusForbidden)
		}
	})
}
