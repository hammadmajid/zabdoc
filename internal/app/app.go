package app

import (
	"html/template"
	"log"
	"os"
)

type App struct {
	Logger   *log.Logger
	Template *template.Template
}

func NewApp() (*App, error) {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	tpl, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		return nil, err
	}

	return &App{
		Logger:   logger,
		Template: tpl,
	}, nil
}
