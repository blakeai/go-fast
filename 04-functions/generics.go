package main

import (
	"fmt"
	"strconv"
)

// Single type parameter with constraint
func dotProduct[F ~float32 | ~float64](v1, v2 []F) F {
	var sum F
	for i, x := range v1 {
		if i < len(v2) {
			sum += x * v2[i]
		}
	}
	return sum
}

// Multiple type parameters
func convert[T any, U any](input T, converter func(T) U) U {
	return converter(input)
}

// Generic slice operations
func contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func mapSlice[T any, U any](slice []T, mapper func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = mapper(v)
	}
	return result
}

// Stack Generic stack implementation
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Numeric Constraint interfaces
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func minExample[T Numeric](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func maxExample[T Numeric](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Stringer Generic with method constraint
type Stringer interface {
	String() string
}

func printStrings[T Stringer](items []T) {
	for _, item := range items {
		fmt.Printf("- %s\n", item.String())
	}
}

// Product Custom type implementing Stringer
type Product struct {
	Name  string
	Price float64
}

func (p Product) String() string {
	return fmt.Sprintf("%s ($%.2f)", p.Name, p.Price)
}

// Type inference examples
func demonstrateTypeInference() {
	fmt.Println("=== Type Inference ===")

	// Explicit type specification
	result1 := dotProduct[float64]([]float64{1.0, 2.0}, []float64{3.0, 4.0})
	fmt.Printf("Explicit: dotProduct[float64] = %.2f\n", result1)

	// Type inference from arguments
	result2 := dotProduct([]float32{1.0, 2.0}, []float32{3.0, 4.0})
	fmt.Printf("Inferred: dotProduct = %.2f\n", result2)

	// Mixed usage
	intToString := func(i int) string { return strconv.Itoa(i) }
	converted := convert(42, intToString)
	fmt.Printf("Converted: %s\n", converted)
}

// Generic function with multiple constraints
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

func sort3[T Ordered](a, b, c T) (T, T, T) {
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}
	return a, b, c
}

func genericsExample() {
	fmt.Println("=== Basic Generics ===")

	// Float vectors
	v1 := []float64{1.0, 2.0, 3.0}
	v2 := []float64{4.0, 5.0, 6.0}
	fmt.Printf("dotProduct(float64): %.2f\n", dotProduct(v1, v2))

	v3 := []float32{1.0, 2.0}
	v4 := []float32{3.0, 4.0}
	fmt.Printf("dotProduct(float32): %.2f\n", dotProduct(v3, v4))

	demonstrateTypeInference()

	fmt.Println("\n=== Generic Slice Operations ===")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("Contains 5: %t\n", contains(numbers, 5))
	fmt.Printf("Contains 15: %t\n", contains(numbers, 15))

	evens := filter(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Printf("Even numbers: %v\n", evens)

	doubled := mapSlice(numbers, func(n int) int { return n * 2 })
	fmt.Printf("Doubled: %v\n", doubled)

	fmt.Println("\n=== Generic Data Structures ===")

	// Integer stack
	intStack := &Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Printf("Stack size: %d\n", intStack.Size())
	for !intStack.IsEmpty() {
		if item, ok := intStack.Pop(); ok {
			fmt.Printf("Popped: %d\n", item)
		}
	}

	// String stack
	stringStack := &Stack[string]{}
	stringStack.Push("hello")
	stringStack.Push("world")

	if item, ok := stringStack.Pop(); ok {
		fmt.Printf("Popped string: %s\n", item)
	}

	fmt.Println("\n=== Numeric Constraints ===")
	fmt.Printf("minExample(5, 3) = %d\n", minExample(5, 3))
	fmt.Printf("maxExample(2.5, 7.1) = %.1f\n", maxExample(2.5, 7.1))

	fmt.Println("\n=== Interface Constraints ===")
	products := []Product{
		{Name: "Laptop", Price: 999.99},
		{Name: "Mouse", Price: 29.99},
		{Name: "Keyboard", Price: 79.99},
	}
	printStrings(products)

	fmt.Println("\n=== Ordered Types ===")
	a, b, c := sort3(3, 1, 2)
	fmt.Printf("sort3(3, 1, 2) = %d, %d, %d\n", a, b, c)

	x, y, z := sort3("zebra", "apple", "banana")
	fmt.Printf("sort3(strings) = %s, %s, %s\n", x, y, z)
}
