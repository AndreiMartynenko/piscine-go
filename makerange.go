package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return []int(nil)
	}
	intS := make([]int, max-min)
	for i := 0; i < len(intS); i++ {
		intS[i] = min + i
	}
	return intS
}
