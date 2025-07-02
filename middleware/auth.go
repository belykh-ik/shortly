package middleware

import (
	"api/shorturl/internal/models"
	"api/shorturl/internal/service/jwt"
	"context"
	"fmt"
	"net/http"
	"strings"
)

type key string

const (
	KEY key = "default"
)

func IsAuth(config *models.Config, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			fmt.Println("Is not auth")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(authorization, "Bareer ")
		ok, email := jwt.NewJWT(config.Secret).Parse(token)
		if !ok {
			fmt.Println("TOKEN_IS_NOT_VALID!")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), KEY, email)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
