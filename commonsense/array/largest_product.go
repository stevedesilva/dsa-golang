package array

import "errors"

var (
	ErrMinimumNumberNotMet = errors.New("minimum array input size is 3")
)

func findLargestProduct(array []int) (int, error) {
	if len(array) < 3 {
		return 0, ErrMinimumInputNotMet
	}
	largestProduct := array[0] * array[1] * array[2]
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			for k := j + 1; k < len(array); k++ {
				curr := array[i] * array[j] * array[k]
				if largestProduct < curr {
					largestProduct = curr
				}
			}
		}
	}
	return largestProduct, nil
}
