package recursion

func doubleNumbersInArray(array []int) []int {

	doubleNumbersInArrayRec(array, 0)
	return array
}

func doubleNumbersInArrayRec(array []int, count int) {
	if count >= len(array) {
		return
	}
	array[count] *= 2
	doubleNumbersInArrayRec(array, count+1)
}
