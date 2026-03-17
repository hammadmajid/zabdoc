package services

import (
	"zabdoc/internal/types"
)

// ValidationService handles validation of form data.
type ValidationService struct{}

func NewValidationService() *ValidationService {
	return &ValidationService{}
}

func (v *ValidationService) ValidateGenerateRequest(data *types.GenerateRequest) error {
	// TODO: Add validation logic here (e.g., required fields)
	return nil
}
