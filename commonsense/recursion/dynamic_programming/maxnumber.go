package dynamic_programming

import "errors"

func FindMaxNumber(numbers []int) (int, error) {
	if len(numbers) < 1 {
		return 0, errors.New("input minimum not met")
	}
	if len(numbers) == 1 {
		return numbers[0], nil
	}
	return 0, nil
}
