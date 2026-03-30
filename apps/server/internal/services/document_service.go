package services

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"zabdoc/internal/types/requests"
)

const templatePath = "templates/cover.docx"

type DocumentService struct{}

func NewDocumentService() *DocumentService {
	return &DocumentService{}
}

func (s *DocumentService) CreateDocument(doc *requests.Document) ([]byte, error) {
	templateBytes, err := os.ReadFile(templatePath)
	if err != nil {
		return nil, fmt.Errorf("read template: %w", err)
	}

	placeholders := buildPlaceholders(doc)

	zr, err := zip.NewReader(bytes.NewReader(templateBytes), int64(len(templateBytes)))
	if err != nil {
		return nil, fmt.Errorf("open template zip: %w", err)
	}

	var buf bytes.Buffer
	zw := zip.NewWriter(&buf)

	for _, f := range zr.File {
		rc, err := f.Open()
		if err != nil {
			return nil, fmt.Errorf("open zip entry %q: %w", f.Name, err)
		}

		content, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			return nil, fmt.Errorf("read zip entry %q: %w", f.Name, err)
		}

		// Only replace placeholders in XML files
		if strings.HasSuffix(f.Name, ".xml") || strings.HasSuffix(f.Name, ".rels") {
			content = replacePlaceholders(content, placeholders)
		}

		fw, err := zw.CreateHeader(&f.FileHeader)
		if err != nil {
			return nil, fmt.Errorf("create zip entry %q: %w", f.Name, err)
		}
		if _, err = fw.Write(content); err != nil {
			return nil, fmt.Errorf("write zip entry %q: %w", f.Name, err)
		}
	}

	if err = zw.Close(); err != nil {
		return nil, fmt.Errorf("close zip writer: %w", err)
	}

	return buf.Bytes(), nil
}

func buildPlaceholders(doc *requests.Document) map[string]string {
	p := map[string]string{
		"{{Class}}":      doc.Class,
		"{{Course}}":     doc.Course,
		"{{CourseCode}}": doc.CourseCode,
		"{{Instructor}}": doc.Instructor,
		"{{DocType}}":    doc.DocType,
		"{{Number}}":     doc.Number,
		"{{Date}}":       doc.Date,
		"{{Marks}}":      doc.Marks,
	}

	// Single student shorthands
	if len(doc.Students) == 1 {
		p["{{StudentName}}"] = doc.Students[0].Name
		p["{{RegNo}}"]       = doc.Students[0].RegNo
	}

	// Multi-student: {{Student1Name}}, {{Student1RegNo}}, etc.
	for i, s := range doc.Students {
		n := fmt.Sprintf("%d", i+1)
		p[fmt.Sprintf("{{Student%sName}}", n)]  = s.Name
		p[fmt.Sprintf("{{Student%sRegNo}}", n)] = s.RegNo
	}

	return p
}

func replacePlaceholders(content []byte, placeholders map[string]string) []byte {
	result := string(content)
	for placeholder, value := range placeholders {
		result = strings.ReplaceAll(result, placeholder, value)
	}
	return []byte(result)
}
