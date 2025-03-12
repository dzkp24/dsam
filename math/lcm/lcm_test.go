package lcm

import (
	"fmt"
	"testing"
)

func TestLcm(t *testing.T) {
	tests := []struct {
		a, b     int64
		expected int64
	}{
		{0, 5, 0},         // LCM(0, 5) = 0
		{7, 0, 0},         // LCM(7, 0) = 0
		{3, 5, 15},        // LCM(3, 5) = 15
		{6, 8, 24},        // LCM(6, 8) = 24
		{21, 6, 42},       // LCM(21, 6) = 42
		{48, 18, 144},     // LCM(48, 18) = 144
		{101, 103, 10403}, // LCM(101, 103) = 10403
	}

	for _, test := range tests {
		t.Run(
			fmt.Sprintf("LCM(%d,%d)", test.a, test.b),
			func(t *testing.T) {
				result := Lcm(test.a, test.b)
				if result != test.expected {
					t.Errorf("LCM(%d, %d) = %d, expected %d", test.a, test.b, result, test.expected)
				}
			},
		)
	}
}
