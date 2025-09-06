package http

import (
	"fmt"
	"log"
	"net/http"

	"zabdoc/internal/api/dto"
	"zabdoc/internal/services"
	"zabdoc/internal/utils"
)

type Handler struct {
	logger   *log.Logger
	services services.Services
}

func NewHandler(logger *log.Logger, sv services.Services) Handler {
	return Handler{
		logger:   logger,
		services: sv,
	}
}

//goland:noinspection GoUnusedParameter
func (h Handler) Root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := utils.GetTemplate(utils.Index).Execute(w, nil); err != nil {
		http.Error(w, "template error", http.StatusInternalServerError)
		h.logger.Printf("template execute: %v", err)
		return
	}
}

func (h Handler) Assignment(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		h.logger.Printf("parse form: %v", err)
		return
	}

	assignment := dto.AssignmentRequest{
		StudentName: r.FormValue("studentName"),
		RegNo:       r.FormValue("regNo"),
		Class:       r.FormValue("class"),
		Course:      r.FormValue("course"),
		Instructor:  r.FormValue("instructor"),
		Number:      r.FormValue("number"),
		Date:        r.FormValue("date"),
	}

	pdf, err := h.services.AssignmentService.Generate(assignment)
	if err != nil {
		http.Error(w, "failed to generate PDF", http.StatusInternalServerError)
		h.logger.Printf("generate PDF: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=assignment.pdf")
	if _, writeErr := w.Write(pdf); writeErr != nil {
		h.logger.Printf("write PDF: %v", writeErr)
	}
}

//goland:noinspection ALL
func (h Handler) Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available")
}
