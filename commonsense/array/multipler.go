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
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			result = append(result, input[i]*input[j])
		}
	}
	return result, nil
}
