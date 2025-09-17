package main

import (
	"fmt"
	"reflect"
)

type SimplePerson struct {
	Name  string
	Age   int
	Email string
}

type Point struct {
	X, Y int
}

type Counter struct {
	value int
}

func (c *Counter) Increment() {
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

type User struct {
	ID       int    `json:"id" db:"user_id"`
	Name     string `json:"name" db:"full_name"`
	Email    string `json:"email,omitempty" db:"email_address"`
	Password string `json:"-" db:"password_hash"`
}

type Container struct {
	Data []int
}

func basicStructDemo() {
	fmt.Println("=== Basic Struct Definition and Initialization ===")

	person1 := SimplePerson{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}
	fmt.Printf("person1 (named fields): %+v\n", person1)

	person2 := SimplePerson{"Bob", 25, "bob@example.com"}
	fmt.Printf("person2 (positional): %+v\n", person2)

	person3 := SimplePerson{
		Name: "Carol",
	}
	fmt.Printf("person3 (partial init): %+v\n", person3)

	person4 := new(SimplePerson)
	person4.Name = "Dave"
	fmt.Printf("person4 (new): %+v\n", *person4)

	person5 := &SimplePerson{
		Name: "Eve",
		Age:  28,
	}
	fmt.Printf("person5 (address of literal): %+v\n", *person5)
}

func structComparisonDemo() {
	fmt.Println("\n=== Struct Comparison and Zero Values ===")

	p1 := Point{1, 2}
	p2 := Point{1, 2}
	p3 := Point{3, 4}

	fmt.Printf("p1: %+v\n", p1)
	fmt.Printf("p2: %+v\n", p2)
	fmt.Printf("p3: %+v\n", p3)
	fmt.Printf("p1 == p2: %t\n", p1 == p2)
	fmt.Printf("p1 == p3: %t\n", p1 == p3)

	var p4 Point
	fmt.Printf("Zero value point: %+v\n", p4)
}

func addressabilityDemo() {
	fmt.Println("\n=== Struct Addressability and Pointer Semantics ===")

	var counter Counter
	fmt.Printf("Initial counter value: %d\n", counter.Value())

	counter.Increment()
	fmt.Printf("After increment: %d\n", counter.Value())

	var person SimplePerson
	person.Name = "Test"
	namePtr := &person.Name
	*namePtr = "Modified"
	fmt.Printf("Modified person name: %s\n", person.Name)

	counterPtr := &Counter{}
	counterPtr.Increment()
	fmt.Printf("Pointer counter value: %d\n", counterPtr.Value())
}

func structTagsDemo() {
	fmt.Println("\n=== Struct Tags for Metadata ===")

	u := User{}
	t := reflect.TypeOf(u)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		dbTag := field.Tag.Get("db")
		fmt.Printf("%s: json='%s', db='%s'\n", field.Name, jsonTag, dbTag)
	}
}

func anonymousStructDemo() {
	fmt.Println("\n=== Anonymous Structs for One-Off Data ===")

	config := struct {
		Host string
		Port int
		SSL  bool
	}{
		Host: "localhost",
		Port: 8080,
		SSL:  false,
	}
	fmt.Printf("Config: %+v\n", config)

	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"zero", 0, "zero"},
		{"positive", 42, "positive"},
		{"negative", -1, "negative"},
	}

	fmt.Println("Test cases:")
	for _, test := range tests {
		result := classify(test.input)
		fmt.Printf("  %s: input=%d, expected=%s, got=%s\n", test.name, test.input, test.expected, result)
	}
}

func classify(n int) string {
	if n == 0 {
		return "zero"
	} else if n > 0 {
		return "positive"
	}
	return "negative"
}

func runStructsExamples() {
	basicStructDemo()
	structComparisonDemo()
	addressabilityDemo()
	structTagsDemo()
	anonymousStructDemo()
}
