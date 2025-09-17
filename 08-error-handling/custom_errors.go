package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
)

type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation failed for field %q (value: %v): %s",
		e.Field, e.Value, e.Message)
}

type DatabaseError struct {
	Operation string
	Table     string
	Cause     error
	Code      int
}

func (e DatabaseError) Error() string {
	return fmt.Sprintf("database %s operation failed on table %q (code: %d): %v",
		e.Operation, e.Table, e.Code, e.Cause)
}

func (e DatabaseError) Unwrap() error {
	return e.Cause
}

type NetworkError struct {
	Op       string
	Addr     string
	Timeout  bool
	Cause    error
	Attempts int
}

func (e NetworkError) Error() string {
	if e.Timeout {
		return fmt.Sprintf("network %s to %s timed out after %d attempts: %v",
			e.Op, e.Addr, e.Attempts, e.Cause)
	}
	return fmt.Sprintf("network %s to %s failed after %d attempts: %v",
		e.Op, e.Addr, e.Attempts, e.Cause)
}

func (e NetworkError) Unwrap() error {
	return e.Cause
}

func (e NetworkError) IsTimeout() bool {
	return e.Timeout
}

type FileSystemError struct {
	Path      string
	Operation string
	Cause     error
}

func (e FileSystemError) Error() string {
	return fmt.Sprintf("filesystem %s operation failed on %q: %v",
		e.Operation, e.Path, e.Cause)
}

func (e FileSystemError) Unwrap() error {
	return e.Cause
}

type MultiError struct {
	Errors []error
}

func (e MultiError) Error() string {
	if len(e.Errors) == 0 {
		return "no errors"
	}
	if len(e.Errors) == 1 {
		return e.Errors[0].Error()
	}

	msgs := make([]string, 0, len(e.Errors))
	for _, err := range e.Errors {
		msgs = append(msgs, err.Error())
	}
	return fmt.Sprintf("multiple errors: %s", strings.Join(msgs, "; "))
}

func (e MultiError) Unwrap() []error {
	return e.Errors
}

func (e MultiError) Is(target error) bool {
	for _, err := range e.Errors {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}

func (e MultiError) As(target interface{}) bool {
	for _, err := range e.Errors {
		if errors.As(err, target) {
			return true
		}
	}
	return false
}

func customErrorTypesDemo() {
	fmt.Println("=== Custom Error Types ===")

	fmt.Println("1. Validation errors:")
	if err := validateAge(-5); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if err := validateAge(200); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if err := validateAge(25); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Age 25 is valid")
	}

	fmt.Println("\n2. Database errors:")
	dbErr := DatabaseError{
		Operation: "SELECT",
		Table:     "users",
		Code:      1045,
		Cause:     errors.New("access denied for user 'app'@'localhost'"),
	}
	fmt.Printf("Database error: %v\n", dbErr)

	fmt.Println("\n3. Network errors:")
	netErr := NetworkError{
		Op:       "dial",
		Addr:     "api.example.com:443",
		Timeout:  true,
		Attempts: 3,
		Cause:    errors.New("connection timeout"),
	}
	fmt.Printf("Network error: %v\n", netErr)
	fmt.Printf("Is timeout: %t\n", netErr.IsTimeout())
}

func validateAge(age int) error {
	if age < 0 {
		return ValidationError{
			Field:   "age",
			Value:   age,
			Message: "must be non-negative",
		}
	}
	if age > 150 {
		return ValidationError{
			Field:   "age",
			Value:   age,
			Message: "unrealistic age value",
		}
	}
	return nil
}

func errorCheckingDemo() {
	fmt.Println("\n=== Error Checking and Type Assertions ===")

	errors := []error{
		ValidationError{Field: "email", Value: "invalid", Message: "must contain @"},
		DatabaseError{Operation: "INSERT", Table: "logs", Code: 2003, Cause: errors.New("can't connect to server")},
		NetworkError{Op: "get", Addr: "localhost:8080", Timeout: true, Attempts: 2, Cause: errors.New("timeout")},
		io.EOF,
		errors.New("generic error"),
	}

	for i, err := range errors {
		fmt.Printf("\nError %d: %v\n", i+1, err)
		handleError(err)
	}
}

func handleError(err error) {
	if err == nil {
		return
	}

	// Check for specific error values
	if errors.Is(err, io.EOF) {
		fmt.Println("  → Reached end of file")
		return
	}

	// Check for specific error types
	var validationErr ValidationError
	if errors.As(err, &validationErr) {
		fmt.Printf("  → Validation failed: %s = %v (%s)\n",
			validationErr.Field, validationErr.Value, validationErr.Message)
		return
	}

	var dbErr DatabaseError
	if errors.As(err, &dbErr) {
		fmt.Printf("  → Database error in %s operation on %s table (code: %d)\n",
			dbErr.Operation, dbErr.Table, dbErr.Code)
		if dbErr.Cause != nil {
			fmt.Printf("    Caused by: %v\n", dbErr.Cause)
		}
		return
	}

	var netErr NetworkError
	if errors.As(err, &netErr) {
		fmt.Printf("  → Network error: %s to %s\n", netErr.Op, netErr.Addr)
		if netErr.IsTimeout() {
			fmt.Println("    This was a timeout error")
		}
		return
	}

	// Generic error handling
	fmt.Printf("  → Unknown error: %v\n", err)
}

func multiErrorDemo() {
	fmt.Println("\n=== Multi-Error Handling ===")

	// Simulate batch validation that collects all errors
	userInputs := []map[string]interface{}{
		{"name": "", "age": -1, "email": "invalid"},
		{"name": "Alice", "age": 25, "email": "alice@example.com"},
		{"name": "Bob", "age": 200, "email": "bob"},
	}

	for i, input := range userInputs {
		fmt.Printf("\nValidating user %d: %v\n", i+1, input)
		if err := validateUser(input); err != nil {
			fmt.Printf("Validation failed: %v\n", err)

			// Check if it's a multi-error
			var multiErr MultiError
			if errors.As(err, &multiErr) {
				fmt.Printf("Found %d validation errors:\n", len(multiErr.Errors))
				for j, subErr := range multiErr.Errors {
					fmt.Printf("  %d. %v\n", j+1, subErr)
				}
			}
		} else {
			fmt.Println("✓ User validation passed")
		}
	}
}

func validateUser(input map[string]interface{}) error {
	var errs []error

	// Validate name
	if name, ok := input["name"].(string); ok {
		if strings.TrimSpace(name) == "" {
			errs = append(errs, ValidationError{
				Field:   "name",
				Value:   name,
				Message: "cannot be empty",
			})
		}
	}

	// Validate age
	if age, ok := input["age"].(int); ok {
		if err := validateAge(age); err != nil {
			errs = append(errs, err)
		}
	}

	// Validate email
	if email, ok := input["email"].(string); ok {
		if !strings.Contains(email, "@") {
			errs = append(errs, ValidationError{
				Field:   "email",
				Value:   email,
				Message: "must contain @ symbol",
			})
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return MultiError{Errors: errs}
}

func complexErrorChainDemo() {
	fmt.Println("\n=== Complex Error Chain ===")

	err := simulateComplexOperation()
	if err != nil {
		fmt.Printf("Complex operation failed: %v\n", err)
		fmt.Println("\nAnalyzing error chain:")
		analyzeErrorChain(err)
	}
}

func simulateComplexOperation() error {
	// Simulate a complex operation with multiple failure points
	if err := connectToDatabase(); err != nil {
		return fmt.Errorf("failed to initialize system: %w", err)
	}
	return nil
}

func connectToDatabase() error {
	if err := establishConnection(); err != nil {
		return DatabaseError{
			Operation: "connect",
			Table:     "system",
			Code:      2003,
			Cause:     err,
		}
	}
	return nil
}

func establishConnection() error {
	return NetworkError{
		Op:       "dial",
		Addr:     "db.example.com:5432",
		Timeout:  true,
		Attempts: 3,
		Cause: &net.OpError{
			Op:   "dial",
			Net:  "tcp",
			Addr: &net.TCPAddr{IP: net.ParseIP("192.168.1.100"), Port: 5432},
			Err:  errors.New("connection refused"),
		},
	}
}

func analyzeErrorChain(err error) {
	depth := 0
	for err != nil {
		indent := strings.Repeat("  ", depth)
		fmt.Printf("%s- %T: %v\n", indent, err, err)

		// Check for specific error types and their properties
		//nolint:errorlint // Educational example showing type switch on errors
		switch e := err.(type) {
		case NetworkError:
			fmt.Printf("%s  Network op: %s, addr: %s, timeout: %t\n",
				indent, e.Op, e.Addr, e.Timeout)
		case DatabaseError:
			fmt.Printf("%s  DB op: %s, table: %s, code: %d\n",
				indent, e.Operation, e.Table, e.Code)
		case *net.OpError:
			fmt.Printf("%s  Net op: %s, network: %s, addr: %v\n",
				indent, e.Op, e.Net, e.Addr)
		}

		err = errors.Unwrap(err)
		depth++
	}
}

func contextualErrorDemo() {
	fmt.Println("\n=== Contextual Error Information ===")

	// Simulate errors with rich context
	operations := []func() error{
		func() error { return processConfigFile("/tmp/config.json") },
		func() error { return validateConfig(map[string]string{"timeout": "invalid"}) },
		func() error { return connectToService("auth-service", 3) },
	}

	for i, op := range operations {
		fmt.Printf("\nOperation %d:\n", i+1)
		if err := op(); err != nil {
			fmt.Printf("Failed: %v\n", err)
			printErrorContext(err)
		} else {
			fmt.Println("Success")
		}
	}
}

func processConfigFile(path string) error {
	return FileSystemError{
		Path:      path,
		Operation: "read",
		Cause:     errors.New("permission denied"),
	}
}

func validateConfig(config map[string]string) error {
	if timeout, exists := config["timeout"]; exists {
		return ValidationError{
			Field:   "timeout",
			Value:   timeout,
			Message: "must be a valid duration (e.g., '30s', '5m')",
		}
	}
	return nil
}

func connectToService(service string, maxRetries int) error {
	return NetworkError{
		Op:       "connect",
		Addr:     service + ".internal:8080",
		Timeout:  false,
		Attempts: maxRetries,
		Cause:    errors.New("service unavailable"),
	}
}

func printErrorContext(err error) {
	//nolint:errorlint // Educational example showing type switch on errors
	switch e := err.(type) {
	case FileSystemError:
		fmt.Printf("  File: %s, Operation: %s\n", e.Path, e.Operation)
	case ValidationError:
		fmt.Printf("  Field: %s, Value: %v\n", e.Field, e.Value)
	case NetworkError:
		fmt.Printf("  Target: %s, Attempts: %d, Timeout: %t\n",
			e.Addr, e.Attempts, e.Timeout)
	}
}

func runCustomErrorExamples() {
	customErrorTypesDemo()
	errorCheckingDemo()
	multiErrorDemo()
	complexErrorChainDemo()
	contextualErrorDemo()
}

func init() {
	fmt.Println("Starting Custom Error Handling Examples")
	fmt.Println(strings.Repeat("=", 60))
}
