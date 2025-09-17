package calculator

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 5},
		{-1, 1, 0},
		{0, 0, 0},
		{100, 200, 300},
	}

	for _, test := range tests {
		result := Add(test.a, test.b)
		if result != test.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{5, 3, 2},
		{1, 1, 0},
		{0, 5, -5},
		{10, 15, -5},
	}

	for _, test := range tests {
		result := Subtract(test.a, test.b)
		if result != test.expected {
			t.Errorf("Subtract(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 6},
		{-2, 3, -6},
		{0, 5, 0},
		{4, 4, 16},
	}

	for _, test := range tests {
		result := Multiply(test.a, test.b)
		if result != test.expected {
			t.Errorf("Multiply(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b     float64
		expected float64
		hasError bool
	}{
		{6.0, 2.0, 3.0, false},
		{5.0, 2.0, 2.5, false},
		{10.0, 0.0, 0.0, true}, // division by zero
		{-8.0, 4.0, -2.0, false},
	}

	for _, test := range tests {
		result, err := Divide(test.a, test.b)

		if test.hasError {
			if err == nil {
				t.Errorf("Divide(%f, %f) expected error but got none", test.a, test.b)
			}
		} else {
			if err != nil {
				t.Errorf("Divide(%f, %f) unexpected error: %v", test.a, test.b, err)
			}
			if result != test.expected {
				t.Errorf("Divide(%f, %f) = %f; want %f", test.a, test.b, result, test.expected)
			}
		}
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		base, exp, expected int
	}{
		{2, 3, 8},
		{5, 0, 1},
		{3, 2, 9},
		{10, 1, 10},
	}

	for _, test := range tests {
		result := Power(test.base, test.exp)
		if result != test.expected {
			t.Errorf("Power(%d, %d) = %d; want %d", test.base, test.exp, result, test.expected)
		}
	}
}

// Test unexported functions (only accessible within package)
func TestMultiplyUnexported(t *testing.T) {
	result := multiply(4, 5)
	if result != 20 {
		t.Errorf("multiply(4, 5) = %d; want 20", result)
	}
}

func TestPowerUnexported(t *testing.T) {
	result := power(2, 4)
	if result != 16 {
		t.Errorf("power(2, 4) = %d; want 16", result)
	}
}

func TestCalculator(t *testing.T) {
	calc := NewCalculator()

	// Test basic operations
	result1 := calc.Add(5, 3)
	if result1 != 8 {
		t.Errorf("Calculator.Add(5, 3) = %d; want 8", result1)
	}

	result2 := calc.Multiply(4, 2)
	if result2 != 8 {
		t.Errorf("Calculator.Multiply(4, 2) = %d; want 8", result2)
	}

	// Test history
	history := calc.GetHistory()
	if len(history) != 2 {
		t.Errorf("Calculator history length = %d; want 2", len(history))
	}

	// Check first operation
	if history[0].Type != "add" || history[0].A != 5 || history[0].B != 3 || history[0].Result != 8 {
		t.Errorf("First operation incorrect: %+v", history[0])
	}

	// Check second operation
	if history[1].Type != "multiply" || history[1].A != 4 || history[1].B != 2 || history[1].Result != 8 {
		t.Errorf("Second operation incorrect: %+v", history[1])
	}

	// Test clear history
	calc.ClearHistory()
	history = calc.GetHistory()
	if len(history) != 0 {
		t.Errorf("Calculator history after clear = %d; want 0", len(history))
	}
}
