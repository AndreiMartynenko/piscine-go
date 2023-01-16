package piscine

/*
func ConvertBase(nbr, baseFrom, baseTo string) string {
	decimal := ConvertToDecimal(nbr, baseFrom)
	return MyConvertByBase(decimal, baseTo)
}

func ConvertToDecimal(nbr string, base string) int {

	if base == "0123456789" {
		return MyRunesToNumber([]rune(nbr), false)
	}

	slRunes := []rune(nbr)

	pow := 0
	resultInt := 0
	baseSize := len(base)
	for i := len(slRunes) - 1; i >= 0; i-- {
		ch := slRunes[i]
		digit := MyIndex(base, string(ch))

		if pow == 0 {
			resultInt = resultInt + digit
			pow = 1
			continue
		}
		if pow == 1 {
			resultInt = resultInt + digit*baseSize
			pow = 2
			continue
		}

		multiplier := baseSize
		for p := 0; p < pow-1; p++ {
			multiplier = multiplier * baseSize
		}

		resultInt = resultInt + digit*multiplier

		pow = pow + 1

	}
	return resultInt
}

func MyIndex(s string, toFind string) int {
	sliceS := []rune(s)
	sliceToFind := []rune(toFind)

	for i := 0; i < len(s); i++ {
		isFound := true
		for iToFind := 0; iToFind < len(toFind); iToFind++ {
			newIndex := i + iToFind
			if newIndex < len(sliceS) {
				if sliceS[newIndex] != sliceToFind[iToFind] {
					isFound = false
				}
			}
		}
		if isFound {
			return i
		}
	}
	return -1
}

func MyIntToSliceRune(nb int) []rune {
	reversed := []rune{}

	for {
		digit := nb % 10
		nb = nb / 10
		reversed = append(reversed, rune(digit+48))
		if nb == 0 {
			break
		}
	}

	result := []rune{}
	for i := len(reversed) - 1; i >= 0; i-- {
		result = append(result, reversed[i])
	}
	return result
}

func MyConvertByBase(nbr int, base string) string {

	baseSorted := make([]rune, len([]rune(base)))
	copy(baseSorted, []rune(base))

	baseRunes := []rune(base)
	mods := []int{}

	for {

		mod := nbr % len(base)
		div := nbr / len(base)
		nbr = div
		mods = append(mods, mod)
		if div == 0 {
			break
		}
	}

	result := ""

	if len(mods) > 0 {
		for i := len(mods) - 1; i >= 0; i-- {
			ch := baseRunes[mods[i]]
			result = result + string(ch)
		}
	}

	return result
}

func MyRunesToNumber(sliceRunes []rune, isNegative bool) int {
	multiplier := 1

	result := 0
	for i := len(sliceRunes) - 1; i >= 0; i-- {
		num := int(sliceRunes[i] - 48)
		result = result + num*multiplier
		multiplier = multiplier * 10
	}
	if isNegative {
		return -1 * result
	}
	return result
}
*/
