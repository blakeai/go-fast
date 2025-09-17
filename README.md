# Go Learning Guide

A comprehensive, hands-on guide to learning Go, written by a Java engineer transitioning to Go. This repository is structured as a progressive course with practical examples, tests, and explanations.

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

### [Chapter 1: Basics](./01-basics/)
- Variables and constants
- Basic types and zero values
- Type inference and declarations

### [Chapter 2: Functions](./02-functions/)
- Function syntax and return values
- Multiple return values
- Variadic functions

### [Chapter 3: Structs](./03-structs/)
- Struct definition and initialization
- Struct composition and embedding
- Field tags

### [Chapter 4: Methods & Receivers](./04-methods-receivers/)
- Value receivers vs pointer receivers
- Method sets and interface satisfaction
- When to use each type

### [Chapter 5: Interfaces](./05-interfaces/)
- Interface definition and implementation
- Empty interface and type assertions
- Interface composition

### [Chapter 6: Concurrency](./06-concurrency/)
- Goroutines and the `go` keyword
- Channels and channel operations
- Select statements and patterns

### [Chapter 7: Error Handling](./07-error-handling/)
- The `error` interface
- Error creation and wrapping
- Error handling patterns

### [Chapter 8: Packages](./08-packages/)
- Package organization
- Visibility rules (exported vs unexported)
- Internal packages

### [Chapter 9: Advanced Topics](./09-advanced/)
- Reflection
- Generics (Go 1.18+)
- Build constraints

## Prerequisites

- Go 1.21+ installed
- Basic programming experience (examples compare to Java concepts)

## Running the Examples

1. Clone this repository
2. Run `go mod tidy` to download dependencies
3. Navigate to any chapter and run the examples
4. All code is tested - use tests as additional documentation

## Contributing

This is a learning resource. If you find errors or want to suggest improvements, please open an issue or PR.