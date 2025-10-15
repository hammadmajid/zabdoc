package pdf

import "github.com/signintech/gopdf"

// Page dimensions (A4 in points: 595x842)
const (
	PageWidth  = 595.0
	PageHeight = 842.0
)

// Margins and spacing
const (
	PageMargin       = 25.0 // mm converted to points
	BorderThickness  = 1.0
	HeaderLogoWidth  = 60.0
	HeaderLogoHeight = 60.0

	// Vertical spacing
	SpacerSmall = 16.0
	SpacerLarge = 44.0
)

// Font sizes
const (
	FontSizeHeader1      = 14.0
	FontSizeHeader2      = 12.0
	FontSizeMarks        = 15.0
	FontSizeCourseTitle  = 30.0
	FontSizeAssignNumber = 25.0
	FontSizeProjectTitle = 22.0
	FontSizeSubmission   = 22.0
	FontSizeInfoTable    = 20.0
	FontSizeFooter       = 14.0
	FontSizeContent      = 12.0
	FontSizeContentH1    = 18.0
	FontSizeContentH2    = 16.0
	FontSizeContentH3    = 14.0
	FontSizeCode         = 10.0
)

// Colors (RGB)
type Color struct {
	R, G, B uint8
}

var (
	ColorBlack      = Color{0, 0, 0}
	ColorDarkGray   = Color{34, 34, 34}
	ColorGray       = Color{51, 51, 51}
	ColorLightGray  = Color{204, 204, 204}
	ColorBorder     = Color{68, 68, 68}
	ColorHeaderBox  = Color{247, 247, 247}
	ColorCodeBg     = Color{245, 245, 245}
	ColorCodeBorder = Color{221, 221, 221}
)

// Font configuration
const (
	FontFamily = "TimesNewRoman"
	FontPath   = "assets/times.ttf"
	LogoPath   = "assets/szabist-logo.png"
)

// Info table dimensions
const (
	InfoTableWidth     = 70.0 // Percentage of page width
	InfoLabelWidth     = 60.0 // mm
	InfoTableMarginBot = 60.0
)

// Cover page layout dimensions
var (
	// A4 page size for gopdf
	PageSizeA4 = &gopdf.Rect{W: PageWidth, H: PageHeight}

	// Content area (inside border)
	ContentX      = PageMargin + BorderThickness
	ContentY      = PageMargin + BorderThickness
	ContentWidth  = PageWidth - (2 * (PageMargin + BorderThickness))
	ContentHeight = PageHeight - (2 * (PageMargin + BorderThickness))
)
