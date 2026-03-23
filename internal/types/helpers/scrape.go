package helpers

import (
	"zabdoc/internal/types/events"
	"zabdoc/internal/types/requests"
)

// ToScrapeRequestWideEvent converts Scrape request to a loggable wide event (excludes password)
func ToScrapeRequestWideEvent(r *requests.Scrape, success bool, errMsg string) events.ScrapeRequestWideEvent {
	return events.ScrapeRequestWideEvent{
		Username: r.Username,
		Success:  success,
		Error:    errMsg,
	}
}

// ToScrapeResultWideEvent converts final scrape result to a loggable wide event
func ToScrapeResultWideEvent(r *requests.Scrape, success bool, data interface{}, errMsg string) events.ScrapeResultWideEvent {
	return events.ScrapeResultWideEvent{
		Username: r.Username,
		Success:  success,
		Data:     data,
		Error:    errMsg,
	}
}
