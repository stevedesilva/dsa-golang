package generics

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

func New[T comparable](value T) *Node[T] {
	return &Node[T]{
		Value: value,
		Next:  nil,
	}
}

func NewWithNext[T comparable](value T, next *Node[T]) *Node[T] {
	return &Node[T]{
		Value: value,
		Next:  next,
	}
}
