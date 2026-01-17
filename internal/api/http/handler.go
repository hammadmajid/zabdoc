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
}

func NewHandler(logger *log.Logger) Handler {
	return Handler{
		logger:            logger,
		validationService: services.NewValidationService(),
		fileService:       services.NewFileService(logger),
		PdfService:        services.NewPDFService(),
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
		IsMultiMode:  r.FormValue("isMultiMode") == "true",
	}

	// Handle multiple students
	if data.IsMultiMode {
		var students []dto.Student
		for i := 1; i <= 6; i++ { // Max 6 students
			nameKey := fmt.Sprintf("student-%d-name", i)
			regNoKey := fmt.Sprintf("student-%d-regNo", i)

			name := r.FormValue(nameKey)
			regNo := r.FormValue(regNoKey)

			if name != "" && regNo != "" {
				students = append(students, dto.Student{
					Name:  name,
					RegNo: regNo,
				})
			}
		}
		data.Students = students
	}

	// Validate form data
	if err := h.validationService.ValidateGenerateRequest(&data); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		h.logger.Printf("validation error: %v", err)
		return
	}

	// Process uploaded images
	if r.MultipartForm != nil && r.MultipartForm.File["images"] != nil {
		images, err := h.fileService.ProcessUploadedImages(r.MultipartForm.File["images"])
		if err != nil {
			http.Error(w, "failed to process images", http.StatusBadRequest)
			h.logger.Printf("file processing error: %v", err)
			return
		}
		data.Images = images
	}

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

	// All dependencies healthy
	fmt.Fprintf(w, "Status is available")
}
