package piscine

func StrRev(s string) string {
	runes := []rune(s)
	for forward, reverse := 0, len(runes)-1; forward < reverse; forward, reverse = forward+1, reverse-1 {
		runes[forward], runes[reverse] = runes[reverse], runes[forward]
	}
	return string(runes)
}
