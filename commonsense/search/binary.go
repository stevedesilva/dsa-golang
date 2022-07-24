package search

func FindInSortedArray(valueToFind int, array []int) (int, error) {
	if len(array) <= 0 {
		return 0, NotFound
	}
	top := len(array) - 1
	bottom := 0

	println("len : ", len(array))
	for bottom <= top {
		mid := (top + bottom) / 2
		valueAtMid := array[mid]
		if valueToFind == valueAtMid {
			return mid, nil
		} else if valueToFind < valueAtMid {
			top = mid - 1
		} else if valueToFind > valueAtMid {
			bottom = mid + 1
		}
	}
	return 0, NotFound
}
