package handlers

import (
	"api/shorturl/internal/models"
	"api/shorturl/internal/service"
	"api/shorturl/middleware"
	"fmt"
	"net/http"
	"strconv"
)

type handlers struct {
	link   *service.LinkDeps
	config *models.Config
}

func RegisterRoutes(router *http.ServeMux, config *models.Config, link *service.LinkDeps) {
	handler := &handlers{
		link:   link,
		config: config,
	}
	router.Handle("GET /link/{hash}", middleware.IsAuth(config, handler.getUrlByHash()))
	router.Handle("GET /links", middleware.IsAuth(config, handler.getAllLinks()))
	router.Handle("POST /create", middleware.IsAuth(config, handler.createUrl()))
	router.Handle("PATCH /update/{id}", middleware.IsAuth(config, handler.updateUrl()))
	router.Handle("DELETE /delete/{id}", middleware.IsAuth(config, handler.deleteLink()))
}

func (h handlers) getUrlByHash() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		hash := req.PathValue("hash")
		originalLink, err := h.link.LinkGet(hash)
		if err != nil {
			service.ResponseJson(w, err, http.StatusBadRequest)
			return
		}
		// http.Redirect(w, req, originalLink.Url, http.StatusPermanentRedirect)
		service.ResponseJson(w, originalLink, http.StatusOK)
	})
}

func (h handlers) getAllLinks() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		limit, err := strconv.Atoi(req.URL.Query().Get("limit"))
		if err != nil {
			service.ResponseJson(w, err, http.StatusBadRequest)
			return
		}
		links := h.link.GetAllLinks(limit)
		service.ResponseJson(w, links, http.StatusOK)
	})
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
		emailAuth := req.Context().Value(middleware.KEY)
		fmt.Println(emailAuth)
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
