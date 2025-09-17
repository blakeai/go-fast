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
var city = "New York"  // string inferred
var population = 8000000  // int inferred

// Short variable declaration (inside functions only)
username := "bob"
count := 42

// Multiple variables
var x, y int = 1, 2
a, b := "hello", "world"
```

### Addressable vs Non-Addressable Values
See [`addressable.go`](./addressable.go) for implementation.

**Addressable values** - you can use `&` to get their address:
```go
var x int = 5
ptr := &x  // OK - variables are addressable

arr := [3]int{1, 2, 3}
ptr := &arr[0]  // OK - array elements are addressable

type S struct { data string }
var s S
ptr := &s  // OK - struct variables are addressable
```

**Non-addressable values** - cannot take address:
```go
// Map values
m := map[int]S{1: {data: "hello"}}
ptr := &m[1]  // ERROR - map values not addressable

// Function return values  
func getS() S { return S{data: "test"} }
ptr := &getS()  // ERROR - function results not addressable

// Literals
ptr := &42  // ERROR - literals not addressable
```

### Why Addressability Matters
```go
type Counter struct {
    value int
}

func (c *Counter) Increment() { c.value++ }  // pointer receiver
func (c Counter) Value() int { return c.value }  // value receiver

// This works:
var counter Counter
counter.Increment()  // OK - counter is addressable

// This fails:
m := map[string]Counter{"main": {value: 0}}
m["main"].Increment()  // ERROR - can't get &m["main"]

// Fix: store pointers in map
m := map[string]*Counter{"main": &Counter{value: 0}}
m["main"].Increment()  // OK - m["main"] is already a pointer
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

## Next Steps

Continue to [Chapter 3: Control Flow](../03-control-flow/)

## References

- [Go Spec - Address operators](https://golang.org/ref/spec#Address_operators)
- [Go Tour - Variables](https://tour.golang.org/basics/8)
- [Effective Go - Variables](https://golang.org/doc/effective_go.html#variables)