package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}
	file := os.Args[1]
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Print(string(contents))
}
