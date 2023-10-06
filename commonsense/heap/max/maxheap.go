package max

import (
	"errors"
	"golang.org/x/exp/constraints"
)

type Heap[T constraints.Ordered] struct {
	data []T
}

func New[T constraints.Ordered]() *Heap[T] {
	return &Heap[T]{}
}

func (h *Heap[T]) Root() (*T, error) {
	var val *T
	if len(h.data) == 0 {
		return val, errors.New("no elements found")
	}
	return &h.data[0], nil
}
