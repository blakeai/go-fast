package main

import (
	"fmt"
	"reflect"
	"time"
)

//goland:noinspection GoBoolExpressions
func basicSwitch() {
	fmt.Println("=== Basic Switch Statements ===")

	day := "friday"

	// Basic switch - no break needed (doesn't fall through by default)
	switch day {
	case "monday":
		fmt.Println("Start of work week")
	case "tuesday", "wednesday", "thursday": // multiple values in one case
		fmt.Println("Midweek")
	case "friday":
		fmt.Println("TGIF!")
	case "saturday", "sunday":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Unknown day")
	}

	// Switch with numbers
	grade := 85
	switch {
	case grade >= 90:
		fmt.Println("Grade A")
	case grade >= 80:
		fmt.Println("Grade B")
	case grade >= 70:
		fmt.Println("Grade C")
	case grade >= 60:
		fmt.Println("Grade D")
	default:
		fmt.Println("Grade F")
	}
}

func switchWithShortDeclaration() {
	fmt.Println("\n=== Switch with Short Variable Declaration ===")

	// Switch with short variable declaration
	switch today := time.Now().Weekday(); today {
	case time.Monday:
		fmt.Println("Monday blues")
	case time.Tuesday, time.Wednesday, time.Thursday:
		fmt.Println("Midweek grind")
	case time.Friday:
		fmt.Println("Almost weekend!")
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend vibes")
	}

	// Get current hour for time-based logic
	switch hour := time.Now().Hour(); {
	case hour < 6:
		fmt.Println("Very early morning")
	case hour < 12:
		fmt.Println("Morning")
	case hour < 18:
		fmt.Println("Afternoon")
	case hour < 22:
		fmt.Println("Evening")
	default:
		fmt.Println("Night")
	}
}

func typeSwitch() {
	fmt.Println("\n=== Type Switch ===")

	// Function to demonstrate type switching
	processValue := func(v interface{}) {
		switch val := v.(type) {
		case nil:
			fmt.Println("Value is nil")
		case int:
			fmt.Printf("Integer: %d (doubled: %d)\n", val, val*2)
		case float64:
			fmt.Printf("Float: %.2f (squared: %.2f)\n", val, val*val)
		case string:
			fmt.Printf("String: %q (length: %d)\n", val, len(val))
		case []int:
			fmt.Printf("Integer slice: %v (length: %d)\n", val, len(val))
		case []string:
			fmt.Printf("String slice: %v (joined: %s)\n", val, fmt.Sprint(val))
		case bool:
			if val {
				fmt.Println("Boolean: true")
			} else {
				fmt.Println("Boolean: false")
			}
		default:
			fmt.Printf("Unknown type: %T with value: %v\n", val, val)
		}
	}

	// Test different types
	values := []interface{}{
		42,
		3.14159,
		"Hello, Go!",
		[]int{1, 2, 3, 4, 5},
		[]string{"a", "b", "c"},
		true,
		nil,
		map[string]int{"key": 42},
	}

	for _, value := range values {
		processValue(value)
	}
}

//goland:noinspection ALL
func switchWithoutExpression() {
	fmt.Println("\n=== Switch Without Expression (replaces if-else chains) ===")

	age := 25
	income := 50000
	hasJob := true

	// Switch without expression - cleaner than long if-else chains
	switch {
	case age < 18:
		fmt.Println("Minor - not eligible for loan")
	case age >= 18 && age < 21 && !hasJob:
		fmt.Println("Young adult without job - high risk")
	case age >= 21 && age < 65 && hasJob && income >= 30000:
		fmt.Println("Eligible for standard loan")
	case age >= 21 && age < 65 && hasJob && income >= 50000:
		fmt.Println("Eligible for premium loan")
	case age >= 65:
		fmt.Println("Senior - special loan terms apply")
	default:
		fmt.Println("Not eligible for loan")
	}

	// Another example with string validation
	email := "user@example.com"
	switch {
	case email == "":
		fmt.Println("Email is required")
	case len(email) < 5:
		fmt.Println("Email too short")
	case !contains(email, "@"):
		fmt.Println("Email must contain @")
	case !contains(email, "."):
		fmt.Println("Email must contain a domain")
	default:
		fmt.Println("Email looks valid")
	}
}

// Helper function for string contains
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func fallThroughExample() {
	fmt.Println("\n=== Explicit Fallthrough ===")

	grade := 'A'

	// Explicit fallthrough when needed (rare in Go)
	switch grade {
	case 'A':
		fmt.Println("Excellent work!")
		fallthrough // explicitly continue to next case
	case 'B':
		fmt.Println("Good job!")
		// No fallthrough here - stops at B
	case 'C':
		fmt.Println("Passing grade")
	case 'D':
		fmt.Println("Below average")
	case 'F':
		fmt.Println("Failing grade")
	default:
		fmt.Println("Invalid grade")
	}

	// Another fallthrough example
	fmt.Println("\nFallthrough with numbers:")
	number := 1
	switch number {
	case 1:
		fmt.Print("One ")
		fallthrough
	case 2:
		fmt.Print("Two ")
		fallthrough
	case 3:
		fmt.Print("Three ")
		// No fallthrough - stops here
	case 4:
		fmt.Print("Four ")
	}
	fmt.Println("(fallthrough demo)")
}

//goland:noinspection GoBoolExpressions
func switchVsIfElse() {
	fmt.Println("\n=== Switch vs If-Else Performance ===")

	value := 5

	// Switch version - often more readable
	switch value {
	case 1, 2, 3:
		fmt.Println("Low value")
	case 4, 5, 6:
		fmt.Println("Medium value")
	case 7, 8, 9:
		fmt.Println("High value")
	default:
		fmt.Println("Out of range")
	}

	// Equivalent if-else version - more verbose
	if value >= 1 && value <= 3 {
		fmt.Println("Low value (if-else)")
	} else if value >= 4 && value <= 6 {
		fmt.Println("Medium value (if-else)")
	} else if value >= 7 && value <= 9 {
		fmt.Println("High value (if-else)")
	} else {
		fmt.Println("Out of range (if-else)")
	}
}

func advancedTypeSwitching() {
	fmt.Println("\n=== Advanced Type Switching ===")

	// Type switching with interfaces
	var shapes []interface{} = []interface{}{
		Circle{radius: 5},
		Rectangle{width: 10, height: 5},
		"not a shape",
		Triangle{base: 6, height: 4},
	}

	for i, shape := range shapes {
		fmt.Printf("Shape %d: ", i+1)
		switch s := shape.(type) {
		case Circle:
			area := 3.14159 * s.radius * s.radius
			fmt.Printf("Circle with radius %.1f, area: %.2f\n", s.radius, area)
		case Rectangle:
			area := s.width * s.height
			fmt.Printf("Rectangle %v×%v, area: %.2f\n", s.width, s.height, area)
		case Triangle:
			area := 0.5 * s.base * s.height
			fmt.Printf("Triangle base:%.1f height:%.1f, area: %.2f\n", s.base, s.height, area)
		case string:
			fmt.Printf("String value: %q (not a shape)\n", s)
		default:
			fmt.Printf("Unknown type: %s\n", reflect.TypeOf(s))
		}
	}
}

// Shape types for advanced example
type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

type Triangle struct {
	base, height float64
}

//goland:noinspection GoBoolExpressions
func practicalSwitchExamples() {
	fmt.Println("\n=== Practical Switch Examples ===")

	// HTTP status code handling
	statusCode := 404
	switch statusCode {
	case 200:
		fmt.Println("✅ OK")
	case 201:
		fmt.Println("✅ Created")
	case 400:
		fmt.Println("❌ Bad Request")
	case 401:
		fmt.Println("❌ Unauthorized")
	case 403:
		fmt.Println("❌ Forbidden")
	case 404:
		fmt.Println("❌ Not Found")
	case 500:
		fmt.Println("❌ Internal Server Error")
	default:
		if statusCode >= 200 && statusCode < 300 {
			fmt.Println("✅ Success")
		} else if statusCode >= 400 && statusCode < 500 {
			fmt.Println("❌ Client Error")
		} else if statusCode >= 500 {
			fmt.Println("❌ Server Error")
		} else {
			fmt.Printf("Unknown status code: %d\n", statusCode)
		}
	}

	// File extension handling
	filename := "document.pdf"
	var ext string
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			ext = filename[i+1:]
			break
		}
	}

	switch ext {
	case "pdf":
		fmt.Println("📄 PDF document")
	case "txt", "md":
		fmt.Println("📝 Text document")
	case "jpg", "jpeg", "png", "gif":
		fmt.Println("🖼️ Image file")
	case "mp4", "avi", "mov":
		fmt.Println("🎥 Video file")
	case "mp3", "wav", "flac":
		fmt.Println("🎵 Audio file")
	default:
		fmt.Printf("❓ Unknown file type: .%s\n", ext)
	}
}

func switchExample() {
	basicSwitch()
	switchWithShortDeclaration()
	typeSwitch()
	switchWithoutExpression()
	fallThroughExample()
	switchVsIfElse()
	advancedTypeSwitching()
	practicalSwitchExamples()

	fmt.Println("\n=== Key Switch Takeaways ===")
	fmt.Println("✅ No break needed - doesn't fall through by default")
	fmt.Println("✅ Multiple values per case: case 1, 2, 3:")
	fmt.Println("✅ Type switches: switch v := x.(type)")
	fmt.Println("✅ Switch without expression replaces if-else chains")
	fmt.Println("✅ Use fallthrough keyword for explicit fall-through")
	fmt.Println("✅ Can have short variable declarations")
}
