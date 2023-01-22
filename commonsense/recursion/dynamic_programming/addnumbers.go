package dynamic_programming

func AddUntilOneHundred(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	tailSum := AddUntilOneHundred(numbers[1:])
	if numbers[0]+tailSum > 100 {
		return tailSum
	} else {
		return numbers[0] + tailSum
	}
}
