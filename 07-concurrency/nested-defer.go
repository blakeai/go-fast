package main

import "fmt"

func simpleNestedDefer() {
	fmt.Println("=== Simple Nested Defer ===")

	defer fmt.Println("A")

	defer func() {
		defer fmt.Println("B")
		defer fmt.Println("C")
	}()

	defer fmt.Println("D")
	fmt.Println("work")
}

func complexExample() {
	fmt.Println("\n=== Complex Nested Defer Example ===")

	defer fmt.Println("A")

	defer func() {
		defer fmt.Println("B")
		defer fmt.Println("C")
	}()

	defer fmt.Println("D")
	fmt.Println("work")
	// Expected output: work, D, C, B, A
}

func multiLevelNesting() {
	fmt.Println("\n=== Multi-Level Nesting ===")

	defer fmt.Println("Level 1 - A")

	defer func() {
		defer fmt.Println("Level 2 - A")

		defer func() {
			defer fmt.Println("Level 3 - A")
			defer fmt.Println("Level 3 - B")
		}()

		defer fmt.Println("Level 2 - B")
	}()

	defer fmt.Println("Level 1 - B")
	fmt.Println("main work")
}

func nestedDeferWithLogic() {
	fmt.Println("\n=== Nested Defer with Logic ===")

	defer fmt.Println("Outer defer 1")

	defer func() {
		fmt.Println("Inner function starts")
		defer fmt.Println("Inner defer 1")
		defer fmt.Println("Inner defer 2")
		fmt.Println("Inner function work")
		// Inner defers execute here when anonymous function returns
	}()

	defer fmt.Println("Outer defer 2")

	fmt.Println("Main function work")
	// All defers execute here when main function returns
}

func nestedDeferWithParameters() {
	fmt.Println("\n=== Nested Defer with Parameters ===")

	x := 1
	defer func(val int) {
		fmt.Printf("Outer defer: x = %d\n", val)
	}(x) // x captured as 1

	defer func() {
		defer func(val int) {
			fmt.Printf("Inner defer: x = %d\n", val)
		}(x) // x will be evaluated when inner anonymous function is called
		x = 100 // This affects the inner defer
	}()

	x = 10
	fmt.Printf("Main: x = %d\n", x)
}

func practicalExample() {
	fmt.Println("\n=== Practical Example: Resource Cleanup ===")

	defer fmt.Println("Final cleanup")

	defer func() {
		defer fmt.Println("Database connection closed")
		defer fmt.Println("Transaction rolled back")

		// Simulate cleanup work
		fmt.Println("Performing nested cleanup...")
	}()

	defer fmt.Println("File handles closed")

	fmt.Println("Doing main work...")
	// All cleanup happens in reverse order
}

func nestedDeferInLoop() {
	fmt.Println("\n=== Nested Defer in Loop ===")

	// Wrong way - accumulates defers
	defer func() {
		fmt.Println("Cleanup for wrong way:")
		for i := 0; i < 3; i++ {
			defer fmt.Printf("defer %d ", i) // All wait until this function ends
		}
	}()

	// Right way - each iteration has its own scope
	fmt.Println("Right way:")
	for i := 0; i < 3; i++ {
		func(i int) {
			defer fmt.Printf("immediate defer %d ", i)
		}(i)
	}
	fmt.Println("\nLoop completed")
}

func impossibleWithFlatDefer() {
	fmt.Println("\n=== Pattern Impossible with Flat Defer ===")

	// This specific output order can only be achieved with nested defer
	defer fmt.Println("1")

	defer func() {
		defer fmt.Println("2")
		fmt.Println("3") // This prints before the inner defer
		defer fmt.Println("4")
	}()

	defer fmt.Println("5")
	fmt.Println("6")
	// Output: 6, 5, 3, 4, 2, 1
	// The "3" between defers can't be achieved with flat defer structure
}

func nestedDeferExample() {
	simpleNestedDefer()
	complexExample()
	multiLevelNesting()
	nestedDeferWithLogic()
	nestedDeferWithParameters()
	practicalExample()
	nestedDeferInLoop()
	impossibleWithFlatDefer()
}
