package piscine

func BasicAtoi2(s string) int {
	byteSlice := []byte(s)

	// Reverse array
	reverseByteSlice := []byte{}

	for i := len(byteSlice) - 1; i >= 0; i-- {
		b := byteSlice[i]
		if !(b >= 48 && b <= 57) {
			return 0
		}
		reverseByteSlice = append(reverseByteSlice, byteSlice[i])
	}

	multilplier := 1

	result := 0
	for _, ch := range reverseByteSlice {
		code := ch - 48
		result = result + multilplier*int(code)
		multilplier = multilplier * 10
	}
	return result
}
