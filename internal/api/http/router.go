package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/hammadmajid/zabcover/internal/app"
	"github.com/hammadmajid/zabcover/internal/middleware"
)

func SetupRoutes(application *app.App) *chi.Mux {
	router := chi.NewRouter()
	handler := NewHandler(application.Logger)

	router.Use(middleware.LoggingMiddleware(application))

	router.Get("/", handler.Root)
	router.Get("/health", handler.Health)
	router.Get("/favicon.svg", handler.Favicon)
	router.Get("/robots.txt", handler.Robots)
	router.Get("/terminal.css", handler.CSS)
	router.Get("/form.js", handler.FormJs)
	router.Get("/storage.js", handler.StorageJs)

	router.Post("/assignment", handler.Assignment)

	return router
}
