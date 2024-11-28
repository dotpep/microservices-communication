package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dotpep/microservices-communication/PlatformServiceGo/internal/converters"
	"github.com/dotpep/microservices-communication/PlatformServiceGo/internal/repositories"
)

type PlatformHandler struct {
	platformRepo repositories.IPlatformRepo
}

func NewPlatformHandler(repo repositories.IPlatformRepo) *PlatformHandler {
	return &PlatformHandler{platformRepo: repo}
}

func (p *PlatformHandler) GetAllPlatforms(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	platforms, err := p.platformRepo.GetAllPlatforms(ctx)
	if err != nil {
		log.Printf("Error retrieving platforms: %v\n", err)
		http.Error(w, "Failed to retrieve platforms", http.StatusInternalServerError)
		return
	}

	response := converters.DatabasePlatformsToPlatforms(platforms)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v\n", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
