package app

import (
	"log"
	"os"

	"zabdoc/internal/api/http"
	"zabdoc/internal/controllers"
)

type App struct {
	Logger     *log.Logger
	Services   controllers.Services
	APIHandler http.Handler
}

func NewApp() (*App, error) {
	l := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	s := controllers.NewServices(l)
	h := http.NewHandler(l, s)

	return &App{
		Logger:     l,
		Services:   s,
		APIHandler: h,
	}, nil
}
