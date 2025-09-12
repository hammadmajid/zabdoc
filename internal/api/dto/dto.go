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

type Task struct {
	Index   int
	Content string // HTML content (converted from markdown)
	Images  []string
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
