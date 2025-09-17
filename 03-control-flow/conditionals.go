package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// User Example struct for demonstration
type User struct {
	Name string
	Age  int
	ID   int
}

// Mock function to simulate getting a user
func getUser(id int) *User {
	users := map[int]*User{
		1: {Name: "Alice", Age: 25, ID: 1},
		2: {Name: "Bob", Age: 30, ID: 2},
	}
	return users[id]
}

// Mock function to simulate an operation that might fail
func doSomething() error {
	// Simulate success or failure
	return nil // Change to fmt.Errorf("something went wrong") to test error case
}

//goland:noinspection GoBoolExpressions
func basicIfStatements() {
	fmt.Println("=== Basic If Statements ===")

	age := 25

	// Standard if
	if age >= 18 {
		fmt.Println("You are an adult")
	}

	// If-else
	if age < 13 {
		fmt.Println("Child")
	} else if age < 20 {
		fmt.Println("Teenager")
	} else {
		fmt.Println("Adult")
	}

	// No parentheses needed around condition (unlike Java/C++)
	temperature := 75
	if temperature > 80 {
		fmt.Println("It's hot!")
	} else if temperature < 60 {
		fmt.Println("It's cold!")
	} else {
		fmt.Println("Nice weather!")
	}
}

func shortVariableDeclarations() {
	fmt.Println("\n=== If with Short Variable Declarations ===")

	// Variable declared and scoped to if block
	if user := getUser(1); user != nil {
		fmt.Printf("Found user: %s (age %d)\n", user.Name, user.Age)
		// user is accessible here
		if user.Age >= 18 {
			fmt.Println("User is an adult")
		}
	} else {
		fmt.Println("User not found")
		// user is also accessible in else block
	}
	// user is NOT accessible here - out of scope

	// Another example with string parsing
	if num, err := strconv.Atoi("123"); err == nil {
		fmt.Printf("Parsed number: %d\n", num)
		fmt.Printf("Number squared: %d\n", num*num)
	} else {
		fmt.Printf("Failed to parse: %v\n", err)
	}

	// Multiple assignment in if
	//goland:noinspection GoBoolExpressions
	if name, age := "Charlie", 35; age > 30 {
		fmt.Printf("%s is over 30 (age: %d)\n", name, age)
	}
}

//goland:noinspection GoUnhandledErrorResult
func errorCheckingPatterns() {
	fmt.Println("\n=== Error Checking Patterns ===")

	// Common pattern for error checking
	if err := doSomething(); err != nil {
		fmt.Printf("Operation failed: %v\n", err)
		return // Early return on error
	}
	fmt.Println("Operation succeeded!")

	// File operation example
	filename := "temp.txt"
	if file, err := os.Create(filename); err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
	} else {
		defer func() {
			file.Close()
			os.Remove(filename) // Clean up
		}()

		if _, err := file.WriteString("Hello, Go!"); err != nil {
			fmt.Printf("Failed to write to file: %v\n", err)
		} else {
			fmt.Println("Successfully wrote to file")
		}
	}

	// Map lookup with ok idiom
	users := map[string]int{"Alice": 25, "Bob": 30}
	if age, exists := users["Alice"]; exists {
		fmt.Printf("Alice is %d years old\n", age)
	} else {
		fmt.Println("Alice not found")
	}
}

//goland:noinspection GoBoolExpressions
func scopeExamples() {
	fmt.Println("\n=== Variable Scope in If Statements ===")

	x := 10
	fmt.Printf("Outer x: %d\n", x)

	if x := 20; x > 15 {
		fmt.Printf("Inner x: %d\n", x) // This shadows outer x
		if y := x * 2; y > 30 {
			fmt.Printf("Nested y: %d\n", y)
			// Both x and y accessible here
		}
		// y is out of scope here, but x is still the inner x
		fmt.Printf("Still inner x: %d\n", x)
	}

	fmt.Printf("Back to outer x: %d\n", x) // Outer x is restored
}

//goland:noinspection GoDfaConstantCondition
func nilChecks() {
	fmt.Println("\n=== Nil Checks and Pointer Safety ===")

	var user *User // nil pointer

	// Check for nil before accessing
	if user != nil { //nolint:nilness // Intentional nil check example
		fmt.Printf("User: %s\n", user.Name)
	} else {
		fmt.Println("User is nil")
	}

	// Safe access with short declaration
	if user := getUser(999); user != nil {
		fmt.Printf("Found user: %s\n", user.Name)
	} else {
		fmt.Println("User with ID 999 not found")
	}

	// Slice nil check
	var numbers []int
	if numbers == nil {
		fmt.Println("Slice is nil")
		numbers = make([]int, 0)
	}

	// Check if slice has elements
	if len(numbers) > 0 {
		fmt.Printf("First element: %d\n", numbers[0])
	} else {
		fmt.Println("Slice is empty")
	}
}

//goland:noinspection GoBoolExpressions,GoDfaConstantCondition
func stringConditions() {
	fmt.Println("\n=== String Conditions ===")

	name := "Alice"

	// String comparison
	if name == "Alice" {
		fmt.Println("Hello, Alice!")
	}

	// String length check
	if len(name) > 3 {
		fmt.Println("Name is longer than 3 characters")
	}

	// String contains
	email := "user@example.com"
	if strings.Contains(email, "@") {
		fmt.Println("Valid email format")
	}

	// String prefix/suffix
	filename := "document.pdf"
	if strings.HasSuffix(filename, ".pdf") {
		fmt.Println("PDF file detected")
	}

	// Empty string check
	var input string
	if input == "" {
		fmt.Println("Input is empty")
	}

	// Alternative: check length
	if len(input) == 0 {
		fmt.Println("Input has zero length")
	}
}

//goland:noinspection GoBoolExpressions
func compoundConditions() {
	fmt.Println("\n=== Compound Conditions ===")

	age := 25
	hasLicense := true
	hasInsurance := true

	// Logical AND
	if age >= 18 && hasLicense {
		fmt.Println("Can drive")
	}

	// Logical OR
	if age < 16 || !hasLicense {
		fmt.Println("Cannot drive alone")
	} else {
		fmt.Println("Can drive alone")
	}

	// Complex condition
	if age >= 18 && hasLicense && hasInsurance {
		fmt.Println("Fully qualified to drive")
	} else {
		fmt.Println("Missing requirements to drive")
	}

	// Multiple conditions with short-circuit evaluation
	user := getUser(1)
	if user != nil && user.Age >= 21 && user.Name != "" {
		fmt.Printf("%s can drink alcohol\n", user.Name)
	}
}

//goland:noinspection GoBoolExpressions
func asiAndBracePlacement() {
	fmt.Println("\n=== ASI and Brace Placement ===")

	condition := true

	// CORRECT - opening brace on same line
	if condition {
		fmt.Println("This works correctly")
	}

	// The following would be WRONG and cause compilation error:
	/*
		if condition
		{  // ASI inserts semicolon after condition
			fmt.Println("This won't work")
		}
	*/

	// Same rule applies to else
	if !condition {
		fmt.Println("Condition is false")
	} else { // brace must be on same line as else
		fmt.Println("Condition is true")
	}

	fmt.Println("✅ Always put opening braces on the same line!")
}

func practicalExamples() {
	fmt.Println("\n=== Practical Examples ===")

	// Configuration validation
	config := map[string]string{
		"host": "localhost",
		"port": "8080",
	}

	if host, exists := config["host"]; !exists || host == "" {
		fmt.Println("❌ Host configuration missing")
	} else if port, exists := config["port"]; !exists || port == "" {
		fmt.Println("❌ Port configuration missing")
	} else {
		fmt.Printf("✅ Server configured: %s:%s\n", host, port)
	}

	// Input validation
	userInput := "42"
	if value, err := strconv.Atoi(userInput); err != nil {
		fmt.Printf("❌ Invalid number: %s\n", userInput)
	} else if value < 0 {
		fmt.Printf("❌ Number must be positive: %d\n", value)
	} else if value > 100 {
		fmt.Printf("❌ Number too large: %d\n", value)
	} else {
		fmt.Printf("✅ Valid input: %d\n", value)
	}
}

func conditionalsExample() {
	basicIfStatements()
	shortVariableDeclarations()
	errorCheckingPatterns()
	scopeExamples()
	nilChecks()
	stringConditions()
	compoundConditions()
	asiAndBracePlacement()
	practicalExamples()

	fmt.Println("\n=== Key Conditional Takeaways ===")
	fmt.Println("✅ No parentheses needed around conditions")
	fmt.Println("✅ Short variable declarations scope to if block")
	fmt.Println("✅ Opening braces must be on same line (ASI)")
	fmt.Println("✅ Use if err != nil pattern for error checking")
	fmt.Println("✅ Check for nil before accessing pointers")
	fmt.Println("✅ Logical operators: && (AND), || (OR), ! (NOT)")
}
