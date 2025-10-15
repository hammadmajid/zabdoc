package pdf

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"
	"zabdoc/internal/api/dto"

	"github.com/signintech/gopdf"
)

// drawSectionContent renders a section's content and images
func (s *Service) drawSectionContent(section dto.Section) error {
	currentY := 30.0
	marginX := 30.0
	contentWidth := PageWidth - (2 * marginX)

	// Parse and render HTML content
	if section.Content != "" {
		var err error
		currentY, err = s.renderHTML(section.Content, marginX, currentY, contentWidth)
		if err != nil {
			return err
		}
	}

	// Render images if any
	if len(section.Images) > 0 {
		currentY += 20.0
		for i, img := range section.Images {
			currentY, _ = s.drawImage(img, marginX, currentY, contentWidth)
			if i < len(section.Images)-1 {
				currentY += 10.0
			}
		}
	}

	return nil
}

// renderHTML parses HTML content and renders it with appropriate styling
func (s *Service) renderHTML(html string, x, y, width float64) (float64, error) {
	currentY := y

	// Remove HTML tags and extract text with basic formatting
	lines := parseHTMLToLines(html)

	s.pdf.SetFont(FontFamily, "", FontSizeContent)
	s.pdf.SetTextColor(ColorBlack.R, ColorBlack.G, ColorBlack.B)

	for _, line := range lines {
		if line.IsH1 {
			currentY += 15.0
			s.pdf.SetFont(FontFamily, "B", FontSizeContentH1)
			s.pdf.SetTextColor(ColorGray.R, ColorGray.G, ColorGray.B)

			s.pdf.SetXY(x, currentY)
			s.pdf.Cell(nil, line.Text)

			// Underline H1
			s.pdf.SetLineWidth(1.0)
			s.pdf.SetStrokeColor(ColorLightGray.R, ColorLightGray.G, ColorLightGray.B)
			textWidth, _ := s.pdf.MeasureTextWidth(line.Text)
			s.pdf.Line(x, currentY+FontSizeContentH1+3.0, x+textWidth, currentY+FontSizeContentH1+3.0)

			currentY += FontSizeContentH1 + 10.0
			s.pdf.SetFont(FontFamily, "", FontSizeContent)
			s.pdf.SetTextColor(ColorBlack.R, ColorBlack.G, ColorBlack.B)
		} else if line.IsH2 {
			currentY += 12.0
			s.pdf.SetFont(FontFamily, "B", FontSizeContentH2)
			s.pdf.SetTextColor(ColorGray.R, ColorGray.G, ColorGray.B)

			s.pdf.SetXY(x, currentY)
			s.pdf.Cell(nil, line.Text)
			currentY += FontSizeContentH2 + 8.0

			s.pdf.SetFont(FontFamily, "", FontSizeContent)
			s.pdf.SetTextColor(ColorBlack.R, ColorBlack.G, ColorBlack.B)
		} else if line.IsH3 {
			currentY += 10.0
			s.pdf.SetFont(FontFamily, "B", FontSizeContentH3)
			s.pdf.SetTextColor(ColorGray.R, ColorGray.G, ColorGray.B)

			s.pdf.SetXY(x, currentY)
			s.pdf.Cell(nil, line.Text)
			currentY += FontSizeContentH3 + 6.0

			s.pdf.SetFont(FontFamily, "", FontSizeContent)
			s.pdf.SetTextColor(ColorBlack.R, ColorBlack.G, ColorBlack.B)
		} else if line.IsCode {
			currentY += 5.0
			// Draw code block background
			codeHeight := FontSizeCode + 16.0
			s.pdf.SetFillColor(ColorCodeBg.R, ColorCodeBg.G, ColorCodeBg.B)
			s.pdf.SetStrokeColor(ColorCodeBorder.R, ColorCodeBorder.G, ColorCodeBorder.B)
			s.pdf.SetLineWidth(1.0)
			s.pdf.RectFromUpperLeftWithStyle(x, currentY, width, codeHeight, "FD")

			// Draw code text
			s.pdf.SetFont("Courier", "", FontSizeCode)
			s.pdf.SetTextColor(ColorBlack.R, ColorBlack.G, ColorBlack.B)
			s.pdf.SetXY(x+8.0, currentY+8.0)
			s.pdf.Cell(nil, line.Text)

			currentY += codeHeight + 5.0
			s.pdf.SetFont(FontFamily, "", FontSizeContent)
		} else if line.IsBlockquote {
			// Draw blockquote with left border
			s.pdf.SetLineWidth(4.0)
			s.pdf.SetStrokeColor(ColorLightGray.R, ColorLightGray.G, ColorLightGray.B)
			s.pdf.Line(x, currentY, x, currentY+FontSizeContent+4.0)

			s.pdf.SetFont(FontFamily, "I", FontSizeContent)
			s.pdf.SetTextColor(102, 102, 102)
			s.pdf.SetXY(x+10.0, currentY)

			// Word wrap for blockquote
			wrappedLines := wrapText(s.pdf, line.Text, width-10.0)
			for _, wLine := range wrappedLines {
				s.pdf.SetXY(x+10.0, currentY)
				s.pdf.Cell(nil, wLine)
				currentY += FontSizeContent + 4.0
			}

			s.pdf.SetFont(FontFamily, "", FontSizeContent)
			s.pdf.SetTextColor(ColorBlack.R, ColorBlack.G, ColorBlack.B)
		} else if line.IsList {
			s.pdf.SetXY(x+15.0, currentY)
			bulletText := "• " + line.Text

			// Word wrap for list items
			wrappedLines := wrapText(s.pdf, bulletText, width-15.0)
			for i, wLine := range wrappedLines {
				if i > 0 {
					s.pdf.SetXY(x+25.0, currentY)
				} else {
					s.pdf.SetXY(x+15.0, currentY)
				}
				s.pdf.Cell(nil, wLine)
				currentY += FontSizeContent + 4.0
			}
		} else if line.Text != "" {
			// Regular paragraph
			wrappedLines := wrapText(s.pdf, line.Text, width)
			for _, wLine := range wrappedLines {
				s.pdf.SetXY(x, currentY)
				s.pdf.Cell(nil, wLine)
				currentY += FontSizeContent + 4.0
			}
			currentY += 4.0
		} else {
			// Empty line
			currentY += FontSizeContent + 2.0
		}
	}

	return currentY, nil
}

// drawImage renders an image from base64 data
func (s *Service) drawImage(img dto.SectionImage, x, y, maxWidth float64) (float64, error) {
	// Decode base64 image
	imageData, err := base64.StdEncoding.DecodeString(img.Data)
	if err != nil {
		return y, fmt.Errorf("failed to decode image: %w", err)
	}

	// Create image holder from bytes
	imgHolder, err := gopdf.ImageHolderByBytes(imageData)
	if err != nil {
		return y, fmt.Errorf("failed to create image holder: %w", err)
	}

	// Calculate image dimensions (max width, maintain aspect ratio)
	imgWidth := maxWidth - 20.0
	imgHeight := 300.0 // Default height, gopdf will maintain aspect ratio

	// Draw image
	if err := s.pdf.ImageByHolderWithOptions(imgHolder, gopdf.ImageOptions{
		X:    x + 10.0,
		Y:    y,
		Rect: &gopdf.Rect{W: imgWidth, H: imgHeight},
	}); err != nil {
		return y, fmt.Errorf("failed to draw image: %w", err)
	}

	return y + imgHeight + 20.0, nil
}

// Line represents a parsed line of HTML content
type Line struct {
	Text         string
	IsH1         bool
	IsH2         bool
	IsH3         bool
	IsCode       bool
	IsBlockquote bool
	IsList       bool
}

// parseHTMLToLines converts HTML to structured lines
func parseHTMLToLines(html string) []Line {
	var lines []Line

	// Remove extra whitespace
	html = strings.TrimSpace(html)

	// Split by common block elements
	re := regexp.MustCompile(`</(p|h1|h2|h3|pre|blockquote|li)>`)
	parts := re.Split(html, -1)

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		line := Line{}

		// Check for heading tags
		if strings.Contains(part, "<h1") {
			line.IsH1 = true
			line.Text = cleanHTML(part)
		} else if strings.Contains(part, "<h2") {
			line.IsH2 = true
			line.Text = cleanHTML(part)
		} else if strings.Contains(part, "<h3") {
			line.IsH3 = true
			line.Text = cleanHTML(part)
		} else if strings.Contains(part, "<pre") || strings.Contains(part, "<code") {
			line.IsCode = true
			line.Text = cleanHTML(part)
		} else if strings.Contains(part, "<blockquote") {
			line.IsBlockquote = true
			line.Text = cleanHTML(part)
		} else if strings.Contains(part, "<li") {
			line.IsList = true
			line.Text = cleanHTML(part)
		} else if strings.Contains(part, "<p") {
			line.Text = cleanHTML(part)
		} else {
			line.Text = cleanHTML(part)
		}

		if line.Text != "" {
			lines = append(lines, line)
		}
	}

	return lines
}

// cleanHTML removes HTML tags and returns clean text
func cleanHTML(html string) string {
	// Remove HTML tags
	re := regexp.MustCompile(`<[^>]*>`)
	text := re.ReplaceAllString(html, "")

	// Decode HTML entities
	text = strings.ReplaceAll(text, "&lt;", "<")
	text = strings.ReplaceAll(text, "&gt;", ">")
	text = strings.ReplaceAll(text, "&amp;", "&")
	text = strings.ReplaceAll(text, "&quot;", "\"")
	text = strings.ReplaceAll(text, "&#39;", "'")
	text = strings.ReplaceAll(text, "&nbsp;", " ")

	return strings.TrimSpace(text)
}
