package dto

type SectionImage struct {
	Data     string // Base64 encoded image data
	MimeType string // MIME type (e.g., "image/jpeg", "image/png")
}

type Section struct {
	Index   int
	Content string         // HTML content (converted from markdown)
	Images  []SectionImage // Base64 encoded images with MIME types
}

type Student struct {
	Name  string
	RegNo string
}

type GenerateRequest struct {
	StudentName  string
	RegNo        string
	Students     []Student // For multi-student mode
	IsMultiMode  bool
	Class        string
	Course       string
	CourseCode   string
	Instructor   string
	DocType      string
	Number       string
	Date         string
	Marks        string
	ProjectTitle string
	Sections     []Section
}
