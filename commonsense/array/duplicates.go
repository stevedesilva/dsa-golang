package array

import "sort"

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

func FindDuplicates(array []string) bool {
	table := make(map[string]bool)
	for _, v := range array {
		if _, ok := table[v]; ok {
			return true
		}
		table[v] = true
	}
	return false
}

func FindDuplicatesWithSort(array []string) bool {
	// sort array
	sort.Strings(array)

	// compare adjacent elements
	for i := 0; i < len(array)-1; i++ {
		if array[i] == array[i+1] {
			return true
		}
	}
	return false
}
