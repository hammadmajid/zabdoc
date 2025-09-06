package app

import (
	"log"
	"os"

	"zabdoc/internal/api/http"
	"zabdoc/internal/services"
)

type App struct {
	Logger     *log.Logger
	Services   services.Services
	APIHandler http.Handler
}

func NewApp() (*App, error) {
	l := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	s := services.NewServices(l)
	h := http.NewHandler(l, s)

	return &App{
		Logger:     l,
		Services:   s,
		APIHandler: h,
	}, nil
}
