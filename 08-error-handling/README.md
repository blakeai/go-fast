# Chapter 8: Error Handling

## Overview

Go uses explicit error values instead of exceptions. Errors are values that implement the `error` interface, and functions return them as additional return values. This approach makes error handling visible in the code flow and prevents unhandled exceptions.

## Key Concepts

- **Errors are values** - implement the `error` interface
- **Explicit error returns** - functions return `(result, error)` tuples
- **Error wrapping** - adding context while preserving the original error
- **Custom error types** - implementing `error` interface for rich error information
- **Error checking patterns** - `if err != nil` is idiomatic

## Examples

### Basic Error Handling Pattern
See [`errors.go`](./errors.go) for implementation.

```go
import (
    "errors"
    "fmt"
    "strconv"
)

// Functions return errors as last return value
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Standard error checking pattern
func example() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Result: %f\n", result)
}

// Multiple error checks in sequence
func processUser(id string) (*User, error) {
    userID, err := strconv.Atoi(id)
    if err != nil {
        return nil, fmt.Errorf("invalid user ID %q: %w", id, err)
    }
    
    user, err := database.GetUser(userID)
    if err != nil {
        return nil, fmt.Errorf("failed to get user %d: %w", userID, err)
    }
    
    if user.Status == "inactive" {
        return nil, fmt.Errorf("user %d is inactive", userID)
    }
    
    return user, nil
}
```

### Error Creation Methods
```go
// Simple error messages
err1 := errors.New("something went wrong")

// Formatted error messages
err2 := fmt.Errorf("failed to process item %d", itemID)

// Error wrapping - preserves original error
originalErr := errors.New("network timeout")
wrappedErr := fmt.Errorf("failed to fetch data: %w", originalErr)

// Check if error wraps another error
if errors.Is(wrappedErr, originalErr) {
    fmt.Println("This error wraps the original network error")
}

// Extract wrapped errors
var netErr *net.OpError
if errors.As(wrappedErr, &netErr) {
    fmt.Printf("Network operation: %s\n", netErr.Op)
}
```

### Custom Error Types
See [`custom_errors.go`](./custom_errors.go) for implementation.

```go
// Simple custom error type
type ValidationError struct {
    Field string
    Value interface{}
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation failed for field %q (value: %v): %s", 
        e.Field, e.Value, e.Message)
}

// Rich error type with multiple fields
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

// Implement Unwrap for error wrapping support
func (e DatabaseError) Unwrap() error {
    return e.Cause
}

// Usage
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
```

### Error Checking and Type Assertions
```go
func handleError(err error) {
    if err == nil {
        return
    }
    
    // Check for specific error values
    if errors.Is(err, io.EOF) {
        fmt.Println("Reached end of file")
        return
    }
    
    // Check for specific error types
    var validationErr ValidationError
    if errors.As(err, &validationErr) {
        fmt.Printf("Validation failed: %s = %v (%s)\n", 
            validationErr.Field, validationErr.Value, validationErr.Message)
        return
    }
    
    var dbErr DatabaseError
    if errors.As(err, &dbErr) {
        fmt.Printf("Database error in %s operation on %s table\n", 
            dbErr.Operation, dbErr.Table)
        return
    }
    
    // Generic error handling
    fmt.Printf("Unknown error: %v\n", err)
}
```

### Error Wrapping and Unwrapping
```go
// Build error context as you return up the stack
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

// Check the error chain
func diagnoseError(err error) {
    // Walk the error chain
    for err != nil {
        fmt.Printf("Error: %v\n", err)
        err = errors.Unwrap(err)
    }
}

// Check if any error in the chain matches
func isTimeoutError(err error) bool {
    for err != nil {
        if netErr, ok := err.(*net.OpError); ok {
            if netErr.Timeout() {
                return true
            }
        }
        err = errors.Unwrap(err)
    }
    return false
}
```

### Sentinel Errors
```go
// Define package-level sentinel errors
var (
    ErrUserNotFound = errors.New("user not found")
    ErrInvalidInput = errors.New("invalid input")
    ErrUnauthorized = errors.New("unauthorized access")
)

func getUser(id int) (*User, error) {
    if id <= 0 {
        return nil, ErrInvalidInput
    }
    
    user := database.FindUser(id)
    if user == nil {
        return nil, ErrUserNotFound
    }
    
    return user, nil
}

// Check for sentinel errors
func handleUserLookup(id int) {
    user, err := getUser(id)
    if err != nil {
        switch {
        case errors.Is(err, ErrUserNotFound):
            fmt.Println("User does not exist")
        case errors.Is(err, ErrInvalidInput):
            fmt.Println("Invalid user ID provided")
        default:
            fmt.Printf("Unexpected error: %v\n", err)
        }
        return
    }
    
    fmt.Printf("Found user: %s\n", user.Name)
}
```

### Error Handling Patterns
```go
// Early return pattern (most common)
func processItems(items []Item) error {
    for _, item := range items {
        if err := processItem(item); err != nil {
            return fmt.Errorf("failed to process item %d: %w", item.ID, err)
        }
    }
    return nil
}

// Accumulate errors pattern
func validateAllFields(data map[string]interface{}) []error {
    var errs []error
    
    for field, value := range data {
        if err := validateField(field, value); err != nil {
            errs = append(errs, err)
        }
    }
    
    return errs
}

// Ignore specific errors pattern
func bestEffortCleanup() {
    if err := deleteTemporaryFiles(); err != nil && !errors.Is(err, os.ErrNotExist) {
        log.Printf("Warning: failed to delete temp files: %v", err)
    }
    
    if err := closeConnections(); err != nil {
        log.Printf("Warning: failed to close connections: %v", err)
    }
}

// Retry with error pattern
func retryableOperation() error {
    const maxRetries = 3
    
    for attempt := 1; attempt <= maxRetries; attempt++ {
        err := doOperation()
        if err == nil {
            return nil
        }
        
        if !isRetryableError(err) {
            return fmt.Errorf("non-retryable error: %w", err)
        }
        
        if attempt == maxRetries {
            return fmt.Errorf("operation failed after %d attempts: %w", maxRetries, err)
        }
        
        time.Sleep(time.Duration(attempt) * time.Second)
    }
    
    return nil
}
```

### Panic and Recovery (Use Sparingly)
```go
// Panic for truly exceptional conditions
func mustParseConfig(filename string) Config {
    config, err := parseConfig(filename)
    if err != nil {
        panic(fmt.Sprintf("failed to parse critical config file: %v", err))
    }
    return config
}

// Recover from panics (usually in web handlers or goroutines)
func safeHandler(w http.ResponseWriter, r *http.Request) {
    defer func() {
        if recover := recover(); recover != nil {
            log.Printf("Handler panicked: %v", recover)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        }
    }()
    
    // Handler code that might panic
    riskyOperation()
}

// Convert panic to error
func convertPanicToError() (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("operation panicked: %v", r)
        }
    }()
    
    // Code that might panic
    mightPanic()
    return nil
}
```

## Running the Code

```bash
go run *.go
go test ./...
```

## Java Developer Notes

- No exceptions - errors are explicit return values
- No try/catch blocks - use `if err != nil` pattern
- Error wrapping ≈ Java's exception chaining with `getCause()`
- `errors.Is()` ≈ Java's `instanceof` for exception types
- `errors.As()` ≈ Java's exception type casting
- Panic/recover ≈ Java exceptions, but used only for truly exceptional cases
- Multiple return values eliminate need for checked exceptions
- Error handling is visible in code flow (can't be ignored like unchecked exceptions)
- Custom error types ≈ Java's custom exception classes

## Next Steps

Continue to [Chapter 9: Packages](../09-packages/)

## References

- [Go Blog - Error handling and Go](https://go.dev/blog/error-handling-and-go)
- [Go Tour - Errors](https://go.dev/tour/methods/19)
- [Effective Go - Errors](https://go.dev/doc/effective_go#errors)