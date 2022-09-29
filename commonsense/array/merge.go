package array

func mergeTwoArrays(a, b []int) []int {
	res := make([]int, 0, len(a)+len(b))
	for i, j := 0, 0; i < len(a) || j < len(b); {

		i2 := a[i]
		i3 := b[j]
		if i2 <= i3 {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}

	}
	return res
}
