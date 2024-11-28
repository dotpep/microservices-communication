package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dotpep/microservices-communication/PlatformServiceGo/internal/database"
)

type AppHandler struct {
	db database.Service
}

func NewAppHandler(db database.Service) *AppHandler {
	return &AppHandler{db: db}
}

func (app *AppHandler) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (app *AppHandler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(app.db.Health())
	_, _ = w.Write(jsonResp)
}
