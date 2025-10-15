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

	// Set up header and footer for all pages
	s.setupHeaderFooter(data)

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

// setupHeaderFooter configures header and footer to appear on all pages
func (s *Service) setupHeaderFooter(data dto.GenerateRequest) {
	// Add header function - same design as cover page
	s.pdf.AddHeader(func() {
		startY := 15.0

		// Load and draw logo
		logoX := ContentX + 20.0
		logoY := startY

		if err := s.pdf.Image(LogoPath, logoX, logoY, &gopdf.Rect{
			W: HeaderLogoWidth,
			H: HeaderLogoHeight,
		}); err != nil {
			// If logo fails to load, continue without it
		}

		// Header text starts after logo
		headerTextX := logoX + HeaderLogoWidth + 20.0
		headerY := startY + 10.0

		// Line 1: Institution name
		s.pdf.SetFont(FontFamily, "", FontSizeHeader1)
		s.pdf.SetTextColor(ColorBlack.R, ColorBlack.G, ColorBlack.B)
		s.pdf.SetXY(headerTextX, headerY)
		s.pdf.Cell(nil, "Shaheed Zulfiqar Ali Bhutto Institute of Science and Technology")

		// Line 2: Department in box
		headerY += 20.0
		boxPadding := 4.0
		boxText := "COMPUTER SCIENCE DEPARTMENT"

		// Calculate box dimensions
		textWidth, _ := s.pdf.MeasureTextWidth(boxText)
		boxWidth := textWidth + (boxPadding * 2) + 10.0
		boxHeight := FontSizeHeader2 + (boxPadding * 2)
		boxX := headerTextX
		boxY := headerY

		// Draw box background
		s.pdf.SetFillColor(ColorHeaderBox.R, ColorHeaderBox.G, ColorHeaderBox.B)
		s.pdf.SetStrokeColor(ColorDarkGray.R, ColorDarkGray.G, ColorDarkGray.B)
		s.pdf.SetLineWidth(2.0)
		s.pdf.RectFromUpperLeftWithStyle(boxX, boxY, boxWidth, boxHeight, "FD")

		// Draw text in box
		s.pdf.SetFont(FontFamily, "", FontSizeHeader2)
		s.pdf.SetTextColor(ColorBlack.R, ColorBlack.G, ColorBlack.B)
		s.pdf.SetXY(boxX+boxPadding+5.0, boxY+boxPadding+2.0)
		s.pdf.Cell(nil, boxText)
	})

	// Add footer function
	s.pdf.AddFooter(func() {
		footerY := PageHeight - 25.0

		s.pdf.SetFont(FontFamily, "", FontSizeFooter)
		s.pdf.SetTextColor(ColorBlack.R, ColorBlack.G, ColorBlack.B)

		// Left: Course code
		s.pdf.SetXY(30.0, footerY)
		s.pdf.Cell(nil, data.CourseCode)

		// Center: Class
		centerX := PageWidth / 2.0
		classWidth, _ := s.pdf.MeasureTextWidth(data.Class)
		s.pdf.SetXY(centerX-(classWidth/2.0), footerY)
		s.pdf.Cell(nil, data.Class)

		// Right: Location
		location := "SZABIST-ISB"
		locWidth, _ := s.pdf.MeasureTextWidth(location)
		s.pdf.SetXY(PageWidth-30.0-locWidth, footerY)
		s.pdf.Cell(nil, location)
	})
}

// generateCoverPage creates the cover page
func (s *Service) generateCoverPage(data dto.GenerateRequest) error {
	s.pdf.AddPage()

	// Draw border
	s.pdf.SetLineWidth(BorderThickness)
	s.pdf.SetStrokeColor(ColorBlack.R, ColorBlack.G, ColorBlack.B)
	s.pdf.RectFromUpperLeft(PageMargin, PageMargin, PageWidth-(2*PageMargin), PageHeight-(2*PageMargin))

	// Start after the automatic header (header height is ~68 points)
	currentY := 80.0

	// Draw marks section
	currentY = s.drawMarks(currentY, data.Marks)
	currentY += SpacerLarge * 2

	// Draw course title section
	currentY = s.drawCourseTitle(currentY, data)
	currentY += SpacerLarge * 2

	// Draw student info
	currentY = s.drawStudentInfo(currentY, data)

	// Footer will be added automatically by gopdf

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
