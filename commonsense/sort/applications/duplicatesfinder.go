package applications

import "github.com/stevedesilva/dsa-golang.git/commonsense/sort"

func HasDuplicates(array []int) bool {
	q := sort.NewQuickSort(array)
	return hasDuplicates(q.Quicksort())
}

func hasDuplicates(sorted []int) bool {
	for i := 0; i < len(sorted)-1; i++ {
		if sorted[i] == sorted[i+1] {
			return true
		}
	}
	return false
}
