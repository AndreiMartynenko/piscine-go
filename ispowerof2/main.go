package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

const (
	True  = "true"
	False = "false"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	nb, err := strconv.Atoi(args[0])
	if err != nil {
		return
	}
	if nb == 1 {
		PrintStringHachatron(True)
		return
	}

	for {
		mod := nb % 2
		if mod != 0 {
			PrintStringHachatron(False)
			return
		}
		nb = nb / 2
		if nb == 1 {
			PrintStringHachatron(True)
			return
		}
	}
}

func PrintStringHachatron(s string) {
	sl := []rune(s)
	for i := 0; i < len(sl); i++ {
		z01.PrintRune(sl[i])
	}
	z01.PrintRune('\n')
}
