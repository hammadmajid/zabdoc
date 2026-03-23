package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/cors"
)

func getAllowedOrigin() string {
	origin := os.Getenv("ALLOWED_ORIGIN")
	if origin == "" {
		panic("env ALLOWED_ORIGIN is not defined")
	}

	return origin
}

var CORSOptions = cors.Options{
	// Allow requests from configured origin only
	AllowedOrigins: []string{getAllowedOrigin()},

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
