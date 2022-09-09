package array

import "errors"

var (
	ErrNoInputProvided = errors.New("no input provided")
)

type CountResult struct {
	Value int
}

func CountOnes(array []int) (*CountResult, error) {
	res := CountResult{}
	if len(array) <= 0 {
		return nil, ErrNoInputSupplied
	}
	res.Value = 0
	return &res, nil
}
