package models

// CourseScrapeData combines attendance and marks for a course.
// "Attendance" keeps the current API spelling expected by clients.
type CourseScrapeData struct {
	Attendance CourseAttendance `json:"attendance"`
	Marks      CourseMarks      `json:"marks"`
}
