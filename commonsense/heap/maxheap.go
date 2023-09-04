package heap

import "golang.org/x/exp/constraints"

type HeapNode[T constraints.Ordered] struct {
	data  *T
	left  *HeapNode[T]
	right *HeapNode[T]
}
