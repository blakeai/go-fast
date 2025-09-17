# Chapter 9: Packages

## Overview

Go organizes code into packages, which are directories containing related `.go` files. Understanding Go's module system, import paths, visibility rules, and the special `internal/` directory is essential for structuring larger applications.

## Key Concepts

- **Module vs Package** - modules contain packages, packages contain code
- **Import paths** - how Go finds and loads packages
- **Visibility rules** - exported (capitalized) vs unexported identifiers
- **Internal packages** - `internal/` restricts import access
- **Package initialization** - `init()` functions and import order

## Examples

### Basic Package Structure
See [`calculator/`](./calculator/) for implementation.

```go
// File: calculator/math.go
package calculator

import "errors"

// Exported function (capitalized)
func Add(a, b int) int {
    return a + b
}

// Exported function with error handling
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// unexported function (lowercase) - only accessible within package
func multiply(a, b int) int {
    return a * b
}

// Exported type
type Calculator struct {
    history []string
}

// Exported method
func (c *Calculator) Add(a, b int) int {
    result := Add(a, b)
    c.recordOperation("add", a, b, result)
    return result
}

// unexported method
func (c *Calculator) recordOperation(op string, a, b, result int) {
    entry := fmt.Sprintf("%s(%d, %d) = %d", op, a, b, result)
    c.history = append(c.history, entry)
}
```

### Import Patterns and Aliases
```go
package main

import (
    "fmt"                    // standard library
    "net/http"              // standard library
    "myapp/calculator"      // local package
    "github.com/gorilla/mux" // external module
    
    // Import aliases
    calc "myapp/calculator"  // shorter alias
    . "strings"             // dot import - use functions without qualifier
    _ "image/png"           // blank import - for side effects only
)

func main() {
    // Using imported packages
    fmt.Println("Hello")
    
    result := calculator.Add(5, 3)           // full package name
    result2 := calc.Add(10, 2)               // using alias
    
    upper := ToUpper("hello")                // dot import - no strings.ToUpper
    
    router := mux.NewRouter()                // external package
}
```

### Package-Level Variables and Constants
```go
package config

import (
    "os"
    "strconv"
)

// Package-level constants (exported)
const (
    DefaultPort = 8080
    MaxRetries  = 3
    AppVersion  = "1.0.0"
)

// Package-level variables (exported)
var (
    DatabaseURL string
    Debug       bool
    Port        int
)

// unexported package-level variables
var (
    initialized bool
    logger     *Logger
)

// Package initialization - runs once when package is imported
func init() {
    // Initialize package-level variables
    DatabaseURL = os.Getenv("DATABASE_URL")
    if DatabaseURL == "" {
        DatabaseURL = "localhost:5432"
    }
    
    Debug = os.Getenv("DEBUG") == "true"
    
    if portStr := os.Getenv("PORT"); portStr != "" {
        if p, err := strconv.Atoi(portStr); err == nil {
            Port = p
        }
    }
    if Port == 0 {
        Port = DefaultPort
    }
    
    initialized = true
}

// Exported function to check if package is ready
func IsInitialized() bool {
    return initialized
}
```

### Internal Packages
Directory structure:
```
myapp/
├── package_demo.go
├── api/
│   ├── handlers.go
│   └── internal/
│       └── auth/
│           └── jwt.go      // only importable by api/ and its subdirectories
├── database/
│   └── models.go
└── internal/
    └── shared/
        └── utils.go        // only importable by myapp/ and its subdirectories
```

```go
// File: api/internal/auth/jwt.go
package auth

// This package can only be imported by:
// - myapp/api/
// - myapp/api/handlers/
// - myapp/api/middleware/
// - etc. (any package within api/)

func GenerateToken(userID int) (string, error) {
    // JWT generation logic
    return "jwt-token", nil
}

// File: internal/shared/utils.go  
package shared

// This package can only be imported by:
// - myapp/ (root package)
// - myapp/api/
// - myapp/database/
// - etc. (any package within myapp/)

func FormatError(err error) string {
    return fmt.Sprintf("Error: %v", err)
}
```

### Module System (go.mod)
```go
// File: go.mod
module myapp

go 1.21

require (
    github.com/gorilla/mux v1.8.0
    github.com/lib/pq v1.10.9
)

require (
    github.com/gorilla/websocket v1.5.0 // indirect
)
```

Import paths in code correspond to module + package path:
```go
import (
    "myapp/api"              // module: myapp, package: api/
    "myapp/api/handlers"     // module: myapp, package: api/handlers/
    "myapp/internal/shared"  // module: myapp, package: internal/shared/
    
    "github.com/gorilla/mux" // external module
)
```

### Package Documentation
```go
// Package calculator provides basic arithmetic operations.
// It supports addition, subtraction, multiplication, and division
// with proper error handling for edge cases.
//
// Example usage:
//   result := calculator.Add(5, 3)
//   quotient, err := calculator.Divide(10, 2)
//   if err != nil {
//       log.Fatal(err)
//   }
package calculator

// Add returns the sum of two integers.
func Add(a, b int) int {
    return a + b
}

// Calculator provides arithmetic operations with history tracking.
type Calculator struct {
    history []Operation
}

// Operation represents a single arithmetic operation.
type Operation struct {
    Type   string  // "add", "subtract", etc.
    A, B   int     // operands
    Result int     // result
}

// NewCalculator creates a new Calculator instance.
func NewCalculator() *Calculator {
    return &Calculator{
        history: make([]Operation, 0),
    }
}
```

### Multiple init() Functions and Order
```go
package database

import (
    "database/sql"
    _ "github.com/lib/pq"  // blank import runs init()
)

var db *sql.DB

// Multiple init functions run in source order
func init() {
    fmt.Println("First init function")
}

func init() {
    fmt.Println("Second init function")
    // Database connection setup
    var err error
    db, err = sql.Open("postgres", "connection-string")
    if err != nil {
        panic(fmt.Sprintf("Failed to connect to database: %v", err))
    }
}

func init() {
    fmt.Println("Third init function")
    // Run migrations or other setup
}

// Package initialization order:
// 1. Import dependencies (their init functions run first)
// 2. Initialize package-level variables
// 3. Run init() functions in source order
```

### Package Testing Structure
```go
// File: calculator/math.go
package calculator

func Add(a, b int) int { return a + b }

// File: calculator/math_test.go
package calculator

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}

// File: calculator/math_integration_test.go
package calculator_test  // external test package

import (
    "testing"
    "myapp/calculator"  // must import explicitly
)

func TestCalculatorIntegration(t *testing.T) {
    // Test package from external perspective
    result := calculator.Add(10, 20)
    if result != 30 {
        t.Errorf("calculator.Add(10, 20) = %d; want 30", result)
    }
    
    // Cannot access unexported functions
    // calculator.multiply(2, 3)  // ERROR - unexported
}
```

### Circular Import Prevention
```go
// WRONG - creates circular import
// package a imports package b
// package b imports package a

// File: a/a.go
package a
import "myapp/b"
func UseB() { b.Function() }

// File: b/b.go  
package b
import "myapp/a"  // ERROR - circular import
func UseA() { a.Function() }

// SOLUTION - extract common interface or create shared package
// File: shared/interfaces.go
package shared
type Service interface { DoWork() }

// File: a/a.go
package a
import "myapp/shared"
type AService struct{}
func (s AService) DoWork() { /* implementation */ }

// File: b/b.go
package b
import "myapp/shared"
func UseService(s shared.Service) { s.DoWork() }
```

## Running the Code

```bash
go run *.go
go test ./...
go mod tidy    # clean up dependencies
go mod download # download dependencies
```

## Java Developer Notes

- Go packages ≈ Java packages, but simpler import system
- No nested packages - each directory is a separate package
- Capitalization determines visibility (no `public`/`private` keywords)
- `internal/` directories ≈ package-private in Java
- Module system ≈ Maven/Gradle dependency management
- `init()` functions ≈ Java static initializer blocks
- No circular imports allowed (stricter than Java)
- Package name doesn't need to match directory name (but should by convention)
- External tests (`package foo_test`) ≈ testing from outside the package

## Next Steps

Continue to [Chapter 9.1: Internal Packages](../09-packages-internal/)

## References

- [Go Modules Reference](https://go.dev/ref/mod)
- [Go Tour - Packages](https://go.dev/tour/basics/1)
- [Effective Go - Package names](https://go.dev/doc/effective_go#package-names)