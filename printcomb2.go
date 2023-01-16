package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for i1 := '0'; i1 <= '9'; i1++ {
		for i2 := '0'; i2 <= '9'; i2++ {
			for i3 := '0'; i3 <= '9'; i3++ {
				for i4 := '0'; i4 <= '9'; i4++ {
					flag := false

					if i3 > i1 {
						flag = true
					}
					if i1 == i3 && i4 > i2 {
						flag = true
					}
					if flag {

						//-----------------Printing--------------
						z01.PrintRune(i1)
						z01.PrintRune(i2)
						z01.PrintRune(' ')
						z01.PrintRune(i3)
						z01.PrintRune(i4)
						if i1 == '9' && i2 == '8' && i3 == '9' && i4 == '9' {
							z01.PrintRune('\n')
							return
						}
						z01.PrintRune(',')
						z01.PrintRune(' ')

						//------------------------------

					}
				}
			}
		}
	}
}
