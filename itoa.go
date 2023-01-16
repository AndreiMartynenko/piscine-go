package piscine

func Itoa(n int) string {
	slDigitsRev := []int{}
	isNegative := n < 0
	for {
		digit := n % 10
		n = n / 10
		slDigitsRev = append(slDigitsRev, digit)
		if n == 0 {
			break
		}
	}
	slRunes := []rune{}
	if isNegative {
		slRunes = append(slRunes, '-')
	}
	for i := len(slDigitsRev) - 1; i >= 0; i-- {
		ch := slDigitsRev[i] + 48
		if isNegative {
			ch = -slDigitsRev[i] + 48
		}
		slRunes = append(slRunes, rune(ch))
	}
	return string(slRunes)
}
