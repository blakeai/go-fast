# Chapter 4: Functions

## Overview

Go functions support multiple return values, method receivers, and the enum pattern using `iota`. Understanding the difference between value and pointer receivers is essential for effective Go programming, and typed constants provide type-safe enumeration patterns.

## Key Concepts

- Function syntax and multiple return values
- **Method receivers** - `(s S)` vs `(s *S)` and method set rules
- Method definition restrictions (same package rule)
- **Enum pattern** - typed constants with `iota`
- Variadic functions
- Function as first-class values

## Examples

### Basic Functions
See [`functions.go`](./functions.go) for implementation.

```go
// Simple function
func add(a, b int) int {
    return a + b
}

// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Named return values
func calculate(a, b int) (sum, product int) {
    sum = a + b
    product = a * b
    return  // naked return
}
```

### Understanding Receivers
See [`receivers.go`](./receivers.go) for implementation.

In Go, parentheses before a function define a **receiver**, transforming a regular function into a method bound to a specific type:

```go
func (r Receiver) methodName() {
    // method body
}
```

The receiver `(r Receiver)` allows you to call the function as `instance.methodName()` rather than a standalone function.

#### Value vs Pointer Receivers

The choice between value and pointer receivers determines whether your method can modify the original instance:

```go
type Counter struct {
    value int
}

// Value receiver - operates on a copy
func (c Counter) increment() {
    c.value++ // Only modifies the copy
}

// Pointer receiver - operates on the original
func (c *Counter) incrementPtr() {
    c.value++ // Modifies the actual instance
}
```

**Value receivers** `(c Counter)`:
- Create a copy of the receiver when called
- Cannot modify the original instance
- Ideal for immutable design patterns
- More efficient for small types
- Safer for concurrent access

**Pointer receivers** `(c *Counter)`:
- Operate directly on the original instance
- Required for mutation
- Avoid copying large structs
- Should be used consistently across all methods on a type

#### Method Set Rules

Go automatically handles conversions between values and pointers:

```go
var v Counter
var p *Counter = &v

v.incrementPtr()  // Go automatically takes the address
p.increment()     // Go automatically dereferences
```

### Method Definition Rules
```go
// You can only define methods on types in the same package:

// This works - Counter is defined in this package
func (c *Counter) Reset() { c.value = 0 }

// This fails - string is not in this package
func (s string) NewMethod() {}     // ERROR

// This fails - time.Time is not in this package  
func (t time.Time) MyMethod() {}   // ERROR

// Workaround - wrap in your own type
type MyString string
func (s MyString) Upper() string {   // OK - MyString is in this package
    return strings.ToUpper(string(s))
}
```

### Value vs Pointer Receivers Decision Tree
```go
type Person struct {
    Name string
    Age  int
}

// Use pointer receiver when:
// 1. Method needs to modify the receiver
func (p *Person) SetAge(age int) {
    p.Age = age
}

// 2. Receiver is large (avoid copying)
func (p *Person) GetFullInfo() string {
    return fmt.Sprintf("%s is %d years old", p.Name, p.Age)  // Use pointer to avoid copy
}

// Use value receiver when:
// 1. Method doesn't modify receiver AND receiver is small
func (p Person) IsAdult() bool {
    return p.Age >= 18
}

// 2. You want to work with both values and pointers seamlessly
var person Person = Person{Name: "Alice", Age: 25}
var personPtr *Person = &person

person.IsAdult()     // works
personPtr.IsAdult()  // also works - Go dereferences automatically
```

### Generics in Functions
See [`generics.go`](./generics.go) for implementation.

```go
// Single type parameter
func dotProduct[F ~float32|~float64](v1, v2 []F) F {
    var sum F
    for i, x := range v1 {
        sum += x * v2[i]
    }
    return sum
}

// Multiple type parameters  
func convert[T any, U any](input T, converter func(T) U) U {
    return converter(input)
}

// Usage
result1 := dotProduct([]float64{1.0, 2.0}, []float64{3.0, 4.0})  // F = float64
result2 := dotProduct([]float32{1.0, 2.0}, []float32{3.0, 4.0})  // F = float32

stringResult := convert(42, strconv.Itoa)  // T=int, U=string
```

## Go's Enum Pattern

Go doesn't have built-in enums but achieves similar functionality using typed constants with `iota`.

### Basic Enum Implementation
See [`enums.go`](./enums.go) for implementation.

```go
type Status int

const (
    Pending Status = iota  // 0
    Running               // 1
    Completed             // 2
    Failed                // 3
)
```

### Understanding iota

`iota` is a predeclared identifier that represents successive untyped integer constants, starting at 0 and incrementing by 1:

```go
const (
    a = iota  // 0
    b         // 1 (implicitly = iota)
    c         // 2
    d         // 3
)
```

The automatic repetition of expressions is a property of const blocks, not `iota` itself:

```go
const (
    x = 42
    y     // y = 42 (repeats the expression)
    z     // z = 42
)
```

`iota` resets to 0 in each new const block and is only valid within const declarations.

### Enhanced Enum Patterns

Add string representation:

```go
func (s Status) String() string {
    switch s {
    case Pending: return "Pending"
    case Running: return "Running"
    case Completed: return "Completed"
    case Failed: return "Failed"
    default: return "Unknown"
    }
}
```

Custom values and expressions:

```go
type Priority int

const (
    Low Priority = iota + 1  // 1
    Medium                   // 2
    High                     // 3
    Critical = 10           // explicit value
)

// Skip values with blank identifier
const (
    _ = iota     // 0 (discarded)
    KB = 1 << (10 * iota)  // 1024
    MB                     // 1048576
    GB                     // 1073741824
)
```

String-based enums:

```go
type Color string

const (
    Red   Color = "red"
    Green Color = "green" 
    Blue  Color = "blue"
)
```

Add validation:

```go
func (s Status) IsValid() bool {
    return s >= Pending && s <= Failed
}
```

## Running the Code

```bash
go run *.go
go test ./...
```

## Java Developer Notes

- No function overloading - use different names or generics
- Methods belong to types, not classes
- Pointer receivers ≈ modifying `this`, value receivers ≈ immutable methods
- Multiple return values eliminate need for wrapper classes
- No constructors - use factory functions or struct literals
- Same package restriction prevents "monkey patching" like Ruby/JavaScript
- **Enums**: No built-in enum keyword - use typed constants with `iota`
- **Receivers**: Similar to instance methods but declared outside the type
- **Method sets**: Go automatically handles value/pointer conversions

## Next Steps

Continue to [Chapter 5: Structs](../05-structs/)

## References

- [Go Tour - Methods](https://tour.golang.org/methods/1)
- [Effective Go - Methods](https://golang.org/doc/effective_go.html#methods)
- [Go Spec - Method declarations](https://golang.org/ref/spec#Method_declarations)
- [Go Spec - Iota](https://golang.org/ref/spec#Iota)
- [Go Blog - Constants](https://blog.golang.org/constants)