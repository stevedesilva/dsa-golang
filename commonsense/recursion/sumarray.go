package recursion

func SumArrayStartingAtZeroIndex(array []int) int {
	return sumArrayStartingAtIndexZeroRec(array, 0)
}

func sumArrayStartingAtIndexZeroRec(array []int, total int) int {
	firstElement := array[0]
	if len(array) == 1 {
		return firstElement
	}
	sliceExcludingFirstElement := array[1:]
	return firstElement + sumArrayStartingAtIndexZeroRec(sliceExcludingFirstElement, total)
}

func SumArrayStartingAtIndexEnd(array []int) int {
	return sumArrayStartingAtIndexEndRec(array, 0)
}

func sumArrayStartingAtIndexEndRec(array []int, total int) int {
	lastElement := array[len(array)-1]
	if len(array) == 1 {
		return lastElement
	}
	sliceExcludingLastElement := array[:len(array)-1]
	return lastElement + sumArrayStartingAtIndexEndRec(sliceExcludingLastElement, total)
}
