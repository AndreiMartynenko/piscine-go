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
	for i := 1; i < len(args); i++ {
		sl := []rune(args[i])
		for j := 0; j < len(sl); j++ {
			ch := sl[j]
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
