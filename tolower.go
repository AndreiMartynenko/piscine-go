package piscine

func ToLower(s string) string {
	Srune := []rune(s)
	for i, v := range Srune {
		if v >= 65 && v <= 90 {
			Srune[i] += 32
		}
	}
	s = string(Srune)
	return s
}
