package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker pool pattern
func workerPoolExample() {
	fmt.Println("=== Worker Pool Pattern ===")

	jobs := make(chan int, 100)
	results := make(chan int, 100)
	var wg sync.WaitGroup

	// Start workers
	numWorkers := 3
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go workerPool(w, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for all workers to complete
	wg.Wait()
	close(results)

	// Collect results
	fmt.Println("Results:")
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}

func workerPool(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(100 * time.Millisecond)
		results <- job * 2
	}
}

// Fan-out, fan-in pattern
func fanOutFanInExample() {
	fmt.Println("\n=== Fan-Out, Fan-In Pattern ===")

	// Input channel
	input := make(chan int)

	// Fan-out to multiple workers
	outputs := fanOut(input, 3)

	// Fan-in to single output
	output := fanIn(outputs...)

	// Send input
	go func() {
		defer close(input)
		for i := 1; i <= 6; i++ {
			input <- i
		}
	}()

	// Read results
	for result := range output {
		fmt.Printf("Final result: %d\n", result)
	}
}

func fanOut(input <-chan int, workers int) []<-chan int {
	outputs := make([]<-chan int, workers)
	for i := 0; i < workers; i++ {
		output := make(chan int)
		outputs[i] = output
		go func(workerID int) {
			defer close(output)
			for n := range input {
				fmt.Printf("Worker %d processing %d\n", workerID, n)
				time.Sleep(50 * time.Millisecond)
				output <- n * n // Square the number
			}
		}(i)
	}
	return outputs
}

func fanIn(inputs ...<-chan int) <-chan int {
	output := make(chan int)
	var wg sync.WaitGroup

	for _, input := range inputs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for value := range ch {
				output <- value
			}
		}(input)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

// Pipeline pattern
func pipelineExample() {
	fmt.Println("\n=== Pipeline Pattern ===")

	// Stage 1: Generate numbers
	numbers := generate(1, 2, 3, 4, 5)

	// Stage 2: Square numbers
	squares := square(numbers)

	// Stage 3: Filter odd numbers
	odds := filterOdd(squares)

	// Consume results
	for result := range odds {
		fmt.Printf("Pipeline result: %d\n", result)
	}
}

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func filterOdd(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n%2 == 1 {
				out <- n
			}
		}
	}()
	return out
}

// Broadcast pattern
func broadcastExample() {
	fmt.Println("\n=== Broadcast Pattern ===")

	input := make(chan string)

	// Create multiple subscribers
	sub1 := subscribe("Subscriber-1", input)
	sub2 := subscribe("Subscriber-2", input)
	sub3 := subscribe("Subscriber-3", input)

	// Send messages
	go func() {
		defer close(input)
		messages := []string{"Hello", "World", "Broadcast", "Pattern"}
		for _, msg := range messages {
			input <- msg
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Wait for all subscribers to finish
	var wg sync.WaitGroup
	wg.Add(3)

	go func() { defer wg.Done(); consume(sub1) }()
	go func() { defer wg.Done(); consume(sub2) }()
	go func() { defer wg.Done(); consume(sub3) }()

	wg.Wait()
	fmt.Println("All subscribers finished")
}

func subscribe(name string, input <-chan string) <-chan string {
	output := make(chan string)
	go func() {
		defer close(output)
		for msg := range input {
			output <- fmt.Sprintf("%s received: %s", name, msg)
		}
	}()
	return output
}

func consume(ch <-chan string) {
	for msg := range ch {
		fmt.Println(msg)
	}
}

// Cancellation pattern
func cancellationExample() {
	fmt.Println("\n=== Cancellation Pattern ===")

	done := make(chan bool)

	// Start a long-running goroutine
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Worker cancelled")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	// Let it work for a while, then cancel
	time.Sleep(800 * time.Millisecond)
	done <- true

	// Give it time to clean up
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Cancellation complete")
}

// Rate limiting pattern
func rateLimitingExample() {
	fmt.Println("\n=== Rate Limiting Pattern ===")

	// Limit to 2 operations per second
	limiter := time.Tick(500 * time.Millisecond)

	requests := []string{"req1", "req2", "req3", "req4", "req5"}

	for _, req := range requests {
		<-limiter // Wait for rate limiter
		fmt.Printf("Processing %s at %s\n", req, time.Now().Format("15:04:05.000"))
	}
}

func patternsExample() {
	workerPoolExample()
	fanOutFanInExample()
	pipelineExample()
	broadcastExample()
	cancellationExample()
	rateLimitingExample()
}
