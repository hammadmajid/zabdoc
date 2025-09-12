package dto

type AssignmentRequest struct {
	RegNo       string
	StudentName string
	Class       string
	Course      string
	CourseCode  string
	Instructor  string
	Number      string
	Date        string
}

type TaskImage struct {
	Data     string // Base64 encoded image data
	MimeType string // MIME type (e.g., "image/jpeg", "image/png")
}

type Task struct {
	Index   int
	Content string      // HTML content (converted from markdown)
	Images  []TaskImage // Base64 encoded images with MIME types
}

type LabTaskRequest struct {
	StudentName string
	RegNo       string
	Class       string
	Course      string
	CourseCode  string
	Instructor  string
	Number      string
	Date        string
	Tasks       []Task
}
