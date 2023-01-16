package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for _, a := range args {
		if a == "01" || a == "galaxy" || a == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
