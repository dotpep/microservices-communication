package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dotpep/microservices-communication/PlatformServiceGo/internal/converters"
	"github.com/dotpep/microservices-communication/PlatformServiceGo/internal/models"
	"github.com/dotpep/microservices-communication/PlatformServiceGo/internal/repositories"
	"github.com/go-chi/chi/v5"
)

type PlatformHandler struct {
	platformRepo repositories.IPlatformRepo
}

func NewPlatformHandler(repo repositories.IPlatformRepo) *PlatformHandler {
	return &PlatformHandler{platformRepo: repo}
}

func (p *PlatformHandler) GetAllPlatformsHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	platforms, err := p.platformRepo.GetAllPlatforms(ctx)
	if err != nil {
		log.Printf("Error retrieving platforms: %v\n", err)
		http.Error(w, "Failed to retrieve platforms", http.StatusInternalServerError)
		return
	}

	response := converters.DatabaseListPlatformsToListPlatforms(platforms)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v\n", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (p *PlatformHandler) GetPlatformByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	platformIDStr := chi.URLParam(r, "platformID")
	platformID, err := strconv.Atoi(platformIDStr)
	if err != nil {
		log.Printf("Couldn't parse platform id: %v", err)
		http.Error(w, "Failed to parse platform ID", http.StatusInternalServerError)
		return
	}

	// TODO: FIX http error sending and make well error and send it to client via REST json
	platform, err := p.platformRepo.GetPlatformByID(ctx, platformID)
	if err != nil {
		log.Printf("Error retrieving specific platform by id: %v\n", err)
		http.Error(w, "Failed to retrieve specific platform by id", http.StatusInternalServerError)
		return
	}

	response := converters.DatabaseOnePlatformToOnePlatform(*platform)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v\n", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (p *PlatformHandler) CreatePlatformHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var platform models.Platform

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&platform)
	if err != nil {
		log.Printf("Error parsing JSON: %v\n", err)
		http.Error(w, "Failed to parse json", http.StatusInternalServerError)
		return
	}

	// TODO: DTO of CreatePlatformInput and also others like (GetPlatformOutput)
	// TODO: Complete with Implementation of CreatePlatformHandler endpoint

	err = p.platformRepo.CreatePlatform(ctx, &platform)
	if err != nil {
		log.Printf("Error with creating platform: %v\n", err)
		http.Error(w, "Failed to create platform", http.StatusInternalServerError)
		return
	}
}
