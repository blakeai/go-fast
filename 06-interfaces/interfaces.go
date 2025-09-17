package main

import (
	"fmt"
	"reflect"
)

// Basic interface example
type Counter interface {
	Increment()
	Value() int
}

type IntCounter int

// Pointer receiver - can modify the original
func (c *IntCounter) Increment() {
	(*c)++
}

func (c *IntCounter) Value() int {
	return int(*c)
}

// Demonstrate why no pointers to interfaces
func demonstrateInterfacePointers() {
	fmt.Println("=== Why No Pointers to Interfaces? ===")

	// Correct - interface as value
	var intCounter IntCounter = 0
	var counter Counter = &intCounter
	counter.Increment() // works perfectly
	fmt.Printf("Counter value: %d\n", counter.Value())

	// Wrong - don't do this (would be confusing)
	// var badCounter *Counter  // pointer to interface
	// This is unnecessary because interface already provides indirection

	fmt.Println("✅ Interface value provides the indirection you need")
}

// Writer interface example
type Writer interface {
	Write([]byte) (int, error)
}

type FileWriter struct {
	filename string
	content  []byte
}

// FileWriter automatically satisfies Writer interface
func (fw *FileWriter) Write(data []byte) (int, error) {
	fw.content = append(fw.content, data...)
	fmt.Printf("Writing %d bytes to %s\n", len(data), fw.filename)
	return len(data), nil
}

func (fw *FileWriter) GetContent() string {
	return string(fw.content)
}

// Multiple interface implementation
type ReadWriter interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
}

type Buffer struct {
	data []byte
	pos  int
}

func (b *Buffer) Read(p []byte) (int, error) {
	if b.pos >= len(b.data) {
		return 0, fmt.Errorf("EOF")
	}

	n := copy(p, b.data[b.pos:])
	b.pos += n
	fmt.Printf("Read %d bytes from buffer\n", n)
	return n, nil
}

func (b *Buffer) Write(p []byte) (int, error) {
	b.data = append(b.data, p...)
	fmt.Printf("Wrote %d bytes to buffer\n", len(p))
	return len(p), nil
}

func (b *Buffer) String() string {
	return string(b.data)
}

// Interface composition example
type Reader interface {
	Read([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type ReadCloser interface {
	Reader // embedded interface
	Closer // embedded interface
}

type File struct {
	name   string
	data   []byte
	pos    int
	closed bool
}

func (f *File) Read(p []byte) (int, error) {
	if f.closed {
		return 0, fmt.Errorf("file is closed")
	}
	if f.pos >= len(f.data) {
		return 0, fmt.Errorf("EOF")
	}

	n := copy(p, f.data[f.pos:])
	f.pos += n
	return n, nil
}

func (f *File) Close() error {
	if f.closed {
		return fmt.Errorf("file already closed")
	}
	f.closed = true
	fmt.Printf("Closed file: %s\n", f.name)
	return nil
}

// Empty interface and type assertions
func demonstrateEmptyInterface() {
	fmt.Println("\n=== Empty Interface and Type Assertions ===")

	// Empty interface - any type satisfies it
	var anything interface{} = 42
	fmt.Printf("anything = %v (type: %T)\n", anything, anything)

	anything = "hello"
	fmt.Printf("anything = %v (type: %T)\n", anything, anything)

	anything = []int{1, 2, 3}
	fmt.Printf("anything = %v (type: %T)\n", anything, anything)

	// Type assertion - extract concrete type
	anything = "world"
	if str, ok := anything.(string); ok {
		fmt.Printf("It's a string: %s\n", str)
	}

	// Type switch
	anything = 123
	switch v := anything.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case []int:
		fmt.Printf("Slice: %v\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

// Nil interfaces vs nil values
func demonstrateNilInterfaces() {
	fmt.Println("\n=== Nil Interfaces vs Nil Values ===")

	var counter Counter // nil interface
	fmt.Printf("counter == nil: %t\n", counter == nil)

	var intCounter *IntCounter        // nil pointer
	var counter2 Counter = intCounter // interface with nil value

	// Check for nil interface
	if counter == nil {
		fmt.Println("counter is a nil interface")
	}

	// Check for nil value inside interface
	if counter2 == nil {
		fmt.Println("this won't print - interface is not nil, value is")
	} else {
		fmt.Println("counter2 interface is not nil, but contains a nil value")
	}

	// Proper nil value check
	if reflect.ValueOf(counter2).IsNil() {
		fmt.Println("counter2 interface contains nil value")
	}

	// Using a non-nil value
	validCounter := IntCounter(5)
	counter3 := Counter(&validCounter)
	fmt.Printf("counter3 value: %d\n", counter3.Value())
}

// Practical interface usage
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return 3.14159 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.radius
}

// Function that works with any shape
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())

	// Type assertion for specific behavior
	switch shape := s.(type) {
	case Rectangle:
		fmt.Printf("  Rectangle: %.1f x %.1f\n", shape.width, shape.height)
	case Circle:
		fmt.Printf("  Circle: radius %.1f\n", shape.radius)
	}
}

func main() {
	fmt.Println("=== Basic Interface Usage ===")

	// Usage - pass interface as value
	var intCounter IntCounter = 0
	var counter Counter = &intCounter                  // interface holds pointer to IntCounter
	counter.Increment()                                // modifies the underlying IntCounter
	fmt.Printf("Counter value: %d\n", counter.Value()) // prints: 1

	demonstrateInterfacePointers()

	fmt.Println("\n=== Interface Satisfaction ===")

	// No explicit "implements" declaration needed
	var w Writer = &FileWriter{filename: "output.txt"}
	w.Write([]byte("Hello, interfaces!"))

	if fw, ok := w.(*FileWriter); ok {
		fmt.Printf("File content: %s\n", fw.GetContent())
	}

	fmt.Println("\n=== Multiple Interface Implementation ===")

	// Buffer satisfies ReadWriter automatically
	var rw ReadWriter = &Buffer{}
	rw.Write([]byte("Hello, buffer!"))

	readBuf := make([]byte, 5)
	n, _ := rw.Read(readBuf)
	fmt.Printf("Read: %s (%d bytes)\n", string(readBuf[:n]), n)

	fmt.Println("\n=== Interface Composition ===")

	// File satisfies ReadCloser through composition
	var rc ReadCloser = &File{
		name: "test.txt",
		data: []byte("Hello, composition!"),
	}

	readBuf = make([]byte, 10)
	n, _ = rc.Read(readBuf)
	fmt.Printf("Read from file: %s\n", string(readBuf[:n]))
	rc.Close()

	demonstrateEmptyInterface()
	demonstrateNilInterfaces()

	fmt.Println("\n=== Practical Example: Shapes ===")

	shapes := []Shape{
		Rectangle{width: 10, height: 5},
		Circle{radius: 3},
		Rectangle{width: 2, height: 8},
		Circle{radius: 1.5},
	}

	for i, shape := range shapes {
		fmt.Printf("Shape %d: ", i+1)
		printShapeInfo(shape)
	}

	fmt.Println("\n=== Key Interface Principles ===")
	fmt.Println("✅ Interfaces define behavior contracts")
	fmt.Println("✅ Implicit satisfaction - no 'implements' keyword")
	fmt.Println("✅ Pass interfaces as values, not pointers")
	fmt.Println("✅ Empty interface accepts any type")
	fmt.Println("✅ Type assertions and switches for discrimination")
	fmt.Println("✅ Composition over inheritance")
}
