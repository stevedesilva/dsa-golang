package array

import "errors"

var (
	ErrInvalidInput = errors.New("minimum input of two items require")
)

func FindTheProduct(input []int) ([]int, error) {
	if len(input) < 2 {
		return nil, ErrInvalidInput
	}
	result := make([]int, 0, len(input))
	return result, nil
}
