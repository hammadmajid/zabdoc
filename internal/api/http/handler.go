package http

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Handler struct {
	logger *log.Logger
}

func NewHandler(logger *log.Logger) Handler {
	return Handler{
		logger: logger,
	}
}

//goland:noinspection GoUnusedParameter
func (h Handler) Root(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("web/templates/index.html"))

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := tpl.Execute(w, nil); err != nil {
		http.Error(w, "template error", http.StatusInternalServerError)
		log.Printf("template execute: %v", err)
		return
	}
}

//goland:noinspection ALL
func (h Handler) Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available")
}

func (h Handler) Favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/static/favicon.svg")
}

func (h Handler) Robots(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/static/robots.txt")
}

func (h Handler) CSS(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/css/terminal.css")
}
