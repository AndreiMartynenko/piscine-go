package piscine

func Sqrt(nb int) int {
	for count := 0; count < 8; count++ {
		if (count * count) == nb {
			return count
		}
	}
	return 0
}
