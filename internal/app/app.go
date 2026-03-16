package app

import (
	"log"
	"zabdoc/internal/api"
)

type App struct {
	Logger     *log.Logger
	APIHandler *api.Handler
}

// NewApp constructs an App with injected dependencies.
func NewApp(logger *log.Logger, handler *api.Handler) *App {
	return &App{
		Logger:     logger,
		APIHandler: handler,
	}
}
