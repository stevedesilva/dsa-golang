package recursion

func GetNumbersFromSlice(array []any) []int {
	result := make([]int, 0)
	return getNumbersFromSliceRec(array, result)
}

func getNumbersFromSliceRec(array []any, res []int) []int {
	for _, v := range array {
		if num, ok := v.(int); ok {
			res = append(res, num)
		} else if arr, ok := v.([]any); ok {
			res = getNumbersFromSliceRec(arr, res)
		} else {
			// error
		}
	}
	return res
}
