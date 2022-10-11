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
	return nil, nil
}
