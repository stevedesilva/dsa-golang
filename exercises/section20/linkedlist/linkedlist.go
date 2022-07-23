package linkedlist

type LinkedList[T comparable] struct {
	Head Node[T]
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}
