package hashmap

import "errors"

var (
	ErrInvalidInput = errors.New("minimum input required")
)

type SubArray interface {
	GetIntersection() ([]int, error)
}

type arrayToMap struct {
	a []int
	b []int
}

func New(a, b []int) SubArray {
	return &arrayToMap{a, b}
}

func (s *arrayToMap) GetIntersection() ([]int, error) {
	if len(s.a) < 1 && len(s.b) < 1 {
		return nil, ErrInvalidInput
	}
	if len(s.a) < 1 || len(s.b) < 1 {
		if len(s.a) < 1 {
			return s.b, nil
		} else {
			return s.a, nil
		}
	}
	var result []int
	if len(s.a) >= len(s.b) {
		result = getSubset(s.a, s.b)
	} else {
		result = getSubset(s.b, s.a)
	}
	return result, nil
}

func getSubset(largeInput, smallInput []int) []int {
	size := len(largeInput)
	values := make(map[int]bool, size)
	res := make([]int, 0, size)

	for _, val := range largeInput {
		values[val] = true
	}

	for _, val := range smallInput {
		if values[val] {
			res = append(res, val)
		}
	}
	return res
}
