package shared

import (
	"fmt"
	"runtime"
	"strings"
)

// WrapError wraps an error with additional context.
// This utility is shared across the module but not exposed externally.
func WrapError(err error, context string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", context, err)
}

// FormatValidationError creates a standardized validation error message.
func FormatValidationError(field, message string) error {
	return fmt.Errorf("validation failed for field %q: %s", field, message)
}

// ChainErrors combines multiple errors into a single error message.
func ChainErrors(errors []error) error {
	if len(errors) == 0 {
		return nil
	}

	if len(errors) == 1 {
		return errors[0]
	}

	var messages []string
	for _, err := range errors {
		if err != nil {
			messages = append(messages, err.Error())
		}
	}

	return fmt.Errorf("multiple errors: %s", strings.Join(messages, "; "))
}

// ErrorWithStack creates an error with stack trace information.
// Useful for debugging internal errors.
func ErrorWithStack(message string) error {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return fmt.Errorf("%s", message)
	}

	// Extract just the filename from the full path
	parts := strings.Split(file, "/")
	if len(parts) > 0 {
		file = parts[len(parts)-1]
	}

	return fmt.Errorf("%s (at %s:%d)", message, file, line)
}

// RecoverError converts a panic into an error.
// Useful for internal error handling in goroutines.
func RecoverError() error {
	if r := recover(); r != nil {
		switch v := r.(type) {
		case error:
			return WrapError(v, "recovered from panic")
		case string:
			return fmt.Errorf("recovered from panic: %s", v)
		default:
			return fmt.Errorf("recovered from panic: %v", v)
		}
	}
	return nil
}
