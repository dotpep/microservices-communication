package handlers

import "exampleFirst/internal/database"

type PlatformHandler struct {
	db database.Service
}

func NewPlatformHandler(db database.Service) *PlatformHandler {
	return &PlatformHandler{db: db}
}

//func (p *PlatformHandler) GetPlatforms(w http.ResponseWriter, r *http.Request) {

//}
