package piscine

func AppendRange(min, max int) []int {
	var intS []int
	for i := 0; i < (max - min); i++ {
		intS = append(intS, i+min)
	}
	return intS
}
