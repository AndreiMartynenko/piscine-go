package piscine

func ActiveBits(n int) int {
	counter := 0
	for {
		mod := n % 2
		if mod == 1 {
			counter++
		}
		n = n / 2
		if n == 0 {
			break
		}
	}
	return counter
}
