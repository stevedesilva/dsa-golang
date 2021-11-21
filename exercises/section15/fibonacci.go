package section15

func GenerateFibSeriesIterative(num int) int {
	res := []int{0, 1}
	for i := 2; i <= num; i++ {
		value := res[(len(res)-2)] + res[(len(res)-1)]
		res = append(res, value)
	}
	return res[(len(res) - 1)]
}

func GenerateFibSeriesIterative2(num int) int {
	res := []int{0, 1}
	for i := 2; i <= num; i++ {
		value := res[i-2] + res[i-1]
		res = append(res, value)
	}
	return res[num]
}

func GenerateFibSeriesRecursive(num int) int {
	res := fibSeriesRecursive(num, 2, []int{0, 1})
	return res[(len(res) - 1)]
}

func fibSeriesRecursive(num, count int, res []int) []int {
	if count > num {
		return res
	}

	value := res[(len(res)-2)] + res[(len(res)-1)]
	res = append(res, value)

	count++
	return fibSeriesRecursive(num, count, res)
}

func GenerateFibSeriesRecursive2(num int) int {
	res := fibSeriesRecursive2(num, 2, []int{0, 1})
	return res[num]
}

func fibSeriesRecursive2(num, count int, res []int) []int {
	if count > num {
		return res
	}

	value := res[count-2] + res[count-1]
	res = append(res, value)

	count++
	return fibSeriesRecursive(num, count, res)
}
