package helpers

import (
	"zabdoc/internal/types/events"
	"zabdoc/internal/types/requests"
)

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
