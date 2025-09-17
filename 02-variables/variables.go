package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strings"
)

var globalCounter int = 0
var (
	globalName   string = "Global Alice"
	globalAge    int    = 30
	globalActive bool   = true
)

const Pi = 3.14159
const MaxUsers = 100
const AppName = "VariablesDemo"

const (
	StatusActive   = 1
	StatusInactive = 0
	StatusPending  = 2
)

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
)

const (
	Red = iota
	Green
	Blue
)

const (
	Small = iota
	Medium
	Large
)

type User struct {
	Name  string
	Email string
}

type Config struct {
	Host string
	Port int
}

func variableDeclarationDemo() {
	fmt.Println("=== Variable Declaration Patterns ===")

	var name string = "Alice"
	var age int = 30
	fmt.Printf("Explicit types: name='%s', age=%d\n", name, age)

	var city = "New York"
	var population = 8000000
	fmt.Printf("Type inference: city='%s' (%s), population=%d (%s)\n",
		city, reflect.TypeOf(city), population, reflect.TypeOf(population))

	username := "bob"
	count := 42
	fmt.Printf("Short declaration: username='%s', count=%d\n", username, count)

	var x, y int = 1, 2
	a, b := "hello", "world"
	fmt.Printf("Multiple variables: x=%d, y=%d, a='%s', b='%s'\n", x, y, a, b)

	var (
		firstName string = "John"
		lastName  string = "Doe"
		salary    int    = 50000
	)
	fmt.Printf("Group declaration: %s %s, salary=$%d\n", firstName, lastName, salary)
}

func shortVsVarDemo() {
	fmt.Println("\n=== Short Declaration vs Var ===")

	name := "Alice"
	var age int = 30
	var height float64
	fmt.Printf("Inside function: name='%s', age=%d, height=%.1f\n", name, age, height)

	name, email := "Bob", "bob@example.com"
	fmt.Printf("Mixed assignment: name='%s' (reassigned), email='%s' (new)\n", name, email)

	fmt.Printf("Global counter: %d\n", globalCounter)

	var buffer strings.Builder
	var users []User
	var config Config
	fmt.Printf("Zero values: buffer=%v, users=%v, config=%+v\n", buffer, users, config)

	items := []string{"item1", "item2"}
	itemName := "test"
	itemCount := len(items)
	fmt.Printf("With initial values: name='%s', count=%d\n", itemName, itemCount)
}

func constantsDemo() {
	fmt.Println("\n=== Constants and Iota ===")

	fmt.Printf("Basic constants: Pi=%.5f, MaxUsers=%d, AppName='%s'\n", Pi, MaxUsers, AppName)

	fmt.Printf("Status constants: Active=%d, Inactive=%d, Pending=%d\n",
		StatusActive, StatusInactive, StatusPending)

	fmt.Printf("Weekdays: Sunday=%d, Monday=%d, Tuesday=%d, Wednesday=%d, Thursday=%d, Friday=%d, Saturday=%d\n",
		Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	fmt.Printf("Bytes: B=%d, KB=%d, MB=%d, GB=%d, TB=%d\n", B, KB, MB, GB, TB)

	fmt.Printf("Colors: Red=%d, Green=%d, Blue=%d\n", Red, Green, Blue)
	fmt.Printf("Sizes: Small=%d, Medium=%d, Large=%d\n", Small, Medium, Large)
}

func typeInferenceDemo() {
	fmt.Println("\n=== Type Inference and Explicit Typing ===")

	var a = 42
	var b = 3.14
	var c = "hello"
	var d = true
	fmt.Printf("Inferred types: a=%d (%s), b=%.2f (%s), c='%s' (%s), d=%t (%s)\n",
		a, reflect.TypeOf(a), b, reflect.TypeOf(b), c, reflect.TypeOf(c), d, reflect.TypeOf(d))

	var smallInt int8 = 42
	var precise float32 = 3.14
	fmt.Printf("Explicit small types: smallInt=%d (%s), precise=%.2f (%s)\n",
		smallInt, reflect.TypeOf(smallInt), precise, reflect.TypeOf(precise))

	var w io.Writer = &bytes.Buffer{}
	fmt.Printf("Interface type: w=%v (%s)\n", w, reflect.TypeOf(w))

	var count int
	var nameStr string
	var active bool
	var items []string
	fmt.Printf("Zero values: count=%d, name='%s', active=%t, items=%v\n", count, nameStr, active, items)

	var i int = 42
	var f float64 = float64(i)
	var j int8 = int8(i)
	fmt.Printf("Type conversions: i=%d -> f=%.1f, i=%d -> j=%d\n", i, f, i, j)

	var r rune = 65
	var s string = string(r)
	var byt byte = 65
	var s2 string = string(byt)
	fmt.Printf("String conversions: rune %d -> '%s', byte %d -> '%s'\n", r, s, byt, s2)
}

func scopeAndShadowingDemo() {
	fmt.Println("\n=== Variable Scope and Shadowing ===")

	outer := "outer"
	fmt.Printf("Before inner scope: outer='%s'\n", outer)

	if true {
		inner := "inner"
		outer := "shadowed" //nolint:govet // Intentional shadowing example
		fmt.Printf("Inside scope: outer='%s', inner='%s'\n", outer, inner)
	}

	fmt.Printf("After inner scope: outer='%s'\n", outer)

	global := "local global"
	fmt.Printf("Shadowed global: local='%s', package='%s'\n", global, globalName)

	fmt.Printf("Package-level globals: name='%s', age=%d, active=%t\n", globalName, globalAge, globalActive)
}

func main() {
	variableDeclarationDemo()
	shortVsVarDemo()
	constantsDemo()
	typeInferenceDemo()
	scopeAndShadowingDemo()

	fmt.Println("\n" + strings.Repeat("=", 60) + "\n")
	runAddressableExamples()
}
