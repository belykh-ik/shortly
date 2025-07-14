package handleMessage

import (
	"api/shorturl/broker/models"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

type HandleMessageDeps struct {
	routrer *http.ServeMux
}

func NewHandleMessageDeps(routrer *http.ServeMux) *HandleMessageDeps {
	return &HandleMessageDeps{
		routrer: routrer,
	}
}

func (h *HandleMessageDeps) HandleMassage(msg []byte) *httptest.ResponseRecorder {
	msgUrl := &models.MessageConsumer{
		Url: string(msg),
	}
	msgByte, err := json.Marshal(msgUrl)
	if err != nil {
		log.Println(err)
	}
	// Create *http.Request
	req := httptest.NewRequest("POST", "/create", bytes.NewReader(msgByte))
	req.Header.Set("Content-Type", "application/json")
	// Add Authorization
	req.Header.Set("Authorization", os.Getenv("TOKEN"))
	// Create *httptest.ResponseRecorder (ResponseWriter interface)
	rr := httptest.NewRecorder()

	h.routrer.ServeHTTP(rr, req)
	fmt.Println(rr.Body.String())
	return rr
}
