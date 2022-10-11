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
	return nil, nil
}
