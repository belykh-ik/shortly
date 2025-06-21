package middleware

import "net/http"

type ModWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *ModWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}
