package models

// CourseScrapeData combines attendance and marks for a course.
// "Attendence" keeps the current API spelling expected by clients.
type CourseScrapeData struct {
	Attendence CourseAttendance `json:"attendence"`
	Marks      CourseMarks      `json:"marks"`
}
