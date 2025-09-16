package services

import "zabdoc/internal/api/dto"

// ValidationService handles validation of form data.
type ValidationService struct{}

func NewValidationService() *ValidationService {
	return &ValidationService{}
}

func (v *ValidationService) ValidateGenerateRequest(data *dto.GenerateRequest) error {
	// TODO: Add validation logic here (e.g., required fields)
	return nil
}
