package main

import (
	"fmt"
	"log"
	"strings"

	"go-fast/09-packages/calculator"
)

func main() {
	packageDemo()
	calculatorDemo()
}

func packageDemo() {
	fmt.Println("=== Package Usage Demo ===")

	// Using exported functions directly
	sum := calculator.Add(10, 5)
	fmt.Printf("calculator.Add(10, 5) = %d\n", sum)

	difference := calculator.Subtract(10, 5)
	fmt.Printf("calculator.Subtract(10, 5) = %d\n", difference)

	product := calculator.Multiply(10, 5)
	fmt.Printf("calculator.Multiply(10, 5) = %d\n", product)

	// Using exported function with error handling
	quotient, err := calculator.Divide(10.0, 5.0)
	if err != nil {
		log.Printf("Division error: %v", err)
	} else {
		fmt.Printf("calculator.Divide(10.0, 5.0) = %.2f\n", quotient)
	}

	// Test error case
	_, err = calculator.Divide(10.0, 0.0)
	if err != nil {
		fmt.Printf("Expected error for division by zero: %v\n", err)
	}

	// Using power function
	powerResult := calculator.Power(2, 8)
	fmt.Printf("calculator.Power(2, 8) = %d\n", powerResult)

	// Note: cannot access unexported functions from outside the package
	// calculator.multiply(2, 3)  // This would cause a compilation error
}

func calculatorDemo() {
	fmt.Println("\n=== Calculator Type Demo ===")

	// Create a new calculator instance
	calc := calculator.NewCalculator()
	fmt.Println("Created new calculator")

	// Perform operations - these will be recorded in history
	calc.Add(15, 25)
	calc.Subtract(50, 10)
	calc.Multiply(6, 7)
	calc.Add(100, 200)

	// Display calculator with history
	fmt.Printf("\n%s\n", calc.String())

	// Get and display history
	history := calc.GetHistory()
	fmt.Printf("History contains %d operations:\n", len(history))
	for i, op := range history {
		fmt.Printf("  %d. Operation{Type: %s, A: %d, B: %d, Result: %d}\n",
			i+1, op.Type, op.A, op.B, op.Result)
	}

	// Clear history and show result
	fmt.Println("\nClearing calculator history...")
	calc.ClearHistory()
	fmt.Printf("%s\n", calc.String())
}

func importPatternsDemo() {
	fmt.Println("\n=== Import Patterns Demo ===")

	// Standard library imports
	fmt.Println("Using fmt package for formatted output")

	// Using strings package functions
	text := "hello go packages"
	upper := strings.ToUpper(text)
	title := toTitle(text) // Custom function to replace deprecated strings.Title

	fmt.Printf("Original: %s\n", text)
	fmt.Printf("Upper: %s\n", upper)
	fmt.Printf("Title: %s\n", title)

	// Local package import
	result := calculator.Add(42, 8)
	fmt.Printf("Using local calculator package: %d\n", result)

	// Note: Other import patterns (aliases, dot imports, blank imports)
	// are demonstrated in the README examples but not executed here
	// to keep this demo focused on basic package usage
}

func visibilityDemo() {
	fmt.Println("\n=== Visibility Rules Demo ===")

	fmt.Println("✓ Can access exported functions:")
	fmt.Printf("  calculator.Add(1, 2) = %d\n", calculator.Add(1, 2))
	fmt.Printf("  calculator.Multiply(3, 4) = %d\n", calculator.Multiply(3, 4))

	fmt.Println("✓ Can access exported types:")
	calc := calculator.NewCalculator()
	fmt.Printf("  Created calculator: %T\n", calc)

	fmt.Println("✓ Can access exported methods:")
	result := calc.Add(5, 6)
	fmt.Printf("  calc.Add(5, 6) = %d\n", result)

	fmt.Println("✗ Cannot access unexported functions:")
	fmt.Println("  calculator.multiply(2, 3)  // Compilation error - unexported")
	fmt.Println("  calculator.power(2, 3)     // Compilation error - unexported")

	fmt.Println("✗ Cannot access unexported methods:")
	fmt.Println("  calc.recordOperation(...)  // Compilation error - unexported")

	fmt.Println("✓ But can access public interface to private functionality:")
	history := calc.GetHistory()
	fmt.Printf("  Got history with %d operations (uses private recordOperation)\n", len(history))
}

// toTitle converts a string to title case, replacing deprecated strings.Title
func toTitle(s string) string {
	if s == "" {
		return s
	}

	result := make([]rune, 0, len(s))
	words := strings.Fields(s)

	for i, word := range words {
		if i > 0 {
			result = append(result, ' ')
		}
		if len(word) > 0 {
			runes := []rune(word)
			upperFirst := []rune(strings.ToUpper(string(runes[0])))
			if len(upperFirst) > 0 {
				runes[0] = upperFirst[0]
			}
			result = append(result, runes...)
		}
	}

	return string(result)
}
