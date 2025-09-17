package main

import (
	"fmt"
	"time"
)

func basicChannels() {
	fmt.Println("=== Basic Channels ===")

	// Unbuffered (synchronous) channel
	ch := make(chan int)

	// Start a goroutine to send data
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch <- 42
		fmt.Println("Sent 42 to channel")
	}()

	// Receive from channel (blocks until data available)
	value := <-ch
	fmt.Printf("Received: %d\n", value)
}

func bufferedChannels() {
	fmt.Println("\n=== Buffered Channels ===")

	// Buffered channel with capacity 3
	buffered := make(chan string, 3)

	// Can send 3 values without blocking
	buffered <- "first"
	buffered <- "second"
	buffered <- "third"
	fmt.Println("Sent 3 values to buffered channel")

	// Receive the values
	fmt.Printf("Received: %s\n", <-buffered)
	fmt.Printf("Received: %s\n", <-buffered)
	fmt.Printf("Received: %s\n", <-buffered)
}

func selectStatement() {
	fmt.Println("\n=== Select Statement ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// Send to channels in goroutines
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "from ch2"
	}()

	// Select waits for first available channel
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Received: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received: %s\n", msg2)
		case <-time.After(200 * time.Millisecond):
			fmt.Println("Timeout!")
		}
	}
}

func selectWithDefault() {
	fmt.Println("\n=== Select with Default ===")

	ch := make(chan int)

	// Non-blocking send/receive with default
	select {
	case ch <- 1:
		fmt.Println("Sent to channel")
	default:
		fmt.Println("Channel not ready for send")
	}

	select {
	case value := <-ch:
		fmt.Printf("Received: %d\n", value)
	default:
		fmt.Println("No data available")
	}
}

func channelDirections() {
	fmt.Println("\n=== Channel Directions ===")

	// Function that only sends
	send := func(ch chan<- string) {
		ch <- "hello"
	}

	// Function that only receives
	receive := func(ch <-chan string) string {
		return <-ch
	}

	ch := make(chan string, 1)
	send(ch)
	msg := receive(ch)
	fmt.Printf("Message: %s\n", msg)
}

func channelsExample() {
	basicChannels()
	bufferedChannels()
	selectStatement()
	selectWithDefault()
	channelDirections()
}
