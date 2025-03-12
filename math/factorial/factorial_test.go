package factorial

import "testing"

type FactorialFunc func(int) int

var implementations = map[string]FactorialFunc{
	"Recursive": Recursive,
	"Iterative": Iterative,
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 1},          // 0! = 1
		{1, 1},          // 1! = 1
		{5, 120},        // 5! = 120
		{7, 5040},       // 7! = 5040
		{10, 3628800},   // 10! = 3628800
		{12, 479001600}, // 12! = 479001600
	}

	for name, function := range implementations {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				result := function(test.n)
				if result != test.expected {
					t.Errorf("%s Factorial(%d) = %d, expected %d", name, test.n, result, test.expected)
				}
			}
		})
	}
}
