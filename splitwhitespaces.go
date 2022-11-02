package piscine

func IsRep(s string) bool {
	runes := []rune(s)
	for _, ch := range runes {
		if ch >= 0 && ch <= 32 {
			return true
		}
	}
	return false
}

func SplitWhiteSpaces(s string) []string {
	result := []string{}
	str := ""
	for i := 0; i < len(s); i++ {
		if IsRep(string(s[i])) && len(str) != 0 {
			result = append(result, str)
			str = ""
		}
		if i == len(s)-1 {
			str += string(s[i])
			result = append(result, str)
		}
		if !IsRep(string(s[i])) {
			str += string(s[i])
		}
	}
	return result
}
