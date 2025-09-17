package main

import (
	"fmt"
	"reflect"
)

// Basic enum implementation
type Status int

const (
	Pending   Status = iota // 0
	Running                 // 1
	Completed               // 2
	Failed                  // 3
)

// String representation
func (s Status) String() string {
	switch s {
	case Pending:
		return "Pending"
	case Running:
		return "Running"
	case Completed:
		return "Completed"
	case Failed:
		return "Failed"
	default:
		return "Unknown"
	}
}

// Validation
func (s Status) IsValid() bool {
	return s >= Pending && s <= Failed
}

// Terminal states
func (s Status) IsTerminal() bool {
	return s == Completed || s == Failed
}

// Understanding iota
func demonstrateIota() {
	fmt.Println("=== Understanding iota ===")

	const (
		a = iota // 0
		b        // 1 (implicitly = iota)
		c        // 2
		d        // 3
	)
	fmt.Printf("a=%d, b=%d, c=%d, d=%d\n", a, b, c, d)

	// Automatic repetition of expressions
	const (
		x = 42
		y // y = 42 (repeats the expression)
		z // z = 42
	)
	fmt.Printf("x=%d, y=%d, z=%d\n", x, y, z)

	// iota resets to 0 in each new const block
	const (
		first  = iota // 0
		second        // 1
	)
	fmt.Printf("first=%d, second=%d\n", first, second)
}

// Custom values and expressions
type Priority int

const (
	Low      Priority   = iota + 1 // 1
	Medium                         // 2
	High                           // 3
	Critical = 10                  // explicit value
	Urgent   = iota + 1            // continues from where iota left off: 4 + 1 = 5
)

func (p Priority) String() string {
	switch p {
	case Low:
		return "Low"
	case Medium:
		return "Medium"
	case High:
		return "High"
	case Critical:
		return "Critical"
	case Urgent:
		return "Urgent"
	default:
		return fmt.Sprintf("Priority(%d)", int(p))
	}
}

// Size constants with bit shifting
const (
	_  = iota             // 0 (discarded)
	KB = 1 << (10 * iota) // 1024
	MB                    // 1048576
	GB                    // 1073741824
)

// String-based enums
type Color string

const (
	Red   Color = "red"
	Green Color = "green"
	Blue  Color = "blue"
)

func (c Color) IsValid() bool {
	return c == Red || c == Green || c == Blue
}

func (c Color) ToHex() string {
	switch c {
	case Red:
		return "#FF0000"
	case Green:
		return "#00FF00"
	case Blue:
		return "#0000FF"
	default:
		return "#000000"
	}
}

// Enum with methods and behavior
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	directions := []string{"North", "East", "South", "West"}
	if d >= 0 && int(d) < len(directions) {
		return directions[d]
	}
	return "Unknown"
}

func (d Direction) Opposite() Direction {
	switch d {
	case North:
		return South
	case South:
		return North
	case East:
		return West
	case West:
		return East
	default:
		return d
	}
}

func (d Direction) TurnRight() Direction {
	return (d + 1) % 4
}

func (d Direction) TurnLeft() Direction {
	return (d + 3) % 4 // Same as (d - 1 + 4) % 4
}

// Flags using bit operations
type Permission int

const (
	Read    Permission = 1 << iota // 1
	Write                          // 2
	Execute                        // 4
)

func (p Permission) String() string {
	var perms []string
	if p&Read != 0 {
		perms = append(perms, "Read")
	}
	if p&Write != 0 {
		perms = append(perms, "Write")
	}
	if p&Execute != 0 {
		perms = append(perms, "Execute")
	}
	return fmt.Sprintf("[%s]", fmt.Sprint(perms))
}

func (p Permission) HasRead() bool    { return p&Read != 0 }
func (p Permission) HasWrite() bool   { return p&Write != 0 }
func (p Permission) HasExecute() bool { return p&Execute != 0 }

// Enum with associated data (using interface)
type LogLevel interface {
	Level() int
	String() string
}

type Debug struct{}
type Info struct{ Message string }
type Warning struct{ Code int }
type Error struct{ Error error }

func (Debug) Level() int     { return 0 }
func (Debug) String() string { return "DEBUG" }

func (i Info) Level() int     { return 1 }
func (i Info) String() string { return fmt.Sprintf("INFO: %s", i.Message) }

func (w Warning) Level() int     { return 2 }
func (w Warning) String() string { return fmt.Sprintf("WARNING (%d)", w.Code) }

func (e Error) Level() int     { return 3 }
func (e Error) String() string { return fmt.Sprintf("ERROR: %v", e.Error) }

func enumsExample() {
	fmt.Println("=== Basic Enum Usage ===")

	status := Running
	fmt.Printf("Status: %s (%d)\n", status, int(status))
	fmt.Printf("Is valid: %t\n", status.IsValid())
	fmt.Printf("Is terminal: %t\n", status.IsTerminal())

	// Test all statuses
	statuses := []Status{Pending, Running, Completed, Failed}
	for _, s := range statuses {
		fmt.Printf("%s: valid=%t, terminal=%t\n", s, s.IsValid(), s.IsTerminal())
	}

	demonstrateIota()

	fmt.Println("\n=== Custom Values ===")
	priorities := []Priority{Low, Medium, High, Critical, Urgent}
	for _, p := range priorities {
		fmt.Printf("%s (%d)\n", p, int(p))
	}

	fmt.Println("\n=== Size Constants ===")
	fmt.Printf("KB: %d bytes\n", KB)
	fmt.Printf("MB: %d bytes\n", MB)
	fmt.Printf("GB: %d bytes\n", GB)

	fmt.Println("\n=== String-based Enums ===")
	color := Red
	fmt.Printf("Color: %s, Hex: %s, Valid: %t\n", color, color.ToHex(), color.IsValid())

	invalidColor := Color("purple")
	fmt.Printf("Invalid color: %s, Valid: %t\n", invalidColor, invalidColor.IsValid())

	fmt.Println("\n=== Direction Enum with Behavior ===")
	dir := North
	fmt.Printf("Direction: %s\n", dir)
	fmt.Printf("Opposite: %s\n", dir.Opposite())
	fmt.Printf("Turn right: %s\n", dir.TurnRight())
	fmt.Printf("Turn left: %s\n", dir.TurnLeft())

	fmt.Println("\n=== Flag Enums ===")
	perm := Read | Write // Combine flags
	fmt.Printf("Permission: %s\n", perm)
	fmt.Printf("Has read: %t\n", perm.HasRead())
	fmt.Printf("Has write: %t\n", perm.HasWrite())
	fmt.Printf("Has execute: %t\n", perm.HasExecute())

	fullPerm := Read | Write | Execute
	fmt.Printf("Full permission: %s\n", fullPerm)

	fmt.Println("\n=== Advanced: Associated Data ===")
	logs := []LogLevel{
		Debug{},
		Info{Message: "System started"},
		Warning{Code: 404},
		Error{Error: fmt.Errorf("connection failed")},
	}

	for _, log := range logs {
		fmt.Printf("Level %d: %s\n", log.Level(), log.String())
	}

	fmt.Println("\n=== Type Safety Demonstration ===")
	// These would cause compile errors:
	// var s Status = 999  // Can't assign int directly
	// fmt.Println(s == 1) // Can't compare with int directly

	// But this works:
	var s Status = Status(999) // Explicit conversion
	fmt.Printf("Invalid status: %s, Valid: %t\n", s, s.IsValid())

	// Type reflection
	fmt.Printf("Status type: %s\n", reflect.TypeOf(status))
	fmt.Printf("Color type: %s\n", reflect.TypeOf(color))
}
