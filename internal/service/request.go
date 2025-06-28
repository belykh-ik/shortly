package service

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func RequestJson[T any](req *http.Request) (*T, error) {
	var request T
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	v := validator.New()
	err = v.Struct(request)
	if err != nil {
		return nil, err
	}
	return &request, nil
}
