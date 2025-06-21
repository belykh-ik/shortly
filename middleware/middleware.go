package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		wrapper := &ModWriter{
			w,
			http.StatusContinue,
		}
		next.ServeHTTP(wrapper, req)
		log.Println(wrapper.statusCode, req.Method, req.URL.Path, time.Since(start))
	})
}
