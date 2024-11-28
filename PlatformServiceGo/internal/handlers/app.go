package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dotpep/microservices-communication/PlatformServiceGo/internal/database"
	"github.com/go-chi/chi/v5"
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

func (app *AppHandler) DatabaseHealthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(
		app.db.Health(),
	)
	_, _ = w.Write(jsonResp)
}

func (app *AppHandler) SayHelloToNameHandler(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	msg := fmt.Sprintf("Hello, greet to see you %v", name)

	resp := make(map[string]string)
	resp["message"] = msg

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
