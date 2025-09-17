# Chapter 3: Control Flow

## Overview

Go's control flow is deliberately minimal - `for` is the only loop construct, `switch` doesn't fall through by default, and brace placement matters due to Automatic Semicolon Insertion (ASI). Understanding these design decisions prevents common mistakes when transitioning from other languages.

## Key Concepts

- **Single loop construct** - `for` handles all iteration patterns
- **ASI and brace placement** - opening braces must be on same line
- **Switch statements** don't fall through (no `break` needed)
- **Short variable declarations** in `if` statements for scoped variables
- **Label breaks** and `goto` for complex control flow

## Examples

### For Loops - The Universal Iterator
See [`loops.go`](./loops.go) for implementation.

```go
// Traditional for loop
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// While-style loop (no while keyword exists)
for condition {
    // loop body
    if someCheck {
        break
    }
}

// Infinite loop
for {
    // runs forever unless break/return
}

// Range over collections
slice := []string{"a", "b", "c"}
for index, value := range slice {
    fmt.Printf("%d: %s\n", index, value)
}

// Range with index only
for i := range slice {
    fmt.Println("Index:", i)
}

// Range with value only (note the blank identifier)
for _, value := range slice {
    fmt.Println("Value:", value)
}
```

### If Statements with Short Declarations
See [`conditionals.go`](./conditionals.go) for implementation.

```go
// Standard if
if age >= 18 {
    fmt.Println("Adult")
}

// If with short variable declaration - variable scoped to if block
if user := getUser(id); user != nil {
    fmt.Printf("Found user: %s\n", user.Name)
    // user is accessible here
} else {
    fmt.Println("User not found")
    // user is also accessible in else block
}
// user is NOT accessible here - out of scope

// Common pattern for error checking
if err := doSomething(); err != nil {
    return fmt.Errorf("operation failed: %w", err)
}
```

### ASI and Brace Placement
```go
// WRONG - ASI inserts semicolon after condition
if condition
{  // This becomes: if condition ; {
    // This is now an unreachable block, not part of if
}

// CORRECT - opening brace on same line
if condition {
    // proper if body
}

// Same rule applies to for loops and functions
for i := 0; i < 10; i++ {  // brace must be here
    fmt.Println(i)
}
```

### Switch Statements - No Fall-Through
See [`switch.go`](./switch.go) for implementation.

```go
// Basic switch - no break needed
switch day {
case "monday":
    fmt.Println("Start of work week")
case "friday":
    fmt.Println("TGIF")
case "saturday", "sunday":  // multiple values in one case
    fmt.Println("Weekend")
default:
    fmt.Println("Midweek")
}

// Switch on type
func processValue(v interface{}) {
    switch val := v.(type) {
    case int:
        fmt.Printf("Integer: %d\n", val)
    case string:
        fmt.Printf("String: %s\n", val)
    case []int:
        fmt.Printf("Slice length: %d\n", len(val))
    default:
        fmt.Printf("Unknown type: %T\n", val)
    }
}

// Switch without expression (replaces if-else chains)
switch {
case age < 13:
    return "child"
case age < 20:
    return "teenager"  
case age < 65:
    return "adult"
default:
    return "senior"
}

// Explicit fall-through when needed
switch grade {
case 'A':
    fmt.Println("Excellent")
    fallthrough  // explicitly continue to next case
case 'B':
    fmt.Println("Good work")
case 'C':
    fmt.Println("Passing")
}
```

### Labeled Breaks and Continue
```go
// Breaking out of nested loops
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            break outer  // breaks out of both loops
        }
        fmt.Printf("(%d,%d) ", i, j)
    }
}

// Continue with labels
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if j == 1 {
            continue outer  // skip to next iteration of outer loop
        }
        fmt.Printf("(%d,%d) ", i, j)
    }
}
```

### Range Behavior Differences
```go
// Slice/Array - index and value
for i, v := range []int{10, 20, 30} {
    // i: 0,1,2  v: 10,20,30
}

// Map - key and value (order not guaranteed)
for key, value := range map[string]int{"a": 1, "b": 2} {
    // key: "a","b"  value: 1,2
}

// String - index and rune (not byte)
for i, r := range "Hello 世界" {
    // i: byte positions  r: Unicode code points
    fmt.Printf("%d: %c\n", i, r)
}

// Channel - value only (blocks until closed)
ch := make(chan int)
go func() {
    ch <- 1
    ch <- 2
    close(ch)
}()

for value := range ch {
    // receives 1, then 2, then exits when channel closed
    fmt.Println(value)
}
```

## Running the Code

```bash
go run *.go
go test ./...
```

## Java Developer Notes

- No `while` or `do-while` - use `for` instead
- `switch` doesn't fall through by default (opposite of Java)
- Enhanced for-loop equivalent is `range`
- No parentheses required around conditions: `if x > 0` not `if (x > 0)`
- Opening braces MUST be on same line due to ASI
- Short variable declarations in `if` statements provide scoped variables
- Labeled breaks work like Java but are used less frequently

## Next Steps

Continue to [Chapter 4: Functions](../04-functions/)

## References

- [Go Tour - Flow control](https://go.dev/tour/flowcontrol/1)
- [Effective Go - Control structures](https://go.dev/doc/effective_go#control-structures)
- [Go Spec - For statements](https://go.dev/ref/spec#For_statements)