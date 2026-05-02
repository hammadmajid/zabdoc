package router

import (
	"zabdoc/internal/application"
	"zabdoc/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *application.Application) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.LoggingMiddleware(app.Logger))

	router.Get("/health", app.Handler.Health)

	router.Post("/document", app.Handler.Document)
	router.Post("/scrape", app.Handler.Scrape)

	return router
}
