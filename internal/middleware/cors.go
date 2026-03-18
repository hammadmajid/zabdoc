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
	// Allow requests from localhost:5173 and zabdoc.xyz (including subdomains)
	AllowedOrigins: []string{getAllowedOrigin()},

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
