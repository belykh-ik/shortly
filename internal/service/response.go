package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, data any) {
	switch v := data.(type) {
	case error:
		fmt.Println(v.Error())
		w.WriteHeader(403)
		json.NewEncoder(w).Encode(v.Error())
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(data)
	}
}
