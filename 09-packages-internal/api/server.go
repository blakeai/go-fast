package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"go-fast/09-packages-internal/api/internal/auth"
	"go-fast/09-packages-internal/api/internal/validation"
	"go-fast/09-packages-internal/internal/shared"
)

// Server represents the API server with internal dependencies.
type Server struct {
	authenticator *auth.Service
	validator     *validation.Service
	logger        func(string, ...interface{})
}

// NewServer creates a new API server instance.
// This demonstrates how internal packages are used within the parent package.
func NewServer() *Server {
	return &Server{
		authenticator: auth.NewService(),
		validator:     validation.NewService(),
		logger:        log.Printf,
	}
}

// LoginRequest represents the login request payload.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents the login response payload.
type LoginResponse struct {
	Token  string `json:"token"`
	UserID int    `json:"user_id"`
}

// HandleLogin handles user authentication requests.
// This demonstrates how the public API uses internal services.
func (s *Server) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		shared.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req LoginRequest
	if err := shared.ParseJSONBody(r, &req); err != nil {
		s.logger("Login parse error: %v", err)
		shared.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Input validation using internal validation service
	if err := s.validator.ValidateCredentials(req.Username, req.Password); err != nil {
		s.logger("Login validation error for user %s: %v", req.Username, err)
		shared.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Authentication using internal auth service
	userID, err := s.authenticator.Authenticate(req.Username, req.Password)
	if err != nil {
		s.logger("Authentication failed for user %s: %v", req.Username, err)
		shared.WriteJSONError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Token generation using internal auth service
	token, err := s.authenticator.GenerateToken(userID)
	if err != nil {
		s.logger("Token generation failed for user ID %d: %v", userID, err)
		shared.WriteJSONError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	// Success response
	response := LoginResponse{
		Token:  token,
		UserID: userID,
	}

	if err := shared.WriteJSONResponse(w, http.StatusOK, response); err != nil {
		s.logger("Failed to write login response: %v", err)
	}

	s.logger("User %s (ID: %d) logged in successfully", req.Username, userID)
}

// HandleValidateToken handles token validation requests.
func (s *Server) HandleValidateToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		shared.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Extract token from Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		shared.WriteJSONError(w, http.StatusBadRequest, "Authorization header required")
		return
	}

	// Simple token extraction (in production, use proper Bearer token parsing)
	token := authHeader
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		token = authHeader[7:]
	}

	// Validate token using internal auth service
	userID, err := s.authenticator.ValidateToken(token)
	if err != nil {
		s.logger("Token validation failed: %v", err)
		shared.WriteJSONError(w, http.StatusUnauthorized, "Invalid or expired token")
		return
	}

	// Success response
	response := map[string]interface{}{
		"valid":   true,
		"user_id": userID,
	}

	if err := shared.WriteJSONResponse(w, http.StatusOK, response); err != nil {
		s.logger("Failed to write validation response: %v", err)
	}
}

// HandleStatus provides server status information.
func (s *Server) HandleStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		shared.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Get internal service status
	tokenCount := s.authenticator.GetTokenCount()

	status := map[string]interface{}{
		"status":        "healthy",
		"timestamp":     time.Now().Format(time.RFC3339),
		"active_tokens": tokenCount,
		"auth_service":  s.authenticator.String(),
	}

	if err := shared.WriteJSONResponse(w, http.StatusOK, status); err != nil {
		s.logger("Failed to write status response: %v", err)
	}
}

// SetupRoutes configures the HTTP routes for the server.
func (s *Server) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Apply logging middleware to all routes
	loggingMiddleware := shared.LoggingMiddleware(s.logger)

	mux.Handle("/login", loggingMiddleware(http.HandlerFunc(s.HandleLogin)))
	mux.Handle("/validate", loggingMiddleware(http.HandlerFunc(s.HandleValidateToken)))
	mux.Handle("/status", loggingMiddleware(http.HandlerFunc(s.HandleStatus)))

	// Add CORS handling
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		shared.SetCORSHeaders(w)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		shared.WriteJSONError(w, http.StatusNotFound, "Endpoint not found")
	}))

	return mux
}

// Start starts the HTTP server on the specified port.
func (s *Server) Start(port int) error {
	mux := s.SetupRoutes()

	addr := fmt.Sprintf(":%d", port)
	s.logger("Starting API server on %s", addr)

	//nolint:gosec // Demo code - in production, use server with timeouts
	return http.ListenAndServe(addr, mux)
}

// Cleanup performs any necessary cleanup operations.
func (s *Server) Cleanup() {
	s.logger("Cleaning up server resources...")

	// Clean up expired tokens
	cleaned := s.authenticator.CleanupExpiredTokens()
	if cleaned > 0 {
		s.logger("Cleaned up %d expired tokens", cleaned)
	}
}
