package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	number := n
	if n < 0 {
		z01.PrintRune('-')
		number = -n
	}

	if n == 0 {
		z01.PrintRune('0')
		return
	}

	sliceOfRunes := []rune{}

	for {

		d := number % 10

		if d < 0 {
			d = -d
		}

		if number == 0 {
			break
		}
		number = (number - d) / 10

		digitAsCode := d + 48

		digitAsRune := rune(digitAsCode)

		sliceOfRunes = append(sliceOfRunes, digitAsRune)
	}

	if len(sliceOfRunes) > 0 {
		for index := len(sliceOfRunes) - 1; index >= 0; index-- {
			z01.PrintRune(sliceOfRunes[index])
		}
	}
}
