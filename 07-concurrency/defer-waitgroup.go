package main

import (
	"fmt"
	"sync"
	"time"
)

func whyDeferDone() {
	fmt.Println("=== Why Defer wg.Done()? ===")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done() // registered, not executed
		fmt.Println("doing work")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("work completed")
		// wg.Done() executes when function returns
	}()

	wg.Wait()
	fmt.Println("All work completed")
}

func deferDoneWithPanic() {
	fmt.Println("\n=== Defer Done() Prevents Deadlock on Panic ===")

	var wg sync.WaitGroup

	// Without defer - this would deadlock
	fmt.Println("Scenario: Worker panics")
	wg.Add(1)
	go func() {
		defer wg.Done() // Still called even if panic occurs
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic: %v\n", r)
			}
		}()

		fmt.Println("Worker starting...")
		panic("something went wrong!")

	}()

	wg.Wait()
	fmt.Println("Worker completed (despite panic)")
}

func deferDoneWithEarlyReturn() {
	fmt.Println("\n=== Defer Done() with Early Return ===")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done() // Guaranteed to execute

		fmt.Println("Checking conditions...")

		// Early return scenario
		if true { // Some condition
			fmt.Println("Early return due to condition")
			return // wg.Done() still called
		}

		fmt.Println("Normal completion")
		// wg.Done() would also be called here
	}()

	wg.Wait()
	fmt.Println("Worker completed")
}

func multipleWorkers() {
	fmt.Println("\n=== Multiple Workers with Defer Done() ===")

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done() // Each worker guarantees cleanup

			fmt.Printf("Worker %d starting\n", id)

			// Simulate different completion scenarios
			switch id {
			case 1:
				panic("worker 1 panics") // Still calls wg.Done()
			case 2:
				time.Sleep(50 * time.Millisecond)
				return // Early return, still calls wg.Done()
			default:
				time.Sleep(100 * time.Millisecond)
				fmt.Printf("Worker %d completed normally\n", id)
			}
		}(i)
	}

	// Recovery for panicking workers
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Main recovered: %v\n", r)
		}
	}()

	wg.Wait()
	fmt.Println("All workers completed")
}

func timingDemonstration() {
	fmt.Println("\n=== Timing: When wg.Done() Actually Executes ===")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("1. Function starts")
		defer func() {
			fmt.Println("4. Defer wg.Done() executes")
			wg.Done()
		}()

		fmt.Println("2. Doing work...")
		time.Sleep(50 * time.Millisecond)
		fmt.Println("3. Function about to return")
		// Defer executes here, after this line
	}()

	wg.Wait()
	fmt.Println("5. Main continues after wg.Wait()")
}

func comparisonWithoutDefer() {
	fmt.Println("\n=== Comparison: Without Defer (Error-Prone) ===")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("Worker starting...")

		// Manual wg.Done() - easy to forget or skip
		if false { // Some error condition
			// OOPS! Forgot wg.Done() here - would cause deadlock
			return
		}

		fmt.Println("Worker doing work...")
		time.Sleep(50 * time.Millisecond)

		// Another place to forget wg.Done()
		if false { // Another condition
			// OOPS! Forgot wg.Done() here too
			return
		}

		wg.Done() // Remember to call this
		fmt.Println("Worker completed")
	}()

	wg.Wait()
	fmt.Println("All workers completed")
}

func deferWaitgroupExample() {
	whyDeferDone()
	deferDoneWithPanic()
	deferDoneWithEarlyReturn()
	multipleWorkers()
	timingDemonstration()
	comparisonWithoutDefer()
}
