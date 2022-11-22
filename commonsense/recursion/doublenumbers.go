package recursion

func doubleNumbersInArray(array []int) []int {

	doubleNumbersInArrayRec(array, 0)
	return array
}

func doubleNumbersInArrayRec(array []int, count int) {
	if count >= len(array) {
		return
	}
	val := array[count] * 2
	array[count] = val
	doubleNumbersInArrayRec(array, count+1)
}
