package factorial

func Iterative(n int) int {
	if n < 0 {
		return 0
	}

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}

	return result
}

func Recursive(n int) int {
	if n < 0 {
		return 0
	}

	if n <= 1 {
		return 1
	}

	return n * Recursive(n-1)
}
