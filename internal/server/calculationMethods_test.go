package server

import "testing"

func TestFibonacciCalculation(t *testing.T) {
	calculator := FibonacciCalculatorData{}

	tests := []struct {
		input    int
		expected int
	}{
		{input: 0, expected: 0},
		{input: 1, expected: 1},
		{input: 2, expected: 1},
		{input: 5, expected: 5},
		{input: 10, expected: 55},
		{input: 20, expected: 6765},
	}

	for _, test := range tests {
		result := calculator.FibonacciCalculation(test.input)
		if result != test.expected {
			t.Errorf("FibonacciCalculation(%d) = %d; want %d", test.input, result, test.expected)
		}
	}
}
