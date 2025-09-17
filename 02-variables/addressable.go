package main

import (
	"fmt"
	"strings"
)

type Person struct {
	name string
	age  int
}

type Counter struct {
	value int
}

func (c *Counter) Increment() {
	c.value++
}

func (c Counter) Value() int {
	return c.value
}

func (c *Counter) Add(n int) {
	c.value += n
}

func (c Counter) String() string {
	return fmt.Sprintf("Counter{value: %d}", c.value)
}

func addressableValuesDemo() {
	fmt.Println("=== Addressable Values ===")

	var x int = 5
	ptr := &x
	fmt.Printf("Variable address: x=%d, ptr=%p, *ptr=%d\n", x, ptr, *ptr)

	arr := [3]int{1, 2, 3}
	arrPtr := &arr
	elemPtr := &arr[0]
	fmt.Printf("Array: arr=%v, &arr=%p, &arr[0]=%p, arr[0]=%d\n", arr, arrPtr, elemPtr, *elemPtr)

	var person Person
	person.name = "Alice"
	personPtr := &person
	namePtr := &person.name
	fmt.Printf("Struct: person=%+v, &person=%p, &person.name=%p\n", person, personPtr, namePtr)

	slice := []int{1, 2, 3}
	slicePtr := &slice
	sliceElemPtr := &slice[1]
	fmt.Printf("Slice: slice=%v, &slice=%p, &slice[1]=%p, slice[1]=%d\n", slice, slicePtr, sliceElemPtr, *sliceElemPtr)
}

func nonAddressableDemo() {
	fmt.Println("\n=== Non-Addressable Values ===")

	m := map[int]Person{1: {name: "Alice", age: 30}}
	fmt.Printf("Map value: m[1]=%+v\n", m[1])
	fmt.Println("Cannot take &m[1] - map values are not addressable")
	fmt.Println("Cannot take &m[1].name - fields of non-addressable values are not addressable")

	fmt.Println("\nFunction return values are not addressable:")
	fmt.Printf("getPerson() returns: %+v\n", getPerson())
	fmt.Println("Cannot take &getPerson() - function results not addressable")

	fmt.Println("\nLiterals are not addressable:")
	fmt.Println("Cannot take &42, &\"hello\", or &Person{name: \"Carol\"}")

	arr := getArray()
	fmt.Printf("Array from function: %v\n", arr)
	fmt.Println("Cannot take &getArray()[0] - elements of non-addressable array")
}

func getPerson() Person {
	return Person{name: "Bob", age: 25}
}

func getArray() [3]int {
	return [3]int{1, 2, 3}
}

func addressabilityMattersDemo() {
	fmt.Println("\n=== Why Addressability Matters ===")

	fmt.Println("1. Variable counter (addressable):")
	var counter Counter
	fmt.Printf("Initial: %s\n", counter.String())
	counter.Increment()
	fmt.Printf("After Increment(): %s\n", counter.String())
	counter.Add(5)
	fmt.Printf("After Add(5): %s\n", counter.String())

	fmt.Println("\n2. Map values (non-addressable) - this would fail:")
	m := map[string]Counter{"main": {value: 0}}
	fmt.Printf("Map counter: %s\n", m["main"].String())
	fmt.Println("m[\"main\"].Increment() would fail - cannot take address")
	fmt.Printf("m[\"main\"].Value() works fine: %d\n", m["main"].Value())

	fmt.Println("\n3. Solutions for map values:")

	fmt.Println("   a) Store pointers in map:")
	m1 := map[string]*Counter{"main": {value: 0}}
	fmt.Printf("   Before: %s\n", m1["main"].String())
	m1["main"].Increment()
	fmt.Printf("   After Increment(): %s\n", m1["main"].String())

	fmt.Println("   b) Extract, modify, put back:")
	tempCounter := m["main"]
	tempCounter.Increment()
	m["main"] = tempCounter
	fmt.Printf("   After extract-modify-putback: %s\n", m["main"].String())
}

func sliceArrayAddressabilityDemo() {
	fmt.Println("\n=== Slice vs Array Addressability ===")

	arr := [3]int{1, 2, 3}
	arrElemPtr := &arr[1]
	fmt.Printf("Array element address: arr[1]=%d, &arr[1]=%p\n", arr[1], arrElemPtr)

	slice := []int{1, 2, 3}
	sliceElemPtr := &slice[1]
	fmt.Printf("Slice element address: slice[1]=%d, &slice[1]=%p\n", slice[1], sliceElemPtr)

	fmt.Println("Non-addressable array from function:")
	returnedArr := getArray()
	fmt.Printf("getArray() returns: %v\n", returnedArr)
	fmt.Println("Cannot take &getArray()[1] - elements of non-addressable array")

	var s []int
	sliceHeaderPtr := &s
	fmt.Printf("Slice header address: &s=%p\n", sliceHeaderPtr)
}

func methodReceiverDemo() {
	fmt.Println("\n=== Method Receivers and Addressability ===")

	fmt.Println("Pointer receiver methods require addressable values:")

	var counter1 Counter
	fmt.Printf("Variable counter: %s\n", counter1.String())
	counter1.Increment()
	fmt.Printf("After Increment() on variable: %s\n", counter1.String())

	counterPtr := &Counter{value: 10}
	fmt.Printf("Pointer to counter: %s\n", counterPtr.String())
	counterPtr.Increment()
	fmt.Printf("After Increment() on pointer: %s\n", counterPtr.String())

	fmt.Println("\nValue receiver methods work on any value:")
	tempCounter := Counter{value: 20}
	fmt.Printf("Temporary counter value: %d\n", tempCounter.Value())
	fmt.Printf("Map counter value: %d\n", map[string]Counter{"test": {value: 30}}["test"].Value())

	fmt.Println("\nDemonstrating automatic address-taking:")
	var autoCounter Counter
	fmt.Printf("Before: %s\n", autoCounter.String())
	autoCounter.Increment()
	fmt.Printf("After: %s\n", autoCounter.String())
	fmt.Println("Go automatically converts autoCounter.Increment() to (&autoCounter).Increment()")
}

func typeConversionDemo() {
	fmt.Println("\n=== Type Conversion Examples ===")

	var i int = 42
	var f float64 = float64(i)
	var i8 int8 = int8(i)
	var i64 int64 = int64(i)

	fmt.Printf("int to other types: i=%d -> f=%.1f, i8=%d, i64=%d\n", i, f, i8, i64)

	var f32 float32 = 3.14
	var f64 float64 = float64(f32)
	var backToInt int = int(f32)

	fmt.Printf("float conversions: f32=%.2f -> f64=%.2f, backToInt=%d\n", f32, f64, backToInt)

	var r rune = 65
	var b byte = 65
	var s1 string = string(r)
	var s2 string = string(b)
	var rs []rune = []rune("Hello")
	var bs []byte = []byte("Hello")

	fmt.Printf("String conversions: rune %d -> '%s', byte %d -> '%s'\n", r, s1, b, s2)
	fmt.Printf("String to slices: \"Hello\" -> runes %v, bytes %v\n", rs, bs)
}

func runAddressableExamples() {
	addressableValuesDemo()
	nonAddressableDemo()
	addressabilityMattersDemo()
	sliceArrayAddressabilityDemo()
	methodReceiverDemo()
	typeConversionDemo()
}

func init() {
	fmt.Println("Starting Variables and Addressability Examples")
	fmt.Println(strings.Repeat("=", 60))
}
