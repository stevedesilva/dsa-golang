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

func FindDuplicates2(array []string) bool {
	// sort array
	sort
	// compare adjacent elements
	return false
}
