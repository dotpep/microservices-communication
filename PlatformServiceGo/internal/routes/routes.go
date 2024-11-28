package routes

import (
	"net/http"

	"github.com/dotpep/microservices-communication/PlatformServiceGo/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	platformHandler *handlers.PlatformHandler
	appHandler      *handlers.AppHandler
}

func NewRouter(
	platformHandler *handlers.PlatformHandler,
	appHandler *handlers.AppHandler,
) *Router {
	return &Router{
		platformHandler: platformHandler,
		appHandler:      appHandler,
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.RegisterRoutes().ServeHTTP(w, req)
}

func (r *Router) RegisterRoutes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Initialize handlers

	// Routes
	router.Get("/", r.appHandler.HelloWorldHandler)
	router.Get("/health", r.appHandler.HealthHandler)

	router.Get("/platforms", r.platformHandler.GetAllPlatforms)

	return router
}
