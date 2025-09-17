package main

import (
	"fmt"
	"time"
)

func traditionalForLoop() {
	fmt.Println("=== Traditional For Loop ===")

	// Classic three-part for loop
	for i := 0; i < 5; i++ {
		fmt.Printf("Count: %d\n", i)
	}

	// Loop with different increment
	fmt.Println("\nCounting by 2s:")
	for i := 0; i < 10; i += 2 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Countdown
	fmt.Println("\nCountdown:")
	for i := 5; i > 0; i-- {
		fmt.Printf("%d... ", i)
	}
	fmt.Println("Blast off!")
}

func whileStyleLoop() {
	fmt.Println("\n=== While-Style Loop ===")

	// No while keyword - use for with condition only
	count := 0
	for count < 3 {
		fmt.Printf("Iteration %d\n", count)
		count++
	}

	// Example with break
	fmt.Println("\nLoop with break:")
	counter := 0
	for counter < 10 {
		if counter == 4 {
			fmt.Println("Breaking at 4")
			break
		}
		fmt.Printf("Counter: %d\n", counter)
		counter++
	}

	// Example with continue
	fmt.Println("\nLoop with continue (skip even numbers):")
	for i := 0; i < 8; i++ {
		if i%2 == 0 {
			continue // skip even numbers
		}
		fmt.Printf("Odd number: %d\n", i)
	}
}

func infiniteLoop() {
	fmt.Println("\n=== Infinite Loop (with break) ===")

	// Infinite loop - be careful!
	iterations := 0
	for {
		iterations++
		fmt.Printf("Iteration %d\n", iterations)

		if iterations >= 3 {
			fmt.Println("Breaking out of infinite loop")
			break
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func rangeOverSlice() {
	fmt.Println("\n=== Range Over Slice ===")

	fruits := []string{"apple", "banana", "cherry", "date"}

	// Range with both index and value
	fmt.Println("With index and value:")
	for index, value := range fruits {
		fmt.Printf("%d: %s\n", index, value)
	}

	// Range with index only
	fmt.Println("\nWith index only:")
	for i := range fruits {
		fmt.Printf("Index %d\n", i)
	}

	// Range with value only (use blank identifier for index)
	fmt.Println("\nWith value only:")
	for _, value := range fruits {
		fmt.Printf("Fruit: %s\n", value)
	}
}

func rangeOverMap() {
	fmt.Println("\n=== Range Over Map ===")

	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	}

	// Range over map - order not guaranteed
	fmt.Println("Map contents (order may vary):")
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// Keys only
	fmt.Println("\nNames only:")
	for name := range ages {
		fmt.Printf("Name: %s\n", name)
	}
}

func rangeOverString() {
	fmt.Println("\n=== Range Over String ===")

	text := "Hello 世界"

	// Range over string yields runes (Unicode code points), not bytes
	fmt.Println("Characters (runes):")
	for i, r := range text {
		fmt.Printf("Position %d: %c (Unicode: %U)\n", i, r, r)
	}

	// Note: position jumps because Unicode characters can be multiple bytes
	fmt.Printf("\nString length in bytes: %d\n", len(text))
	fmt.Printf("String length in runes: %d\n", len([]rune(text)))
}

func rangeOverChannel() {
	fmt.Println("\n=== Range Over Channel ===")

	// Create a channel
	ch := make(chan int)

	// Send values in a goroutine
	go func() {
		defer close(ch) // Important: close the channel when done
		for i := 1; i <= 5; i++ {
			ch <- i * i // Send squares
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// Range over channel - blocks until channel is closed
	fmt.Println("Receiving from channel:")
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
	fmt.Println("Channel closed, loop ended")
}

func nestedLoops() {
	fmt.Println("\n=== Nested Loops ===")

	// Simple nested loops
	fmt.Println("Multiplication table (3x3):")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d×%d=%d  ", i, j, i*j)
		}
		fmt.Println()
	}
}

func labeledBreakContinue() {
	fmt.Println("\n=== Labeled Break and Continue ===")

	// Labeled break - break out of outer loop
	fmt.Println("Labeled break example:")
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Printf("Breaking at (%d,%d)\n", i, j)
				break outer // breaks out of both loops
			}
			fmt.Printf("(%d,%d) ", i, j)
		}
		fmt.Println()
	}
	fmt.Println("Exited both loops")

	// Labeled continue - continue outer loop
	fmt.Println("\nLabeled continue example:")
outerContinue:
	for i := 0; i < 3; i++ {
		fmt.Printf("Row %d: ", i)
		for j := 0; j < 4; j++ {
			if j == 2 {
				fmt.Printf("[skipping rest] ")
				continue outerContinue // skip to next iteration of outer loop
			}
			fmt.Printf("%d ", j)
		}
		fmt.Println() // This won't be reached when continue is triggered
	}
}

func loopPerformanceTips() {
	fmt.Println("\n=== Loop Performance Tips ===")

	// Cache slice length if not changing
	numbers := make([]int, 1000)
	for i := 0; i < len(numbers); i++ {
		numbers[i] = i
	}

	// Use range when you need both index and value
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Printf("Sum using range: %d\n", sum)

	// Use traditional for when you only need index
	sum2 := 0
	for i := 0; i < len(numbers); i++ {
		sum2 += numbers[i]
	}
	fmt.Printf("Sum using traditional for: %d\n", sum2)
}

func loopsExample() {
	traditionalForLoop()
	whileStyleLoop()
	infiniteLoop()
	rangeOverSlice()
	rangeOverMap()
	rangeOverString()
	rangeOverChannel()
	nestedLoops()
	labeledBreakContinue()
	loopPerformanceTips()

	fmt.Println("\n=== Key Loop Takeaways ===")
	fmt.Println("✅ for is the only loop construct in Go")
	fmt.Println("✅ range gives you index and value (or just one)")
	fmt.Println("✅ range over string yields runes, not bytes")
	fmt.Println("✅ range over channel blocks until closed")
	fmt.Println("✅ Use labeled break/continue for nested loops")
	fmt.Println("✅ for {} creates an infinite loop")
}
