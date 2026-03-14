package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"zabdoc/internal/api/dto"
	"zabdoc/internal/services"
)

type Handler struct {
	logger            *log.Logger
	validationService *services.ValidationService
	fileService       *services.FileService
	PdfService        *services.PDFService
	scraper           *services.Scraper
}

func NewHandler(logger *log.Logger) Handler {
	return Handler{
		logger:            logger,
		validationService: services.NewValidationService(),
		fileService:       services.NewFileService(logger),
		PdfService:        services.NewPDFService(),
		scraper:           services.NewScraper(),
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
		Class:      r.FormValue("class"),
		Course:     r.FormValue("course"),
		CourseCode: r.FormValue("courseCode"),
		Instructor: r.FormValue("instructor"),
		DocType:    r.FormValue("type"),
		Number:     r.FormValue("number"),
		Date:       r.FormValue("date"),
		Marks:      r.FormValue("marks"),
	}

	// Check if multi-student mode (student-1-name exists)
	if r.FormValue("student-1-name") != "" {
		var students []dto.Student
		for i := 1; i <= 6; i++ { // Max 6 students
			name := r.FormValue(fmt.Sprintf("student-%d-name", i))
			regNo := r.FormValue(fmt.Sprintf("student-%d-regNo", i))
			if name != "" && regNo != "" {
				students = append(students, dto.Student{Name: name, RegNo: regNo})
			}
		}
		data.Students = students
	} else {
		// Single student mode
		data.Students = []dto.Student{{
			Name:  r.FormValue("studentName"),
			RegNo: r.FormValue("regNo"),
		}}
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

	// Log wide event with all request info (images as metadata only)
	wideEvent := data.ToWideEvent()
	if eventJSON, err := json.Marshal(wideEvent); err == nil {
		h.logger.Printf("[WIDE_EVENT] %s", eventJSON)
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

// sendJSON sends a JSON response
func (h *Handler) sendJSON(w http.ResponseWriter, statusCode int, response dto.JSONResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logger.Printf("Error encoding JSON: %v", err)
	}
}

// sendError sends an error JSON response
func (h *Handler) sendError(w http.ResponseWriter, statusCode int, message string) {
	h.sendJSON(w, statusCode, dto.JSONResponse{
		Success: false,
		Error:   message,
	})
}

// Scrap handles the JSON API request for fetching attendance and marks(TODO)
func (h *Handler) Scrap(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Parse JSON request body
	var reqBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		h.sendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if reqBody.Username == "" || reqBody.Password == "" {
		h.sendError(w, http.StatusBadRequest, "Username and password are required")
		return
	}

	data, err := h.scraper.ScrapeCourseData(reqBody.Username, reqBody.Password)
	if err != nil {
		h.logger.Printf("Scraping error: %v", err)
		h.sendError(w, http.StatusInternalServerError, "Failed to fetch attendance and marks data")
		return
	}

	h.sendJSON(w, http.StatusOK, dto.JSONResponse{
		Success: true,
		Data:    data,
	})
}

// BuildInfo returns git commit information from GitHub API
//
//goland:noinspection GoUnusedParameter
func (h *Handler) BuildInfo(w http.ResponseWriter, r *http.Request) {
	const githubAPI = "https://api.github.com/repos/hammadmajid/zabscrap/commits/main"

	var buildInfo struct {
		Hash      string `json:"hash"`
		Message   string `json:"message"`
		TimeAgo   string `json:"timeAgo"`
		Available bool   `json:"available"`
	}

	// Fetch latest commit from GitHub
	resp, err := http.Get(githubAPI)
	if err != nil {
		h.logger.Printf("Failed to fetch GitHub commit info: %v", err)
		buildInfo.Available = false
		h.sendJSON(w, http.StatusOK, dto.JSONResponse{
			Success: true,
			Data:    buildInfo,
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		h.logger.Printf("GitHub API returned status: %d", resp.StatusCode)
		buildInfo.Available = false
		h.sendJSON(w, http.StatusOK, dto.JSONResponse{
			Success: true,
			Data:    buildInfo,
		})
		return
	}

	var githubResponse struct {
		SHA    string `json:"sha"`
		Commit struct {
			Message string `json:"message"`
			Author  struct {
				Date string `json:"date"`
			} `json:"author"`
		} `json:"commit"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&githubResponse); err != nil {
		h.logger.Printf("Failed to decode GitHub response: %v", err)
		buildInfo.Available = false
		h.sendJSON(w, http.StatusOK, dto.JSONResponse{
			Success: true,
			Data:    buildInfo,
		})
		return
	}

	// Extract short hash (first 7 characters)
	hash := githubResponse.SHA
	if len(hash) > 7 {
		hash = hash[:7]
	}

	// Parse commit time
	commitTime, err := time.Parse(time.RFC3339, githubResponse.Commit.Author.Date)
	if err != nil {
		h.logger.Printf("Failed to parse commit time: %v", err)
		buildInfo.Available = false
		h.sendJSON(w, http.StatusOK, dto.JSONResponse{
			Success: true,
			Data:    buildInfo,
		})
		return
	}

	// Calculate time ago
	duration := time.Since(commitTime)
	var timeAgo string

	if duration.Hours() < 1 {
		minutes := int(duration.Minutes())
		if minutes == 1 {
			timeAgo = "1 minute ago"
		} else {
			timeAgo = fmt.Sprintf("%d minutes ago", minutes)
		}
	} else if duration.Hours() < 24 {
		hours := int(duration.Hours())
		if hours == 1 {
			timeAgo = "1 hour ago"
		} else {
			timeAgo = fmt.Sprintf("%d hours ago", hours)
		}
	} else {
		days := int(duration.Hours() / 24)
		if days == 1 {
			timeAgo = "1 day ago"
		} else {
			timeAgo = fmt.Sprintf("%d days ago", days)
		}
	}

	buildInfo.Hash = hash
	buildInfo.Message = githubResponse.Commit.Message
	buildInfo.TimeAgo = timeAgo
	buildInfo.Available = true

	h.sendJSON(w, http.StatusOK, dto.JSONResponse{
		Success: true,
		Data:    buildInfo,
	})
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
