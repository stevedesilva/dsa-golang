package dynamic_programming

import "errors"

func FindMaxNumber(numbers []int) (int, error) {
	if len(numbers) < 1 {
		return 0, errors.New("input minimum not met")
	}
	return findMaxNumber(numbers), nil
}

func findMaxNumber(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}
	max := findMaxNumber(numbers[1:])
	if numbers[0] > max {
		return numbers[0]
	} else {
		return max
	}
}
