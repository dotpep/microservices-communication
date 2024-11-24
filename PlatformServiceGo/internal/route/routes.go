package route

import (
	"encoding/json"
	"exampleFirst/internal/database"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	db database.Service
}

func NewRouter() *Router {
	return &Router{}
}

func (ro *Router) RegisterRoutes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", ro.HelloWorldHandler)

	router.Get("/health", ro.healthHandler)

	return router
}

func (ro *Router) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (ro *Router) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(ro.db.Health())
	_, _ = w.Write(jsonResp)
}
