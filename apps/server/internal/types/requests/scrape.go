package requests

// Scrape represents the request body for the scrape endpoint
type Scrape struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Url      string `json:"url"`
}
