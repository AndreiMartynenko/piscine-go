package main

import (
	"fmt"
	"os"
)

func main() {
	toOrder := false
	insertParam := ""
	param := ""

	commandInsert := "--insert="
	commandI := "-i="
	commandOrder := "--order"
	commandO := "-o"
	commandHelp := "--help"
	commandH := "-h"

	args := os.Args

	if len(args) <= 1 {
		PrintHelp()
		return
	}

	for i := 1; i < len(args); i++ {
		if args[i] == commandHelp || args[i] == commandH {
			PrintHelp()
			return
		}
		if args[i] == commandOrder || args[i] == commandO {
			toOrder = true
		}

		indexOfInsert := -1
		indexOfI := -1
		indexOfInsert = Index(args[i], commandInsert)
		indexOfI = Index(args[i], commandI)
		if indexOfInsert == 0 || indexOfI == 0 {
			sl := []rune(args[i])
			if indexOfInsert == 0 {
				insertParam = string(sl[len(commandInsert):])
			}
			if indexOfI == 0 {
				insertParam = string(sl[len(commandI):])
			}
		}
	}

	param = args[len(args)-1]

	output := param + insertParam
	if toOrder {
		output = string(SortRunesAsc([]rune(output)))
	}

	fmt.Println(output)
}

func PrintHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func Index(s string, toFind string) int {
	sliceS := []rune(s)
	sliceToFind := []rune(toFind)

	for i := 0; i < len(s); i++ {
		isFound := true
		for iToFind := 0; iToFind < len(toFind); iToFind++ {
			newIndex := i + iToFind
			if newIndex < len(sliceS) {
				if sliceS[newIndex] != sliceToFind[iToFind] {
					isFound = false
				}
			}
		}
		if isFound {
			return i
		}
	}
	return -1
}

func SortRunesAsc(sliceInput []rune) []rune {
	if len(sliceInput) <= 1 {
		return sliceInput
	}
	isSorted := false
	for !isSorted {
		isSorted = true
		for index := 0; index < len(sliceInput)-1; index++ {
			item1 := sliceInput[index]
			item2 := sliceInput[index+1]
			if item1 > item2 {
				isSorted = false
				temp := sliceInput[index]
				temp1 := sliceInput[index+1]
				sliceInput[index] = temp1
				sliceInput[index+1] = temp
			}
		}
	}
	return sliceInput
}
