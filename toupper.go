package piscine

func ToUpper(s string) string {
	Srune := []rune(s)
	for i, v := range Srune {
		if v >= 97 && v <= 122 {
			Srune[i] -= 32
		}
	}
	s = string(Srune)
	return s
}
