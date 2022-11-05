package recursion

func countTo(number int) []int {
	res := count(number, make([]int, 0))
	return res
}

func count(number int, acc []int) {
	if number == 0 {
		return
	}
	neoNum := number - 1
	acc = append(acc, neoNum)
	count(neoNum, acc)
}
