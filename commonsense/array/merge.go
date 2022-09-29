package array

func mergeTwoArrays(a, b []int) []int {
	res := make([]int, 0, len(a)+len(b))
	for i, j := 0, 0; i < len(a) || j < len(b); {
		if i < len(a) && j < len(b) {
			if a[i] <= b[j] {
				res = append(res, a[i])
				i++
			} else {
				res = append(res, b[j])
				j++
			}
		} else if i < len(a) && j >= len(b) {
			res = append(res, a[i])
			i++
		} else if j < len(b) && i >= len(a) {
			res = append(res, b[j])
			j++
		}

	}
	return res
}
