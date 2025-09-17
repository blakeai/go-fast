package main

import (
	"fmt"
	"strings"
)

type Counter struct {
	value int
}

// Value receiver - operates on a copy
func (c Counter) increment() {
	c.value++ // Only modifies the copy
	fmt.Printf("Inside increment (value receiver): %d\n", c.value)
}

// Pointer receiver - operates on the original
func (c *Counter) incrementPtr() {
	c.value++ // Modifies the actual instance
	fmt.Printf("Inside incrementPtr (pointer receiver): %d\n", c.value)
}

// Value receiver for reading
func (c Counter) getValue() int {
	return c.value
}

// Pointer receiver for writing
func (c *Counter) setValue(value int) {
	c.value = value
}

// Demonstrate method set rules
func demonstrateMethodSets() {
	fmt.Println("=== Method Set Rules ===")

	var v Counter = Counter{value: 10}
	var p *Counter = &Counter{value: 20}

	fmt.Printf("Value v: %d\n", v.getValue())
	fmt.Printf("Pointer p: %d\n", p.getValue())

	// Go automatically takes address for pointer receiver methods
	v.incrementPtr() // Equivalent to (&v).incrementPtr()
	fmt.Printf("v after incrementPtr: %d\n", v.getValue())

	// Go automatically dereferences for value receiver methods
	p.increment() // Equivalent to (*p).increment()
	fmt.Printf("p after increment: %d\n", p.getValue())
}

// Example with larger struct
type Person struct {
	Name string
	Age  int
	City string
}

// Pointer receiver for modification
func (p *Person) SetAge(age int) {
	p.Age = age
}

// Pointer receiver to avoid copying large struct
func (p *Person) GetFullInfo() string {
	return fmt.Sprintf("%s is %d years old and lives in %s", p.Name, p.Age, p.City)
}

// Value receiver for small, immutable operations
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// Value receiver that returns modified copy (immutable pattern)
func (p Person) WithAge(age int) Person {
	p.Age = age // Modifies the copy
	return p    // Returns the modified copy
}

// Same package rule demonstration
type MyString string

func (s MyString) Upper() string {
	return strings.ToUpper(string(s))
}

func (s MyString) Reverse() string {
	runes := []rune(string(s))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Consistency example - all methods should use same receiver type
type BankAccount struct {
	balance float64
}

// All methods use pointer receivers for consistency
func (b *BankAccount) Deposit(amount float64) {
	b.balance += amount
}

func (b *BankAccount) Withdraw(amount float64) error {
	if amount > b.balance {
		return fmt.Errorf("insufficient funds")
	}
	b.balance -= amount
	return nil
}

// Even read-only methods use pointer receivers for consistency
func (b *BankAccount) GetBalance() float64 {
	return b.balance
}

func receiversExample() {
	fmt.Println("=== Value vs Pointer Receivers ===")

	counter := Counter{value: 5}
	fmt.Printf("Initial counter: %d\n", counter.getValue())

	// Value receiver - doesn't modify original
	counter.increment()
	fmt.Printf("After increment (value receiver): %d\n", counter.getValue())

	// Pointer receiver - modifies original
	counter.incrementPtr()
	fmt.Printf("After incrementPtr (pointer receiver): %d\n", counter.getValue())

	demonstrateMethodSets()

	fmt.Println("\n=== Large Struct Example ===")
	person := Person{Name: "Alice", Age: 25, City: "New York"}
	fmt.Printf("IsAdult: %t\n", person.IsAdult())
	fmt.Println(person.GetFullInfo())

	person.SetAge(30)
	fmt.Println(person.GetFullInfo())

	// Immutable pattern
	youngerPerson := person.WithAge(20)
	fmt.Printf("Original person age: %d\n", person.Age)
	fmt.Printf("Younger person age: %d\n", youngerPerson.Age)

	fmt.Println("\n=== Custom Type Methods ===")
	text := MyString("hello world")
	fmt.Printf("Original: %s\n", text)
	fmt.Printf("Upper: %s\n", text.Upper())
	fmt.Printf("Reverse: %s\n", text.Reverse())

	fmt.Println("\n=== Consistent Receiver Types ===")
	account := &BankAccount{balance: 100.0}
	fmt.Printf("Initial balance: $%.2f\n", account.GetBalance())

	account.Deposit(50.0)
	fmt.Printf("After deposit: $%.2f\n", account.GetBalance())

	err := account.Withdraw(200.0)
	if err != nil {
		fmt.Printf("Withdrawal error: %v\n", err)
	}

	err = account.Withdraw(30.0)
	if err != nil {
		return
	}
	fmt.Printf("Final balance: $%.2f\n", account.GetBalance())
}
