package recursion

func sumArrayStartingAtIndexZero(array []int) int {
	return sumArrayStartingAtIndexZeroRec(array, 0)
}

func sumArrayStartingAtIndexZeroRec(array []int, total int) int {
	if len(array) == 1 {
		return array[0]
	}
	return array[0] + sumArrayStartingAtIndexZeroRec(array[1:], total)
}
