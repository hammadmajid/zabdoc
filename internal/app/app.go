package app

import (
	"log"
	"os"

	"zabdoc/internal/api/http"
)

type App struct {
	Logger     *log.Logger
	APIHandler http.Handler
}

func NewApp() (*App, error) {
	l := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	h := http.NewHandler(l)

	return &App{
		Logger:     l,
		APIHandler: h,
	}, nil
}
