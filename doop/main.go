package main

import (
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		return
	}

	nb1, err := strconv.Atoi(args[0])
	if err != nil {
		return
	}

	nb2, err2 := strconv.Atoi(args[2])
	if err2 != nil {
		return
	}

	if len(args[1]) != 1 {
		return
	}

	switch args[1] {

	case "+":
		result := nb1 + nb2

		if nb1 > 0 && nb2 > 0 && result < 0 {
			return
		}
		if nb1 < 0 && nb2 < 0 && result > 0 {
			return
		}
		PrintNbr(result)

	case "-":

		result := nb1 - nb2

		if nb1 < 0 && nb2 > 0 && result > 0 {
			return
		}

		if nb1 > 0 && nb2 < 0 && result < 0 {
			return
		}
		PrintNbr(result)

	case "*":

		result := nb1 * nb2

		if nb1 != 0 && nb2 != 0 {
			if nb1 != result/nb2 {
				return
			}
		}
		PrintNbr(result)

	case "/":
		if nb2 == 0 {
			PrintStr("No division by 0")
			return
		}
		PrintNbr(nb1 / nb2)
	case "%":
		if nb2 == 0 {
			PrintStr("No modulo by 0")
			return
		}
		PrintNbr(nb1 % nb2)
	}
}

func PrintNbr(n int) {
	number := n

	toPrint := []rune{}

	if n < 0 {
		toPrint = append(toPrint, '-')
		number = -n
	}

	if n == 0 {
		PrintStr("0")
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
			toPrint = append(toPrint, sliceOfRunes[index])
		}
		PrintStr(string(toPrint))
	}
}

func PrintStr(s string) {
	os.Stdout.Write([]byte(s))
	os.Stdout.Write([]byte("\n"))
}
