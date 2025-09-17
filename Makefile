# Go Learning Guide - Makefile
# Standard Go tooling for formatting, linting, and testing

.PHONY: help fmt lint test check clean install-tools

# Default target
help:
	@echo "Available targets:"
	@echo "  fmt          - Format all Go code with goimports"
	@echo "  lint         - Run golangci-lint on all packages"
	@echo "  test         - Run all tests"
	@echo "  check        - Run fmt, lint, and test (CI pipeline)"
	@echo "  clean        - Clean temporary files"
	@echo "  install-tools - Install required tools (goimports, golangci-lint)"
	@echo "  help         - Show this help message"

# Format all Go code
fmt:
	@echo "🔧 Formatting Go code with goimports..."
	@if command -v goimports >/dev/null 2>&1; then \
		goimports -w .; \
		echo "✅ Code formatted successfully"; \
	else \
		echo "❌ goimports not found. Run 'make install-tools' first"; \
		exit 1; \
	fi

# Run linters
lint:
	@echo "🔍 Running golangci-lint..."
	@golangci-lint run --config .golangci.yml --fix && echo "✅ Linting completed" || (echo "❌ golangci-lint failed" && exit 1)
# Run tests
test:
	@echo "🧪 Running tests..."
	@go test ./... -v
	@echo "✅ Tests completed"

# Run all checks (for CI)
check: fmt lint test
	@echo "✅ All checks passed!"

# Clean temporary files
clean:
	@echo "🧹 Cleaning temporary files..."
	@find . -name "*.tmp" -delete
	@find . -name "temp.txt" -delete
	@go clean -testcache
	@echo "✅ Cleanup completed"

# Install required tools
install-tools:
	@echo "📦 Installing Go tools..."
	@echo "Installing goimports..."
	@go install golang.org/x/tools/cmd/goimports@latest
	@echo "Installing golangci-lint..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "✅ Tools installed successfully"
	@echo ""
	@echo "Verify installation:"
	@echo "  goimports version: $$(goimports --help 2>&1 | head -1 || echo 'not found')"
	@echo "  golangci-lint version: $$(golangci-lint --version 2>/dev/null || echo 'not found')"

# Run go vet on all packages
vet:
	@echo "🔍 Running go vet..."
	@go vet ./...
	@echo "✅ go vet completed"

# Show formatting diff without applying changes
fmt-check:
	@echo "🔍 Checking formatting..."
	@if command -v goimports >/dev/null 2>&1; then \
		if [ -n "$$(goimports -l .)" ]; then \
			echo "❌ Files need formatting:"; \
			goimports -l .; \
			exit 1; \
		else \
			echo "✅ All files are properly formatted"; \
		fi \
	else \
		echo "❌ goimports not found. Run 'make install-tools' first"; \
		exit 1; \
	fi

# Quick format and test for development
dev-check: fmt test
	@echo "✅ Development checks passed!"

# Build all examples to ensure they compile
build:
	@echo "🔨 Building all examples..."
	@for dir in 01-basics 02-variables 03-control-flow 04-functions 05-structs 06-interfaces 07-concurrency 08-error-handling 09-packages 10-advanced; do \
		if [ -d "$$dir" ]; then \
			echo "Building $$dir..."; \
			(cd "$$dir" && go build .) || exit 1; \
		fi \
	done
	@echo "✅ All examples built successfully"

# Run a specific chapter's examples
run-chapter:
	@if [ -z "$(CHAPTER)" ]; then \
		echo "❌ Please specify CHAPTER=XX (e.g., make run-chapter CHAPTER=03)"; \
		exit 1; \
	fi
	@if [ ! -d "$(CHAPTER)-"* ]; then \
		echo "❌ Chapter $(CHAPTER) not found"; \
		exit 1; \
	fi
	@for dir in $(CHAPTER)-*; do \
		if [ -d "$$dir" ]; then \
			echo "🚀 Running examples in $$dir..."; \
			(cd "$$dir" && go run .); \
		fi \
	done
