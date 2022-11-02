package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}

	var a, b int

	b = 1

	for index--; index > 0; index-- {
		a += b
		a, b = b, a
	}

	return b
}
