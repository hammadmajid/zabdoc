package api

import (
	"encoding/json"
	"net/http"
	"zabdoc/internal/types/responses"
)

// sendJSON sends a JSON response
func (h *Handler) sendJSON(w http.ResponseWriter, statusCode int, response responses.JSONResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logger.Printf("Error encoding JSON: %v", err)
	}
}

// sendError sends an error JSON response
func (h *Handler) sendError(w http.ResponseWriter, statusCode int, message string) {
	h.sendJSON(w, statusCode, responses.JSONResponse{
		Success: false,
		Error:   message,
	})
}
