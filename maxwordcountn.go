package piscine

import (
	"fmt"
	"sort"
)

func MaxWordCountN(text string, n int) map[string]int {

	arr := GetWordsHackatron(text)
	sort.Strings(arr)
	fmt.Println(arr)

	return map[string]int{}

}

func GetWordsHackatron(s string) []string {
	result := []string{}

	slRunes := []rune(s)

	for i := 0; i < len(slRunes); i++ {
		ch := slRunes[i]

		if ch == ' ' || ch == 9 || ch == 10 || ch == 13 {
			word := string(slRunes[:i])
			if word != " " && word != "\n" && word != "\r" && word != "\t" && word != "" {
				result = append(result, word)
			}
			slRunes = slRunes[i+1:]
			i = -1
		}
	}

	word := string(slRunes)

	if word != " " && word != "\n" && word != "\r" && word != "\t" && word != "" {
		result = append(result, word)
	}
	return result
}
