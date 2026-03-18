package helpers

import (
	"zabdoc/internal/types/events"
	"zabdoc/internal/types/requests"
)

// IsMultiMode returns true if there are multiple students
func IsMultiMode(r *requests.Generate) bool {
	return len(r.Students) > 1
}

// FirstStudent returns the first student (for single mode templates)
func FirstStudent(r *requests.Generate) requests.Student {
	if len(r.Students) > 0 {
		return r.Students[0]
	}
	return requests.Student{}
}

// ToGenerateRequestWideEvent converts the request to a loggable GenerateRequestWideEvent (excludes image data)
func ToGenerateRequestWideEvent(r *requests.Generate) events.GenerateRequestWideEvent {
	imageMetas := make([]events.ImageMeta, len(r.Images))
	for i, img := range r.Images {
		imageMetas[i] = events.ImageMeta{
			MimeType:      img.MimeType,
			OriginalBytes: img.OriginalBytes,
			Base64Length:  len(img.Data),
		}
	}

	return events.GenerateRequestWideEvent{
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
