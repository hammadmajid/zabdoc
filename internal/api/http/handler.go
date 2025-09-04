package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hammadmajid/zabcover/internal/services"
	"github.com/hammadmajid/zabcover/internal/utils"
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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := utils.TemplateFiles[utils.Index].Execute(w, nil); err != nil {
		http.Error(w, "template error", http.StatusInternalServerError)
		h.logger.Printf("template execute: %v", err)
		return
	}
}

func (h Handler) Assignment(w http.ResponseWriter, r *http.Request) {
	pdf, err := services.Generate(r.Context())
	if err != nil {
		http.Error(w, "failed to generate PDF", http.StatusInternalServerError)
		h.logger.Printf("generate PDF: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=assignment.pdf")
	w.Write(pdf)
}

//goland:noinspection ALL
func (h Handler) Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available")
}

func (h Handler) Favicon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	http.ServeFile(w, r, "web/static/favicon.svg")
}

func (h Handler) Robots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	http.ServeFile(w, r, "web/static/robots.txt")
}

func (h Handler) CSS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	http.ServeFile(w, r, "web/css/terminal.css")
}

func (h Handler) FormJs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	http.ServeFile(w, r, "web/js/form.js")
}
