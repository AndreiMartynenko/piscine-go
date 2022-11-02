package piscine

import "fmt"

func QuadB(x, y int) {
	height := y
	width := x

	for h := 1; h <= height; h++ {
		for w := 1; w <= width; w++ {
			if h >= 2 && h < height && w == 1 || h >= 2 && h < height && w == width {
				fmt.Print("*")
			} else if w == h && w == 1 {
				fmt.Print("/")
			} else if w == width && h == 1 {
				fmt.Print("\\")
			} else if w == width && h == height {
				fmt.Print("/")
			} else if w == 1 || h == width {
				fmt.Print("\\")
			} else if h == 1 || h == height {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
