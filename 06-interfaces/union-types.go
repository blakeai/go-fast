package main

import (
	"fmt"
	"strconv"
	"time"
)

// Handler Basic Handler interface - replaces EmailHandler | SMSHandler union
type Handler interface {
	Handle() error
}

type EmailHandler struct {
	recipient string
	subject   string
	body      string
}

type SMSHandler struct {
	phoneNumber string
	message     string
}

func (e EmailHandler) Handle() error {
	fmt.Printf("📧 Sending email to %s: %s\n", e.recipient, e.subject)
	time.Sleep(50 * time.Millisecond) // Simulate work
	return nil
}

func (s SMSHandler) Handle() error {
	fmt.Printf("📱 Sending SMS to %s: %s\n", s.phoneNumber, s.message)
	time.Sleep(30 * time.Millisecond) // Simulate work
	return nil
}

// Process any handler - works with both types
func process(h Handler) error {
	return h.Handle()
}

// Type discrimination with type switches
func processWithDetails(h Handler) error {
	// Execute the common behavior first
	if err := h.Handle(); err != nil {
		return err
	}

	// Then handle type-specific logic
	switch handler := h.(type) {
	case EmailHandler:
		fmt.Printf("  ✓ Email delivered to %s\n", handler.recipient)
	case SMSHandler:
		fmt.Printf("  ✓ SMS delivered to %s\n", handler.phoneNumber)
	default:
		fmt.Println("  ? Unknown handler type")
	}

	return nil
}

// More complex example: Payment methods
type PaymentMethod interface {
	Process(amount float64) error
	GetFee(amount float64) float64
}

type CreditCard struct {
	number     string
	expiryDate string
}

type PayPal struct {
	email string
}

type BankTransfer struct {
	accountNumber string
	routingNumber string
}

func (c CreditCard) Process(amount float64) error {
	fmt.Printf("💳 Processing $%.2f via Credit Card ending in %s\n",
		amount, c.number[len(c.number)-4:])
	return nil
}

func (c CreditCard) GetFee(amount float64) float64 {
	return amount * 0.029 // 2.9% fee
}

func (p PayPal) Process(amount float64) error {
	fmt.Printf("🅿️ Processing $%.2f via PayPal (%s)\n", amount, p.email)
	return nil
}

func (p PayPal) GetFee(amount float64) float64 {
	return amount * 0.034 // 3.4% fee
}

func (b BankTransfer) Process(amount float64) error {
	fmt.Printf("🏦 Processing $%.2f via Bank Transfer (%s)\n",
		amount, b.accountNumber)
	return nil
}

func (b BankTransfer) GetFee(amount float64) float64 {
	return 0.50 // Flat $0.50 fee
}

// Process payment with any method
func processPayment(method PaymentMethod, amount float64) error {
	fee := method.GetFee(amount)
	total := amount + fee

	fmt.Printf("Processing payment: Amount=$%.2f, Fee=$%.2f, Total=$%.2f\n",
		amount, fee, total)

	return method.Process(total)
}

// Document types - more complex union type scenario
type Document interface {
	Render() string
	GetMetadata() map[string]string
}

type PDFDocument struct {
	pages  int
	title  string
	author string
}

type HTMLDocument struct {
	content string
	title   string
	charset string
}

type MarkdownDocument struct {
	content string
	title   string
	tags    []string
}

func (p PDFDocument) Render() string {
	return fmt.Sprintf("PDF: %s by %s (%d pages)", p.title, p.author, p.pages)
}

func (p PDFDocument) GetMetadata() map[string]string {
	return map[string]string{
		"type":   "pdf",
		"title":  p.title,
		"author": p.author,
		"pages":  strconv.Itoa(p.pages),
	}
}

func (h HTMLDocument) Render() string {
	return fmt.Sprintf("HTML: %s (charset: %s)", h.title, h.charset)
}

func (h HTMLDocument) GetMetadata() map[string]string {
	return map[string]string{
		"type":    "html",
		"title":   h.title,
		"charset": h.charset,
	}
}

func (m MarkdownDocument) Render() string {
	return fmt.Sprintf("Markdown: %s (tags: %v)", m.title, m.tags)
}

func (m MarkdownDocument) GetMetadata() map[string]string {
	metadata := map[string]string{
		"type":  "markdown",
		"title": m.title,
	}
	for i, tag := range m.tags {
		metadata[fmt.Sprintf("tag_%d", i)] = tag
	}
	return metadata
}

// Process documents with type-specific optimizations
func processDocument(doc Document) {
	fmt.Printf("Document: %s\n", doc.Render())

	// Type-specific processing
	switch d := doc.(type) {
	case PDFDocument:
		if d.pages > 100 {
			fmt.Println("  📄 Large PDF detected - enabling compression")
		}
	case HTMLDocument:
		if d.charset != "UTF-8" {
			fmt.Printf("  🌐 Converting from %s to UTF-8\n", d.charset)
		}
	case MarkdownDocument:
		if len(d.tags) > 0 {
			fmt.Printf("  🏷️ Indexing %d tags\n", len(d.tags))
		}
	}

	// Common processing
	metadata := doc.GetMetadata()
	fmt.Printf("  📋 Metadata: %v\n", metadata)
}

// Error types - another union type use case
type AppError interface {
	Error() string
	Code() int
	IsRetryable() bool
}

type NetworkError struct {
	message string
	url     string
}

type ValidationError struct {
	field   string
	message string
}

type DatabaseError struct {
	query   string
	message string
}

func (n NetworkError) Error() string {
	return fmt.Sprintf("network error: %s (URL: %s)", n.message, n.url)
}

func (n NetworkError) Code() int { return 500 }

func (n NetworkError) IsRetryable() bool { return true }

func (v ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s (field: %s)", v.message, v.field)
}

func (v ValidationError) Code() int { return 400 }

func (v ValidationError) IsRetryable() bool { return false }

func (d DatabaseError) Error() string {
	return fmt.Sprintf("database error: %s (query: %s)", d.message, d.query)
}

func (d DatabaseError) Code() int { return 503 }

func (d DatabaseError) IsRetryable() bool { return true }

// Handle errors with type-specific logic
func handleError(err AppError) {
	fmt.Printf("❌ Error %d: %s\n", err.Code(), err.Error())

	if err.IsRetryable() {
		fmt.Println("  🔄 This error is retryable")
	} else {
		fmt.Println("  🛑 This error is not retryable")
	}

	// Type-specific handling
	switch e := err.(type) {
	case NetworkError:
		fmt.Printf("  🌐 Consider checking network connectivity to %s\n", e.url)
	case ValidationError:
		fmt.Printf("  📝 Fix the '%s' field and try again\n", e.field)
	case DatabaseError:
		fmt.Printf("  💾 Query optimization may be needed: %s\n", e.query)
	}
}

//goland:noinspection SqlNoDataSourceInspection,SqlNoDataSourceInspection
func unionTypesExample() {
	fmt.Println("=== Basic Handler Union Types ===")

	handlers := []Handler{
		EmailHandler{
			recipient: "user@example.com",
			subject:   "Welcome!",
			body:      "Thanks for signing up!",
		},
		SMSHandler{
			phoneNumber: "+1-555-0123",
			message:     "Your verification code is 123456",
		},
	}

	// Process all handlers uniformly
	for _, handler := range handlers {
		err := process(handler)
		if err != nil {
			return
		}
	}

	fmt.Println("\n=== Handler Union Types with Type Discrimination ===")

	for _, handler := range handlers {
		err := processWithDetails(handler)
		if err != nil {
			return
		}
	}

	fmt.Println("\n=== Payment Method Union Types ===")

	paymentMethods := []PaymentMethod{
		CreditCard{number: "4532-1234-5678-9876", expiryDate: "12/25"},
		PayPal{email: "user@example.com"},
		BankTransfer{accountNumber: "123456789", routingNumber: "987654321"},
	}

	amount := 100.0
	for _, method := range paymentMethods {
		err := processPayment(method, amount)
		if err != nil {
			return
		}
		fmt.Println()
	}

	fmt.Println("=== Document Union Types ===")

	documents := []Document{
		PDFDocument{pages: 150, title: "Go Programming Guide", author: "Gopher"},
		HTMLDocument{content: "<html>...</html>", title: "Web Page", charset: "ISO-8859-1"},
		MarkdownDocument{content: "# Title\nContent", title: "README", tags: []string{"docs", "go"}},
	}

	for _, doc := range documents {
		processDocument(doc)
		fmt.Println()
	}

	fmt.Println("=== Error Union Types ===")

	errors := []AppError{
		NetworkError{message: "connection timeout", url: "https://api.example.com"},
		ValidationError{field: "email", message: "invalid format"},
		DatabaseError{query: "SELECT * FROM users", message: "connection pool exhausted"},
	}

	for _, err := range errors {
		handleError(err)
		fmt.Println()
	}

	fmt.Println("=== Interface vs Union Types Summary ===")
	fmt.Println("✅ Go interfaces provide:")
	fmt.Println("  - Implicit satisfaction (no 'implements' keyword)")
	fmt.Println("  - Focus on behavior over data shape")
	fmt.Println("  - Runtime type discrimination via type switches")
	fmt.Println("  - Composition through interface embedding")
	fmt.Println("  - Flexibility and extensibility")
}
