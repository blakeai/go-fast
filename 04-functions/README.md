# Chapter 4: Functions

## Overview

Go functions support multiple return values, method receivers, and the enum pattern using `iota`. Understanding the difference between value and pointer receivers is essential for effective Go programming, and typed constants provide type-safe enumeration patterns.

## Key Concepts

- Function syntax and multiple return values
- **Method receivers** - `(s S)` vs `(s *S)` and method set rules
- Method definition restrictions (same package rule)
- **Enum pattern** - typed constants with `iota`
- **Closures** - functions that capture variables from their enclosing scope
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

### Closures in Go
See [`closures.go`](./closures.go) for basic examples and [`advanced_closures.go`](./advanced_closures.go) for advanced patterns.

A **closure** is a function value that captures variables from its surrounding lexical scope. Even after the outer function returns, the inner function can continue to access and modify those captured variables.

#### Key Concepts

1. **Functions are first-class values** – you can assign functions to variables, pass them around, or return them
2. **Closures capture variables** – an inner function can reference variables outside its body
3. **State preservation** – closures let you maintain state across calls without global variables
4. **Each closure maintains its own state** – multiple closures from the same factory function are independent

#### Basic Closure Example

```go
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x  // captures and modifies 'sum' from outer scope
        return sum
    }
}

func main() {
    posSum := adder()
    fmt.Println(posSum(3))  // 3
    fmt.Println(posSum(5))  // 8  (sum persists between calls)
    fmt.Println(posSum(10)) // 18

    another := adder()      // independent closure with its own 'sum'
    fmt.Println(another(2)) // 2
}
```

#### Closure Factory Pattern

Closures are excellent for creating specialized functions with pre-configured behavior:

```go
func makeMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor  // 'factor' is captured from outer scope
    }
}

double := makeMultiplier(2)  // creates a function that doubles
triple := makeMultiplier(3)  // creates a function that triples

fmt.Println(double(5))  // 10
fmt.Println(triple(5))  // 15
```

#### Closures in Loops (Historical Note)

**Note**: Go 1.22+ fixed the classic closure-in-loop issue. The following now works correctly:

```go
// Go 1.22+ - works correctly (each iteration gets its own variable)
var funcs []func() int
for i := 0; i < 3; i++ {
    funcs = append(funcs, func() int {
        return i  // each closure captures its own 'i'
    })
}
// prints 0, 1, 2

// Historical pattern for older Go versions (still valid)
var correctFuncs []func() int
for i := 0; i < 3; i++ {
    correctFuncs = append(correctFuncs, func(val int) func() int {
        return func() int {
            return val  // captures the parameter value
        }
    }(i))
}
```

**For Go versions before 1.22**: All closures would capture the same variable reference, causing them to all return the final value of the loop variable.

#### Advanced Use Cases

**Validation Functions:**
```go
func createValidator(minLength int, maxLength int) func(string) bool {
    return func(input string) bool {
        length := len(input)
        return length >= minLength && length <= maxLength
    }
}

passwordValidator := createValidator(8, 20)
fmt.Println(passwordValidator("short"))    // false
fmt.Println(passwordValidator("good_pass")) // true
```

**Stateful Objects with Multiple Methods:**
```go
func createAccount(initial float64) (deposit func(float64), withdraw func(float64), balance func() float64) {
    currentBalance := initial

    deposit = func(amount float64) {
        currentBalance += amount
    }

    withdraw = func(amount float64) {
        if amount <= currentBalance {
            currentBalance -= amount
        }
    }

    balance = func() float64 {
        return currentBalance
    }

    return deposit, withdraw, balance
}
```

#### Advanced Closure Patterns

Go's closures enable sophisticated programming patterns commonly used in production applications. See [`advanced_closures.go`](./advanced_closures.go) for complete implementations.

##### 1. Middleware Pattern
Create composable middleware chains using closures:

```go
func middleware() func(string) string {
    return func(input string) string {
        logger := func(next func(string) string) func(string) string {
            return func(s string) string {
                fmt.Printf("[LOG] Processing: %s\n", s)
                result := next(s)
                fmt.Printf("[LOG] Result: %s\n", result)
                return result
            }
        }

        timer := func(next func(string) string) func(string) string {
            return func(s string) string {
                start := time.Now()
                result := next(s)
                fmt.Printf("[TIMER] Took %v\n", time.Since(start))
                return result
            }
        }

        handler := func(s string) string {
            return "Processed: " + s
        }

        // Build the chain: logger wraps timer wraps handler
        chain := logger(timer(handler))
        return chain(input)
    }
}
```

##### 2. Rate Limiter with State Management
Thread-safe rate limiting using closures and mutex:

```go
func rateLimiter(requests int, duration time.Duration) func() bool {
    tokens := requests
    lastReset := time.Now()
    mu := sync.Mutex{}

    return func() bool {
        mu.Lock()
        defer mu.Unlock()

        now := time.Now()
        if now.Sub(lastReset) >= duration {
            tokens = requests
            lastReset = now
        }

        if tokens > 0 {
            tokens--
            return true
        }
        return false
    }
}
```

##### 3. Memoization for Performance
Cache expensive function results using closures:

```go
func memoize(fn func(int) int) func(int) int {
    cache := make(map[int]int)
    return func(x int) int {
        if val, exists := cache[x]; exists {
            return val // Cache hit
        }
        result := fn(x)
        cache[x] = result
        return result
    }
}

// Usage
memoFib := memoize(fibonacci)
result := memoFib(40) // Computed once
result = memoFib(40)  // Returns cached result instantly
```

##### 4. Event Emitter Pattern
Implement event-driven programming with closures:

```go
func createEventEmitter() (on func(string, func()), emit func(string)) {
    listeners := make(map[string][]func())

    on = func(event string, handler func()) {
        listeners[event] = append(listeners[event], handler)
    }

    emit = func(event string) {
        if handlers, exists := listeners[event]; exists {
            for _, handler := range handlers {
                handler()
            }
        }
    }

    return on, emit
}
```

##### 5. Iterator Generators
Generate infinite sequences using closures:

```go
func fibonacciGenerator() func() int {
    a, b := 0, 1
    return func() int {
        result := a
        a, b = b, a+b
        return result
    }
}

// Usage
fibGen := fibonacciGenerator()
for i := 0; i < 10; i++ {
    fmt.Print(fibGen(), " ") // 0 1 1 2 3 5 8 13 21 34
}
```

##### 6. Builder Pattern with Fluent Interface
Create fluent APIs using closures:

```go
type QueryBuilder struct {
    table  string
    wheres []string
    limit  int
}

func (qb *QueryBuilder) Where(condition string) *QueryBuilder {
    qb.wheres = append(qb.wheres, condition)
    return qb
}

func (qb *QueryBuilder) Limit(n int) *QueryBuilder {
    qb.limit = n
    return qb
}

// Usage
query := createQueryBuilder()("users").
    Where("age > 18").
    Where("city = 'NYC'").
    Limit(10).
    Build()
```

##### 7. Dynamic Sorting with Comparators
Create flexible sorting logic using closures:

```go
func sortByAge(ascending bool) func(i, j int) bool {
    return func(i, j int) bool {
        if ascending {
            return people[i].Age < people[j].Age
        }
        return people[i].Age > people[j].Age
    }
}

sort.Slice(people, sortByAge(true))  // Ascending
sort.Slice(people, sortByAge(false)) // Descending
```

##### 8. Retry Logic with Exponential Backoff
Implement robust error handling patterns:

```go
func retryWithBackoff(maxAttempts int) func(func() error) error {
    return func(operation func() error) error {
        var lastErr error
        backoff := 100 * time.Millisecond

        for attempt := 1; attempt <= maxAttempts; attempt++ {
            if err := operation(); err == nil {
                return nil // Success
            } else {
                lastErr = err
                if attempt < maxAttempts {
                    time.Sleep(backoff)
                    backoff *= 2 // Exponential backoff
                }
            }
        }

        return fmt.Errorf("all %d attempts failed: %w", maxAttempts, lastErr)
    }
}
```

##### 9. Function Pipeline/Composition
Chain transformations using closures:

```go
func pipeline(funcs ...func(int) int) func(int) int {
    return func(x int) int {
        result := x
        for _, fn := range funcs {
            result = fn(result)
        }
        return result
    }
}

// Usage
double := func(x int) int { return x * 2 }
addTen := func(x int) int { return x + 10 }
square := func(x int) int { return x * x }

transform := pipeline(double, addTen, square)
result := transform(5) // (5*2 + 10)^2 = 400
```

##### 10. Debounce and Throttle Patterns
Control function execution timing:

```go
func debounce(fn func(), delay time.Duration) func() {
    var timer *time.Timer
    var mu sync.Mutex

    return func() {
        mu.Lock()
        defer mu.Unlock()

        if timer != nil {
            timer.Stop()
        }
        timer = time.AfterFunc(delay, fn)
    }
}

func throttle(fn func(), limit time.Duration) func() {
    var lastCall time.Time
    var mu sync.Mutex

    return func() {
        mu.Lock()
        defer mu.Unlock()

        now := time.Now()
        if now.Sub(lastCall) >= limit {
            fn()
            lastCall = now
        }
    }
}
```

These patterns demonstrate how closures enable elegant solutions for:
- **Middleware**: Composable request/response processing
- **State Management**: Encapsulating mutable state without globals
- **Caching**: Transparent performance optimization
- **Event Systems**: Decoupled component communication
- **Functional Programming**: Composition and transformation pipelines
- **Concurrency Control**: Rate limiting and timing patterns

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
- **Closures**: Similar to Java lambdas but with lexical variable capture (like JavaScript closures)

## Next Steps

Continue to [Chapter 5: Structs](../05-structs/)

## References

- [Go Tour - Methods](https://tour.golang.org/methods/1)
- [Effective Go - Methods](https://golang.org/doc/effective_go.html#methods)
- [Go Spec - Method declarations](https://golang.org/ref/spec#Method_declarations)
- [Go Spec - Iota](https://golang.org/ref/spec#Iota)
- [Go Blog - Constants](https://blog.golang.org/constants)