package main

import (
	"github.com/01-edu/z01"
)

func repeatAlpha(s string) {
	letter := []rune(s)
	count := 0

	for i := 0; i < len(letter); i++ {
		count = int(letter[i] - 96)
		for j := 0; j < count; j++ {
			z01.PrintRune(letter[i])
		}
	}
	z01.PrintRune('\n')
}

func main() {
	str := "abC"
	repeatAlpha(str)
}
