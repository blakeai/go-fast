package main

import (
	"fmt"
	"reflect"
)

var (
	globalName   string = "Alice"
	globalAge    int    = 30
	globalActive bool   = true
)

type Counter struct {
	value int
	name  string
}

func zeroValuesDemo() {
	fmt.Println("=== Zero Values ===")

	var i int
	var f float64
	var b bool
	var s string
	var p *int
	var slice []int
	var m map[string]int
	var ch chan int

	fmt.Printf("int zero value: %d\n", i)
	fmt.Printf("float64 zero value: %f\n", f)
	fmt.Printf("bool zero value: %t\n", b)
	fmt.Printf("string zero value: '%s'\n", s)
	fmt.Printf("pointer zero value: %v\n", p)
	fmt.Printf("slice zero value: %v (len=%d, cap=%d)\n", slice, len(slice), cap(slice))
	fmt.Printf("map zero value: %v\n", m)
	fmt.Printf("channel zero value: %v\n", ch)

	var counter Counter
	fmt.Printf("struct zero value: %+v\n", counter)
	fmt.Printf("Counter is immediately usable: value=%d, name='%s'\n", counter.value, counter.name)
}

func basicDataTypesDemo() {
	fmt.Println("\n=== Basic Data Types ===")

	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	var i int = 42

	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	var u uint = 42

	var f32 float32 = 3.14159
	var f64 float64 = 3.141592653589793

	var c64 complex64 = 3 + 4i
	var c128 complex128 = complex(3.0, 4.0)

	var str string = "Hello, 世界"
	var flag bool = true

	var b byte = 65
	var r rune = '世'

	fmt.Printf("Signed integers: i8=%d, i16=%d, i32=%d, i64=%d, i=%d\n", i8, i16, i32, i64, i)
	fmt.Printf("Unsigned integers: u8=%d, u16=%d, u32=%d, u64=%d, u=%d\n", u8, u16, u32, u64, u)
	fmt.Printf("Floating point: f32=%f, f64=%f\n", f32, f64)
	fmt.Printf("Complex: c64=%v, c128=%v\n", c64, c128)
	fmt.Printf("String: %s\n", str)
	fmt.Printf("Boolean: %t\n", flag)
	fmt.Printf("Byte (as char): %c, Rune (as char): %c\n", b, r)

	fmt.Printf("Type of i: %s\n", reflect.TypeOf(i))
	fmt.Printf("Type of f64: %s\n", reflect.TypeOf(f64))
	fmt.Printf("Type of str: %s\n", reflect.TypeOf(str))
}

func multipleDeclarationsDemo() {
	fmt.Println("\n=== Multiple Variable Declarations ===")

	fmt.Printf("Global variables: name=%s, age=%d, active=%t\n", globalName, globalAge, globalActive)

	var x, y, z int = 1, 2, 3
	fmt.Printf("Multiple same type: x=%d, y=%d, z=%d\n", x, y, z)

	var a, b = "hello", 42
	fmt.Printf("Mixed types with inference: a=%s, b=%d\n", a, b)

	name := "Bob"
	age := 25
	fmt.Printf("Short variable declaration: name=%s, age=%d\n", name, age)
}

func blankIdentifierDemo() {
	fmt.Println("\n=== Blank Identifier Demo ===")

	unused := "This would normally cause an error"
	_ = unused
	fmt.Println("Blank identifier prevents unused variable error")

	result, err := divide(15, 3)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("15 / 3 = %.2f\n", result)
	}

	result2, _ := divide(20, 4)
	fmt.Printf("20 / 4 = %.2f (ignoring error with blank identifier)\n", result2)
}

type ExportedType struct {
	PublicField  string
	privateField int
}

func ExportedFunction() string {
	return "This function is exported"
}

func privateFunction() string {
	return "This function is not exported"
}

func visibilityDemo() {
	fmt.Println("\n=== Visibility Rules Demo ===")

	exported := ExportedType{
		PublicField:  "accessible from other packages",
		privateField: 42,
	}

	fmt.Printf("ExportedType: %+v\n", exported)
	fmt.Printf("Exported function result: %s\n", ExportedFunction())
	fmt.Printf("Private function result: %s\n", privateFunction())

	fmt.Println("Note: privateField and privateFunction would not be accessible from other packages")
}

func runBasicsExamples() {
	zeroValuesDemo()
	basicDataTypesDemo()
	multipleDeclarationsDemo()
	blankIdentifierDemo()
	visibilityDemo()
}
