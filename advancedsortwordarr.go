package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	SortStringsByFunction(a, f)
}

func SortStringsByFunction(input []string, f func(a, b string) int) {
	if len(input) < 2 {
		return
	}
	isSorted := false
	for !isSorted {
		isSorted = true
		for index := 0; index < len(input)-1; index++ {
			s1 := input[index]
			s2 := input[index+1]
			if f(s1, s2) == 1 {
				isSorted = false
				temp := s1
				input[index] = input[index+1]
				input[index+1] = temp
			}
		}
	}
}

func CompareStringsQuest9(a, b string) int {
	sliceA := []rune(a)
	sliceB := []rune(b)
	smallerLenth := len(sliceA)
	if len(sliceB) < smallerLenth {
		smallerLenth = len(sliceB)
	}
	for i := 0; i < smallerLenth; i++ {
		if sliceA[i] > sliceB[i] {
			return -1
		}
		if sliceA[i] < sliceB[i] {
			return 1
		}
	}
	if len(sliceA) > len(sliceB) {
		return -1
	}
	if len(sliceA) < len(sliceB) {
		return 1
	}
	return 0
}
