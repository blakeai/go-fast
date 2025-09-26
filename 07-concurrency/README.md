# Chapter 7: Concurrency

## Overview

Go's concurrency model relies heavily on channels for communication and defer statements for cleanup. Combined with sync.WaitGroup, these primitives form the backbone of robust concurrent Go programs. This chapter covers channels, defer mechanics, and WaitGroups working together.

## Key Concepts

- **Channels**: CSP-based communication primitives
- **Defer statements**: Guaranteed cleanup with LIFO execution
- **WaitGroups**: Goroutine coordination and synchronization
- **Select statements**: Non-blocking channel operations
- **Buffered vs unbuffered channels**

## The Channel Arrow Operator `<-`

Before diving into channel examples, let's understand Go's **channel arrow operator** `<-` - a special two-character symbol made from "less than" and "hyphen": `<` + `-`

### The Arrow Shows Data Flow Direction

```go
ch <- 42    // Arrow points INTO channel (sending)
x := <-ch   // Arrow points FROM channel (receiving)
```

Think of it like a pipe where the arrow shows which way data flows:

```go
// Sending: data flows INTO the channel
ch <- 42    // Read as: "send 42 into ch"
            // Arrow points from value (42) to channel (ch)

// Receiving: data flows OUT OF the channel
x := <-ch   // Read as: "receive from ch into x"
            // Arrow points from channel (ch) outward
```

### Visual Analogy

```
SENDING:     value --> channel
             42    ->  ch
             ch <- 42

RECEIVING:   channel --> variable
             ch      ->  x
             x := <-ch
```

### Channel Direction in Function Signatures

The arrow also indicates channel direction in types:

```go
func send(ch chan<- int) {    // Send-only channel
    ch <- 42                   // Can only send
}

func receive(ch <-chan int) { // Receive-only channel
    x := <-ch                  // Can only receive
}

func both(ch chan int) {       // Bidirectional channel
    ch <- 42                   // Can send
    x := <-ch                  // Can receive
}
```

### Java Comparison

If you're coming from Java, think of:
- `ch <- 42` as similar to `queue.put(42)` or `blockingQueue.offer(42)`
- `x := <-ch` as similar to `x = queue.take()` or `x = blockingQueue.poll()`

The `<-` is just Go's syntax for these operations - it's not a "backwards" arrow, but rather shows the direction data moves relative to the channel!

## Examples

### Channels: The Communication Highway
See [`channels.go`](./channels.go) for implementation.

Channels implement Communicating Sequential Processes (CSP) and serve as Go's primary synchronization primitive:

```go
// Unbuffered (synchronous) channel
ch := make(chan int)

// Buffered channel with capacity
buffered := make(chan string, 3)

// Select for non-blocking operations
select {
case msg := <-ch1:
    // handle message
case <-time.After(1*time.Second):
    // timeout case
default:
    // non-blocking fallback
}
```

Key mechanics:
- `make(chan T)` creates unbuffered (synchronous) channels
- `make(chan T, n)` creates buffered channels with capacity n
- Unbuffered channels block until both sender and receiver are ready
- Buffered channels only block when full (send) or empty (receive)

### Defer: Guaranteed Cleanup
See [`defer.go`](./defer.go) for implementation.

The defer statement ensures code executes when a function returns, regardless of how it exits:

```go
func example() {
    defer fmt.Println("first")
    defer fmt.Println("second") 
    defer fmt.Println("third")
    fmt.Println("work")
}
// Output: work, third, second, first
```

**Critical detail:** defer captures argument values at registration time, not execution time.

### WaitGroups: Coordination Made Simple
See [`waitgroups.go`](./waitgroups.go) for implementation.

sync.WaitGroup coordinates goroutine completion:

```go
var wg sync.WaitGroup
for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done() // guaranteed execution
        // do work
    }(i)
}
wg.Wait()
```

### Why Defer Done()?
See [`defer-waitgroup.go`](./defer-waitgroup.go) for implementation.

`defer wg.Done()` **registers** the cleanup but doesn't **execute** until the function returns:

```go
go func() {
    defer wg.Done() // registered, not executed
    fmt.Println("doing work")
    // work happens here
    // wg.Done() executes when function returns
}()
```

This pattern ensures Done() is called even if the goroutine panics.

### Nested Defer Behavior
See [`nested-defer.go`](./nested-defer.go) for implementation.

Defer statements can be nested within anonymous functions:

```go
func complexExample() {
    defer fmt.Println("A")
    
    defer func() {
        defer fmt.Println("B")
        defer fmt.Println("C")
    }()
    
    defer fmt.Println("D")
    fmt.Println("work")
}
// Output: work, D, C, B, A
```

### Practical Concurrency Patterns
See [`patterns.go`](./patterns.go) for implementation.

```go
// Worker pool pattern
func workerPool(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        results <- job * 2
    }
}

// Fan-out, fan-in pattern
func fanOut(input <-chan int, workers int) []<-chan int {
    outputs := make([]<-chan int, workers)
    for i := 0; i < workers; i++ {
        output := make(chan int)
        outputs[i] = output
        go func() {
            defer close(output)
            for n := range input {
                output <- process(n)
            }
        }()
    }
    return outputs
}
```

## Running the Code

```bash
go run *.go
go test ./...
```

## Java Developer Notes

- **Channels vs BlockingQueue**: Channels transfer ownership rather than sharing mutable state
- **Defer vs try-finally**: Defer guarantees cleanup even on panic, similar to finally blocks
- **WaitGroup vs CountDownLatch**: Similar concept but WaitGroup can be incremented dynamically
- **Select vs Thread.interrupt()**: Select provides non-blocking operations without polling
- **Goroutines vs Threads**: Much lighter weight - can have thousands of goroutines

Key philosophy difference: "Don't communicate by sharing memory; share memory by communicating."

## Next Steps

Continue to [Chapter 8: Error Handling](../08-error-handling/)

## References

- [Go Concurrency Patterns](https://talks.golang.org/2012/concurrency.slide)
- [Effective Go - Concurrency](https://golang.org/doc/effective_go.html#concurrency)
- [Go Blog - Share Memory by Communicating](https://blog.golang.org/codelab-share)