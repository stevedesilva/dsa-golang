package sort

func InsertionSort(array []int) []int {
	// 1 select n + 1
	// 2 take select out of array
	// 3 compare n + 1 with left
	// 4 if value less than selected number then swap - else move to next value in array
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
