package fibonacci

import "testing"

type FibonacciFunc func(int) int

var implementations = map[string]FibonacciFunc{
	"Recursive": Recursive,
	"Iterative": Iterative,
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"Test Fibonacci 0", 0, 0},
		{"Test Fibonacci 1", 1, 1},
		{"Test Fibonacci 2", 2, 1},
		{"Test Fibonacci 3", 3, 2},
		{"Test Fibonacci 4", 4, 3},
		{"Test Fibonacci 5", 5, 5},
		{"Test Fibonacci 6", 6, 8},
		{"Test Fibonacci 7", 7, 13},
		{"Test Fibonacci 8", 8, 21},
		{"Test Fibonacci 9", 9, 34},
	}

	for name, fn := range implementations {
		for _, tt := range tests {
			t.Run(name+"_"+tt.name, func(t *testing.T) {
				got := fn(tt.input)
				if got != tt.expected {
					t.Errorf("%s: for input %d, expected %d but got %d", name, tt.input, tt.expected, got)
				}
			})
		}
	}
}
