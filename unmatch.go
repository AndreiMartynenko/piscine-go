package piscine

func Unmatch(a []int) int {
	aCopy := []int{}
	for i := 0; i < len(a); i++ {
		aCopy = append(aCopy, a[i])
	}
	if len(aCopy) == 1 {
		return aCopy[0]
	}
	if len(aCopy) == 0 {
		return -1
	}
	for {
		pairFound := false
		for i := 1; i < len(aCopy); i++ {
			if aCopy[0] == aCopy[i] {
				aCopy = RemoveAt(aCopy, i)
				aCopy = aCopy[1:]
				pairFound = true
				break
			}
		}
		if !pairFound {
			return aCopy[0]
		}
		if len(aCopy) == 0 {
			return -1
		}
		if len(aCopy) == 1 {
			return aCopy[0]
		}
	}
}

func RemoveAt(arr []int, i int) []int {
	if i < 0 || i > len(arr) {
		return arr
	}
	if i == len(arr) {
		return arr[:len(arr)-1]
	}
	left := arr[:i]
	right := arr[i+1:]
	aa := []int(left)
	aa = append(aa, right...)
	return aa
}
