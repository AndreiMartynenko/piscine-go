package piscine

func Split(s, sep string) []string {
	result := []string{}

	for {
		index := Index(s, sep)

		if index == -1 {
			break
		}
		slRunes := []rune(s)
		word := string(slRunes[:index])
		result = append(result, word)
		startingIndex := index + len(sep)
		s = string(slRunes[startingIndex:])
	}

	result = append(result, s)

	return result
}
