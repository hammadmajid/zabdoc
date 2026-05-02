package requests

type Student struct {
	Name  string
	RegNo string
}

type Document struct {
	Students   []Student // Always used - single student or multiple
	Class      string
	Course     string
	CourseCode string
	Instructor string
	DocType    string
	Number     string
	Date       string
	Marks      string
}

// IsMultiMode returns true if more than one student is provided in the document request.
func (g Document) IsMultiMode() bool {
	return len(g.Students) > 1
}

// FirstStudent returns the first student, or zero Student if none.
func (g Document) FirstStudent() Student {
	if len(g.Students) > 0 {
		return g.Students[0]
	}
	return Student{}
}
