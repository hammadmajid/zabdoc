package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hammadmajid/zabcover/internal/app"
)

func SetupRoutes(application *app.App) *chi.Mux {
	router := chi.NewRouter()
	handler := NewHandler(application.Logger)

	router.Get("/", handler.Root)
	router.Get("/health", HealthCheck)

	return router
}

//goland:noinspection ALL
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available")
}
