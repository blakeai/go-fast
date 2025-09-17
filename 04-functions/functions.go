package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Simple function
func add(a, b int) int {
	return a + b
}

// Multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Named return values
func calculate(a, b int) (sum, product int) {
	sum = a + b
	product = a * b
	return // naked return
}

// Variadic function
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Function as parameter
func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// Function as return value
func getOperator(operation string) func(int, int) int {
	switch operation {
	case "add":
		return func(a, b int) int { return a + b }
	case "multiply":
		return func(a, b int) int { return a * b }
	default:
		return func(a, b int) int { return 0 }
	}
}

// Anonymous functions
func demonstrateAnonymousFunctions() {
	// Immediate invocation
	result := func(x, y int) int {
		return x * y
	}(5, 3)

	fmt.Printf("Anonymous function result: %d\n", result)

	// Stored in variable
	multiply := func(a, b int) int {
		return a * b
	}

	fmt.Printf("Stored anonymous function: %d\n", multiply(4, 6))
}

// Closure example
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Error handling patterns
func processData(data string) (int, error) {
	if data == "" {
		return 0, errors.New("empty data")
	}

	value, err := strconv.Atoi(data)
	if err != nil {
		return 0, fmt.Errorf("invalid data '%s': %w", data, err)
	}

	return value * 2, nil
}

func main() {
	fmt.Println("=== Basic Functions ===")
	fmt.Printf("add(3, 5) = %d\n", add(3, 5))

	result, err := divide(10, 3)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("divide(10, 3) = %.2f\n", result)
	}

	// Test division by zero
	_, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}

	fmt.Println("\n=== Named Returns ===")
	s, p := calculate(4, 5)
	fmt.Printf("calculate(4, 5) = sum: %d, product: %d\n", s, p)

	fmt.Println("\n=== Variadic Functions ===")
	fmt.Printf("sum(1, 2, 3, 4, 5) = %d\n", sum(1, 2, 3, 4, 5))

	numbers := []int{10, 20, 30}
	fmt.Printf("sum(10, 20, 30) = %d\n", sum(numbers...))

	fmt.Println("\n=== Functions as Parameters ===")
	fmt.Printf("applyOperation(8, 3, add) = %d\n", applyOperation(8, 3, add))

	fmt.Println("\n=== Functions as Return Values ===")
	addFunc := getOperator("add")
	multiplyFunc := getOperator("multiply")
	fmt.Printf("addFunc(7, 8) = %d\n", addFunc(7, 8))
	fmt.Printf("multiplyFunc(7, 8) = %d\n", multiplyFunc(7, 8))

	fmt.Println("\n=== Anonymous Functions ===")
	demonstrateAnonymousFunctions()

	fmt.Println("\n=== Closures ===")
	counter1 := createCounter()
	counter2 := createCounter()
	fmt.Printf("counter1: %d, %d, %d\n", counter1(), counter1(), counter1())
	fmt.Printf("counter2: %d, %d\n", counter2(), counter2())

	fmt.Println("\n=== Error Handling ===")
	data := []string{"123", "abc", "", "456"}
	for _, d := range data {
		result, err := processData(d)
		if err != nil {
			fmt.Printf("processData('%s') error: %v\n", d, err)
		} else {
			fmt.Printf("processData('%s') = %d\n", d, result)
		}
	}
}
