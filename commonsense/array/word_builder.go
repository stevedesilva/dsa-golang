package array

func BuildTwoLevelWordArray(array []rune) []string {
	res := make([]string, 0, len(array))
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if i != j {
				res = append(res, string(array[i])+string(array[j]))
			}
		}
	}
	return res
}
func BuildThreeLevelWordArray(array []rune) []string {
	res := make([]string, 0, len(array))
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			for k := 0; k < len(array); k++ {
				if i != j && j != k && i != k {
					res = append(res, string(array[i])+string(array[j])+string(array[k]))
				}
			}
		}
	}
	return res
}
