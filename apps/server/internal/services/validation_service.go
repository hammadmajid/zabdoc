package services

import (
	"fmt"
	"strings"
	"zabdoc/internal/types/requests"
)

// ValidationService handles validation of form data.
type ValidationService struct{}

func NewValidationService() *ValidationService {
	return &ValidationService{}
}

func (v *ValidationService) ValidateDocumentRequest(data *requests.Document) error {
	// Validate students only if provided
	for i, s := range data.Students {
		if s.Name == "" {
			return fmt.Errorf("student %d: name is required", i+1)
		}
		if s.RegNo == "" {
			return fmt.Errorf("student %d: registration number is required", i+1)
		}
	}

	if strings.TrimSpace(data.Class) == "" {
		return fmt.Errorf("class is required")
	}
	if strings.TrimSpace(data.Course) == "" {
		return fmt.Errorf("course is required")
	}
	if strings.TrimSpace(data.CourseCode) == "" {
		return fmt.Errorf("course code is required")
	}
	if strings.TrimSpace(data.Instructor) == "" {
		return fmt.Errorf("instructor is required")
	}

	validDocTypes := map[string]bool{
		"assignment": true,
		"task": true, // Lab task
		"project": true, // Lab project
	}
	if !validDocTypes[strings.ToLower(data.DocType)] {
		return fmt.Errorf("invalid document type: %q", data.DocType)
	}

	if strings.TrimSpace(data.Number) == "" {
		return fmt.Errorf("number is required")
	}
	if strings.TrimSpace(data.Date) == "" {
		return fmt.Errorf("date is required")
	}

	return nil
}

func (ValidationService) ValidateScrapeRequest(data *requests.Scrape) error {
	if strings.TrimSpace(data.Username) == "" {
		return fmt.Errorf("username is required")
	}
	if strings.TrimSpace(data.Password) == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}
