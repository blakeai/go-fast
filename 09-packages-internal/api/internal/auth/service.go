package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

// Service provides authentication functionality.
// This is internal to the api package and cannot be imported by external packages.
type Service struct {
	secretKey []byte
	tokenTTL  time.Duration
	tokens    map[string]tokenInfo // In-memory token storage for demo
}

// tokenInfo holds information about a generated token.
type tokenInfo struct {
	UserID    int
	CreatedAt time.Time
	ExpiresAt time.Time
}

// NewService creates a new authentication service.
func NewService() *Service {
	return &Service{
		secretKey: []byte("demo-secret-key"),
		tokenTTL:  time.Hour,
		tokens:    make(map[string]tokenInfo),
	}
}

// NewServiceWithTTL creates a new authentication service with custom token TTL.
func NewServiceWithTTL(ttl time.Duration) *Service {
	service := NewService()
	service.tokenTTL = ttl
	return service
}

// Authenticate validates user credentials and returns a user ID.
// In a real implementation, this would check against a database.
func (s *Service) Authenticate(username, password string) (int, error) {
	// Demo authentication logic
	validUsers := map[string]struct {
		userID   int
		password string
	}{
		"alice": {userID: 1, password: "password123"},
		"bob":   {userID: 2, password: "secret456"},
		"admin": {userID: 100, password: "admin789"},
	}

	user, exists := validUsers[username]
	if !exists {
		return 0, fmt.Errorf("user %q not found", username)
	}

	if user.password != password {
		return 0, fmt.Errorf("invalid password for user %q", username)
	}

	return user.userID, nil
}

// GenerateToken creates a new authentication token for the given user ID.
func (s *Service) GenerateToken(userID int) (string, error) {
	// Generate a random token
	tokenBytes := make([]byte, 16)
	if _, err := rand.Read(tokenBytes); err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	token := hex.EncodeToString(tokenBytes)

	// Store token information
	s.tokens[token] = tokenInfo{
		UserID:    userID,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(s.tokenTTL),
	}

	return token, nil
}

// ValidateToken validates a token and returns the associated user ID.
func (s *Service) ValidateToken(token string) (int, error) {
	info, exists := s.tokens[token]
	if !exists {
		return 0, fmt.Errorf("invalid token")
	}

	if time.Now().After(info.ExpiresAt) {
		// Clean up expired token
		delete(s.tokens, token)
		return 0, fmt.Errorf("token expired")
	}

	return info.UserID, nil
}

// RevokeToken revokes (deletes) a token.
func (s *Service) RevokeToken(token string) error {
	if _, exists := s.tokens[token]; !exists {
		return fmt.Errorf("token not found")
	}

	delete(s.tokens, token)
	return nil
}

// CleanupExpiredTokens removes all expired tokens from memory.
func (s *Service) CleanupExpiredTokens() int {
	now := time.Now()
	cleaned := 0

	for token, info := range s.tokens {
		if now.After(info.ExpiresAt) {
			delete(s.tokens, token)
			cleaned++
		}
	}

	return cleaned
}

// GetTokenCount returns the number of active tokens.
func (s *Service) GetTokenCount() int {
	return len(s.tokens)
}

// isValidSecret checks if the service has a valid secret key.
// This is an unexported method for internal use and testing.
func (s *Service) isValidSecret() bool {
	return len(s.secretKey) > 0
}

// String returns a string representation of the service (without sensitive data).
func (s *Service) String() string {
	return fmt.Sprintf("AuthService{TokenTTL: %v, ActiveTokens: %d}", s.tokenTTL, len(s.tokens))
}
