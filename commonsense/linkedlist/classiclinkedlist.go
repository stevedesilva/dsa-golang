package linkedlist

import (
	"errors"
	"github.com/stevedesilva/dsa-golang.git/commonsense/node"
)

type LinkedList[T comparable] struct {
	head *node.Node[T]
}

type ClassicLinkedList[T comparable] interface {
	Add(value T)
	AddByIndex(index int, value T) error
	Read(index int) (T, error)
	Search(value T) (int, error)
	Delete(index int) error
	Head() *node.Node[T]
}

func NewClassicLinkedList[T comparable]() ClassicLinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Head() *node.Node[T] {
	return l.head
}

func (l *LinkedList[T]) Add(value T) {
	node := &node.Node[T]{
		Data: value,
		Next: nil,
	}
	if l.head == nil {
		l.head = node
		return
	}
	current := l.head
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
		node.Next = l.head
		l.head = node
		return nil
	}

	current := l.head
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
	current := l.head
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
	current := l.head
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
	current := l.head
	if index == 0 {
		l.head = l.head.Next
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
