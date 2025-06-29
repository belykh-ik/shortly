package middleware

import (
	"api/shorturl/internal/models"
	"api/shorturl/internal/service/jwt"
	"fmt"
	"net/http"
	"strings"
)

func IsAuth(config *models.Config, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			fmt.Println("Is not auth")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		token := strings.TrimPrefix(authorization, "Bareer ")
		if !jwt.NewJWT(config.Secret).Parse(token) {
			fmt.Println("TOKEN_IS_NOT_VALID!")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
