package http

import (
	"fmt"
	"log"
	"net/http"

	"zabdoc/internal/api/dto"
	"zabdoc/internal/services"
)

type Handler struct {
	logger            *log.Logger
	validationService *services.ValidationService
	fileService       *services.FileService
	PdfService        *services.PDFService
	markdownService   *services.MarkdownService
}

func NewHandler(logger *log.Logger) Handler {
	return Handler{
		logger:            logger,
		validationService: services.NewValidationService(),
		fileService:       services.NewFileService(logger),
		PdfService:        services.NewPDFService(),
		markdownService:   services.NewMarkdownService(),
	}
}

const maxMultipartMemory int64 = 100 << 20 // 100 MBs

func (h Handler) Generate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(maxMultipartMemory); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		h.logger.Printf("parse multipart form: %v", err)
		return
	}

	data := dto.GenerateRequest{
		StudentName:  r.FormValue("studentName"),
		RegNo:        r.FormValue("regNo"),
		Class:        r.FormValue("class"),
		Course:       r.FormValue("course"),
		CourseCode:   r.FormValue("courseCode"),
		Instructor:   r.FormValue("instructor"),
		DocType:      r.FormValue("type"),
		Number:       r.FormValue("number"),
		Date:         r.FormValue("date"),
		Marks:        r.FormValue("marks"),
		ProjectTitle: r.FormValue("projectTitle"),
	}

	// Validate form data
	if err := h.validationService.ValidateGenerateRequest(&data); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		h.logger.Printf("validation error: %v", err)
		return
	}

	var sections []dto.Section
	for i := range 15 { // max 15 sections as per frontend
		contentKey := fmt.Sprintf("section-%d-content", i)
		filesKey := fmt.Sprintf("section-%d-files", i)

		content := r.FormValue(contentKey)
		if content == "" {
			continue
		}

		html, err := h.markdownService.ParseAndSanitize(content)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to parse markdown for section %d: %v", i, err), http.StatusBadRequest)
			h.logger.Printf("markdown parse (section %d): %v", i, err)
			return
		}

		section := dto.Section{
			Index:   i,
			Content: html,
		}

		// Process uploaded images using FileService
		if r.MultipartForm != nil && r.MultipartForm.File[filesKey] != nil {
			images, err := h.fileService.ProcessUploadedImages(r.MultipartForm.File[filesKey], filesKey)
			if err != nil {
				h.logger.Printf("file processing error: %v", err)
			} else {
				section.Images = append(section.Images, images...)
			}
		}

		sections = append(sections, section)
	}

	data.Sections = sections

	pdf, err := h.PdfService.GeneratePDF(data)
	if err != nil {
		http.Error(w, "failed to generate PDF", http.StatusInternalServerError)
		h.logger.Printf("generate PDF: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=document.pdf")
	if _, writeErr := w.Write(pdf); writeErr != nil {
		h.logger.Printf("write PDF: %v", writeErr)
	}
}

func (h Handler) Health(w http.ResponseWriter, r *http.Request) {
	if h.validationService == nil {
		http.Error(w, "ValidationService unavailable", http.StatusServiceUnavailable)
		h.logger.Println("Health check failed: ValidationService unavailable")
		return
	}
	if h.fileService == nil {
		http.Error(w, "FileService unavailable", http.StatusServiceUnavailable)
		h.logger.Println("Health check failed: FileService unavailable")
		return
	}
	if h.PdfService == nil {
		http.Error(w, "PDFService unavailable", http.StatusServiceUnavailable)
		h.logger.Println("Health check failed: PDFService unavailable")
		return
	}
	if h.markdownService == nil {
		http.Error(w, "MarkdownService unavailable", http.StatusServiceUnavailable)
		h.logger.Println("Health check failed: MarkdownService unavailable")
		return
	}

	// All dependencies healthy
	fmt.Fprintf(w, "Status is available")
}
