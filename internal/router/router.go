package router

import (
	"zabdoc/internal/app"
	"zabdoc/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(application *app.App) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.LoggingMiddleware(application.Logger))

	router.Get("/health", application.APIHandler.Health)

	router.Post("/assignment", application.APIHandler.AssignmentPdf)

	return router
}
