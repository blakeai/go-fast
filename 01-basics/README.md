# Chapter 1: Go Basics

## Overview

Welcome to Go! This chapter introduces fundamental Go concepts including syntax, basic data types, and program structure. Go is designed for simplicity, readability, and efficient compilation, making it an excellent choice for modern software development.

## Key Concepts

- Go syntax and program structure
- Basic data types (int, string, bool, float)
- Variable declaration and initialization
- Package system and imports
- The `main` function and program entry point

## Examples

### Example 1: Hello World Program
See [`hello.go`](./hello.go) for implementation.

Key points:
- Every Go program starts with a `package` declaration
- The `main` function is the entry point
- `fmt.Println()` for output (similar to `System.out.println()` in Java)

### Example 2: Variable Declarations
See [`variables.go`](./variables.go) for implementation.

Key points:
- Multiple ways to declare variables: `var`, `:=`, and explicit typing
- Go infers types when possible
- Zero values for uninitialized variables

## Running the Code

```bash
go run *.go
go test ./...
```

## Java Developer Notes

- No semicolons required (unlike Java)
- Package names are lowercase by convention
- No classes - Go uses functions and structs instead
- Explicit error handling instead of exceptions
- `main()` function doesn't return anything (vs `public static void main` in Java)

## Next Steps

Continue to [Chapter 2: Variables](../02-variables/)

## References

- [Go Tour - Basics](https://tour.golang.org/basics/1)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go by Example - Hello World](https://gobyexample.com/hello-world)