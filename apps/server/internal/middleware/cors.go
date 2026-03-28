package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/cors"
)

func getAllowedOrigin() []string {
	// TODO: parse comma-separated list
	// origin := os.Getenv("ALLOWED_ORIGIN")
	// if origin == "" {
	// 	panic("env ALLOWED_ORIGIN is not defined")
	// }

	return []string{"https://zabdoc.xyz", "http://localhost:8080", "http://localhost:5173"}
}

var CORSOptions = cors.Options{
	// Allow requests from configured origin only
	AllowedOrigins: getAllowedOrigin(),

	// Only allow methods actually used by frontend (/generate and /scrap endpoints)
	AllowedMethods: []string{
		http.MethodPost,    // Used for API requests
		http.MethodOptions, // Required for CORS preflight
	},

	// Only allow Content-Type header (needed for JSON requests)
	AllowedHeaders: []string{"Content-Type"},

	// Allow credentials (cookies, authorization headers, etc.)
	AllowCredentials: true,

	// No custom headers exposed to the client
	ExposedHeaders: []string{},

	// Cache preflight requests for 12 hours
	MaxAge: int(12 * time.Hour / time.Second),
}
