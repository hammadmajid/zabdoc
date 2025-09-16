package app

import (
	"log"

	"zabdoc/internal/api/http"
)

type App struct {
	Logger     *log.Logger
	APIHandler http.Handler
}

// NewApp constructs an App with injected dependencies.
func NewApp(logger *log.Logger, handler http.Handler) *App {
	return &App{
		Logger:     logger,
		APIHandler: handler,
	}
}
