package array

import "errors"

var (
	ErrMinimumInputNotMet = errors.New("minimum input of 2 values required")
)

func OppositeValuesSumToHundred(numbers []int) (bool, error) {
	if len(numbers) < 2 {
		return false, ErrMinimumInputNotMet
	}
	return false, nil
}
