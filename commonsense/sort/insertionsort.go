package sort

/*
InsertionSort

	For every element in the array:
	for current element
	sort everything to the left of the element
*/
func InsertionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		value := array[i]
		leastPos := i
		for j := i - 1; j >= 0; j-- {
			if array[j] > value {
				array[leastPos] = array[j]
				leastPos = j
			} else {
				break
			}
		}
		if i != leastPos {
			array[leastPos] = value
		}

	}
	return array
}
