package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	if len(base) <= 1 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	for _, v := range base {
		if v == '-' || v == '+' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
	}

	baseSorted := make([]rune, len([]rune(base)))
	copy(baseSorted, []rune(base))
	baseSorted = SortRunesAsc(baseSorted)

	for i := 1; i < len(baseSorted); i++ {
		if baseSorted[i-1] == baseSorted[i] {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
	}

	baseRunes := []rune(base)
	mods := []int{}

	isNegative := false
	if nbr < 0 {
		isNegative = true
	}

	for {

		mod := nbr % len(base)
		if mod < 0 {
			mod = -mod
		}

		div := nbr / len(base)
		if div < 0 {
			div = -div
		}

		nbr = div
		mods = append(mods, mod)
		if div == 0 {
			break
		}
	}

	if len(mods) > 0 {
		if isNegative {
			z01.PrintRune('-')
		}
		for i := len(mods) - 1; i >= 0; i-- {
			ch := baseRunes[mods[i]]
			z01.PrintRune(ch)
		}
	}
}
