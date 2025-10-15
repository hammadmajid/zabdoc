package pdf

import (
	"zabdoc/internal/api/dto"
)

// drawMarks draws the marks section (right-aligned)
func (s *Service) drawMarks(startY float64, marks string) float64 {
	s.pdf.SetFont(FontFamily, "", FontSizeMarks)
	s.pdf.SetTextColor(ColorBlack.R, ColorBlack.G, ColorBlack.B)

	rightX := PageWidth - ContentX - 100.0
	currentY := startY

	// Total Marks line
	s.pdf.SetXY(rightX-80.0, currentY)
	s.pdf.Cell(nil, "Total Marks:")

	// Draw underline for marks value
	underlineX := rightX + 20.0
	underlineWidth := 60.0
	s.pdf.SetXY(underlineX, currentY)
	s.pdf.Cell(nil, marks)

	// Draw underline
	s.pdf.SetLineWidth(2.0)
	s.pdf.SetStrokeColor(ColorBorder.R, ColorBorder.G, ColorBorder.B)
	s.pdf.Line(underlineX, currentY+FontSizeMarks+2.0, underlineX+underlineWidth, currentY+FontSizeMarks+2.0)

	currentY += FontSizeMarks + 12.0

	// Obtained Marks line
	s.pdf.SetXY(rightX-80.0, currentY)
	s.pdf.Cell(nil, "Obtained Marks:")

	// Draw empty underline for obtained marks
	s.pdf.Line(underlineX, currentY+FontSizeMarks+2.0, underlineX+underlineWidth, currentY+FontSizeMarks+2.0)

	return currentY + FontSizeMarks + 6.0
}

// drawCourseTitle draws the course and assignment information
func (s *Service) drawCourseTitle(startY float64, data dto.GenerateRequest) float64 {
	centerX := PageWidth / 2.0
	currentY := startY

	// Course title (centered)
	s.pdf.SetFont(FontFamily, "", FontSizeCourseTitle)
	s.pdf.SetTextColor(ColorBlack.R, ColorBlack.G, ColorBlack.B)

	courseWidth, _ := s.pdf.MeasureTextWidth(data.Course)
	s.pdf.SetXY(centerX-(courseWidth/2.0), currentY)
	s.pdf.Cell(nil, data.Course)
	currentY += FontSizeCourseTitle + 6.0

	// Assignment number or project type
	s.pdf.SetFont(FontFamily, "", FontSizeAssignNumber)
	var assignText string
	if data.ProjectTitle != "" {
		assignText = data.DocType
	} else {
		assignText = data.DocType + " #" + data.Number
	}

	assignWidth, _ := s.pdf.MeasureTextWidth(assignText)
	s.pdf.SetXY(centerX-(assignWidth/2.0), currentY)
	s.pdf.Cell(nil, assignText)
	currentY += FontSizeAssignNumber + 4.0

	// Submission date
	s.pdf.SetFont(FontFamily, "", FontSizeSubmission)
	dateText := "Submission date: " + data.Date
	dateWidth, _ := s.pdf.MeasureTextWidth(dateText)
	s.pdf.SetXY(centerX-(dateWidth/2.0), currentY)
	s.pdf.Cell(nil, dateText)
	currentY += FontSizeSubmission

	// Project title if exists
	if data.ProjectTitle != "" {
		currentY += SpacerSmall
		s.pdf.SetFont(FontFamily, "", FontSizeProjectTitle)
		s.pdf.SetTextColor(ColorGray.R, ColorGray.G, ColorGray.B)

		projectWidth, _ := s.pdf.MeasureTextWidth(data.ProjectTitle)
		s.pdf.SetXY(centerX-(projectWidth/2.0), currentY)
		s.pdf.Cell(nil, data.ProjectTitle)
		currentY += FontSizeProjectTitle + 8.0
	}

	return currentY
}

// drawStudentInfo draws the student information table
func (s *Service) drawStudentInfo(startY float64, data dto.GenerateRequest) float64 {
	centerX := PageWidth / 2.0
	tableWidth := (PageWidth * InfoTableWidth) / 100.0
	tableX := centerX - (tableWidth / 2.0)
	currentY := startY

	s.pdf.SetFont(FontFamily, "", FontSizeInfoTable)
	s.pdf.SetTextColor(ColorBlack.R, ColorBlack.G, ColorBlack.B)

	if data.IsMultiMode {
		// Shared info table
		currentY = s.drawInfoRow(tableX, currentY, tableWidth, "Submitted to:", data.Instructor)
		currentY = s.drawInfoRow(tableX, currentY, tableWidth, "Class/Section:", data.Class)

		currentY += 20.0

		// Student table header
		s.pdf.SetFont(FontFamily, "B", FontSizeInfoTable)
		headerY := currentY

		// Draw header row
		s.pdf.SetXY(tableX, headerY)
		s.pdf.Cell(nil, "Student Name")

		s.pdf.SetXY(tableX+(tableWidth/2.0), headerY)
		s.pdf.Cell(nil, "Reg. Number")

		// Underline header
		s.pdf.SetLineWidth(2.0)
		s.pdf.SetStrokeColor(ColorGray.R, ColorGray.G, ColorGray.B)
		s.pdf.Line(tableX, headerY+FontSizeInfoTable+4.0, tableX+tableWidth, headerY+FontSizeInfoTable+4.0)

		currentY = headerY + FontSizeInfoTable + 12.0

		// Draw student rows
		s.pdf.SetFont(FontFamily, "", FontSizeInfoTable)
		for _, student := range data.Students {
			s.pdf.SetXY(tableX+(tableWidth/4.0)-30.0, currentY)
			s.pdf.Cell(nil, student.Name)

			s.pdf.SetXY(tableX+(tableWidth*3.0/4.0)-30.0, currentY)
			s.pdf.Cell(nil, student.RegNo)

			currentY += FontSizeInfoTable + 10.0
		}
	} else {
		// Single student mode
		currentY = s.drawInfoRow(tableX, currentY, tableWidth, "Student Name:", data.StudentName)
		currentY = s.drawInfoRow(tableX, currentY, tableWidth, "Reg. Number:", data.RegNo)
		currentY = s.drawInfoRow(tableX, currentY, tableWidth, "Submitted to:", data.Instructor)
		currentY = s.drawInfoRow(tableX, currentY, tableWidth, "Class/Section:", data.Class)
	}

	return currentY
}

// drawInfoRow draws a single row in the info table
func (s *Service) drawInfoRow(x, y, width float64, label, value string) float64 {
	labelWidth := InfoLabelWidth

	// Draw label (right-aligned)
	s.pdf.SetXY(x, y)
	s.pdf.Cell(nil, label)

	// Draw value (centered with underline)
	valueX := x + labelWidth + 10.0
	valueWidth := width - labelWidth - 10.0

	valueTextWidth, _ := s.pdf.MeasureTextWidth(value)
	valueCenterX := valueX + (valueWidth / 2.0) - (valueTextWidth / 2.0)

	s.pdf.SetXY(valueCenterX, y)
	s.pdf.Cell(nil, value)

	// Draw underline
	s.pdf.SetLineWidth(1.0)
	s.pdf.SetStrokeColor(ColorLightGray.R, ColorLightGray.G, ColorLightGray.B)
	s.pdf.Line(valueX, y+FontSizeInfoTable+2.0, valueX+valueWidth, y+FontSizeInfoTable+2.0)

	return y + FontSizeInfoTable + 10.0
}
