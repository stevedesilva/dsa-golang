package sort

func HasDuplicates(array []int) bool {
	q := NewQuickSort(array)
	q.Quicksort()
	return hasDuplicates(q.array)
}

func hasDuplicates(sorted []int) bool {
	for i := 0; i < len(sorted)-1; i++ {
		if sorted[i] == sorted[i+1] {
			return true
		}
	}
	return false
}
