package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidInput = errors.New("invalid input")
	ErrUnauthorized = errors.New("unauthorized access")
)

type User struct {
	ID     int
	Name   string
	Status string
}

type Item struct {
	ID   int
	Name string
}

func basicErrorHandlingDemo() {
	fmt.Println("=== Basic Error Handling Patterns ===")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	result2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Printf("Error: %v\n", err2)
	} else {
		fmt.Printf("Result: %.2f\n", result2)
	}

	user, err := processUser("42")
	if err != nil {
		fmt.Printf("Process user error: %v\n", err)
	} else {
		fmt.Printf("Processed user: %s\n", user.Name)
	}

	user2, err := processUser("invalid")
	if err != nil {
		fmt.Printf("Process user error: %v\n", err)
	} else {
		fmt.Printf("Processed user: %s\n", user2.Name)
	}
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func processUser(id string) (*User, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID %q: %w", id, err)
	}

	user, err := getUser(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user %d: %w", userID, err)
	}

	if user.Status == "inactive" {
		return nil, fmt.Errorf("user %d is inactive", userID)
	}

	return user, nil
}

func getUser(id int) (*User, error) {
	if id <= 0 {
		return nil, ErrInvalidInput
	}

	mockUsers := map[int]*User{
		42: {ID: 42, Name: "Alice", Status: "active"},
		99: {ID: 99, Name: "Bob", Status: "inactive"},
	}

	user, exists := mockUsers[id]
	if !exists {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func errorCreationDemo() {
	fmt.Println("\n=== Error Creation Methods ===")

	err1 := errors.New("something went wrong")
	fmt.Printf("Simple error: %v\n", err1)

	itemID := 123
	err2 := fmt.Errorf("failed to process item %d", itemID)
	fmt.Printf("Formatted error: %v\n", err2)

	originalErr := errors.New("network timeout")
	wrappedErr := fmt.Errorf("failed to fetch data: %w", originalErr)
	fmt.Printf("Wrapped error: %v\n", wrappedErr)

	if errors.Is(wrappedErr, originalErr) {
		fmt.Println("✓ Wrapped error contains the original network error")
	}

	fmt.Printf("Unwrapping: %v\n", errors.Unwrap(wrappedErr))
}

func errorWrappingDemo() {
	fmt.Println("\n=== Error Wrapping and Unwrapping ===")

	err := processFile("nonexistent.txt")
	if err != nil {
		fmt.Printf("File processing error: %v\n", err)
		fmt.Println("\nError chain:")
		diagnoseError(err)
	}
}

func processFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %q: %w", filename, err)
	}
	defer file.Close()

	data, err := parseFile(file)
	if err != nil {
		return fmt.Errorf("failed to parse file %q: %w", filename, err)
	}

	if err := validateData(data); err != nil {
		return fmt.Errorf("invalid data in file %q: %w", filename, err)
	}

	return nil
}

func parseFile(file *os.File) ([]byte, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file contents: %w", err)
	}
	return data, nil
}

func validateData(data []byte) error {
	if len(data) == 0 {
		return errors.New("file is empty")
	}
	return nil
}

func diagnoseError(err error) {
	for err != nil {
		fmt.Printf("  - %v\n", err)
		err = errors.Unwrap(err)
	}
}

func sentinelErrorsDemo() {
	fmt.Println("\n=== Sentinel Errors ===")

	handleUserLookup(42)
	handleUserLookup(99)
	handleUserLookup(-1)
	handleUserLookup(999)
}

func handleUserLookup(id int) {
	user, err := getUser(id)
	if err != nil {
		switch {
		case errors.Is(err, ErrUserNotFound):
			fmt.Printf("User %d does not exist\n", id)
		case errors.Is(err, ErrInvalidInput):
			fmt.Printf("Invalid user ID: %d\n", id)
		case errors.Is(err, ErrUnauthorized):
			fmt.Printf("Unauthorized access for user %d\n", id)
		default:
			fmt.Printf("Unexpected error for user %d: %v\n", id, err)
		}
		return
	}

	fmt.Printf("Found user %d: %s (status: %s)\n", user.ID, user.Name, user.Status)
}

func errorHandlingPatternsDemo() {
	fmt.Println("\n=== Error Handling Patterns ===")

	fmt.Println("1. Early return pattern:")
	items := []Item{{ID: 1, Name: "item1"}, {ID: 2, Name: "item2"}}
	if err := processItems(items); err != nil {
		fmt.Printf("Failed to process items: %v\n", err)
	} else {
		fmt.Println("All items processed successfully")
	}

	fmt.Println("\n2. Accumulate errors pattern:")
	data := map[string]interface{}{
		"age":   -5,
		"name":  "",
		"email": "invalid-email",
	}
	if errs := validateAllFields(data); len(errs) > 0 {
		fmt.Printf("Validation errors:\n")
		for _, err := range errs {
			fmt.Printf("  - %v\n", err)
		}
	}

	fmt.Println("\n3. Best effort cleanup:")
	bestEffortCleanup()

	fmt.Println("\n4. Retry pattern:")
	if err := retryableOperation(); err != nil {
		fmt.Printf("Operation failed: %v\n", err)
	}
}

func processItems(items []Item) error {
	for _, item := range items {
		if err := processItem(item); err != nil {
			return fmt.Errorf("failed to process item %d: %w", item.ID, err)
		}
	}
	return nil
}

func processItem(item Item) error {
	if item.Name == "" {
		return fmt.Errorf("item %d has empty name", item.ID)
	}
	fmt.Printf("  Processed item %d: %s\n", item.ID, item.Name)
	return nil
}

func validateAllFields(data map[string]interface{}) []error {
	var errs []error

	for field, value := range data {
		if err := validateField(field, value); err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

func validateField(field string, value interface{}) error {
	switch field {
	case "age":
		if age, ok := value.(int); ok {
			if age < 0 {
				return fmt.Errorf("age must be non-negative, got %d", age)
			}
			if age > 150 {
				return fmt.Errorf("age must be realistic, got %d", age)
			}
		}
	case "name":
		if name, ok := value.(string); ok {
			if strings.TrimSpace(name) == "" {
				return fmt.Errorf("name cannot be empty")
			}
		}
	case "email":
		if email, ok := value.(string); ok {
			if !strings.Contains(email, "@") {
				return fmt.Errorf("invalid email format: %s", email)
			}
		}
	}
	return nil
}

func bestEffortCleanup() {
	fmt.Println("  Performing cleanup operations...")

	if err := deleteTemporaryFiles(); err != nil && !errors.Is(err, os.ErrNotExist) {
		fmt.Printf("  Warning: failed to delete temp files: %v\n", err)
	} else {
		fmt.Println("  ✓ Temporary files cleaned up")
	}

	if err := closeConnections(); err != nil {
		fmt.Printf("  Warning: failed to close connections: %v\n", err)
	} else {
		fmt.Println("  ✓ Connections closed")
	}
}

func deleteTemporaryFiles() error {
	return nil
}

func closeConnections() error {
	return nil
}

var operationAttempts int

func retryableOperation() error {
	const maxRetries = 3
	operationAttempts = 0

	for attempt := 1; attempt <= maxRetries; attempt++ {
		err := doOperation()
		if err == nil {
			fmt.Printf("  ✓ Operation succeeded on attempt %d\n", attempt)
			return nil
		}

		if !isRetryableError(err) {
			return fmt.Errorf("non-retryable error: %w", err)
		}

		fmt.Printf("  Attempt %d failed: %v\n", attempt, err)

		if attempt == maxRetries {
			return fmt.Errorf("operation failed after %d attempts: %w", maxRetries, err)
		}

		fmt.Printf("  Retrying in %d second(s)...\n", attempt)
		time.Sleep(time.Duration(attempt) * 100 * time.Millisecond) // Faster for demo
	}

	return nil
}

func doOperation() error {
	operationAttempts++
	if operationAttempts < 3 {
		return errors.New("temporary network error")
	}
	return nil
}

func isRetryableError(err error) bool {
	return strings.Contains(err.Error(), "network") || strings.Contains(err.Error(), "timeout")
}

func panicRecoveryDemo() {
	fmt.Println("\n=== Panic and Recovery ===")

	fmt.Println("1. Converting panic to error:")
	if err := convertPanicToError(); err != nil {
		fmt.Printf("Caught panic as error: %v\n", err)
	} else {
		fmt.Println("No panic occurred")
	}

	fmt.Println("\n2. Safe operation with recovery:")
	safeOperation()
}

func convertPanicToError() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("operation panicked: %v", r)
		}
	}()

	mightPanic(true)
	return nil
}

func mightPanic(shouldPanic bool) {
	if shouldPanic {
		panic("something went wrong!")
	}
	fmt.Println("Operation completed successfully")
}

func safeOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	fmt.Println("Starting risky operation...")
	riskyOperation()
	fmt.Println("Operation completed safely")
}

func riskyOperation() {
	fmt.Println("This operation is safe")
}

func main() {
	basicErrorHandlingDemo()
	errorCreationDemo()
	errorWrappingDemo()
	sentinelErrorsDemo()
	errorHandlingPatternsDemo()
	panicRecoveryDemo()

	fmt.Println("\n" + strings.Repeat("=", 60) + "\n")
	runCustomErrorExamples()
}
