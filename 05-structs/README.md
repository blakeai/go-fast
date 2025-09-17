# Chapter 5: Structs

## Overview

Go structs are the primary way to group data and define custom types. Unlike classes, structs don't have inheritance - instead, Go uses composition and embedding. Understanding struct initialization patterns, embedding mechanics, and when struct fields are addressable is crucial for effective Go programming.

## Key Concepts

- **No inheritance** - use composition and struct embedding instead
- **Multiple initialization patterns** - literal, new(), and field-by-field
- **Embedded structs** vs **embedded fields** - anonymous composition
- **Addressability of struct fields** - when you can use `&structVar.field`
- **Value vs pointer semantics** - copying behavior and method receivers

## Examples

### Basic Struct Definition and Initialization
See [`structs.go`](./structs.go) for implementation.

```go
type Person struct {
    Name string
    Age  int
    Email string
}

// Struct literal with field names (recommended)
person1 := Person{
    Name:  "Alice",
    Age:   30,
    Email: "alice@example.com",
}

// Struct literal with positional arguments (fragile, avoid)
person2 := Person{"Bob", 25, "bob@example.com"}

// Partial initialization - unspecified fields get zero values
person3 := Person{
    Name: "Carol",
    // Age: 0, Email: "" (zero values)
}

// Using new() - returns pointer to zero-valued struct
person4 := new(Person)  // *Person
person4.Name = "Dave"   // Go automatically dereferences

// Taking address of struct literal
person5 := &Person{
    Name: "Eve",
    Age:  28,
}
```

### Struct Embedding - Composition Over Inheritance
See [`embedding.go`](./embedding.go) for implementation.

```go
type Address struct {
    Street string
    City   string
    State  string
}

type Contact struct {
    Phone string
    Email string
}

// Struct embedding - Address and Contact are promoted to Person level
type Person struct {
    Name    string
    Age     int
    Address // embedded struct - no field name
    Contact // embedded struct - no field name
}

// Usage - embedded fields are promoted
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

// Can access embedded fields directly
fmt.Println(person.Street)  // promoted from person.Address.Street
fmt.Println(person.Phone)   // promoted from person.Contact.Phone

// Or access the embedded struct explicitly
fmt.Println(person.Address.Street)
fmt.Println(person.Contact.Phone)
```

### Embedded Fields vs Embedded Structs
```go
// Embedded struct (anonymous field)
type Employee struct {
    Person  // field name is the type name: "Person"
    Salary  int
    Title   string
}

// Named struct field (composition, not embedding)
type Manager struct {
    Info   Person  // explicit field name
    Reports []Employee
}

// Usage differences:
emp := Employee{}
emp.Name = "Alice"        // promoted from embedded Person
emp.Person.Name = "Alice" // explicit access to embedded struct

mgr := Manager{}
mgr.Info.Name = "Bob"     // must use explicit field name - no promotion
// mgr.Name = "Bob"       // ERROR - Name not promoted
```

### Method Promotion with Embedding
```go
type Writer struct {
    data []string
}

func (w *Writer) Write(s string) {
    w.data = append(w.data, s)
}

func (w *Writer) String() string {
    return strings.Join(w.data, "\n")
}

// Logger embeds Writer - gets all Writer methods
type Logger struct {
    Writer   // embedded
    prefix   string
}

func (l *Logger) Log(message string) {
    // Can call promoted method directly
    l.Write(fmt.Sprintf("[%s] %s", l.prefix, message))
}

// Usage
logger := Logger{prefix: "INFO"}
logger.Log("Application started")  // calls Logger.Log
logger.Write("Direct write")        // calls promoted Writer.Write
fmt.Println(logger.String())        // calls promoted Writer.String
```

### Struct Tags for Metadata
```go
type User struct {
    ID       int    `json:"id" db:"user_id"`
    Name     string `json:"name" db:"full_name"`
    Email    string `json:"email,omitempty" db:"email_address"`
    Password string `json:"-" db:"password_hash"`  // "-" means ignore in JSON
}

// Tags are used by reflection-based libraries
func printStructTags() {
    u := User{}
    t := reflect.TypeOf(u)
    
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        jsonTag := field.Tag.Get("json")
        dbTag := field.Tag.Get("db")
        fmt.Printf("%s: json='%s', db='%s'\n", field.Name, jsonTag, dbTag)
    }
}
```

### Struct Addressability and Pointer Semantics
```go
type Counter struct {
    value int
}

func (c *Counter) Increment() { c.value++ }

// Struct variables are addressable
var counter Counter
counter.Increment()  // Go takes address automatically: (&counter).Increment()

// Struct literals are not addressable
// Counter{}.Increment()  // ERROR - can't take address of literal

// But pointer to struct literal works
(&Counter{}).Increment()  // OK - explicit address

// Fields of addressable structs are addressable
var person Person
namePtr := &person.Name  // OK - person is addressable, so person.Name is too

// Fields of non-addressable structs are not addressable
// namePtr := &Person{Name: "Alice"}.Name  // ERROR - Person{} not addressable
```

### Anonymous Structs for One-Off Data
```go
// Anonymous struct for temporary grouping
config := struct {
    Host string
    Port int
    SSL  bool
}{
    Host: "localhost",
    Port: 8080,
    SSL:  false,
}

// Common in table-driven tests
tests := []struct {
    name     string
    input    int
    expected string
}{
    {"zero", 0, "zero"},
    {"positive", 42, "positive"},
    {"negative", -1, "negative"},
}

for _, test := range tests {
    t.Run(test.name, func(t *testing.T) {
        result := classify(test.input)
        if result != test.expected {
            t.Errorf("got %s, want %s", result, test.expected)
        }
    })
}
```

### Struct Comparison and Zero Values
```go
type Point struct {
    X, Y int
}

// Structs are comparable if all fields are comparable
p1 := Point{1, 2}
p2 := Point{1, 2}
fmt.Println(p1 == p2)  // true

// Zero value is struct with all fields zero-valued
var p3 Point  // {0, 0}

// Struct with non-comparable fields (slices, maps, functions) cannot be compared
type Container struct {
    Data []int  // slice makes struct non-comparable
}

// var c1, c2 Container
// fmt.Println(c1 == c2)  // ERROR - cannot compare
```

## Running the Code

```bash
go run *.go
go test ./...
```

## Java Developer Notes

- No classes or inheritance - structs + embedding replace class hierarchies
- No constructors - use factory functions or struct literals
- All fields are public if capitalized, private if lowercase (no `public`/`private` keywords)
- Embedded structs ≈ multiple inheritance but simpler and more explicit
- No `super` keyword - access embedded structs by type name
- Struct comparison works automatically (unlike Java object reference comparison)
- No automatic getters/setters - access fields directly or write explicit methods
- `new(Type)` ≈ Java's `new Type()` but returns pointer to zero value

## Next Steps

Continue to [Chapter 6: Interfaces](../06-interfaces/)

## References

- [Go Tour - Structs](https://go.dev/tour/moretypes/2)
- [Effective Go - Composite literals](https://go.dev/doc/effective_go#composite_literals)
- [Go Spec - Struct types](https://go.dev/ref/spec#Struct_types)