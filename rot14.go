package piscine

func Rot14(s string) string {
	Srune := []rune(s)
	for i, l := range Srune {
		if l >= 'a' && l <= 'z' {
			Srune[i] = Srune[i] + 14

			if Srune[i] > 'z' {
				Srune[i] = Srune[i] - 26
			}
		}
		if l >= 'A' && l <= 'Z' {
			Srune[i] = Srune[i] + 14

			if Srune[i] > 'Z' {
				Srune[i] = Srune[i] - 26
			}
		}
		s = string(Srune)
	}
	return s
}
