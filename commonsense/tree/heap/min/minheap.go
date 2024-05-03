package min

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

func (h *Heap[T]) Insert(value T) error {
	// insert at last node
	h.data = append(h.data, value)
	// trickle up
	currIdx := len(h.data) - 1
	parentIdx := parentIndex(currIdx)
	for currIdx > 0 && h.data[currIdx] < h.data[parentIdx] { // changed > to <
		h.data[currIdx], h.data[parentIdx] = h.data[parentIdx], h.data[currIdx]
		currIdx = parentIdx
		parentIdx = parentIndex(currIdx)
	}
	return nil
}

func parentIndex(index int) int {
	return (index - 1) / 2
}

func (h *Heap[T]) Delete() (*T, error) {
	// delete top node
	if len(h.data) == 0 {
		return nil, errors.New("no value found")
	}
	val := h.data[0]
	// swap last child with root
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	// trickle down
	currIdx := 0
	if len(h.data) > 0 {
		// trickle to root node down into its proper place
		for h.hasChildren(currIdx) {
			// swap curr with smaller indexOfSmallerChild
			indexOfSmallerChild := h.indexOfSmallerChild(currIdx)
			if h.data[indexOfSmallerChild] < h.data[currIdx] { // changed > to <
				h.data[currIdx], h.data[indexOfSmallerChild] = h.data[indexOfSmallerChild], h.data[currIdx]
				currIdx = indexOfSmallerChild
			} else {
				break
			}

		}
	}
	return &val, nil
}

func (h *Heap[T]) hasChildren(index int) bool {
	return childLeftIndex(index) <= len(h.data)-1
}

func (h *Heap[T]) indexOfSmallerChild(parentIdx int) int {
	rightIndex := childRightIndex(parentIdx)
	totalLen := len(h.data) - 1
	if rightIndex >= totalLen {
		return childLeftIndex(parentIdx)
	}
	if h.data[childLeftIndex(parentIdx)] < h.data[childRightIndex(parentIdx)] { // changed > to <
		return childLeftIndex(parentIdx)
	} else {
		return childRightIndex(parentIdx)
	}
}

func childLeftIndex(parentIndex int) int {
	return (parentIndex * 2) + 1
}
func childRightIndex(parentIndex int) int {
	return (parentIndex * 2) + 2
}
