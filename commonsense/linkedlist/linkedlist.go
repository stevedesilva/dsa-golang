package linkedlist

import (
	"errors"
	"github.com/stevedesilva/dsa-golang.git/commonsense/node"
)

type LinkedList[T comparable] struct {
	Head *node.Node[T]
}

//type ClassicLinkedList[T comparable] interface {
//	Read(index int) (T, error)
//	Add(value T)
//	Search(index T) (int, error)
//	Delete(index int) error
//}

func NewClassicLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Add(value T) {
	node := &node.Node[T]{
		Data: value,
		Next: nil,
	}
	if l.Head == nil {
		l.Head = node
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
}

func (l *LinkedList[T]) AddByIndex(index int, value T) error {
	node := &node.Node[T]{
		Data: value,
		Next: nil,
	}
	if index == 0 {
		node.Next = l.Head
		l.Head = node
		return nil
	}

	current := l.Head
	currentIdx := 0

	for currentIdx < index-1 {
		current = current.Next
		currentIdx++
		if current == nil {
			return errors.New("index not found")
		}
	}
	node.Next = current.Next
	current.Next = node
	return nil
}

func (l *LinkedList[T]) Read(index int) (T, error) {
	var value T
	currentIdx := 0
	current := l.Head
	for currentIdx < index {
		current = current.Next
		currentIdx++
		if current == nil {
			return value, errors.New("index not found")
		}
	}
	return current.Data, nil
}

func (l *LinkedList[T]) Search(value T) (int, error) {
	currentIdx := 0
	current := l.Head
	for current != nil {
		if current.Data == value {
			return currentIdx, nil
		}
		current = current.Next
		currentIdx++
	}
	return -1, errors.New("value not found")
}

func (l *LinkedList[T]) Delete(index int) error {
	currentIdx := 0
	current := l.Head
	if index == 0 {
		l.Head = l.Head.Next
	}
	for currentIdx < index-1 {
		current = current.Next
		currentIdx++
		if current == nil {
			return errors.New("index not found")
		}
	}
	current.Next = current.Next.Next
	return nil
}
