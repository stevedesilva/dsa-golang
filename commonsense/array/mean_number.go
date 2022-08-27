package array

import "errors"

var (
	ErrNotFound = errors.New("no even numbers found")
)

func FindMeanForEvenNumbers(array []int) (float64, error) {
	var sum float64 = 0
	var count float64 = 0
	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			count++
			sum = sum + float64(array[i])
		}
	}
	if count < 1 {
		return 0, ErrNotFound
	}
	res := sum / count
	return res, nil
}
