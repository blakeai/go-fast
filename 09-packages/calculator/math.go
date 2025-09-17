// Package calculator provides basic arithmetic operations.
// It supports addition, subtraction, multiplication, and division
// with proper error handling for edge cases.
//
// Example usage:
//
//	result := calculator.Add(5, 3)
//	quotient, err := calculator.Divide(10, 2)
//	if err != nil {
//	    log.Fatal(err)
//	}
package calculator

import (
	"errors"
	"fmt"
)

// Add returns the sum of two integers.
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference of two integers.
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of two integers using the unexported multiply function.
func Multiply(a, b int) int {
	return multiply(a, b)
}

// Divide returns the quotient of two float64 numbers.
// It returns an error if the divisor is zero.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// multiply is an unexported function that can only be used within the calculator package.
func multiply(a, b int) int {
	return a * b
}

// power is another unexported helper function.
func power(base, exp int) int {
	if exp == 0 {
		return 1
	}
	result := 1
	for i := 0; i < exp; i++ {
		result = multiply(result, base)
	}
	return result
}

// Power returns base raised to the power of exp using unexported helper functions.
func Power(base, exp int) int {
	return power(base, exp)
}

// Operation represents a single arithmetic operation.
type Operation struct {
	Type   string // "add", "subtract", "multiply", "divide"
	A, B   int    // operands (for float operations, these are converted)
	Result int    // result (for float operations, this is truncated)
}

// Calculator provides arithmetic operations with history tracking.
type Calculator struct {
	history []Operation
}

// NewCalculator creates a new Calculator instance.
func NewCalculator() *Calculator {
	return &Calculator{
		history: make([]Operation, 0),
	}
}

// Add performs addition and records the operation in history.
func (c *Calculator) Add(a, b int) int {
	result := Add(a, b)
	c.recordOperation("add", a, b, result)
	return result
}

// Subtract performs subtraction and records the operation in history.
func (c *Calculator) Subtract(a, b int) int {
	result := Subtract(a, b)
	c.recordOperation("subtract", a, b, result)
	return result
}

// Multiply performs multiplication and records the operation in history.
func (c *Calculator) Multiply(a, b int) int {
	result := Multiply(a, b)
	c.recordOperation("multiply", a, b, result)
	return result
}

// GetHistory returns a copy of the operation history.
func (c *Calculator) GetHistory() []Operation {
	historyCopy := make([]Operation, len(c.history))
	copy(historyCopy, c.history)
	return historyCopy
}

// ClearHistory clears the operation history.
func (c *Calculator) ClearHistory() {
	c.history = c.history[:0]
}

// recordOperation is an unexported method that records operations in the history.
func (c *Calculator) recordOperation(op string, a, b, result int) {
	c.history = append(c.history, Operation{
		Type:   op,
		A:      a,
		B:      b,
		Result: result,
	})
}

// String returns a string representation of the calculator's history.
func (c *Calculator) String() string {
	if len(c.history) == 0 {
		return "Calculator with no operations"
	}

	result := fmt.Sprintf("Calculator with %d operations:\n", len(c.history))
	for i, op := range c.history {
		result += fmt.Sprintf("  %d. %s(%d, %d) = %d\n", i+1, op.Type, op.A, op.B, op.Result)
	}
	return result
}
