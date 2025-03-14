package fibonacci

func Iterative(x int) int {
	if x == 0 {
		return 0
	}

	a, b := 0, 1
	for i := 2; i <= x; i++ {
		a, b = b, a+b
	}

	return b
}

func Recursive(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	}

	return Recursive(x-1) + Recursive(x-2)
}
