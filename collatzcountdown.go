package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	if start == 1 {
		return 0
	}
	term := start
	counter := 0
	for {
		if term%2 == 0 {
			term = term / 2
		} else {
			term = term*3 + 1
		}
		counter++
		if term == 1 {
			return counter
		}
	}
}
