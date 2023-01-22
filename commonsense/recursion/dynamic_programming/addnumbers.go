package dynamic_programming

func AddUntilOneHundred(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	if numbers[0]+AddUntilOneHundred(numbers[1:]) > 100 {
		return AddUntilOneHundred(numbers[1:])
	} else {
		return numbers[0] + AddUntilOneHundred(numbers[1:])
	}
}
