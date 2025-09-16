package services

import (
	"bytes"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
)

// MarkdownService handles markdown parsing and sanitization.
type MarkdownService struct {
	policy *bluemonday.Policy
}

func NewMarkdownService() *MarkdownService {
	return &MarkdownService{
		policy: bluemonday.UGCPolicy(),
	}
}

// ParseAndSanitize converts markdown to sanitized HTML.
func (m *MarkdownService) ParseAndSanitize(markdown string) (string, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(markdown), &buf); err != nil {
		return "", err
	}
	return m.policy.Sanitize(buf.String()), nil
}
