package shared

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// HTTPError represents a structured HTTP error response.
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// WriteJSONError writes a JSON error response to the HTTP response writer.
// This utility is shared across HTTP handlers but not exposed externally.
func WriteJSONError(w http.ResponseWriter, statusCode int, message string) {
	WriteJSONErrorWithDetails(w, statusCode, message, "")
}

// WriteJSONErrorWithDetails writes a JSON error response with additional details.
func WriteJSONErrorWithDetails(w http.ResponseWriter, statusCode int, message, details string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := HTTPError{
		Code:    statusCode,
		Message: message,
		Details: details,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		// Fallback to plain text if JSON encoding fails
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Error: %s", message)
	}
}

// WriteJSONResponse writes a JSON response to the HTTP response writer.
func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

// ParseJSONBody parses the JSON request body into the provided destination.
func ParseJSONBody(r *http.Request, dst interface{}) error {
	if r.Body == nil {
		return fmt.Errorf("request body is empty")
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Strict parsing

	if err := decoder.Decode(dst); err != nil {
		return WrapError(err, "failed to parse JSON body")
	}

	return nil
}

// SetCORSHeaders sets common CORS headers for API responses.
func SetCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Max-Age", "86400") // 24 hours
}

// LoggingMiddleware creates a middleware that logs HTTP requests.
// Returns a function that can wrap HTTP handlers.
func LoggingMiddleware(logger func(string, ...interface{})) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Wrap the response writer to capture the status code
			wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

			next.ServeHTTP(wrapped, r)

			duration := time.Since(start)
			logger("HTTP %s %s - %d - %v", r.Method, r.URL.Path, wrapped.statusCode, duration)
		})
	}
}

// responseWriter wraps http.ResponseWriter to capture the status code.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
