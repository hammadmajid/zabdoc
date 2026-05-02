package helpers

import (
	"zabdoc/internal/types/events"
	"zabdoc/internal/types/requests"
)

// ToDocumentRequestWideEvent converts the request to a loggable DocumentRequestWideEvent (excludes image data)
func ToDocumentRequestWideEvent(r *requests.Document) events.DocumentRequestWideEvent {

	return events.DocumentRequestWideEvent{
		Students:   r.Students,
		Class:      r.Class,
		Course:     r.Course,
		CourseCode: r.CourseCode,
		Instructor: r.Instructor,
		DocType:    r.DocType,
		Number:     r.Number,
		Date:       r.Date,
		Marks:      r.Marks,
	}
}
