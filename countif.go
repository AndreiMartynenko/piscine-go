package piscine

func CountIf(f func(string) bool, tab []string) int {
	el := 0
	for _, s := range tab {
		if f(s) == true {
			el++
		}
	}
	return el
}
