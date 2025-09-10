package router

import (
	"zabdoc/internal/app"
	"zabdoc/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func SetupRoutes(application *app.App) *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(middleware.CORSOptions))
	router.Use(middleware.LoggingMiddleware(application.Logger))

	router.Get("/health", application.APIHandler.Health)

	router.Post("/assignment", application.APIHandler.AssignmentPdf)

	return router
}
