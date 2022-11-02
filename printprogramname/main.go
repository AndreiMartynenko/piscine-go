package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Myprogramname := []rune(os.Args[0])
	j := 0
	for i := (len(Myprogramname) - 1); i >= 0; i-- {
		if Myprogramname[i] == '/' {
			break
		} else {
			j++
		}
	}
	for starter := (len(Myprogramname) - j); starter < (starter + j); starter++ {
		z01.PrintRune(Myprogramname[starter])

		j--

	}
	z01.PrintRune('\n')
}
