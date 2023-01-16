package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	switch len(args) {

	case 0:
		//reader := os.Stdin
		//data, err := reader.ReadFrom(os.Stdin)
		//if err == nil {
		//	PrintString(string(data))
		//}
		data, err := ioutil.ReadAll(os.Stdin)
		if err == nil {
			PrintString(string(data))
		}

		return

	case 1:
		PrintFiles([]string{args[0]})

	case 2:
		if args[0] == "<<EOF>" {
			data, err := ioutil.ReadAll(os.Stdin)
			if err == nil {

				f, err := os.Create(args[1])
				if err == nil {
					_, err2 := f.WriteString(string(data))
					if err2 != nil {
						os.Exit(1)
					}
				}
				defer f.Close()
			}
			return
		}

		fileNames := []string{}
		for _, fn := range args {
			fileNames = append(fileNames, fn)
		}
		PrintFiles(fileNames)
	}
}

func PrintFiles(fileNames []string) {
	for i := 0; i < len(fileNames); i++ {
		output := []byte{}
		data, err := os.ReadFile(fileNames[i])
		if err != nil {
			PrintString("ERROR: " + err.Error())
			os.Exit(1)
		}
		data = data[:len(data)-1]

		for _, b := range data {
			output = append(output, b)
		}
		PrintString(string(output))
	}
}

func PrintString(s string) {
	slRunes := []rune(s)

	for _, ch := range slRunes {
		z01.PrintRune(ch)
	}
	if len(slRunes) > 0 {
		if slRunes[len(slRunes)-1] != '\n' {
			z01.PrintRune('\n')
		}
	}
}
