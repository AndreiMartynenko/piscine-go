package piscine

import (
	"github.com/01-edu/z01"
)

func ToUpperRune(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		return ch - 32
	}
	return ch
}

func ToLowerRune(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return ch + 32
	}
	return ch
}

func IsAlphaNumeric(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= 48 && ch <= 57)
}

func IsNumericRune(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func RunesToNumber(sliceRunes []rune, isNegative bool) int {
	multiplier := 1

	result := 0
	for i := len(sliceRunes) - 1; i >= 0; i-- {
		num := int(sliceRunes[i] - 48)
		result = result + num*multiplier
		multiplier = multiplier * 10
	}
	if isNegative {
		return -1 * result
	}
	return result
}

func SortRunesAsc(sliceInput []rune) []rune {
	if len(sliceInput) <= 1 {
		return sliceInput
	}
	isSorted := false
	for !isSorted {
		isSorted = true
		for index := 0; index < len(sliceInput)-1; index++ {
			item1 := sliceInput[index]
			item2 := sliceInput[index+1]
			if item1 > item2 {
				isSorted = false
				temp := sliceInput[index]
				temp1 := sliceInput[index+1]
				sliceInput[index] = temp1
				sliceInput[index+1] = temp
			}
		}
	}
	return sliceInput
}

func IntToSliceRune(nb int) []rune {
	reversed := []rune{}

	for {
		digit := nb % 10
		nb = nb / 10
		reversed = append(reversed, rune(digit+48))
		if nb == 0 {
			break
		}
	}

	result := []rune{}
	for i := len(reversed) - 1; i >= 0; i-- {
		result = append(result, reversed[i])
	}
	return result
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

func CompareStrings(a, b string) int {
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
