package main

import (
	"fmt"
	"os"
)

func basicDefer() {
	fmt.Println("=== Basic Defer ===")

	defer fmt.Println("This runs last")
	fmt.Println("This runs first")
	fmt.Println("This runs second")
}

func deferOrder() {
	fmt.Println("\n=== Defer Order (LIFO) ===")

	defer fmt.Println("first defer")
	defer fmt.Println("second defer")
	defer fmt.Println("third defer")
	fmt.Println("normal execution")
	// Output: normal execution, third defer, second defer, first defer
}

func deferWithArguments() {
	fmt.Println("\n=== Defer with Arguments ===")

	x := 10
	defer fmt.Printf("Deferred: x = %d\n", x) // x captured as 10

	x = 20
	fmt.Printf("Current: x = %d\n", x)
	// Deferred function will print x = 10, not 20
}

func deferWithPointers() {
	fmt.Println("\n=== Defer with Pointers ===")

	x := 10
	defer func() {
		fmt.Printf("Deferred (closure): x = %d\n", x) // x evaluated when defer executes
	}()

	defer fmt.Printf("Deferred (direct): x = %d\n", x) // x captured immediately

	x = 20
	fmt.Printf("Current: x = %d\n", x)
}

func deferForCleanup() {
	fmt.Println("\n=== Defer for Cleanup ===")

	file, err := os.Create("temp.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer func() {
		err := file.Close()
		if err != nil {
			return
		}
		err = os.Remove("temp.txt")
		if err != nil {
			return
		} // Clean up
		fmt.Println("File cleaned up")
	}()

	// Do work with file
	_, err = file.WriteString("Hello, defer!")
	if err != nil {
		return
	}
	fmt.Println("Work with file completed")
}

func deferWithPanic() {
	fmt.Println("\n=== Defer with Panic Recovery ===")

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	defer fmt.Println("This defer runs even with panic")

	fmt.Println("About to panic...")
	panic("something went wrong")

}

func deferInLoop() {
	fmt.Println("\n=== Defer in Loop (Common Mistake) ===")

	fmt.Println("Wrong way - defers accumulate:")
	func() {
		for i := 0; i < 3; i++ {
			defer fmt.Printf("defer %d ", i) // All defers wait until function ends
		}
		fmt.Println("loop done")
	}()

	fmt.Println("\nRight way - use anonymous function:")
	for i := 0; i < 3; i++ {
		func(i int) {
			defer fmt.Printf("defer %d ", i) // Each defer executes immediately
		}(i)
	}
	fmt.Println("loop done")
}

func example() {
	defer fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("third")
	fmt.Println("work")
}

func deferExample() {
	basicDefer()
	deferOrder()
	deferWithArguments()
	deferWithPointers()
	deferForCleanup()
	deferWithPanic()
	deferInLoop()

	fmt.Println("\n=== Example from README ===")
	example()
}
