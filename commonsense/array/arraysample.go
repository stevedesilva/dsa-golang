package array

import "errors"

type result struct {
	start, mid, end int
}

var errNoInputSupplied = errors.New("no input arguments provided")

func findStartMidEndOfArray(array []int) (*result, error) {
	if len(array) <= 0 {
		return nil, errNoInputSupplied
	}
	start := array[0]
	mid := array[len(array)/2]
	end := array[len(array)-1]
	return &result{start, mid, end}, nil
}
