package pdf

import (
	"fmt"
	"zabdoc/internal/api/dto"

	"github.com/signintech/gopdf"
)

// Service handles PDF generation
type Service struct {
	pdf *gopdf.GoPdf
}

// NewService creates a new PDF service instance
func NewService() *Service {
	return &Service{
		pdf: &gopdf.GoPdf{},
	}
}

// GeneratePDF generates a PDF document from the provided data
func (s *Service) GeneratePDF(data dto.GenerateRequest) ([]byte, error) {
	// Initialize PDF with A4 page size
	s.pdf.Start(gopdf.Config{
		PageSize: *PageSizeA4,
		Unit:     gopdf.UnitPT,
	})

	// Load font
	if err := s.pdf.AddTTFFont(FontFamily, FontPath); err != nil {
		return nil, fmt.Errorf("failed to load font: %w", err)
	}

	// Generate cover page
	if err := s.generateCoverPage(data); err != nil {
		return nil, fmt.Errorf("failed to generate cover page: %w", err)
	}

	// Generate content pages
	if len(data.Sections) > 0 {
		if err := s.generateContentPages(data.Sections); err != nil {
			return nil, fmt.Errorf("failed to generate content pages: %w", err)
		}
	}

	// Get PDF bytes
	pdfBytes, err := s.pdf.GetBytesPdfReturnErr()
	if err != nil {
		return nil, fmt.Errorf("failed to get PDF bytes: %w", err)
	}

	return pdfBytes, nil
}

// generateCoverPage creates the cover page
func (s *Service) generateCoverPage(data dto.GenerateRequest) error {
	s.pdf.AddPage()

	// Draw border
	s.pdf.SetLineWidth(BorderThickness)
	s.pdf.SetStrokeColor(ColorBlack.R, ColorBlack.G, ColorBlack.B)
	s.pdf.RectFromUpperLeft(PageMargin, PageMargin, PageWidth-(2*PageMargin), PageHeight-(2*PageMargin))

	currentY := ContentY + 5.0

	// Draw header
	currentY = s.drawHeader(currentY)
	currentY += SpacerSmall

	// Draw marks section
	currentY = s.drawMarks(currentY, data.Marks)
	currentY += SpacerLarge * 2

	// Draw course title section
	currentY = s.drawCourseTitle(currentY, data)
	currentY += SpacerLarge * 2

	// Draw student info
	currentY = s.drawStudentInfo(currentY, data)

	// Draw footer
	s.drawFooter(data)

	return nil
}

// generateContentPages creates content pages from sections
func (s *Service) generateContentPages(sections []dto.Section) error {
	for _, section := range sections {
		s.pdf.AddPage()

		if err := s.drawSectionContent(section); err != nil {
			return err
		}
	}
	return nil
}
