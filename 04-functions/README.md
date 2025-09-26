# Chapter 4: Functions

## Overview

Go functions are first-class citizens that support multiple return values, method receivers, and powerful closure patterns. Understanding function mechanics, method receivers, and Go's enum pattern using `iota` is essential for effective Go programming.

## Key Concepts

- **Function syntax** - multiple return values, named returns, variadic parameters
- **Method receivers** - `(s S)` vs `(s *S)` and when to use each
- **Closures** - functions that capture and maintain state from their lexical scope
- **Enum pattern** - typed constants with `iota` for enumeration
- **Generics** - type parameters for reusable functions

## Function Basics

Functions in Go can return multiple values and support named return parameters for clarity:

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

See [`functions.go`](./functions.go) for comprehensive examples of basic function patterns.

## Method Receivers

Methods in Go are functions with a special receiver argument. The receiver appears between the `func` keyword and the method name:

```go
func (c *Counter) increment() {
    c.value++
}
```

### Value vs Pointer Receivers

**Use pointer receivers when:**
- Method needs to modify the receiver
- Receiver is a large struct (avoid copying)
- Consistency across all methods on a type

**Use value receivers when:**
- Method doesn't modify the receiver
- Receiver is small
- You want immutable semantics

See [`receivers.go`](./receivers.go) for detailed examples and the decision-making process.

## Closures

Closures are functions that capture variables from their enclosing scope, creating powerful patterns for state management and functional programming:

```go
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x  // captures 'sum' from outer scope
        return sum
    }
}
```

### Advanced Closure Patterns

Go's closures enable sophisticated programming patterns:

- **Middleware** - Composable request/response processing chains
- **Rate Limiting** - Token bucket algorithms with state management
- **Memoization** - Caching expensive function results
- **Event Systems** - Observer pattern implementations
- **Generators** - Infinite sequence producers
- **Retry Logic** - Exponential backoff for resilience

See [`closures.go`](./closures.go) for foundational closure examples and [`advanced_closures.go`](./advanced_closures.go) for production-ready patterns including middleware chains, rate limiters, memoization, event emitters, and more.

## Generics

Go's generics enable type-safe, reusable functions:

```go
func dotProduct[F ~float32|~float64](v1, v2 []F) F {
    var sum F
    for i, x := range v1 {
        sum += x * v2[i]
    }
    return sum
}
```

See [`generics.go`](./generics.go) for complete generic function examples.

## Enums with iota

Go implements enums using typed constants with `iota`:

```go
type Status int

const (
    Pending Status = iota  // 0
    Running               // 1
    Completed             // 2
    Failed                // 3
)
```

The `iota` identifier provides successive integer constants, resetting to 0 in each const block.

See [`enums.go`](./enums.go) for enum implementations with string methods, validation, and advanced patterns.

## Running Examples

Each file contains a main function demonstrating its concepts:

```bash
go run functions.go
go run receivers.go
go run closures.go
go run advanced_closures.go
go run generics.go
go run enums.go

# Or run all examples:
go run main.go
```

## Java Developer Notes

- No function overloading - use different names or generics
- Methods belong to types, not classes
- Pointer receivers ≈ modifying `this`, value receivers ≈ immutable methods
- Multiple return values eliminate wrapper classes
- Closures similar to lambdas but with lexical variable capture
- No built-in enum keyword - use typed constants with `iota`

## Next Steps

Continue to [Chapter 5: Structs](../05-structs/)