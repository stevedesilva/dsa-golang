package recursion

func CountNumberOfSteps(number int) int {
	return countNumberOfSteps(number)
}

func countNumberOfSteps(number int) int {
	if number <= 0 {
		return 0
	}
	if number == 1 {
		return 1
	}
	if number == 2 {
		return 2
	}
	if number == 3 {
		return 4
	}

	return countNumberOfSteps(number-3) + countNumberOfSteps(number-2) + countNumberOfSteps(number-1)
}
