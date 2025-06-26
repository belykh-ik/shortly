package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func IsAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			fmt.Println("Is not auth")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		auth := strings.TrimPrefix(authorization, "Bareer ")
		fmt.Println(auth)
		next.ServeHTTP(w, r)
	})
}
