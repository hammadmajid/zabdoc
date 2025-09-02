package app

import (
	"log"
	"os"
)

type App struct {
	Logger *log.Logger
}

func NewApp() (*App, error) {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	return &App{
		Logger: logger,
	}, nil
}
