package pdf

import (
	"strings"
	"zabdoc/internal/api/dto"

	"github.com/signintech/gopdf"
)

// wrapText splits text into multiple lines based on available width
func wrapText(pdf *gopdf.GoPdf, text string, maxWidth float64) []string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return []string{}
	}

	var lines []string
	var currentLine strings.Builder

	for i, word := range words {
		testLine := currentLine.String()
		if testLine != "" {
			testLine += " " + word
		} else {
			testLine = word
		}

		width, err := pdf.MeasureTextWidth(testLine)
		if err != nil {
			// If measurement fails, add word anyway
			if currentLine.Len() > 0 {
				currentLine.WriteString(" ")
			}
			currentLine.WriteString(word)
			continue
		}

		if width <= maxWidth {
			// Word fits on current line
			if currentLine.Len() > 0 {
				currentLine.WriteString(" ")
			}
			currentLine.WriteString(word)
		} else {
			// Word doesn't fit, start new line
			if currentLine.Len() > 0 {
				lines = append(lines, currentLine.String())
				currentLine.Reset()
			}
			currentLine.WriteString(word)
		}

		// Add last line
		if i == len(words)-1 && currentLine.Len() > 0 {
			lines = append(lines, currentLine.String())
		}
	}

	if len(lines) == 0 && currentLine.Len() > 0 {
		lines = append(lines, currentLine.String())
	}

	return lines
}

// centerText calculates X position to center text within given width
func centerText(pdf *gopdf.GoPdf, text string, pageWidth float64) float64 {
	textWidth, err := pdf.MeasureTextWidth(text)
	if err != nil {
		return pageWidth / 2.0
	}
	return (pageWidth - textWidth) / 2.0
}

// rightAlignText calculates X position to right-align text
func rightAlignText(pdf *gopdf.GoPdf, text string, pageWidth, rightMargin float64) float64 {
	textWidth, err := pdf.MeasureTextWidth(text)
	if err != nil {
		return pageWidth - rightMargin
	}
	return pageWidth - rightMargin - textWidth
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
