package array

import "errors"

var (
	ErrNoInputProvided = errors.New("no input provided")
)

type CountResult struct {
	Value int
}

func CountOnes(array [][]int) (*CountResult, error) {
	if len(array) <= 0 {
		return nil, ErrNoInputProvided
	}
	count := 0
	for _, innerArray := range array {
		for _, j := range innerArray {
			if j == 1 {
				count++
			}
		}
	}
	return &CountResult{
		Value: count,
	}, nil
}
