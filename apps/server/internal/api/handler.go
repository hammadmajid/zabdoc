package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"zabdoc/internal/services"
	"zabdoc/internal/types/helpers"
	"zabdoc/internal/types/requests"
	"zabdoc/internal/types/responses"
)

type Handler struct {
	logger            *log.Logger
	validationService *services.ValidationService
	scraper           *services.ScraperService
}

func NewHandler(logger *log.Logger) *Handler {
	return &Handler{
		logger:            logger,
		validationService: services.NewValidationService(),
		scraper:           services.NewScraperService(),
	}
}

// Document handles document generation requests
func (h *Handler) Document(w http.ResponseWriter, r *http.Request) {
	var data requests.Document
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		h.logger.Printf("decode JSON body: %v", err)
		return
	}

	// Log wide event with all request info
	wideEvent := helpers.ToDocumentRequestWideEvent(&data)
	if eventJSON, err := json.Marshal(wideEvent); err == nil {
		h.logger.Printf("[WIDE_EVENT] %s", eventJSON)
	}

	// Validate the data
	if err := h.validationService.ValidateDocumentRequest(&data); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		h.logger.Printf("validation error: %v", err)
		return
	}

	docx := []byte("") // TODO: read from file system

	docxMimeType := "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	w.Header().Set("Content-Type", docxMimeType)
	w.Header().Set("Content-Disposition", "attachment; filename=document.docx") // TODO: filename should be UUID
	if _, writeErr := w.Write(docx); writeErr != nil {
		h.logger.Printf("write docx: %v", writeErr)
	}
}

// Scrape handles the JSON API request for fetching attendance and marks(TODO)
func (h *Handler) Scrape(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Parse JSON request body
	var reqBody requests.Scrape

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		h.sendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if reqBody.Username == "" || reqBody.Password == "" {
		h.logScrapeEvent(&reqBody, false, "username and password are required")
		h.logScrapeResultEvent(&reqBody, false, nil, "username and password are required")
		h.sendError(w, http.StatusBadRequest, "Username and password are required")
		return
	}

	data, err := h.scraper.ScrapeCourseData(&reqBody)
	if err != nil {
		h.logScrapeEvent(&reqBody, false, err.Error())
		h.logScrapeResultEvent(&reqBody, false, nil, err.Error())
		h.logger.Printf("Scraping error: %v", err)
		h.sendError(w, http.StatusInternalServerError, "Failed to fetch attendance and marks data")
		return
	}

	h.logScrapeEvent(&reqBody, true, "")
	h.logScrapeResultEvent(&reqBody, true, data, "")
	h.sendJSON(w, http.StatusOK, responses.JSONResponse{
		Success: true,
		Data:    data,
	})
}

//goland:noinspection GoUnusedParameter
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	if h.validationService == nil {
		http.Error(w, "ValidationService unavailable", http.StatusServiceUnavailable)
		h.logger.Println("Health check failed: ValidationService unavailable")
		return
	}

	if h.scraper == nil {
		http.Error(w, "Scraper unavailable", http.StatusServiceUnavailable)
		h.logger.Println("Health check failed: Scraper unavailable")
		return
	}

	// All dependencies healthy
	_, _ = fmt.Fprintf(w, "Status is available")
}
