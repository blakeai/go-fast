# Chapter 9.1: Internal Packages

## Overview

Go's `internal/` directory provides a mechanism for creating packages that can only be imported by specific parts of your codebase. This enables better encapsulation and prevents external dependencies on implementation details, similar to package-private visibility in Java.

## Key Concepts

- **Internal package visibility rules** - import restrictions based on directory structure
- **Strategic code organization** - when and where to use `internal/` directories
- **API boundaries** - controlling what external code can access
- **Refactoring safety** - internal packages can change without breaking external users

## Examples

### Basic Internal Package Structure
```
myapp/
├── cmd/
│   └── server/
│       └── main.go
├── api/
│   ├── handlers.go
│   ├── middleware.go
│   └── internal/
│       ├── auth/
│       │   └── jwt.go       # Only api/ can import this
│       └── validation/
│           └── rules.go     # Only api/ can import this
├── database/
│   ├── models.go
│   └── internal/
│       └── migrations/
│           └── runner.go    # Only database/ can import this
└── internal/
    ├── config/
    │   └── loader.go        # Only myapp/ root can import this
    └── shared/
        └── utils.go         # Only myapp/ root can import this
```

### Internal Package Import Rules
See [`internal/`](./internal/) directory structure for implementation.

```go
// File: api/internal/auth/jwt.go
package auth

import "time"

// This package can ONLY be imported by:
// - myapp/api/ (parent directory)
// - myapp/api/handlers/ (sibling directories under api/)
// - myapp/api/middleware/ (sibling directories under api/)

func GenerateToken(userID int, expiry time.Duration) string {
    // JWT implementation - internal to api package
    return "jwt-token"
}

func ValidateToken(token string) (int, error) {
    // Token validation - internal to api package
    return 123, nil
}
```

### Legal and Illegal Imports
```go
// File: api/handlers.go
package handlers

import (
    "myapp/api/internal/auth"      // ✓ LEGAL - same parent directory
    "myapp/api/internal/validation" // ✓ LEGAL - same parent directory
)

func LoginHandler() {
    token := auth.GenerateToken(123, time.Hour) // ✓ Can use internal auth
}

// File: database/models.go  
package database

import (
    "myapp/api/internal/auth"  // ✗ ILLEGAL - different parent directory
)

// File: cmd/server/main.go
package main

import (
    "myapp/api"                    // ✓ LEGAL - public API
    "myapp/api/internal/auth"      // ✗ ILLEGAL - internal to api package
    "myapp/internal/config"        // ✓ LEGAL - same module root
)
```

### Strategic API Design with Internal Packages
```go
// File: api/handlers.go (public API)
package api

import (
    "myapp/api/internal/auth"
    "myapp/api/internal/validation"
)

// Public API - stable interface for external users
type Server struct {
    // internal implementation hidden
    authenticator *auth.Service
    validator     *validation.Service
}

func NewServer() *Server {
    return &Server{
        authenticator: auth.NewService(),
        validator:     validation.NewService(),
    }
}

// Public methods expose clean interface
func (s *Server) HandleLogin(username, password string) (string, error) {
    // Input validation (internal)
    if err := s.validator.ValidateCredentials(username, password); err != nil {
        return "", err
    }
    
    // Authentication logic (internal)
    userID, err := s.authenticator.Authenticate(username, password)
    if err != nil {
        return "", err
    }
    
    // Token generation (internal)
    token := s.authenticator.GenerateToken(userID)
    return token, nil
}

// File: api/internal/auth/service.go (internal implementation)
package auth

// Internal implementation - can change without breaking external users
type Service struct {
    secretKey []byte
    tokenTTL  time.Duration
}

func NewService() *Service {
    return &Service{
        secretKey: []byte("secret"),
        tokenTTL:  time.Hour,
    }
}

func (s *Service) Authenticate(username, password string) (int, error) {
    // Complex authentication logic hidden from external users
    return 123, nil
}

func (s *Service) GenerateToken(userID int) string {
    // JWT generation logic hidden from external users
    return "token"
}
```

### Shared Internal Utilities
```go
// File: internal/shared/errors.go
package shared

// Shared utilities available to entire module but not external users
func WrapError(err error, context string) error {
    if err == nil {
        return nil
    }
    return fmt.Errorf("%s: %w", context, err)
}

func FormatValidationError(field, message string) error {
    return fmt.Errorf("validation failed for field %q: %s", field, message)
}

// File: internal/shared/http.go  
package shared

func WriteJSONError(w http.ResponseWriter, statusCode int, message string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(map[string]string{
        "error": message,
    })
}

func ParseJSONBody(r *http.Request, dst interface{}) error {
    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields()
    return decoder.Decode(dst)
}
```

### Module-Level Internal Organization
```go
// File: internal/config/config.go
package config

import "os"

// Module-wide configuration - internal to prevent external dependency
type Config struct {
    DatabaseURL string
    APIKey      string
    Debug       bool
}

func Load() (*Config, error) {
    return &Config{
        DatabaseURL: getEnvOrDefault("DATABASE_URL", "localhost:5432"),
        APIKey:      os.Getenv("API_KEY"),
        Debug:       os.Getenv("DEBUG") == "true",
    }, nil
}

func getEnvOrDefault(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

// File: cmd/server/main.go
package main

import "myapp/internal/config"  // ✓ Can import module-level internal

func main() {
    cfg, err := config.Load()    // ✓ Use internal config
    if err != nil {
        log.Fatal(err)
    }
    
    // Start server with config...
}
```

### Testing Internal Packages
```go
// File: api/internal/auth/jwt_test.go
package auth

import "testing"

// Internal package tests have full access to unexported functions
func TestGenerateToken(t *testing.T) {
    service := NewService()
    token := service.GenerateToken(123)
    
    if token == "" {
        t.Error("Expected non-empty token")
    }
    
    // Can test internal functions
    if !service.isValidSecret() {  // unexported method
        t.Error("Service secret should be valid")
    }
}

// File: api/internal/auth/integration_test.go
package auth_test

import (
    "testing"
    "myapp/api/internal/auth"  // ✓ Test package can import
)

// External tests still work with internal packages (within same scope)
func TestAuthServiceIntegration(t *testing.T) {
    service := auth.NewService()
    token := service.GenerateToken(456)
    
    userID, err := service.ValidateToken(token)
    if err != nil {
        t.Fatalf("Token validation failed: %v", err)
    }
    
    if userID != 456 {
        t.Errorf("Expected userID 456, got %d", userID)
    }
}
```

### Refactoring with Internal Packages
```go
// Before: everything public, hard to change
package calculator

// Public - external users depend on this
func Add(a, b int) int { return a + b }

// Public but should be internal - hard to change without breaking users
func FormatResult(result int) string { return fmt.Sprintf("Result: %d", result) }
func ValidateInput(a, b int) error { /* validation */ }

// After: clean public API with internal implementation
// File: calculator/calculator.go
package calculator

import "myapp/calculator/internal/formatting"
import "myapp/calculator/internal/validation"

// Public API - stable interface
func Add(a, b int) (string, error) {
    if err := validation.ValidateNumbers(a, b); err != nil {
        return "", err
    }
    
    result := a + b
    return formatting.FormatResult(result), nil
}

// File: calculator/internal/formatting/formatter.go
package formatting

// Internal - can change implementation without breaking external users
func FormatResult(result int) string {
    // Can change format without breaking external API
    return fmt.Sprintf("Answer: %d", result)
}

// File: calculator/internal/validation/validator.go
package validation  

// Internal - validation rules can evolve independently
func ValidateNumbers(a, b int) error {
    if a < 0 || b < 0 {
        return errors.New("negative numbers not supported")
    }
    return nil
}
```

### When to Use Internal Packages
```go
// ✓ Use internal/ for:
// - Implementation details that might change
// - Shared utilities specific to your module
// - Complex logic that should be hidden behind clean APIs
// - Package-private functionality

// ✗ Don't use internal/ for:
// - Code that truly needs to be reusable across modules
// - Stable APIs that external users should access
// - Simple packages with no implementation hiding needed

// Example: Good use of internal/
mywebapp/
├── api/           # Public HTTP API
├── internal/
│   ├── auth/      # Authentication logic (might change)
│   ├── storage/   # Database abstractions (implementation detail)
│   └── email/     # Email sending (implementation detail)
└── cmd/
    └── server/    # Application entry point
```

## Running the Code

```bash
go run *.go
go test ./...
go test ./internal/...  # Test internal packages
```

## Java Developer Notes

- `internal/` ≈ package-private visibility in Java
- Better than Java's package-private - enforced at import level, not just access level
- No equivalent to Java's `protected` - Go prefers composition over inheritance
- Internal packages prevent "implementation creep" where internals become de facto public APIs
- Similar to Java's module system (Project Jigsaw) but simpler and directory-based
- Helps maintain backwards compatibility by clearly separating public and private APIs

## Next Steps

Continue to [Chapter 10: Advanced Topics](../10-advanced/)

## References

- [Go Command - Internal directories](https://pkg.go.dev/cmd/go#hdr-Internal_Directories)
- [Go Blog - Package organization](https://go.dev/blog/organizing-go-code)
- [Effective Go - Package structure](https://go.dev/doc/effective_go#package-names)