# Chapter 6: Interfaces

## Overview

Go interfaces define behavior contracts and are satisfied implicitly. They serve as Go's solution to union types, allowing multiple types to satisfy the same interface. A key principle: pass interfaces as values, not pointers - the interface already handles indirection when needed.

## Key Concepts

- Interface definition and implicit satisfaction
- **No pointers to interfaces** - `SomeInterface` not `*SomeInterface`
- **Union types via interfaces** - Go's alternative to `TypeA | TypeB`
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
var intCounter IntCounter = 0
var counter Counter = &intCounter  // interface holds pointer to IntCounter
counter.Increment() // modifies the underlying IntCounter
fmt.Println(counter.Value()) // prints: 1
```

### Why No Pointers to Interfaces?
```go
// Correct - interface as value
var intCounter IntCounter = 0
var counter Counter = &intCounter
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

## Go's Interface Solution to Union Types
See [`union-types.go`](./union-types.go) for implementation.

Union types let you express that a value can be one of several different types. Languages like TypeScript make this explicit with syntax like `string | number`, giving you compile-time guarantees about what types you're working with.

Go doesn't have union types, but it solves similar problems using interfaces and implicit satisfaction.

### The Interface Approach

Instead of declaring `EmailHandler | SMSHandler`, you define a common interface:

```go
type Handler interface {
    Handle() error
}

type EmailHandler struct {
    recipient string
    subject   string
    body      string
}

type SMSHandler struct {
    phoneNumber string
    message     string
}

func (e EmailHandler) Handle() error {
    // Send email logic
    return nil
}

func (s SMSHandler) Handle() error {
    // Send SMS logic  
    return nil
}
```

### Implicit Interface Satisfaction

The key difference from Java or C# is that Go uses implicit interface satisfaction. There's no `implements` keyword. The moment `EmailHandler` and `SMSHandler` define methods matching the `Handler` interface signature, they automatically satisfy it.

```go
func process(h Handler) error {
    return h.Handle()
}

// Both work automatically
var emailHandler Handler = EmailHandler{...}
var smsHandler Handler = SMSHandler{...}
```

### Type Discrimination

When you need type discrimination, Go provides type assertions and type switches:

```go
switch h := handler.(type) {
case EmailHandler:
    // Handle email-specific logic
    fmt.Printf("Sending email to %s\n", h.recipient)
case SMSHandler:
    // Handle SMS-specific logic
    fmt.Printf("Sending SMS to %s\n", h.phoneNumber)
default:
    fmt.Println("Unknown handler type")
}
```

### Trade-offs

This approach forces you to think about behavior rather than just data shape, which often leads to better design. The downside is you lose some compile-time type information that explicit union types would preserve.

Go's implicit interfaces offer flexibility through composition rather than explicit type unions, embracing the language's philosophy of simplicity over feature completeness.

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