package handlers

import (
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
	router.HandleFunc("DELETE /login/{id}", deleteLink)
	router.HandleFunc("POST /create/{url}", handler.createUrl)
}

func deleteLink(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	fmt.Println(id)
}

func (h handlers) createUrl(w http.ResponseWriter, req *http.Request) {
	url := req.PathValue("url")
	fmt.Println(url)
	h.link.Create()
}
