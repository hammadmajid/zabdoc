package application

import (
	"log"
	"zabdoc/internal/api"
)

type Application struct {
	Logger  *log.Logger
	Handler *api.Handler
}

// NewApplication constructs an Application with injected dependencies.
func NewApplication(logger *log.Logger, handler *api.Handler) *Application {
	return &Application{
		Logger:  logger,
		Handler: handler,
	}
}
