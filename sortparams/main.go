package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		return
	}
	args = SortStrings(args[1:])
	PrintStrings(args)
}

func SortStrings(input []string) []string {
	if len(input) < 2 {
		return input
	}

	isSorted := false

	for !isSorted {
		isSorted = true
		for index := 0; index < len(input)-1; index++ {
			s1 := input[index]
			s2 := input[index+1]
			if Compare(s1, s2) == 1 {
				isSorted = false
				temp := s1
				input[index] = input[index+1]
				input[index+1] = temp
			}
		}
	}
	return input
}

func Compare(a, b string) int {
	sliceA := []rune(a)
	sliceB := []rune(b)

	smallerLenth := len(sliceA)
	if len(sliceB) < smallerLenth {
		smallerLenth = len(sliceB)
	}

	for i := 0; i < smallerLenth; i++ {
		if sliceA[i] > sliceB[i] {
			return 1
		}
		if sliceA[i] < sliceB[i] {
			return -1
		}
	}

	if len(sliceA) > len(sliceB) {
		return 1
	}

	if len(sliceA) < len(sliceB) {
		return -1
	}

	return 0
}

func PrintStrings(input []string) {
	for _, v := range input {

		sl := []rune(v)

		for i := 0; i < len(sl); i++ {
			ch := sl[i]
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
