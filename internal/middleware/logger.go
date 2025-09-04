package middleware

import (
	"net/http"
	"time"

	"github.com/hammadmajid/zabcover/internal/app"
)

// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

// LoggingMiddleware logs method, path, status, and duration for each request
func LoggingMiddleware(logger *app.App) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rw := &responseWriter{ResponseWriter: w, status: 200}
			next.ServeHTTP(rw, r)
			duration := time.Since(start).Milliseconds()
			logger.Logger.Printf("[%s] %s %d %dms", r.Method, r.URL.Path, rw.status, duration)
		})
	}
}
