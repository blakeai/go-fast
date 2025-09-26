package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

// 1. Middleware Pattern - HTTP-like middleware chain
func middleware() func(string) string {
	// Chain of responsibility using closures
	return func(input string) string {
		// Logger middleware
		logger := func(next func(string) string) func(string) string {
			return func(s string) string {
				fmt.Printf("[LOG] Processing: %s\n", s)
				result := next(s)
				fmt.Printf("[LOG] Result: %s\n", result)
				return result
			}
		}

		// Timer middleware
		timer := func(next func(string) string) func(string) string {
			return func(s string) string {
				start := time.Now()
				result := next(s)
				fmt.Printf("[TIMER] Took %v\n", time.Since(start))
				return result
			}
		}

		// Core handler
		handler := func(s string) string {
			time.Sleep(10 * time.Millisecond) // Simulate work
			return "Processed: " + s
		}

		// Build the chain
		chain := logger(timer(handler))
		return chain(input)
	}
}

// 2. Rate Limiter - Closure with time-based state
func rateLimiter(requests int, duration time.Duration) func() bool {
	tokens := requests
	lastReset := time.Now()
	mu := sync.Mutex{}

	return func() bool {
		mu.Lock()
		defer mu.Unlock()

		now := time.Now()
		// Reset tokens if duration has passed
		if now.Sub(lastReset) >= duration {
			tokens = requests
			lastReset = now
		}

		if tokens > 0 {
			tokens--
			return true
		}
		return false
	}
}

// 3. Memoization - Cache function results
func memoize(fn func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(x int) int {
		if val, exists := cache[x]; exists {
			fmt.Printf("Cache hit for %d: %d\n", x, val)
			return val
		}
		result := fn(x)
		cache[x] = result
		fmt.Printf("Computed and cached %d: %d\n", x, result)
		return result
	}
}

// Example expensive function for memoization
func fibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

// 4. Event Emitter Pattern
func createEventEmitter() (on func(string, func()), emit func(string)) {
	listeners := make(map[string][]func())

	on = func(event string, handler func()) {
		listeners[event] = append(listeners[event], handler)
		fmt.Printf("Registered handler for event '%s'\n", event)
	}

	emit = func(event string) {
		if handlers, exists := listeners[event]; exists {
			fmt.Printf("Emitting event '%s' to %d handlers\n", event, len(handlers))
			for _, handler := range handlers {
				handler()
			}
		}
	}

	return on, emit
}

// 5. Iterator Generator - Closure that produces sequences
func fibonacciGenerator() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func primeGenerator() func() int {
	current := 2

	isPrime := func(n int) bool {
		if n < 2 {
			return false
		}
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	return func() int {
		for !isPrime(current) {
			current++
		}
		result := current
		current++
		return result
	}
}

// 6. Builder Pattern with Fluent Interface
type QueryBuilder struct {
	table  string
	wheres []string
	limit  int
}

func createQueryBuilder() func(string) *QueryBuilder {
	return func(table string) *QueryBuilder {
		return &QueryBuilder{table: table}
	}
}

// Add methods to QueryBuilder
func (qb *QueryBuilder) Where(condition string) *QueryBuilder {
	qb.wheres = append(qb.wheres, condition)
	return qb
}

func (qb *QueryBuilder) Limit(n int) *QueryBuilder {
	qb.limit = n
	return qb
}

func (qb *QueryBuilder) Build() string {
	query := fmt.Sprintf("SELECT * FROM %s", qb.table)
	if len(qb.wheres) > 0 {
		query += " WHERE "
		for i, where := range qb.wheres {
			if i > 0 {
				query += " AND "
			}
			query += where
		}
	}
	if qb.limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", qb.limit)
	}
	return query
}

// 7. Sorting with Custom Comparators
func sortWithClosures() {
	type Person struct {
		Name string
		Age  int
		City string
	}

	people := []Person{
		{"Alice", 30, "NYC"},
		{"Bob", 25, "LA"},
		{"Charlie", 35, "NYC"},
		{"Diana", 28, "Chicago"},
	}

	// Create sorting functions using closures
	sortByAge := func(ascending bool) func(i, j int) bool {
		return func(i, j int) bool {
			if ascending {
				return people[i].Age < people[j].Age
			}
			return people[i].Age > people[j].Age
		}
	}

	// Multi-field sorting with closure
	sortByMultiple := func() func(i, j int) bool {
		return func(i, j int) bool {
			// First by city, then by age
			if people[i].City != people[j].City {
				return people[i].City < people[j].City
			}
			return people[i].Age < people[j].Age
		}
	}

	fmt.Println("\n--- Sorting with Closures ---")
	fmt.Println("Original:", people)

	sort.Slice(people, sortByAge(true))
	fmt.Println("By age (ascending):", people)

	sort.Slice(people, sortByAge(false))
	fmt.Println("By age (descending):", people)

	sort.Slice(people, sortByMultiple())
	fmt.Println("By city then age:", people)
}

// 8. Retry Logic with Exponential Backoff
func retryWithBackoff(maxAttempts int) func(func() error) error {
	return func(operation func() error) error {
		var lastErr error
		backoff := 100 * time.Millisecond

		for attempt := 1; attempt <= maxAttempts; attempt++ {
			fmt.Printf("Attempt %d/%d\n", attempt, maxAttempts)

			if err := operation(); err == nil {
				fmt.Println("Success!")
				return nil
			} else {
				lastErr = err
				fmt.Printf("Failed: %v\n", err)

				if attempt < maxAttempts {
					fmt.Printf("Waiting %v before retry...\n", backoff)
					time.Sleep(backoff)
					backoff *= 2 // Exponential backoff
				}
			}
		}

		return fmt.Errorf("all %d attempts failed: %w", maxAttempts, lastErr)
	}
}

// 9. Pipeline/Chain of Transformations
func pipeline(funcs ...func(int) int) func(int) int {
	return func(x int) int {
		result := x
		for _, fn := range funcs {
			result = fn(result)
		}
		return result
	}
}

// 10. Debounce/Throttle Pattern
func debounce(fn func(), delay time.Duration) func() {
	var timer *time.Timer
	var mu sync.Mutex

	return func() {
		mu.Lock()
		defer mu.Unlock()

		if timer != nil {
			timer.Stop()
		}

		timer = time.AfterFunc(delay, fn)
	}
}

func throttle(fn func(), limit time.Duration) func() {
	var lastCall time.Time
	var mu sync.Mutex

	return func() {
		mu.Lock()
		defer mu.Unlock()

		now := time.Now()
		if now.Sub(lastCall) >= limit {
			fn()
			lastCall = now
		}
	}
}

// Demo function for advanced closure patterns
func advancedClosuresExample() {
	fmt.Println("\n=== Advanced Closure Examples ===")

	// 1. Middleware
	fmt.Println("\n--- Middleware Pattern ---")
	process := middleware()
	process("hello")

	// 2. Rate Limiter
	fmt.Println("\n--- Rate Limiter ---")
	limiter := rateLimiter(3, 1*time.Second)
	for i := 0; i < 5; i++ {
		if limiter() {
			fmt.Printf("Request %d: Allowed\n", i+1)
		} else {
			fmt.Printf("Request %d: Rate limited\n", i+1)
		}
	}

	// 3. Memoization
	fmt.Println("\n--- Memoization ---")
	memoFib := memoize(fibonacciRecursive)
	fmt.Println(memoFib(10))
	fmt.Println(memoFib(10)) // Cache hit
	fmt.Println(memoFib(15))

	// 4. Event Emitter
	fmt.Println("\n--- Event Emitter ---")
	on, emit := createEventEmitter()
	on("login", func() { fmt.Println("User logged in!") })
	on("login", func() { fmt.Println("Send welcome email") })
	on("logout", func() { fmt.Println("User logged out!") })
	emit("login")
	emit("logout")

	// 5. Generators
	fmt.Println("\n--- Generators ---")
	fibGen := fibonacciGenerator()
	fmt.Print("First 10 Fibonacci numbers: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fibGen())
	}
	fmt.Println()

	primeGen := primeGenerator()
	fmt.Print("First 10 Prime numbers: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", primeGen())
	}
	fmt.Println()

	// 6. Query Builder
	fmt.Println("\n--- Query Builder ---")
	newQuery := createQueryBuilder()
	query := newQuery("users").
		Where("age > 18").
		Where("city = 'NYC'").
		Limit(10).
		Build()
	fmt.Println("Query:", query)

	// 7. Sorting
	sortWithClosures()

	// 8. Retry with Backoff
	fmt.Println("\n--- Retry with Backoff ---")
	retry := retryWithBackoff(3)
	attempts := 0
	err := retry(func() error {
		attempts++
		if attempts < 3 {
			return fmt.Errorf("simulated failure")
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Final error: %v\n", err)
	}

	// 9. Pipeline
	fmt.Println("\n--- Pipeline ---")
	double := func(x int) int { return x * 2 }
	addTen := func(x int) int { return x + 10 }
	square := func(x int) int { return x * x }

	transform := pipeline(double, addTen, square)
	result := transform(5) // (5*2 + 10)^2 = 400
	fmt.Printf("Pipeline(5) = %d\n", result)

	// 10. Debounce/Throttle
	fmt.Println("\n--- Debounce/Throttle ---")

	saveAction := func() {
		fmt.Printf("Saved at %v\n", time.Now().Format("15:04:05.000"))
	}

	debouncedSave := debounce(saveAction, 500*time.Millisecond)
	throttledSave := throttle(saveAction, 1*time.Second)

	fmt.Println("Debounced calls (only last executes):")
	for i := 0; i < 3; i++ {
		debouncedSave()
		time.Sleep(100 * time.Millisecond)
	}
	time.Sleep(600 * time.Millisecond) // Wait for debounce to execute

	fmt.Println("\nThrottled calls (rate limited):")
	for i := 0; i < 5; i++ {
		throttledSave()
		time.Sleep(300 * time.Millisecond)
	}
}
