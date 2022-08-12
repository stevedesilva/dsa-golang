package array

func HasDuplicates(array []int) bool {
	foundValues := make([]bool, len(array))
	for i := 0; i < len(array); i++ {
		for j := i; j < len(array); j++ {
			if foundValues[array[j]] {
				return true
			} else {
				foundValues[array[j]] = true
			}

		}
	}
	return false
}
