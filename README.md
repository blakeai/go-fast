# Go Fast

> ### A learning guide for engineers who want to learn Go fast

---

This repository is a comprehensive, hands-on guide to learning Go, written from the perspective of a Java engineer
transitioning to Go. We cover essential concepts, common pitfalls, and idiomatic patterns with practical examples and tests.

**Status: Work in Progress** - Content is being actively developed and refined.

### How to Use This Guide

Each chapter is self-contained with examples and tests. Run tests with:
```bash
go test ./...
```

Or test a specific chapter:
```bash
go test ./01-basics/
```

## Table of Contents

### [Chapter 1: Go Basics](./01-basics/)
- Go syntax and program structure
- Basic data types and zero values
- Package system and the main function
- **Automatic Semicolon Insertion** - why brace placement matters

### [Chapter 2: Variables](./02-variables/)
- Variable declaration patterns (`var`, `:=`, explicit typing)
- Constants and `iota`
- **Addressable vs Non-addressable values** - why `&m[key]` fails
- Zero values and type inference

### [Chapter 3: Control Flow](./03-control-flow/)
- **Single loop construct** - `for` handles all iteration patterns
- **ASI and brace placement** - opening braces must be on same line
- **Switch statements** don't fall through (no `break` needed)
- **Short variable declarations** in `if` statements for scoped variables
- **Label breaks** and `goto` for complex control flow

### [Chapter 4: Functions](./04-functions/)
- Function syntax and multiple return values
- **Method Receivers** - `(s S)` vs `(s *S)` and method set rules
- Method definition restrictions (same package rule)
- **Enum pattern** - typed constants with `iota`
- **Generics syntax** - `[T constraint]` and type parameters

### [Chapter 5: Structs](./05-structs/)
- Struct definition and initialization
- Struct composition and embedding
- Working with addressable struct fields

### [Chapter 6: Interfaces](./06-interfaces/)
- Interface definition and implicit satisfaction
- **No pointers to interfaces** - pass `SomeInterface`, not `*SomeInterface`
- **Union types via interfaces** - Go's alternative to `TypeA | TypeB`
- Empty interface and type assertions
- Interface composition and nil interface gotchas

### [Chapter 7: Concurrency](./07-concurrency/)
- **Channels**: CSP-based communication primitives (buffered vs unbuffered)
- **Defer statements**: LIFO execution and guaranteed cleanup
- **WaitGroups**: Goroutine coordination with `sync.WaitGroup`
- **defer wg.Done()** pattern - why it goes at the top
- Select statements and concurrency patterns

### [Chapter 8: Error Handling](./08-error-handling/)
- The `error` interface
- Error creation and wrapping
- Error handling patterns

### [Chapter 9: Packages](./09-packages/)
- Package organization
- Visibility rules (exported vs unexported)
- Internal packages

### [Chapter 10: Advanced Topics](./10-advanced/)
- **Lexer and tokenization** - how Go parses source code
- **HTTP ServeMux vs Server** - understanding web architecture
- **Generics deep dive** - constraints and type inference
- Reflection basics

### [Chapter 11: String Formatting](./11-string-formatting/)
- Format verbs (`%s`, `%d`, `%v`, `%+v`, `%T`)
- Precision and padding control
- **Printf vs Sprintf** - when to use each
- Performance considerations for string building

## Key Insights from This Guide

This guide emphasizes concepts that trip up developers coming from other languages:

### **Lexer & ASI (Automatic Semicolon Insertion)**
Go's lexer automatically inserts semicolons, which explains why brace placement matters:
```go
// Breaks due to ASI
if condition
{  // becomes: if condition ; {

// Works
if condition {
```

### **Addressability Rules**
Not everything can have its address taken:
```go
m := map[string]Person{"key": person}
ptr := &m["key"]  // ERROR - map values not addressable
```

### **Method Receivers**
The parentheses before function names define method receivers:
```go
func (s *S) Write(data string) { s.data = data }  // pointer receiver - can modify
func (s S) Read() string { return s.data }        // value receiver - read-only
```

### **Interface Guidelines**
Pass interfaces as values, not pointers:
```go
var counter Counter = &IntCounter{0}  // interface holds pointer
var badCounter *Counter               // don't do this
```

### **HTTP Architecture**
ServeMux is just routing, not a complete server:
```go
mux := http.NewServeMux()      // Just routing
server := &http.Server{        // Complete server
    Handler: mux,              // ServeMux plugs into server
}
```

### **Generics Syntax**
Square brackets for type parameters:
```go
func dotProduct[F ~float32|~float64](v1, v2 []F) F {
    // F can be float32, float64, or types with those as underlying types
    // All uses of F in one call must be same concrete type
}
```

### **Defer and WaitGroups**
Defer executes in LIFO order and captures arguments immediately:
```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done() // registers cleanup, executes when func returns
    // work happens here
    // wg.Done() guaranteed to execute even if panic occurs
}()
```

### **Enum Pattern with iota**
Go uses typed constants for enums, with `iota` for auto-incrementing:
```go
type Status int
const (
    Pending Status = iota  // 0
    Running               // 1 
    Completed             // 2
)
```

## Prerequisites

- Go 1.21+ installed
- Basic programming experience (examples compare to Java concepts)

## Running the Examples

1. Clone this repository
2. Run `go mod tidy` to download dependencies
3. Navigate to any chapter and run the examples
4. All code is tested - use tests as additional documentation

## Code Quality and Linting

This repository uses standard Go tooling for consistent code quality:

### Tools Used
- **goimports** - Canonical Go formatter (includes gofmt + import management)
- **golangci-lint** - Meta-linter running ~50 linters including staticcheck, govet, revive
- **go vet** - Built-in Go error detection

### Setup
```bash
# Install tools
make install-tools

# Or manually:
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### Usage
```bash
# Format all code
make fmt

# Run all linters
make lint

# Run tests
make test

# Run all checks (fmt + lint + test)
make check

# Development workflow
make dev-check  # fmt + test (faster)
```

### Configuration
- **`.golangci.yml`** - Comprehensive linter configuration with educational-friendly settings
- **`Makefile`** - Standard targets for code quality checks
- **Editor integration** - Configure your editor to run goimports on save

### CI Integration
The `make check` target runs all quality checks and is perfect for CI pipelines.

## Contributing

This is a learning resource in active development. If you find errors or want to suggest improvements, please open an issue or PR.
