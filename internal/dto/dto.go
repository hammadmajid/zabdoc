package dto

type Image struct {
	Data          string // Base64 encoded image data
	MimeType      string // MIME type (e.g., "image/jpeg", "image/png")
	OriginalBytes int    // Original file size in bytes
}

// ImageMeta contains metadata about an image for logging (no actual data)
type ImageMeta struct {
	MimeType      string `json:"mimeType"`
	OriginalBytes int    `json:"originalBytes"`
	Base64Length  int    `json:"base64Length"`
}

// WideEvent represents a structured log entry for a request
type WideEvent struct {
	Students   []Student   `json:"students"`
	Class      string      `json:"class"`
	Course     string      `json:"course"`
	CourseCode string      `json:"courseCode"`
	Instructor string      `json:"instructor"`
	DocType    string      `json:"docType"`
	Number     string      `json:"number"`
	Date       string      `json:"date"`
	Marks      string      `json:"marks"`
	ImageCount int         `json:"imageCount"`
	Images     []ImageMeta `json:"images"`
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

// ToWideEvent converts the request to a loggable WideEvent (excludes image data)
func (r *GenerateRequest) ToWideEvent() WideEvent {
	imageMetas := make([]ImageMeta, len(r.Images))
	for i, img := range r.Images {
		imageMetas[i] = ImageMeta{
			MimeType:      img.MimeType,
			OriginalBytes: img.OriginalBytes,
			Base64Length:  len(img.Data),
		}
	}

	return WideEvent{
		Students:   r.Students,
		Class:      r.Class,
		Course:     r.Course,
		CourseCode: r.CourseCode,
		Instructor: r.Instructor,
		DocType:    r.DocType,
		Number:     r.Number,
		Date:       r.Date,
		Marks:      r.Marks,
		ImageCount: len(r.Images),
		Images:     imageMetas,
	}
}

// JSONResponse represents the API response structure
type JSONResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// AttendanceRecord represents a single attendance entry for a lecture
type AttendanceRecord struct {
	Lecture string `json:"lecture"`
	Date    string `json:"date"`
	Status  string `json:"status"`
}

// CourseAttendance represents attendance data for a course
type CourseAttendance struct {
	CourseName string             `json:"courseName"`
	Instructor string             `json:"instructor"`
	Records    []AttendanceRecord `json:"records"`
}

// MarkEntry represents a single mark row for a course
type MarkEntry struct {
	Head     string `json:"head"`
	Max      string `json:"max"`
	Obtained string `json:"obtained"`
}

// CourseMarks represents marks data for a course
type CourseMarks struct {
	Entries []MarkEntry `json:"entries"`
}

// CourseScrapeData combines attendance and marks for a course.
// "Attendence" keeps the current API spelling expected by clients.
type CourseScrapeData struct {
	Attendence CourseAttendance `json:"attendence"`
	Marks      CourseMarks      `json:"marks"`
}
