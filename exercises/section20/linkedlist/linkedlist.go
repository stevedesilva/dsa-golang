package linkedlist

import . "github.com/stevedesilva/dsa-golang.git/exercises/section20/generics"

type LinkedList[T comparable] struct {
	Head Node[T]
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}
