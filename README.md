# Go Learning Guide

A comprehensive, hands-on guide to learning Go, written by a Java engineer transitioning to Go. This repository covers essential concepts, common pitfalls, and idiomatic patterns with practical examples and tests.

**Status: Work in Progress** - Content is being actively developed and refined.

## How to Use This Guide

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
- if/else statements and brace placement rules
- for loops and range
- switch statements

### [Chapter 4: Functions](./04-functions/)
- Function syntax and multiple return values
- **Method Receivers** - `(s S)` vs `(s *S)`
- Method definition restrictions (same package rule)
- **Generics syntax** - `[T constraint]` and type parameters

### [Chapter 5: Structs](./05-structs/)
- Struct definition and initialization
- Struct composition and embedding
- Working with addressable struct fields

### [Chapter 6: Interfaces](./06-interfaces/)
- Interface definition and implicit satisfaction
- **No pointers to interfaces** - pass `SomeInterface`, not `*SomeInterface`
- Empty interface and type assertions
- Interface composition and nil interface gotchas

### [Chapter 7: Concurrency](./07-concurrency/)
- Goroutines and the `go` keyword
- Channels and channel operations
- Select statements and patterns

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

## Prerequisites

- Go 1.21+ installed
- Basic programming experience (examples compare to Java concepts)

## Running the Examples

1. Clone this repository
2. Run `go mod tidy` to download dependencies
3. Navigate to any chapter and run the examples
4. All code is tested - use tests as additional documentation

## Contributing

This is a learning resource in active development. If you find errors or want to suggest improvements, please open an issue or PR.