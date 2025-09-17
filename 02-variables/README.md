# Chapter 2: Variables

## Overview

Go offers multiple ways to declare and initialize variables with strong type safety and automatic type inference. Understanding addressable vs non-addressable values is crucial for working with pointers and method receivers.

## Key Concepts

- Variable declaration patterns: `var`, `:=`, and explicit typing
- Zero values and initialization
- Constants and `iota`
- **Addressable values** - what can you take the address of?
- Type inference vs explicit typing

## Examples

### Variable Declaration Patterns
See [`variables.go`](./variables.go) for implementation.

```go
// Explicit type declaration
var name string = "Alice"
var age int = 30

// Type inference
var city = "New York"     // string inferred
var population = 8000000  // int inferred

// Short variable declaration (inside functions only)
username := "bob"
count := 42

// Multiple variables
var x, y int = 1, 2
a, b := "hello", "world"

// Group declaration
var (
    firstName string = "John"
    lastName  string = "Doe" 
    salary    int    = 50000
)
```

### Short Declaration vs Var
```go
// Inside functions - short declaration preferred
func example() {
    name := "Alice"        // preferred - concise
    var age int = 30       // verbose but sometimes needed
    var height float64     // zero value initialization
    
    // Short declaration can mix new and existing variables
    name, email := "Bob", "bob@example.com"  // name reassigned, email new
}

// Package level - must use var
var globalCounter int = 0
// globalCounter := 0  // ERROR - := not allowed at package level

// When to use var inside functions:
func whenToUseVar() {
    var buffer strings.Builder  // zero value needed
    var users []User           // nil slice preferred over empty slice
    var config Config          // zero value struct
    
    // vs short declaration when you have initial values
    name := "Alice"
    count := len(items)
}
```

### Constants and Iota
```go
// Basic constants
const Pi = 3.14159
const MaxUsers = 100
const AppName = "MyApp"

// Grouped constants
const (
    StatusActive   = 1
    StatusInactive = 0
    StatusPending  = 2
)

// iota for auto-incrementing constants
const (
    Sunday = iota  // 0
    Monday         // 1
    Tuesday        // 2
    Wednesday      // 3
    Thursday       // 4
    Friday         // 5
    Saturday       // 6
)

// iota with expressions
const (
    B  = 1 << (10 * iota)  // 1
    KB                     // 1024
    MB                     // 1048576
    GB                     // 1073741824
    TB                     // 1099511627776
)

// iota resets in each const block
const (
    Red = iota    // 0
    Green         // 1
    Blue          // 2
)
const (
    Small = iota  // 0 (reset)
    Medium        // 1
    Large         // 2
)
```

### Addressable vs Non-Addressable Values
See [`addressable.go`](./addressable.go) for implementation.

**Addressable values** - you can use `&` to get their address:
```go
var x int = 5
ptr := &x  // OK - variables are addressable

arr := [3]int{1, 2, 3}
ptr := &arr[0]  // OK - array elements are addressable

type Person struct { name string }
var person Person
ptr := &person       // OK - struct variables are addressable
namePtr := &person.name  // OK - fields of addressable structs are addressable
```

**Non-addressable values** - cannot take address:
```go
// Map values are never addressable
m := map[int]Person{1: {name: "Alice"}}
ptr := &m[1]       // ERROR - map values not addressable
namePtr := &m[1].name  // ERROR - can't access fields of non-addressable values

// Function return values are not addressable
func getPerson() Person { return Person{name: "Bob"} }
ptr := &getPerson()    // ERROR - function results not addressable

// Literals are not addressable
ptr := &42             // ERROR - numeric literals not addressable
ptr := &"hello"        // ERROR - string literals not addressable
ptr := &Person{name: "Carol"}  // ERROR - struct literals not addressable
```

### Why Addressability Matters
```go
type Counter struct {
    value int
}

func (c *Counter) Increment() { c.value++ }  // pointer receiver
func (c Counter) Value() int { return c.value }  // value receiver

// This works - counter variable is addressable
var counter Counter
counter.Increment()  // Go automatically takes address: (&counter).Increment()

// This fails - map values not addressable
m := map[string]Counter{"main": {value: 0}}
m["main"].Increment()  // ERROR - can't get &m["main"]

// Solutions:
// 1. Store pointers in map
m1 := map[string]*Counter{"main": &Counter{value: 0}}
m1["main"].Increment()  // OK - m1["main"] is already a pointer

// 2. Extract to variable, modify, put back
counter = m["main"]
counter.Increment()
m["main"] = counter

// 3. Use value receiver methods only
fmt.Println(m["main"].Value())  // OK - value receiver works fine
```

### Slice vs Array Addressability
```go
// Array elements are addressable
arr := [3]int{1, 2, 3}
ptr := &arr[1]  // OK - array elements are addressable

// Slice elements are addressable
slice := []int{1, 2, 3}
ptr := &slice[1]  // OK - slice elements are addressable

// But slice from non-addressable array...
func getArray() [3]int { return [3]int{1, 2, 3} }
ptr := &getArray()[1]  // ERROR - elements of non-addressable array

// Slice header itself
var s []int
headerPtr := &s  // OK - slice variable is addressable
```

### Type Inference and Explicit Typing
```go
// Type inference
var a = 42          // int
var b = 3.14        // float64
var c = "hello"     // string
var d = true        // bool

// Sometimes you need explicit typing
var smallInt int8 = 42      // without explicit type, would be int
var precise float32 = 3.14  // without explicit type, would be float64

// Interface variables often need explicit typing
var w io.Writer = &bytes.Buffer{}  // explicit interface type needed

// Zero value initialization requires explicit type
var count int        // 0
var name string      // ""
var active bool      // false
var items []string   // nil

// Type conversion (not type coercion - must be explicit)
var i int = 42
var f float64 = float64(i)  // explicit conversion required
var j int8 = int8(i)        // explicit conversion required

// String conversions
var r rune = 65
var s string = string(r)    // "A"
var b byte = 65
var s2 string = string(b)   // "A"
```

### Variable Scope and Shadowing
```go
var global = "global"

func scopeExample() {
    var outer = "outer"
    
    if true {
        var inner = "inner"
        var outer = "shadowed"  // shadows outer variable
        fmt.Println(outer)      // prints "shadowed"
        fmt.Println(inner)      // prints "inner"
    }
    
    fmt.Println(outer)  // prints "outer" - original outer restored
    // fmt.Println(inner)  // ERROR - inner out of scope
    
    // Short declaration can shadow too
    global := "local global"  // shadows package-level global
    fmt.Println(global)       // prints "local global"
}
```

## Running the Code

```bash
go run *.go
go test ./...
```

## Java Developer Notes

- No `final` keyword - use `const` for compile-time constants
- Zero values replace null - uninitialized `int` is 0, `string` is `""`
- Short variable declaration (`:=`) only works inside functions
- Map values are never addressable (unlike Java object references)
- Go's addressability rules prevent many runtime errors
- Type inference is stronger than Java's diamond operator
- No implicit type conversion - all casts must be explicit
- Variable shadowing works similarly to Java but with different scoping rules

## Next Steps

Continue to [Chapter 3: Control Flow](../03-control-flow/)

## References

- [Go Spec - Address operators](https://go.dev/ref/spec#Address_operators)
- [Go Tour - Variables](https://go.dev/tour/basics/8)
- [Effective Go - Variables](https://go.dev/doc/effective_go#variables)