# Chapter 4: Functions

## Overview

Go functions support multiple return values, method receivers, and have specific rules about where methods can be defined. Understanding the difference between value and pointer receivers is essential for effective Go programming.

## Key Concepts

- Function syntax and multiple return values
- **Method receivers** - `(s S)` vs `(s *S)`
- Method definition restrictions (same package rule)
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

### Method Receivers
See [`receivers.go`](./receivers.go) for implementation.

The parentheses before the function name define a method receiver:

```go
type Counter struct {
    value int
}

// Value receiver - gets a copy, can't modify original
func (c Counter) Value() int {
    return c.value
}

// Pointer receiver - gets pointer, can modify original
func (c *Counter) Increment() {
    c.value++
}

func (c *Counter) Add(n int) {
    c.value += n
}

// Usage
var counter Counter
counter.Increment()    // Go automatically takes address: (&counter).Increment()
fmt.Println(counter.Value())  // prints: 1
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

## Next Steps

Continue to [Chapter 5: Structs](../05-structs/)

## References

- [Go Tour - Methods](https://tour.golang.org/methods/1)
- [Effective Go - Methods](https://golang.org/doc/effective_go.html#methods)
- [Go Spec - Method declarations](https://golang.org/ref/spec#Method_declarations)