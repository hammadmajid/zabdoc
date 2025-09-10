package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/cors"
)

var CORSOptions = cors.Options{
	// Allow requests from localhost:5173 and zabdoc.me (including subdomains)
	AllowedOrigins: []string{
		"http://localhost:5173",
		"https://localhost:5173",
		"https://zabdoc.me",
		"https://*.zabdoc.me",
		"http://zabdoc.me",
		"http://*.zabdoc.me",
	},

	// Allow all HTTP methods
	AllowedMethods: []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodOptions,
		http.MethodHead,
	},

	// Allow all headers
	AllowedHeaders: []string{"*"},

	// Allow credentials (cookies, authorization headers, etc.)
	AllowCredentials: true,

	// Expose all headers to the client
	ExposedHeaders: []string{"*"},

	// Cache preflight requests for 12 hours
	MaxAge: int(12 * time.Hour / time.Second),
}
