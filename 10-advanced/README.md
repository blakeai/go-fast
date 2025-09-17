# Chapter 6: Interfaces

## Overview

Go interfaces define behavior contracts and are satisfied implicitly. A key principle: pass interfaces as values, not pointers - the interface already handles indirection when needed.

## Key Concepts

- Interface definition and implicit satisfaction
- **No pointers to interfaces** - `SomeInterface` not `*SomeInterface`
- Empty interface and type assertions
- Interface composition
- Interface values and nil interfaces

## Examples

### Basic Interface Usage
See [`interfaces.go`](./interfaces.go) for implementation.

```go
type Counter interface {
    Increment()
    Value() int
}

type IntCounter int

// Pointer receiver - can modify the original  
func (c *IntCounter) Increment() { 
    (*c)++ 
}

func (c *IntCounter) Value() int { 
    return int(*c) 
}

// Usage - pass interface as value
var counter Counter = &IntCounter{0}  // interface holds pointer to IntCounter
counter.Increment() // modifies the underlying IntCounter
fmt.Println(counter.Value()) // prints: 1
```

### Why No Pointers to Interfaces?
```go
// Correct - interface as value
var counter Counter = &IntCounter{0}
counter.Increment() // works perfectly

// Wrong - don't do this
var badCounter *Counter  // pointer to interface
// badCounter = &counter // confusing and unnecessary

// The interface value already provides the indirection you need
```

### Interface Satisfaction
```go
type Writer interface {
    Write([]byte) (int, error)
}

type FileWriter struct {
    filename string
}

// FileWriter automatically satisfies Writer interface
func (fw *FileWriter) Write(data []byte) (int, error) {
    // implementation here
    return len(data), nil
}

// No explicit "implements" declaration needed
var w Writer = &FileWriter{filename: "output.txt"}
```

### Multiple Interface Implementation
```go
type ReadWriter interface {
    Read([]byte) (int, error)
    Write([]byte) (int, error)
}

type Buffer struct {
    data []byte
}

func (b *Buffer) Read(p []byte) (int, error) {
    // implementation
    return 0, nil
}

func (b *Buffer) Write(p []byte) (int, error) {
    b.data = append(b.data, p...)
    return len(p), nil
}

// Buffer satisfies ReadWriter automatically
var rw ReadWriter = &Buffer{}
```

### Interface Composition
```go
// Interfaces can embed other interfaces
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

type ReadWriter interface {
    Reader  // embedded interface
    Writer  // embedded interface
}

// Equivalent to:
// type ReadWriter interface {
//     Read([]byte) (int, error)
//     Write([]byte) (int, error)
// }
```

### Empty Interface and Type Assertions
```go
// Empty interface - any type satisfies it
var anything interface{} = 42
anything = "hello"
anything = []int{1, 2, 3}

// Type assertion - extract concrete type
if str, ok := anything.(string); ok {
    fmt.Println("It's a string:", str)
}

// Type switch
switch v := anything.(type) {
case int:
    fmt.Println("Integer:", v)
case string:
    fmt.Println("String:", v)
case []int:
    fmt.Println("Slice:", v)
default:
    fmt.Println("Unknown type")
}
```

### Nil Interfaces vs Nil Values
```go
var counter Counter  // nil interface

var intCounter *IntCounter  // nil pointer
var counter2 Counter = intCounter  // interface with nil value

// Check for nil interface
if counter == nil {
    fmt.Println("nil interface")
}

// Check for nil value inside interface
if counter2 == nil {
    fmt.Println("this won't print - interface is not nil, value is")
}

// Proper nil value check
if reflect.ValueOf(counter2).IsNil() {
    fmt.Println("interface contains nil value")
}
```

## Running the Code

```bash
go run *.go
go test ./...
```

## Java Developer Notes

- No explicit `implements` declaration - interfaces satisfied implicitly
- Similar to Java interfaces but more flexible
- Empty interface `interface{}` ≈ Java's `Object`
- Type assertions ≈ casting, but safer with ok idiom
- Interface values can hold any type that satisfies the interface
- No inheritance - use composition and interface embedding

## Next Steps

Continue to [Chapter 7: Concurrency](../07-concurrency/)

## References

- [Go Tour - Interfaces](https://tour.golang.org/methods/9)
- [Effective Go - Interfaces](https://golang.org/doc/effective_go.html#interfaces)
- [Go Blog - Interface Values](https://blog.golang.org/laws-of-reflection)