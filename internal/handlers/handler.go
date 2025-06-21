package handlers

import (
	"api/shorturl/internal/models"
	"api/shorturl/internal/service"
	"fmt"
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
	router.HandleFunc("GET /test", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Hello World")
	})
	router.HandleFunc("GET /link/{hash}", handler.getUrlByHash)
	router.HandleFunc("POST /create", handler.createUrl)
	router.HandleFunc("PATCH /update/{id}", handler.updateUrl)
	router.HandleFunc("DELETE /delete/{id}", handler.deleteLink)
}

func (h handlers) getUrlByHash(w http.ResponseWriter, req *http.Request) {
	hash := req.PathValue("hash")
	originalLink := h.link.LinkGet(hash)
	http.Redirect(w, req, originalLink.Url, http.StatusPermanentRedirect)
}

func (h handlers) createUrl(w http.ResponseWriter, req *http.Request) {
	link, err := service.RequestJson[models.Link](req)
	if err != nil {
		service.ResponseJson(w, err)
	}
	h.link.LinkCreate(link)
	service.ResponseJson(w, link)
}

func (h handlers) updateUrl(w http.ResponseWriter, req *http.Request) {
	link, err := service.RequestJson[models.Link](req)
	if err != nil {
		service.ResponseJson(w, err)
	}
	idString := req.PathValue("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		service.ResponseJson(w, err)
	}
	h.link.LinkUpdate(link, &id)
	service.ResponseJson(w, link)
}

func (h handlers) deleteLink(w http.ResponseWriter, req *http.Request) {
	idString := req.PathValue("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		service.ResponseJson(w, err)
	}
	err = h.link.LinkDelete(&id)
	if err != nil {
		service.ResponseJson(w, err)
	}
}
