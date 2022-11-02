package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var cur_arg string

	for i := len(os.Args) - 1; i > 0; i-- {
		cur_arg = os.Args[i]
		for j := 0; j < len(cur_arg); j++ {
			z01.PrintRune(rune(cur_arg[j]))
		}
		z01.PrintRune('\n')
	}
}
