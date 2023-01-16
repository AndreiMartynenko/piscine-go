package piscine

func Abort(a, b, c, d, e int) int {
	slInput := []int{a, b, c, d, e}
	slSorted := SortIntsHackatron(slInput)
	if len(slSorted)%2 == 1 {
		return slSorted[(len(slSorted)-1)/2]
	} else {
		m1 := slSorted[len(slSorted)/2]
		m2 := slSorted[len(slSorted)/2+1]
		return (m1 + m2) / 2
	}
}

func SortIntsHackatron(input []int) []int {
	if len(input) < 2 {
		return input
	}
	for {
		isSorted := true

		for i := 1; i < len(input); i++ {
			n1 := input[i-1]
			n2 := input[i]
			if n1 > n2 {
				isSorted = false
				temp := n1
				input[i-1] = n2
				input[i] = temp
			}
		}
		if isSorted {
			return input
		}
	}
}
