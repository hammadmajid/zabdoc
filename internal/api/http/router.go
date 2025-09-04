package http

import (
	"net/http"

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

	// Serve static files
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	router.Handle("/css/*", http.StripPrefix("/css/", http.FileServer(http.Dir("web/css"))))
	router.Handle("/js/*", http.StripPrefix("/js/", http.FileServer(http.Dir("web/js"))))

	router.Post("/assignment", handler.Assignment)

	return router
}
