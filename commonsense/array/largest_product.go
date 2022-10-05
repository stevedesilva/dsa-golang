package array

import "errors"

var (
	ErrMinimumNumberNotMet = errors.New("minimum array input size is 3")
)

func findLargestProduct(array []int) (int, error) {
	if len(array) < 3 {
		return 0, ErrMinimumInputNotMet
	}
	return 0, nil
}
