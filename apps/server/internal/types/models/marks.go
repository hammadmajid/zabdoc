package models

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
