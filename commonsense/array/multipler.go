package array

import (
	"errors"
	"math"
)

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

func FindTheProductFromTwoArrays(inputA []int, inputB []int) []int {
	max := math.Max(float64(len(inputA)), float64(len(inputB)))
	result := make([]int, 0, int(math.Round(max)))
	for i := 0; i < len(inputA); i++ {
		for j := 0; j < len(inputB); j++ {
			result = append(result, inputA[i]*inputB[j])
		}
	}
	return result
}
