package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Secret []byte
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: []byte(secret),
	}
}

func (j *JWT) Create(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})
	tokenString, err := token.SignedString(j.Secret)
	if err != nil {
		return "", errors.New("ERROR")
	}
	return tokenString, nil
}
