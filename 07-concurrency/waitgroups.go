package main

import (
	"fmt"
	"sync"
	"time"
)

func basicWaitGroup() {
	fmt.Println("=== Basic WaitGroup ===")

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1) // Increment counter
		go func(id int) {
			defer wg.Done() // Decrement counter when done
			fmt.Printf("Worker %d starting\n", id)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Worker %d done\n", id)
		}(i)
	}

	wg.Wait() // Block until counter reaches zero
	fmt.Println("All workers completed")
}

func waitGroupWithoutDefer() {
	fmt.Println("\n=== WaitGroup Without Defer (Dangerous) ===")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("Worker starting...")

		// Simulate potential panic or early return
		if true { // Could be some condition
			wg.Done() // Must remember to call Done()
			return
		}

		// More work...
		fmt.Println("Worker finishing...")
		wg.Done() // Easy to forget this!
	}()

	wg.Wait()
	fmt.Println("Worker completed")
}

func waitGroupWithDefer() {
	fmt.Println("\n=== WaitGroup With Defer (Safe) ===")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done() // Always called, no matter how function exits

		fmt.Println("Worker starting...")

		// Simulate potential panic or early return
		if true { // Could be some condition
			return // wg.Done() still called via defer
		}

		// More work...
		fmt.Println("Worker finishing...")
		// wg.Done() called automatically
	}()

	wg.Wait()
	fmt.Println("Worker completed")
}

func waitGroupCommonMistakes() {
	fmt.Println("\n=== Common WaitGroup Mistakes ===")

	var wg sync.WaitGroup

	// MISTAKE 1: Adding inside goroutine
	fmt.Println("Mistake 1: Race condition")
	for i := 0; i < 2; i++ {
		go func(id int) {
			//nolint:govet,staticcheck // Intentional bad example
			wg.Add(1) // WRONG: Race condition!
			defer wg.Done()
			fmt.Printf("Worker %d\n", id)
		}(i)
	}
	// wg.Wait() // This might return before all goroutines are added

	// CORRECT WAY: Add before launching goroutine
	fmt.Println("\nCorrect way:")
	for i := 0; i < 2; i++ {
		wg.Add(1) // CORRECT: Add before launching
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Worker %d\n", id)
		}(i)
	}
	wg.Wait()
}

func waitGroupWithError() {
	fmt.Println("\n=== WaitGroup with Error Handling ===")

	var wg sync.WaitGroup
	errors := make(chan error, 3)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Simulate work that might fail
			if id == 1 {
				errors <- fmt.Errorf("worker %d failed", id)
				return
			}

			fmt.Printf("Worker %d succeeded\n", id)
			errors <- nil
		}(i)
	}

	// Wait for all workers to complete
	wg.Wait()
	close(errors)

	// Check for errors
	fmt.Println("Checking errors:")
	for err := range errors {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func waitGroupWithContext() {
	fmt.Println("\n=== WaitGroup with Worker Pool ===")

	var wg sync.WaitGroup
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Start workers
	numWorkers := 3
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				fmt.Printf("Worker %d processing job %d\n", workerID, job)
				time.Sleep(50 * time.Millisecond)
				results <- job * 2
			}
		}(w)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for workers to complete
	wg.Wait()
	close(results)

	// Collect results
	fmt.Println("Results:")
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}

func waitgroupsExample() {
	basicWaitGroup()
	waitGroupWithoutDefer()
	waitGroupWithDefer()
	waitGroupCommonMistakes()
	waitGroupWithError()
	waitGroupWithContext()
}
