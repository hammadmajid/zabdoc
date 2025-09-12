package dto

type AssignmentRequest struct {
	StudentName string `json:"studentName"`
	RegNo       string `json:"regNo"`
	Class       string `json:"class"`
	Course      string `json:"course"`
	CourseCode  string `json:"courseCode"`
	Instructor  string `json:"instructor"`
	Number      string `json:"number"`
	Date        string `json:"date"`
}

type LabTaskRequest struct {
	StudentName string `json:"studentName"`
	RegNo       string `json:"regNo"`
	Class       string `json:"class"`
	Course      string `json:"course"`
	CourseCode  string `json:"courseCode"`
	Instructor  string `json:"instructor"`
	Number      string `json:"number"`
	Date        string `json:"date"`
}
