package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/hammadmajid/zabcover/internal/app"
)

func SetupRoutes(application *app.App) *chi.Mux {
	router := chi.NewRouter()
	handler := NewHandler(application.Logger, application.Template)

	router.Get("/", handler.Root)
	router.Get("/health", handler.Health)
	router.Get("/favicon.svg", handler.Favicon)
	router.Get("/robots.txt", handler.Robots)
	router.Get("/terminal.css", handler.CSS)
	router.Get("/form.js", handler.FormJs)

	return router
}
