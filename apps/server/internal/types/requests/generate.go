package requests

type Image struct {
	Data          string // Base64 encoded image data
	MimeType      string // MIME type (e.g., "image/jpeg", "image/png")
	OriginalBytes int    // Original file size in bytes
}

type Student struct {
	Name  string
	RegNo string
}

type Generate struct {
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

// IsMultiMode returns true if more than one student is provided in the generate request.
func (g Generate) IsMultiMode() bool {
	return len(g.Students) > 1
}

// FirstStudent returns the first student, or zero Student if none.
func (g Generate) FirstStudent() Student {
	if len(g.Students) > 0 {
		return g.Students[0]
	}
	return Student{}
}
