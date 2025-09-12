package http

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

				// Read file content
				fileData, err := io.ReadAll(file)
				if err != nil {
					h.logger.Printf("Error reading file: %v", err)
					continue
				}

				// Create temporary file
				ext := filepath.Ext(fileHeader.Filename)
				if ext == "" {
					ext = ".jpg" // default extension
				}
				tmpFile, err := os.CreateTemp("", "task_image_*"+ext)
				if err != nil {
					h.logger.Printf("Error creating temp file: %v", err)
					continue
				}

				// Write image data to temp file
				if _, err := tmpFile.Write(fileData); err != nil {
					h.logger.Printf("Error writing temp file: %v", err)
					os.Remove(tmpFile.Name())
					tmpFile.Close()
					continue
				}
				tmpFile.Close()

				// Store file path for later use in template
				task.Images = append(task.Images, "file://"+tmpFile.Name())
			}
		}

		tasks = append(tasks, task)
	}

	labTask.Tasks = tasks

	h.logger.Printf("Received form data: %+v", labTask)

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
