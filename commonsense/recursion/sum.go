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

func SumNumberBetweenRangeNoTailRec(start, end int) int {
	if start >= end {
		return end
	}
	return start + SumNumberBetweenRangeNoTailRec(start+1, end)
}
