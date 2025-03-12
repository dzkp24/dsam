package lcm

import "github.com/dzkp24/dsam/math/gcd"

func Lcm(a, b int64) int64 {
	if a == 0 || b == 0 {
		return 0
	}

	return (a / gcd.Recursive(a, b)) * b
}
