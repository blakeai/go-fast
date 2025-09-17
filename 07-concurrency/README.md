# Chapter 7: Concurrency

## Overview

Go's concurrency model relies heavily on channels for communication and defer statements for cleanup. Combined with sync.WaitGroup, these primitives form the backbone of robust concurrent Go programs. This chapter covers channels, defer mechanics, and WaitGroups working together.

## Key Concepts

- **Channels**: CSP-based communication primitives
- **Defer statements**: Guaranteed cleanup with LIFO execution
- **WaitGroups**: Goroutine coordination and synchronization
- **Select statements**: Non-blocking channel operations
- **Buffered vs unbuffered channels**

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