package services

import (
	"zabdoc/internal/api/dto"
	"zabdoc/internal/pdf"
)

// PDFService handles PDF generation.
type PDFService struct{}

// NewPDFService creates a new PDF service instance
func NewPDFService() *PDFService {
	return &PDFService{}
}

// GeneratePDF generates a PDF document from the provided data
func (ps *PDFService) GeneratePDF(data dto.GenerateRequest) ([]byte, error) {
	service := pdf.NewService()
	return service.GeneratePDF(data)
}

// Shutdown is a no-op for compatibility (no longer needed)
func (ps *PDFService) Shutdown() {
	// No cleanup needed with gopdf
}
