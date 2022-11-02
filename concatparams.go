package piscine

func ConcatParams(args []string) string {
	var str string
	for i, v := range args {
		str = str + v
		if i < len(args)-1 {
			str = str + "\n"
		}
	}
	return str
}
