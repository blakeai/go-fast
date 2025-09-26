package main

import "fmt"

// adder creates a closure that maintains a running sum.
// Each call to the returned function adds the argument to the accumulated sum.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// createCalculator returns a closure that maintains both a total and operation count.
// It supports basic arithmetic operations and provides operation tracking.
func createCalculator() func(string, int) int {
	total := 0
	operationCount := 0

	return func(operation string, value int) int {
		operationCount++
		switch operation {
		case "add":
			total += value
		case "subtract":
			total -= value
		case "multiply":
			total *= value
		case "reset":
			total = 0
		}
		fmt.Printf("Operation %d: %s %d, total: %d\n", operationCount, operation, value, total)
		return total
	}
}

// makeMultiplier is a closure factory that returns multiplier functions.
// Each returned function multiplies its input by the captured factor.
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// closureInLoop demonstrates the classic closure-in-loop pitfall and its solutions.
// Shows both the problem and correct patterns for capturing loop variables.
func closureInLoop() {
	fmt.Println("\n=== Closure in Loop Examples ===")

	// In Go 1.22+, this now works correctly (each iteration gets its own variable)
	fmt.Println("Common mistake (all capture same variable):")
	var funcs []func() int
	for i := 0; i < 3; i++ {
		funcs = append(funcs, func() int {
			return i // captures the loop variable 'i'
		})
	}
	for idx, f := range funcs {
		fmt.Printf("funcs[%d]() = %d\n", idx, f()) // Go 1.22+ prints 0, 1, 2 correctly
	}

	// Historical pattern for older Go versions (still valid)
	fmt.Println("\nCorrect pattern (capture by parameter):")
	var correctFuncs []func() int
	for i := 0; i < 3; i++ {
		correctFuncs = append(correctFuncs, func(val int) func() int {
			return func() int {
				return val // captures the parameter value
			}
		}(i))
	}
	for idx, f := range correctFuncs {
		fmt.Printf("correctFuncs[%d]() = %d\n", idx, f()) // prints 0, 1, 2
	}
}

// createValidator creates a closure that validates string length within bounds.
// The returned function captures the min/max parameters for reuse.
func createValidator(minLength, maxLength int) func(string) bool {
	return func(input string) bool {
		length := len(input)
		return length >= minLength && length <= maxLength
	}
}

// createAccount demonstrates closure-based encapsulation by creating an account with deposit,
// withdraw, and balance functions that share access to the same currentBalance variable.
// This pattern provides data privacy similar to object-oriented encapsulation.
func createAccount(initialBalance float64) (deposit func(float64), withdraw func(float64), balance func() float64) {
	currentBalance := initialBalance

	deposit = func(amount float64) {
		if amount > 0 {
			currentBalance += amount
			fmt.Printf("Deposited $%.2f, balance: $%.2f\n", amount, currentBalance)
		}
	}

	withdraw = func(amount float64) {
		if amount > 0 && amount <= currentBalance {
			currentBalance -= amount
			fmt.Printf("Withdrew $%.2f, balance: $%.2f\n", amount, currentBalance)
		} else {
			fmt.Printf("Cannot withdraw $%.2f (insufficient funds or invalid amount)\n", amount)
		}
	}

	balance = func() float64 {
		return currentBalance
	}

	return deposit, withdraw, balance
}

func closuresExample() {
	fmt.Println("=== Closures in Go ===")

	// Basic adder example
	fmt.Println("\n--- Basic Closure (Adder) ---")
	posSum := adder()
	fmt.Printf("posSum(3) = %d\n", posSum(3))   // 3
	fmt.Printf("posSum(5) = %d\n", posSum(5))   // 8
	fmt.Printf("posSum(10) = %d\n", posSum(10)) // 18

	// Each closure maintains its own state
	another := adder()
	fmt.Printf("another(2) = %d\n", another(2)) // 2
	fmt.Printf("posSum(1) = %d\n", posSum(1))   // 19 (continues from previous state)

	// Calculator with multiple captured variables
	fmt.Println("\n--- Calculator Closure ---")
	calc := createCalculator()
	calc("add", 10)     // Operation 1: add 10, total: 10
	calc("subtract", 3) // Operation 2: subtract 3, total: 7
	calc("multiply", 2) // Operation 3: multiply 2, total: 14

	// Multiplier factory
	fmt.Println("\n--- Multiplier Factory ---")
	double := makeMultiplier(2)
	triple := makeMultiplier(3)
	fmt.Printf("double(5) = %d\n", double(5)) // 10
	fmt.Printf("triple(5) = %d\n", triple(5)) // 15

	// Closure in loop pitfalls
	closureInLoop()

	// Validator closure
	fmt.Println("\n--- Validator Closure ---")
	passwordValidator := createValidator(8, 20)
	usernameValidator := createValidator(3, 15)

	passwords := []string{"short", "good_password", "this_password_is_way_too_long"}
	for _, pwd := range passwords {
		fmt.Printf("Password '%s' valid: %t\n", pwd, passwordValidator(pwd))
	}

	usernames := []string{"ab", "user123", "validusername"}
	for _, user := range usernames {
		fmt.Printf("Username '%s' valid: %t\n", user, usernameValidator(user))
	}

	// Account closure with multiple functions
	fmt.Println("\n--- Account Closure ---")
	deposit, withdraw, balance := createAccount(100.0)
	fmt.Printf("Initial balance: $%.2f\n", balance())
	deposit(50.0)
	withdraw(30.0)
	withdraw(200.0) // Should fail
	fmt.Printf("Final balance: $%.2f\n", balance())
}
