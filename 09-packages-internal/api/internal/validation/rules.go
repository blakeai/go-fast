package validation

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// Service provides input validation functionality.
// This is internal to the api package and cannot be imported by external packages.
type Service struct {
	emailRegex *regexp.Regexp
}

// NewService creates a new validation service.
func NewService() *Service {
	// Compile email validation regex once
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return &Service{
		emailRegex: emailRegex,
	}
}

// ValidateCredentials validates username and password for authentication.
func (s *Service) ValidateCredentials(username, password string) error {
	if err := s.ValidateUsername(username); err != nil {
		return fmt.Errorf("username validation failed: %w", err)
	}

	if err := s.ValidatePassword(password); err != nil {
		return fmt.Errorf("password validation failed: %w", err)
	}

	return nil
}

// ValidateUsername validates a username according to business rules.
func (s *Service) ValidateUsername(username string) error {
	username = strings.TrimSpace(username)

	if username == "" {
		return fmt.Errorf("username cannot be empty")
	}

	if len(username) < 3 {
		return fmt.Errorf("username must be at least 3 characters long")
	}

	if len(username) > 50 {
		return fmt.Errorf("username must be no more than 50 characters long")
	}

	// Check for valid characters (alphanumeric and underscore only)
	for _, char := range username {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) && char != '_' {
			return fmt.Errorf("username can only contain letters, numbers, and underscores")
		}
	}

	// Username must start with a letter
	if !unicode.IsLetter(rune(username[0])) {
		return fmt.Errorf("username must start with a letter")
	}

	return nil
}

// ValidatePassword validates a password according to security requirements.
func (s *Service) ValidatePassword(password string) error {
	if password == "" {
		return fmt.Errorf("password cannot be empty")
	}

	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	if len(password) > 128 {
		return fmt.Errorf("password must be no more than 128 characters long")
	}

	// Check for required character types
	var hasLower, hasUpper, hasDigit, hasSpecial bool

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasLower {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}

	if !hasUpper {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}

	if !hasDigit {
		return fmt.Errorf("password must contain at least one digit")
	}

	if !hasSpecial {
		return fmt.Errorf("password must contain at least one special character")
	}

	return nil
}

// ValidateEmail validates an email address format.
func (s *Service) ValidateEmail(email string) error {
	email = strings.TrimSpace(email)

	if email == "" {
		return fmt.Errorf("email cannot be empty")
	}

	if len(email) > 254 {
		return fmt.Errorf("email must be no more than 254 characters long")
	}

	if !s.emailRegex.MatchString(email) {
		return fmt.Errorf("email format is invalid")
	}

	return nil
}

// ValidateUserInput validates a complete user input structure.
type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// ValidateUserInput validates all fields in a user input structure.
func (s *Service) ValidateUserInput(input UserInput) []error {
	var errors []error

	if err := s.ValidateUsername(input.Username); err != nil {
		errors = append(errors, err)
	}

	if err := s.ValidatePassword(input.Password); err != nil {
		errors = append(errors, err)
	}

	if err := s.ValidateEmail(input.Email); err != nil {
		errors = append(errors, err)
	}

	return errors
}

// ValidateRequired checks if a value is not empty (for string fields).
func (s *Service) ValidateRequired(fieldName, value string) error {
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("field %q is required", fieldName)
	}
	return nil
}

// ValidateLength checks if a string is within specified length bounds.
func (s *Service) ValidateLength(fieldName, value string, min, max int) error {
	length := len(value)

	if length < min {
		return fmt.Errorf("field %q must be at least %d characters long, got %d", fieldName, min, length)
	}

	if length > max {
		return fmt.Errorf("field %q must be no more than %d characters long, got %d", fieldName, max, length)
	}

	return nil
}

// SanitizeInput removes potentially dangerous characters from input.
func (s *Service) SanitizeInput(input string) string {
	// Remove leading/trailing whitespace
	sanitized := strings.TrimSpace(input)

	// Remove null bytes
	sanitized = strings.ReplaceAll(sanitized, "\x00", "")

	// Remove other control characters except newlines and tabs
	var result strings.Builder
	for _, char := range sanitized {
		if char == '\n' || char == '\t' || !unicode.IsControl(char) {
			result.WriteRune(char)
		}
	}

	return result.String()
}
