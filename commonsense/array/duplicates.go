package array

func HasDuplicates(array []int) bool {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if i != j && array[i] == array[j] {
				return true
			}
		}
	}
	return false
}
