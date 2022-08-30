package array

import "errors"

type Result struct {
	start, mid, end int
}

var ErrNoInputSupplied = errors.New("no input arguments provided")

func findStartMidEndOfArray(array []int) (*Result, error) {
	if len(array) <= 0 {
		return nil, ErrNoInputSupplied
	}
	return nil, nil
}
