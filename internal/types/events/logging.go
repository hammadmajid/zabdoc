package events

import "zabdoc/internal/types/requests"

// ImageMeta contains metadata about an image for logging (no actual data)
type ImageMeta struct {
	MimeType      string `json:"mimeType"`
	OriginalBytes int    `json:"originalBytes"`
	Base64Length  int    `json:"base64Length"`
}

// GenerateRequestWideEvent represents a structured log entry for a request
type GenerateRequestWideEvent struct {
	Students   []requests.Student `json:"students"`
	Class      string             `json:"class"`
	Course     string             `json:"course"`
	CourseCode string             `json:"courseCode"`
	Instructor string             `json:"instructor"`
	DocType    string             `json:"docType"`
	Number     string             `json:"number"`
	Date       string             `json:"date"`
	Marks      string             `json:"marks"`
	ImageCount int                `json:"imageCount"`
	Images     []ImageMeta        `json:"images"`
}
