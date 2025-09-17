# Chapter 1: Go Basics

## Overview

Go prioritizes simplicity and explicit behavior. Key differences from other languages: no semicolons required, package-based organization instead of classes, and Automatic Semicolon Insertion (ASI) that affects brace placement. Understanding these fundamentals prevents common syntax errors.

## Key Concepts

- **Package-based organization** - every file starts with `package declaration`
- **Automatic Semicolon Insertion (ASI)** - why brace placement matters
- **Zero values** - uninitialized variables get meaningful defaults
- **Exported vs unexported** - capitalization determines visibility
- **Import declarations** and unused import errors

## Examples

### Basic Program Structure
See [`hello.go`](./hello.go) for implementation.

```go
package main  // executable packages must be named "main"

import (
    "fmt"     // standard library package
    "strings" // another standard library package
)

// main() is the entry point for executable programs
func main() {
    message := "Hello, Go!"
    fmt.Println(strings.ToUpper(message))
}
```

### Package Declaration Rules
```go
// Executable program
package main
func main() { /* entry point */ }

// Library package (can be any name)
package calculator
func Add(a, b int) int { return a + b }  // Exported (capitalized)
func subtract(a, b int) int { return a - b }  // unexported (lowercase)
```

### ASI and Brace Placement
```go
// WRONG - ASI inserts semicolon after main()
func main()
{  // This becomes: func main(); {
    fmt.Println("This won't work")
}

// CORRECT - opening brace on same line
func main() {
    fmt.Println("This works")
}

// Same rule applies to all control structures
if condition {  // brace must be here, not next line
    // code
}
```

### Import Patterns
```go
// Single import
import "fmt"

// Multiple imports (preferred style)
import (
    "fmt"
    "strings"
    "net/http"
)

// Import aliases
import (
    "fmt"
    str "strings"  // alias 'strings' as 'str'
    . "math"       // dot import - use functions without qualifier
    _ "image/png"  // blank import - only for side effects
)

// Usage with aliases
func example() {
    fmt.Println("Regular import")
    str.ToUpper("aliased import")  // using alias
    Pi                             // dot import - no math.Pi needed
}
```

### Zero Values
```go
var i int        // 0
var f float64    // 0.0
var b bool       // false
var s string     // ""
var p *int       // nil
var slice []int  // nil (but len=0, cap=0)
var m map[string]int  // nil
var ch chan int  // nil

// Zero values make structs immediately usable
type Counter struct {
    value int
    name  string
}

var counter Counter  // {value: 0, name: ""}
// counter is immediately usable - no initialization required
```

### Basic Data Types
```go
// Numeric types
var i8 int8 = 127
var i16 int16 = 32767
var i32 int32 = 2147483647
var i64 int64 = 9223372036854775807
var i int = 42        // platform-dependent size (32 or 64 bit)

var u8 uint8 = 255    // also known as 'byte'
var u16 uint16 = 65535
var u32 uint32 = 4294967295
var u64 uint64 = 18446744073709551615
var u uint = 42       // platform-dependent size

var f32 float32 = 3.14159
var f64 float64 = 3.141592653589793  // default float type

var c64 complex64 = 3 + 4i
var c128 complex128 = complex(3.0, 4.0)

// String and boolean
var str string = "Hello, 世界"  // UTF-8 by default
var flag bool = true

// Byte and rune aliases
var b byte = 65        // uint8
var r rune = '世'       // int32 - represents Unicode code point
```

### Visibility Rules
```go
package mypackage

// Exported (public) - starts with capital letter
type User struct {
    Name  string  // exported field
    Email string  // exported field
    age   int     // unexported field - only accessible within package
}

func NewUser(name, email string) *User {  // exported constructor
    return &User{Name: name, Email: email}
}

func (u *User) GetAge() int {  // exported method
    return u.age
}

func (u *User) setAge(age int) {  // unexported method
    u.age = age
}

// From another package:
// user := mypackage.NewUser("Alice", "alice@example.com")  // OK
// user.Name = "Bob"     // OK - Name is exported
// user.age = 30         // ERROR - age is unexported
// user.setAge(30)       // ERROR - setAge is unexported
```

### Unused Imports and Variables
```go
package main

import (
    "fmt"
    "strings"  // ERROR if not used - Go enforces this
)

func main() {
    message := "hello"
    unused := "world"  // ERROR if not used
    
    fmt.Println(message)
    // strings package and unused variable will cause compilation errors
}

// Workarounds for temporary unused items:
func development() {
    unused := "temporary"
    _ = unused  // blank identifier prevents unused variable error
    
    // For unused imports during development:
    // _ = strings.ToUpper  // prevents unused import error
}
```

### Multiple Variable Declarations
```go
// Group declaration
var (
    name    string = "Alice"
    age     int    = 30
    active  bool   = true
)

// Multiple assignment
var x, y, z int = 1, 2, 3
var a, b = "hello", 42  // mixed types with inference

// Multiple return values (common pattern)
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

result, err := divide(10, 2)
if err != nil {
    fmt.Printf("Error: %v\n", err)
} else {
    fmt.Printf("Result: %f\n", result)
}
```

## Running the Code

```bash
go run *.go
go test ./...
```

## Java Developer Notes

- No semicolons required (ASI handles them)
- Package declaration replaces Java's `package` + class structure
- `main()` function replaces `public static void main(String[] args)`
- No `public`/`private` keywords - capitalization determines visibility
- Unused imports/variables are compilation errors (stricter than Java)
- Zero values replace Java's default initialization patterns
- No classes - packages contain functions, types, and variables directly

## Next Steps

Continue to [Chapter 2: Variables](../02-variables/)

## References

- [Go Tour - Packages](https://go.dev/tour/basics/1)
- [Effective Go - Names](https://go.dev/doc/effective_go#names)
- [Go Spec - Lexical elements](https://go.dev/ref/spec#Lexical_elements)