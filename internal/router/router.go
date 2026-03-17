package router

import (
	"zabdoc/internal/application"
	"zabdoc/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func SetupRoutes(app *application.Application) *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(middleware.CORSOptions))
	router.Use(middleware.LoggingMiddleware(app.Logger))

	router.Get("/health", app.Handler.Health)

	router.Post("/generate", app.Handler.Generate)
	router.Post("/scrap", app.Handler.Scrap)

	return router
}
