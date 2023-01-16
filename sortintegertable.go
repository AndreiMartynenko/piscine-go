package piscine

func SortIntegerTable(table []int) {
	isSorted := false

	if len(table) <= 1 {
		return
	}

	for !isSorted {
		isSorted = true

		for index := 0; index < len(table)-1; index++ {
			item1 := table[index]
			item2 := table[index+1]
			if item1 > item2 {
				isSorted = false
				temp := table[index]
				temp1 := table[index+1]
				table[index] = temp1
				table[index+1] = temp
			}
		}
	}
}
