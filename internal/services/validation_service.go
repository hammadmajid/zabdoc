package services

import (
	"zabdoc/internal/types/requests"
)

// ValidationService handles validation of form data.
type ValidationService struct{}

func NewValidationService() *ValidationService {
	return &ValidationService{}
}

func (v *ValidationService) ValidateGenerateRequest(data *requests.Generate) error {
	// TODO: Add validation logic here (e.g., required fields)
	return nil
}
