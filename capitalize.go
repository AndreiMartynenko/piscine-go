package piscine

func Capitalize(s string) string {
	Srune := []rune(s)
	New_Count := 1

	for i := 0; i < len(Srune); i++ {
		if New_Count == 1 {
			if (Srune[i] >= 48 && Srune[i] <= 57) ||
				(Srune[i] >= 65 && Srune[i] <= 90) {
				New_Count = -1
				continue
			} else if Srune[i] >= 97 && Srune[i] <= 122 {
				New_Count = -1
				Srune[i] -= 32
				continue
			}
		}

		if New_Count == -1 {
			if !(Srune[i] >= 48 && Srune[i] <= 57) &&
				!(Srune[i] >= 97 && Srune[i] <= 122) &&
				!(Srune[i] >= 65 && Srune[i] <= 90) {
				New_Count = 1
				continue
			} else if Srune[i] >= 65 && Srune[i] <= 90 {
				New_Count = -1
				Srune[i] += 32
				continue
			}
		}

		if !(Srune[i] >= 48 && Srune[i] <= 57) &&
			!(Srune[i] >= 97 && Srune[i] <= 122) &&
			!(Srune[i] >= 65 && Srune[i] <= 90) {
			New_Count = 1
			continue
		}
	}
	s = string(Srune)
	return s
}
