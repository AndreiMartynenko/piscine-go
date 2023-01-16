package piscine

func Compact(ptr *[]string) int {
	for i := 0; i < len(*ptr); i++ {
		if (*ptr)[i] == "" {
			newSlice := []string{}
			newSlice = append(newSlice, (*ptr)[:i]...)
			newSlice = append(newSlice, (*ptr)[i+1:]...)
			*ptr = newSlice
			i = -1
		}
	}
	return len(*ptr)
}
