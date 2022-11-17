package recursion

func SumNumberBetweenRange(start, end int) int {
	return sumNumberBetweenRangeRec(start, end, 0)
}

func sumNumberBetweenRangeRec(start, end int, acc int) int {
	if start > end {
		return acc
	}
	acc = acc + start
	return sumNumberBetweenRangeRec(start+1, end, acc)
}
