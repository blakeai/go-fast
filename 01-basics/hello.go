package main

import (
	"fmt"
	"strings"
)

func main() {
	message := "Hello, Go!"
	fmt.Println(strings.ToUpper(message))

	basicProgramStructureDemo()
	importPatternsDemo()
	multipleReturnDemo()

	fmt.Println("\n" + strings.Repeat("=", 60) + "\n")
	runBasicsExamples()
}

func basicProgramStructureDemo() {
	fmt.Println("\n=== Basic Program Structure ===")

	greeting := "Welcome to Go!"
	fmt.Printf("Original: %s\n", greeting)
	fmt.Printf("Uppercase: %s\n", strings.ToUpper(greeting))
	fmt.Printf("Lowercase: %s\n", strings.ToLower(greeting))
	fmt.Printf("Title case: %s\n", toTitle(greeting))
}

func importPatternsDemo() {
	fmt.Println("\n=== Import Patterns Demo ===")

	fmt.Println("Standard fmt import works")

	aliasedDemo()

	_ = "This demonstrates the blank identifier to avoid unused variable errors"
}

func aliasedDemo() {
	str := strings.ToUpper("this shows how to use aliased imports")
	fmt.Printf("Using regular import: %s\n", str)
}

func multipleReturnDemo() {
	fmt.Println("\n=== Multiple Return Values ===")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	result2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Printf("Error: %v\n", err2)
	} else {
		fmt.Printf("Result: %.2f\n", result2)
	}
}

func toTitle(s string) string {
	if len(s) == 0 {
		return s
	}
	result := strings.ToLower(s)
	words := strings.Fields(result)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + word[1:]
		}
	}
	return strings.Join(words, " ")
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
