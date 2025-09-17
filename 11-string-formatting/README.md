# Chapter 11: String Formatting

## Overview

Go's `fmt` package provides powerful string formatting through format verbs. Understanding the key verbs and when to use `Sprintf` vs `Printf` is essential for effective string manipulation and debugging.

## Key Concepts

- Format verbs for type-specific formatting
- Precision and padding control
- `Sprintf` vs `Printf` usage patterns
- Universal formatting with `%v`

## Examples

### Basic Format Verbs
See [`formatting.go`](./formatting.go) for implementation.

```go
name := "Alice"
age := 30
height := 5.8
active := true

// Core format verbs
fmt.Sprintf("Name: %s", name)           // strings
fmt.Sprintf("Age: %d", age)             // integers  
fmt.Sprintf("Height: %f", height)       // floats
fmt.Sprintf("Active: %t", active)       // booleans

// Universal formatter - works with any type
fmt.Sprintf("Person: %v", person)       // basic representation
fmt.Sprintf("Person: %+v", person)      // includes struct field names
fmt.Sprintf("Type: %T", person)         // prints the type itself
```

### Precision and Padding Control
```go
price := 19.99
id := 42
percentage := 0.8567

// Decimal precision
fmt.Sprintf("Price: $%.2f", price)      // "Price: $19.99"
fmt.Sprintf("Rate: %.1f%%", percentage*100) // "Rate: 85.7%"

// Integer padding
fmt.Sprintf("ID: %04d", id)             // "ID: 0042" (zero-padded)
fmt.Sprintf("ID: %4d", id)              // "ID:   42" (space-padded)

// String width and alignment
fmt.Sprintf("|%-10s|", "Go")           // "|Go        |" (left-aligned)
fmt.Sprintf("|%10s|", "Go")            // "|        Go|" (right-aligned)
```

### Advanced Format Verbs
```go
data := []byte("Hello")
ptr := &name

// Hexadecimal and binary
fmt.Sprintf("Hex: %x", data)            // "Hex: 48656c6c6f"
fmt.Sprintf("Binary: %b", 42)           // "Binary: 101010"

// Pointers and memory addresses  
fmt.Sprintf("Pointer: %p", ptr)         // "Pointer: 0xc000010240"

// Quoted strings (useful for debugging)
fmt.Sprintf("Quoted: %q", "hello\nworld") // "Quoted: \"hello\\nworld\""
```

### Complex Types
```go
type Person struct {
Name string
Age  int
}

person := Person{Name: "Bob", Age: 25}
people := []Person{person, {Name: "Carol", Age: 28}}
scores := map[string]int{"math": 95, "english": 87}

// Structs
fmt.Sprintf("%v", person)               // "{Bob 25}"
fmt.Sprintf("%+v", person)              // "{Name:Bob Age:25}"

// Slices and maps
fmt.Sprintf("%v", people)               // "[{Bob 25} {Carol 28}]"  
fmt.Sprintf("%v", scores)               // "map[english:87 math:95]"
```

### Printf vs Sprintf Usage Patterns
```go
// Use Sprintf when you need the string for further processing
func buildLogMessage(level string, msg string) string {
timestamp := time.Now().Format("2006-01-02 15:04:05")
return fmt.Sprintf("[%s] %s: %s", timestamp, level, msg)
}

// Use Printf when outputting immediately
func debugPrint(obj interface{}) {
fmt.Printf("Debug: %+v\n", obj)
}

// Common pattern: build complex strings piece by piece
var buffer strings.Builder
for i, item := range items {
line := fmt.Sprintf("Item %d: %s (%.2f)\n", i+1, item.Name, item.Price)
buffer.WriteString(line)
}
result := buffer.String()
```

### Error Formatting
```go
// Format errors with context
func processFile(filename string) error {
if _, err := os.Open(filename); err != nil {
return fmt.Errorf("failed to open file %q: %w", filename, err)
}
return nil
}

// Debug complex error chains
err := processFile("missing.txt")
fmt.Printf("Error: %v\n", err)          // basic error message
fmt.Printf("Error: %+v\n", err)         // may include stack trace with some error types
```

### Performance Considerations
```go
// For simple concatenation, prefer + or strings.Builder
simple := "Hello " + name              // faster for simple cases

// Use Sprintf for complex formatting
complex := fmt.Sprintf("User %s (%d) logged in at %v", name, id, time.Now())

// For repeated string building, use strings.Builder
var builder strings.Builder
for _, item := range items {
builder.WriteString(fmt.Sprintf("Item: %s\n", item))
}
result := builder.String()
```

## Running the Code

```bash
go run *.go
go test ./...
```

## Java Developer Notes

- Format verbs replace Java's `%s`, `%d` printf-style formatting
- `%v` is more flexible than Java's `toString()` - works universally
- No `String.format()` equivalent - use `fmt.Sprintf()`
- `%+v` for structs ≈ Java's automatic `toString()` with field names
- Error wrapping with `%w` ≈ Java's exception chaining

## Next Steps

Continue exploring Go's standard library packages for text processing and I/O operations.

## References

- [Go fmt package documentation](https://pkg.go.dev/fmt)
- [Go by Example - String Formatting](https://gobyexample.com/string-formatting)
- [Effective Go - Formatting](https://golang.org/doc/effective_go.html#formatting)