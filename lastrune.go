package piscine

func LastRune(s string) rune {
	sentence := []rune(s)

	positioniwant := len(sentence) - 1

	lastrune := sentence[positioniwant]

	return lastrune
}
