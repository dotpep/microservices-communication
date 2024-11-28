package config

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"github.com/dotpep/microservices-communication/PlatformServiceGo/internal/database"
)

type Server struct {
	port int

	db database.Service
}

func NewServer(handler http.Handler) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("API_PORT"))

	return &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}
