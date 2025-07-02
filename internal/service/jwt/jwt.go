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

func (j *JWT) Parse(token string) (bool, string) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, ""
	}
	email, _ := t.Claims.(jwt.MapClaims)["email"].(string)
	return t.Valid, email
}
