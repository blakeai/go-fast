package main

import (
	"fmt"
	"strings"
)

type Address struct {
	Street string
	City   string
	State  string
}

func (a Address) String() string {
	return fmt.Sprintf("%s, %s, %s", a.Street, a.City, a.State)
}

type Contact struct {
	Phone string
	Email string
}

func (c Contact) String() string {
	return fmt.Sprintf("Phone: %s, Email: %s", c.Phone, c.Email)
}

type Person struct {
	Name string
	Age  int
	Address
	Contact
}

type Employee struct {
	Person
	Salary int
	Title  string
}

type Manager struct {
	Info    Person
	Reports []Employee
}

type Writer struct {
	data []string
}

func (w *Writer) Write(s string) {
	w.data = append(w.data, s)
}

func (w *Writer) String() string {
	return strings.Join(w.data, "\n")
}

type Logger struct {
	Writer
	prefix string
}

func (l *Logger) Log(message string) {
	l.Write(fmt.Sprintf("[%s] %s", l.prefix, message))
}

func embeddingBasicsDemo() {
	fmt.Println("=== Struct Embedding - Composition Over Inheritance ===")

	person := Person{
		Name: "Alice",
		Age:  30,
		Address: Address{
			Street: "123 Main St",
			City:   "Portland",
			State:  "OR",
		},
		Contact: Contact{
			Phone: "555-1234",
			Email: "alice@example.com",
		},
	}

	fmt.Printf("Person: %+v\n", person)
	fmt.Printf("Direct access to Street: %s\n", person.Street)
	fmt.Printf("Direct access to Phone: %s\n", person.Phone)
	fmt.Printf("Explicit access to Address.Street: %s\n", person.Address.Street)
	fmt.Printf("Explicit access to Contact.Phone: %s\n", person.Contact.Phone)

	fmt.Printf("Address String(): %s\n", person.Address.String())
	fmt.Printf("Contact String(): %s\n", person.Contact.String())
}

func embeddedVsNamedFieldsDemo() {
	fmt.Println("\n=== Embedded Fields vs Named Fields ===")

	emp := Employee{
		Person: Person{
			Name: "Bob",
			Age:  25,
			Address: Address{
				Street: "456 Oak Ave",
				City:   "Seattle",
				State:  "WA",
			},
		},
		Salary: 75000,
		Title:  "Software Engineer",
	}

	fmt.Printf("Employee: %+v\n", emp)
	fmt.Printf("Promoted field access - emp.Name: %s\n", emp.Name)
	fmt.Printf("Explicit access - emp.Person.Name: %s\n", emp.Person.Name)
	fmt.Printf("Nested promotion - emp.Street: %s\n", emp.Street)

	mgr := Manager{
		Info: Person{
			Name: "Carol",
			Age:  35,
		},
		Reports: []Employee{emp},
	}

	fmt.Printf("\nManager: %+v\n", mgr)
	fmt.Printf("Named field access - mgr.Info.Name: %s\n", mgr.Info.Name)
}

func methodPromotionDemo() {
	fmt.Println("\n=== Method Promotion with Embedding ===")

	logger := Logger{
		Writer: Writer{},
		prefix: "INFO",
	}

	logger.Log("Application started")
	logger.Write("Direct write to Writer")
	logger.Log("Another log message")

	fmt.Printf("Logger output:\n%s\n", logger.String())

	debugLogger := Logger{
		Writer: Writer{},
		prefix: "DEBUG",
	}

	debugLogger.Log("Debug message 1")
	debugLogger.Log("Debug message 2")
	fmt.Printf("\nDebug logger output:\n%s\n", debugLogger.String())
}

func embeddingConflictsDemo() {
	fmt.Println("\n=== Handling Embedding Conflicts ===")

	type A struct {
		Value int
	}

	type B struct {
		Value string
	}

	type C struct {
		A
		B
	}

	c := C{
		A: A{Value: 42},
		B: B{Value: "hello"},
	}

	fmt.Printf("c.A.Value (int): %d\n", c.A.Value)
	fmt.Printf("c.B.Value (string): %s\n", c.B.Value)
}

func main() {
	runStructsExamples()
	fmt.Println("\n" + strings.Repeat("=", 60) + "\n")
	embeddingBasicsDemo()
	embeddedVsNamedFieldsDemo()
	methodPromotionDemo()
	embeddingConflictsDemo()
}
