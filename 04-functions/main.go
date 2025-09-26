package main

import (
	"fmt"
	"strings"
)

// main runs all function-related examples in the correct order.
// This demonstrates various Go function concepts from basic to advanced.
func main() {
	fmt.Println("Running Go Functions Examples...")
	fmt.Println("=================================")

	// Basic function concepts
	functionsExample()
	fmt.Println("\n" + strings.Repeat("=", 50))

	// Method receivers and object-oriented patterns
	receiversExample()
	fmt.Println("\n" + strings.Repeat("=", 50))

	// Generics and type parameters
	genericsExample()
	fmt.Println("\n" + strings.Repeat("=", 50))

	// Enum patterns in Go
	enumsExample()
	fmt.Println("\n" + strings.Repeat("=", 50))

	// Closure fundamentals
	closuresExample()
	fmt.Println("\n" + strings.Repeat("=", 50))

	// Advanced closure patterns
	advancedClosuresExample()

	fmt.Println("\n=================================")
	fmt.Println("All function examples completed!")
}
