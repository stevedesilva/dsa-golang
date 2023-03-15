package applications

import "errors"

var (
	ErrMinimumValue = errors.New("minimum values of slice not supplied")
)

func LargestValueN(numbers []int) (*int, error) {
	if len(numbers) < 1 {
		return nil, ErrMinimumValue
	}
	maxValue := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if maxValue < numbers[i] {
			maxValue = numbers[i]
		}
	}
	return intPtr(maxValue), nil
}

func LargestValueNLogN(numbers []int) (*int, error) {
	if len(numbers) < 1 {
		return nil, ErrMinimumValue
	}
	return intPtr(0), nil
}

func LargestValueN2(numbers []int) (*int, error) {
	if len(numbers) < 1 {
		return nil, ErrMinimumValue
	}
	return intPtr(0), nil
}

func intPtr(n int) *int {

	return &n
}
