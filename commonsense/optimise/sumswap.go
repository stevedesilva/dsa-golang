package optimise

func SumSwap(a, b []int) (int, int) {
	sumA, sumB := 0, 0
	for _, v := range a {
		sumA += v
	}
	for _, v := range b {
		sumB += v
	}
	diff := sumA - sumB
	if diff%2 != 0 {
		return 0, 0
	}
	diff /= 2
	for _, v := range a {
		for _, w := range b {
			if v-w == diff {
				return v, w
			}
		}
	}
	return 0, 0
}

func SwapToMakeEqual(a1, a2 []int) []int {
	return []int{0, 0}
}
