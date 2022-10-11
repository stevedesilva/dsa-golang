package hashmap

import "errors"

var (
	ErrInvalidInput = errors.New("minimum input required")
)

type SubArray interface {
	findSubArray(a, b []int) ([]int, error)
}

type arrayToMap struct {
	a []int
	b []int
}

func New(a, b []int) SubArray {
	return &arrayToMap{a, b}
}

func (s *arrayToMap) findSubArray(a, b []int) ([]int, error) {
	return nil, nil
}
