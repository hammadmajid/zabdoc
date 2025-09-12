package http

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"path/filepath"

	"zabdoc/internal/api/dto"
	"zabdoc/internal/controllers"

	"github.com/yuin/goldmark"
)

type Handler struct {
	logger   *log.Logger
	services controllers.Services
}

func NewHandler(logger *log.Logger, sv controllers.Services) Handler {
	return Handler{
		logger:   logger,
		services: sv,
	}
}

func (h Handler) AssignmentPdf(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		h.logger.Printf("parse multipart form: %v", err)
		return
	}

	assignment := dto.AssignmentRequest{
		StudentName: r.FormValue("studentName"),
		RegNo:       r.FormValue("regNo"),
		Class:       r.FormValue("class"),
		Course:      r.FormValue("course"),
		CourseCode:  r.FormValue("courseCode"),
		Instructor:  r.FormValue("instructor"),
		Number:      r.FormValue("number"),
		Date:        r.FormValue("date"),
	}

	h.logger.Printf("Received form data: %+v", assignment)

	pdf, err := h.services.AssignmentCtrl.Generate(assignment)
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

func (h Handler) LabTask(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		h.logger.Printf("parse multipart form: %v", err)
		return
	}

	labTask := dto.LabTaskRequest{
		StudentName: r.FormValue("studentName"),
		RegNo:       r.FormValue("regNo"),
		Class:       r.FormValue("class"),
		Course:      r.FormValue("course"),
		CourseCode:  r.FormValue("courseCode"),
		Instructor:  r.FormValue("instructor"),
		Number:      r.FormValue("number"),
		Date:        r.FormValue("date"),
	}

	// Process tasks
	var tasks []dto.Task
	for i := 0; i < 15; i++ { // Max 15 tasks as per frontend
		contentKey := fmt.Sprintf("task-%d-content", i)
		filesKey := fmt.Sprintf("task-%d-files", i)

		content := r.FormValue(contentKey)
		if content == "" {
			continue // Skip empty tasks
		}

		// Convert markdown to HTML
		var buf bytes.Buffer
		if err := goldmark.Convert([]byte(content), &buf); err != nil {
			panic(err)
		}

		task := dto.Task{
			Index:   i,
			Content: buf.String(),
		}

		// Process uploaded images
		if files := r.MultipartForm.File[filesKey]; files != nil {
			for _, fileHeader := range files {
				file, err := fileHeader.Open()
				if err != nil {
					h.logger.Printf("Error opening file: %v", err)
					continue
				}
				defer file.Close()

				// Detect MIME type
				ext := filepath.Ext(fileHeader.Filename)
				mimeType := mime.TypeByExtension(ext)
				if mimeType == "" {
					mimeType = "image/jpeg" // fallback
				}

				// Read file content
				fileData, err := io.ReadAll(file)
				if err != nil {
					h.logger.Printf("Error reading file: %v", err)
					continue
				}

				if len(fileData) == 0 {
					h.logger.Printf("Empty image file: %s", fileHeader.Filename)
					continue
				}

				// Encode as base64
				base64Data := base64.StdEncoding.EncodeToString(fileData)

				// Store as TaskImage with base64 data and MIME type
				taskImage := dto.TaskImage{
					Data:     base64Data,
					MimeType: mimeType,
				}
				task.Images = append(task.Images, taskImage)

				h.logger.Printf("Image processed: %s (%d bytes → %d base64 chars)",
					mimeType, len(fileData), len(base64Data))
			}
		}

		tasks = append(tasks, task)
	}

	labTask.Tasks = tasks

	pdf, err := h.services.LabTaskCtrl.Generate(labTask)
	if err != nil {
		http.Error(w, "failed to generate PDF", http.StatusInternalServerError)
		h.logger.Printf("generate PDF: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=lab-task.pdf")
	if _, writeErr := w.Write(pdf); writeErr != nil {
		h.logger.Printf("write PDF: %v", writeErr)
	}
}

//goland:noinspection ALL
func (h Handler) Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available")
}
