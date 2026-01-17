package dto

type Image struct {
	Data     string // Base64 encoded image data
	MimeType string // MIME type (e.g., "image/jpeg", "image/png")
}

type Student struct {
	Name  string
	RegNo string
}

type GenerateRequest struct {
	Students   []Student // Always used - single student or multiple
	Class      string
	Course     string
	CourseCode string
	Instructor string
	DocType    string
	Number     string
	Date       string
	Marks      string
	Images     []Image
}

// IsMultiMode returns true if there are multiple students
func (r *GenerateRequest) IsMultiMode() bool {
	return len(r.Students) > 1
}

// FirstStudent returns the first student (for single mode templates)
func (r *GenerateRequest) FirstStudent() Student {
	if len(r.Students) > 0 {
		return r.Students[0]
	}
	return Student{}
}
