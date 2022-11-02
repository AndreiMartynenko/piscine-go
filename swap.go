package piscine

func Swap(a *int, b *int) {
	var swap4a int
	var swap4b int

	swap4a = *a // writing into temporary container value of pointer a
	swap4b = *b

	*a = swap4b // actual swap - writing b value into the pointer for a
	*b = swap4a
}
