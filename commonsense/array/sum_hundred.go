package array

import "errors"

var (
	ErrMinimumInputNotMet = errors.New("minimum input of 2 values required")
)

func OppositeValuesSumToHundred(numbers []int) (bool, error) {
	if len(numbers) < 2 {
		return false, ErrMinimumInputNotMet
	}

	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		if sum := numbers[i] + numbers[j]; sum == 100 {
			return true, nil
		}
	}
	return false, nil
}
