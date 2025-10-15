package pdf

import (
	"strings"

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
