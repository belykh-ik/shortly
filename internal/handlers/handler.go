package handlers

import (
	"api/shorturl/internal/models"
	"api/shorturl/internal/service"
	"fmt"
	"net/http"
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
	router.HandleFunc("DELETE /login/{id}", deleteLink)
}

func (h handlers) getUrlByHash(w http.ResponseWriter, req *http.Request) {
	hash := req.PathValue("hash")
	originalLink := h.link.LinkGet(hash)
	http.Redirect(w, req, originalLink.Url, http.StatusPermanentRedirect)
}

func deleteLink(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	fmt.Println(id)
}

func (h handlers) createUrl(w http.ResponseWriter, req *http.Request) {
	link, err := service.RequestJson[models.Link](req)
	if err != nil {
		service.ResponseJson(w, err)
	}
	fmt.Println(link)
	h.link.LinkCreate(link)
}
