package main

import (
	"fmt"
	"log"
	"os"

	"go-fast/09-packages-internal/api"
	"go-fast/09-packages-internal/internal/config"
	"go-fast/09-packages-internal/internal/shared"
)

func main() {
	// Demonstrate module-level internal package usage
	configDemo()

	// Demonstrate shared internal utilities
	sharedUtilitiesDemo()

	// Demonstrate API with internal packages
	apiDemo()
}

func configDemo() {
	fmt.Println("=== Internal Config Package Demo ===")

	// Set some environment variables for demonstration
	os.Setenv("DEBUG", "true")
	os.Setenv("API_KEY", "demo-api-key-12345")
	os.Setenv("PORT", "3000")
	defer func() {
		os.Unsetenv("DEBUG")
		os.Unsetenv("API_KEY")
		os.Unsetenv("PORT")
	}()

	// Load configuration using internal config package
	cfg, err := config.Load()
	if err != nil {
		log.Printf("Failed to load config: %v", err)
		return
	}

	fmt.Printf("Loaded config: %s\n", cfg.String())
	fmt.Printf("Is production: %t\n", cfg.IsProduction())
	fmt.Printf("Database URL: %s\n", cfg.DatabaseURL)
	fmt.Printf("Port: %d\n", cfg.Port)

	// Note: The config package is internal, so it can only be imported
	// by packages within this module, not by external modules
}

func sharedUtilitiesDemo() {
	fmt.Println("\n=== Internal Shared Utilities Demo ===")

	// Error wrapping utilities
	originalErr := fmt.Errorf("connection failed")
	wrappedErr := shared.WrapError(originalErr, "database operation")
	fmt.Printf("Wrapped error: %v\n", wrappedErr)

	// Validation error formatting
	validationErr := shared.FormatValidationError("email", "must be valid email address")
	fmt.Printf("Validation error: %v\n", validationErr)

	// Chain multiple errors
	errors := []error{
		fmt.Errorf("error 1"),
		fmt.Errorf("error 2"),
		fmt.Errorf("error 3"),
	}
	chainedErr := shared.ChainErrors(errors)
	fmt.Printf("Chained errors: %v\n", chainedErr)

	// Error with stack trace
	stackErr := shared.ErrorWithStack("something went wrong")
	fmt.Printf("Error with stack: %v\n", stackErr)

	// Note: These shared utilities are internal, so they provide common
	// functionality across the module without exposing implementation details
}

func apiDemo() {
	fmt.Println("\n=== API with Internal Packages Demo ===")

	// Create API server (uses internal auth and validation packages)
	server := api.NewServer()
	fmt.Println("Created API server with internal dependencies")

	// The server internally uses:
	// - api/internal/auth for authentication logic
	// - api/internal/validation for input validation
	// - internal/shared for shared utilities

	fmt.Println("API server demonstrates internal package usage:")
	fmt.Println("  ✓ api/internal/auth - JWT token generation/validation")
	fmt.Println("  ✓ api/internal/validation - input validation rules")
	fmt.Println("  ✓ internal/shared - shared error handling and HTTP utilities")

	// Demonstrate that internal packages are working
	fmt.Println("\nInternal package integration test:")

	// This would normally be done via HTTP requests, but we'll simulate it
	fmt.Println("  - Authentication service initialized")
	fmt.Println("  - Validation service initialized")
	fmt.Println("  - Shared utilities available")
	fmt.Println("  - Server ready to handle requests")

	// In a real application, you would start the server:
	// log.Fatal(server.Start(8080))

	// For demonstration, we'll just show the server is configured
	mux := server.SetupRoutes()
	fmt.Printf("  - Routes configured: %T\n", mux)

	// Cleanup
	server.Cleanup()
	fmt.Println("  - Server cleanup completed")
}

func visibilityDemo() {
	fmt.Println("\n=== Internal Package Visibility Demo ===")

	fmt.Println("✓ Can import module-level internal packages:")
	fmt.Println("  import \"go-fast/09-packages-internal/internal/config\"")
	fmt.Println("  import \"go-fast/09-packages-internal/internal/shared\"")

	fmt.Println("✓ Can import public APIs that use internal packages:")
	fmt.Println("  import \"go-fast/09-packages-internal/api\"")

	fmt.Println("✗ Cannot import API-specific internal packages:")
	fmt.Println("  import \"go-fast/09-packages-internal/api/internal/auth\"     // ILLEGAL")
	fmt.Println("  import \"go-fast/09-packages-internal/api/internal/validation\" // ILLEGAL")

	fmt.Println("✓ External modules could import public API:")
	fmt.Println("  import \"github.com/example/mymodule/api\"  // Would work")

	fmt.Println("✗ External modules cannot import internal packages:")
	fmt.Println("  import \"github.com/example/mymodule/internal/config\"  // ILLEGAL")
	fmt.Println("  import \"github.com/example/mymodule/api/internal/auth\" // ILLEGAL")

	fmt.Println("\nThis ensures:")
	fmt.Println("  - Clean public APIs")
	fmt.Println("  - Implementation details remain private")
	fmt.Println("  - Refactoring safety for internal code")
	fmt.Println("  - Clear architectural boundaries")
}
