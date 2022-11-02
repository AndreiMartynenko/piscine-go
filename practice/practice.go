// func main() {
// aString := "Hello"
// aStringChangeable := []byte(aString)
// aStringChangeable[0] = 'A'
// aStringFinalized := string(aStringChangeable)

// fmt.Println(aString)

// fmt.Println(aStringChangeable)

// fmt.Println(aStringFinalized)

// Range Loop

// 	slice := []string{
// 		"Word1",
// 		"Word2",
// 		"Word3",
// 		"Word4",
// 		"Word5",
// 		"Word6",
// 	}

// 	for index, word := range slice {
// 		fmt.Printf("The index is: %v the element is: %v\n", index, word)
// 		fmt.Printf("the element is: %v\n", word)
// 	}
// }

// Pointer

// func main() {
// 	a := 10

// 	fmt.Println(a)
// 	fmt.Printf("THis is the address of a : %v", &a)
// 	fmt.Println()

// 	var pointer *int = &a

// 	fmt.Println(pointer)
// }

// func main() {
// 	a := 15
// 	b := 10

// 	division := a / b

// 	c := 15
// 	d := 10

// 	module := c % d

// 	fmt.Println(division)
// 	fmt.Println(module)
// }

// Os.Args

package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args

	fmt.Printf("Number of arguments: %v\n", len(arguments))

	if len(arguments) >= 3 {
		fmt.Printf("arguments# 3 is: %v", os.Args[2])
	}
}
