package router

import (
	"net/http"

	"zabdoc/internal/app"
	"zabdoc/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(application *app.App) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.LoggingMiddleware(application.Logger))

	router.Get("/", application.APIHandler.Root)
	router.Get("/about", application.APIHandler.About)
	router.Get("/health", application.APIHandler.Health)

	// Serve static files
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	router.Handle("/css/*", http.StripPrefix("/css/", http.FileServer(http.Dir("web/css"))))
	router.Handle("/js/*", http.StripPrefix("/js/", http.FileServer(http.Dir("web/js"))))

	router.Post("/assignment", application.APIHandler.Assignment)

	return router
}
