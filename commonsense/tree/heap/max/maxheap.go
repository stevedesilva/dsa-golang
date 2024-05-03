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

func (h *Heap[T]) Insert(value T) error {
	// insert at last node
	h.data = append(h.data, value)
	// trickle up
	currIdx := len(h.data) - 1
	parentIdx := parentIndex(currIdx)
	for currIdx > 0 && h.data[currIdx] > h.data[parentIdx] {
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
			// swap curr with larger indexOfGreaterChild
			indexOfGreaterChild := h.indexOfGreaterChild(currIdx)
			if h.data[indexOfGreaterChild] > h.data[currIdx] {
				h.data[currIdx], h.data[indexOfGreaterChild] = h.data[indexOfGreaterChild], h.data[currIdx]
				currIdx = indexOfGreaterChild
			} else {
				break
			}

		}
	}
	//var val T
	return &val, nil
}

func (h *Heap[T]) hasChildren(index int) bool {
	// no need to check right child, if left child is not present (all fill from left to right)
	return childLeftIndex(index) <= len(h.data)-1
}

/*
//0,1,2,3,4,5,6,7,8,9,10

	//           0
	//      /          \
	//     1            2
	//    /  \        /  \
	//   3    4      5    6
	//  / \  / \    / \   / \
	// 7  8  9  10 11 12 12 14
*/
func (h *Heap[T]) indexOfGreaterChild(parentIdx int) int {
	// if no right child, return left child
	rightIndex := childRightIndex(parentIdx)
	totalLen := len(h.data) - 1
	if rightIndex >= totalLen {
		index := childLeftIndex(parentIdx)
		return index
	}
	// if left child is greater than right child, return left child
	if h.data[childLeftIndex(parentIdx)] > h.data[childRightIndex(parentIdx)] {
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
