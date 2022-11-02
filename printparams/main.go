package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var word string

	for i := 1; i < len(os.Args); i++ {
		word = os.Args[i]
		for j := 0; j < len(word); j++ {
			z01.PrintRune(rune(word[j]))
		}
		z01.PrintRune('\n')
	}
}
