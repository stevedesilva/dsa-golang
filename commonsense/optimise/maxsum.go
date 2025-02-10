package optimise

func MaxSum(array []int) int {
	max, curr := 0, 0

	for _, v := range array {
		curr += v
		if curr < 0 {
			curr = 0
		}
		if max < curr {
			max = curr
		}
	}
	return max

}
