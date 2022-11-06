package recursion

func getFactorial(num int) int {
	return factorial(num, 1)
}

func factorial(num int, acc int) int {
	if num == 0 {
		return acc
	}
	return factorial(num-1, acc*num)
}
