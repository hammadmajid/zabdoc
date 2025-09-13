package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/cors"
)

var CORSOptions = cors.Options{
	// Allow requests from localhost:5173 and zabdoc.xyz (including subdomains)
	AllowedOrigins: []string{
		"http://zabdoc.localhost:5173",
		"https://zabdoc.localhost:5173",
		"https://zabdoc.xyz",
		"https://*.zabdoc.xyz",
		"http://zabdoc.xyz",
		"http://*.zabdoc.xyz",
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
