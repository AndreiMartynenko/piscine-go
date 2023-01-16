package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	al := "abcdefghijklmnopqrstuvwxyz"
	alCap := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	args := os.Args

	if len(args) <= 1 {
		return
	}

	isCapital := false
	if len(args) > 2 && args[1] == "--upper" {
		isCapital = true
	}

	for i := 1; i < len(args); i++ {
		sl := []rune(args[i])

		index := -1

		if len(sl) == 1 {
			index = -2
			ch := sl[0]
			if ch >= '1' && ch <= '9' {
				index = int(ch) - 48
			}
		}

		if len(sl) == 2 {
			index = -2
			ch1 := sl[0]
			ch2 := sl[1]
			if (ch1 == '1' || ch1 == '2') && (ch2 >= '0' && ch2 <= '9') {
				index = (int(ch1)-48)*10 + (int(ch2) - 48)
			}
		}
		if len(sl) > 2 {
			index = -2
		}

		if index > 26 {
			index = -2
		}

		if index > 0 {
			ch := al[index-1]

			if isCapital {
				ch = alCap[index-1]
			}
			z01.PrintRune(rune(ch))
		}
		if index == -2 && !isCapital {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
