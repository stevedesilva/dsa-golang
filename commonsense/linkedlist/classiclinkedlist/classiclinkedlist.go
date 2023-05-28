package classiclinkedlist

import (
	"errors"
	"fmt"
	"strings"
)

type Node[T comparable] struct {
	data T
	next *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
}

type ClassicLinkedList[T comparable] interface {
	Add(value T)
	AddByIndex(index int, value T) error
	Read(index int) (T, error)
	Search(value T) (int, error)
	Delete(index int) error
	DeleteItems(predicate func(T) bool)
	GetHead() *Node[T]
	SetHead(*Node[T])
	PrintItems() string
	PrintItemsInReverse() string
	ReadLastItem() (T, error)
}

func NewClassicLinkedList[T comparable]() ClassicLinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) GetHead() *Node[T] {
	return l.head
}
func (l *LinkedList[T]) SetHead(head *Node[T]) {
	l.head = head
}

func (l *LinkedList[T]) Add(value T) {
	node := &Node[T]{
		data: value,
		next: nil,
	}
	if l.head == nil {
		l.head = node
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
}

func (l *LinkedList[T]) AddByIndex(index int, value T) error {
	node := &Node[T]{
		data: value,
		next: nil,
	}
	if index == 0 {
		node.next = l.head
		l.head = node
		return nil
	}

	current := l.head
	currentIdx := 0

	for currentIdx < index-1 {
		current = current.next
		currentIdx++
		if current == nil {
			return errors.New("index not found")
		}
	}
	node.next = current.next
	current.next = node
	return nil
}

func (l *LinkedList[T]) Read(index int) (T, error) {
	var value T
	currentIdx := 0
	current := l.head
	for currentIdx < index {
		current = current.next
		currentIdx++
		if current == nil {
			return value, errors.New("index not found")
		}
	}
	return current.data, nil
}

func (l *LinkedList[T]) Search(value T) (int, error) {
	currentIdx := 0
	current := l.head
	for current != nil {
		if current.data == value {
			return currentIdx, nil
		}
		current = current.next
		currentIdx++
	}
	return -1, errors.New("value not found")
}

func (l *LinkedList[T]) Delete(index int) error {
	currentIdx := 0
	current := l.head
	if index == 0 {
		l.head = l.head.next
	}
	for currentIdx < index-1 {
		current = current.next
		currentIdx++
		if current == nil {
			return errors.New("index not found")
		}
	}
	current.next = current.next.next
	return nil
}

func (l *LinkedList[T]) DeleteItems(predicate func(T) bool) {
	var previous *Node[T]
	current := l.head
	for current != nil {
		if predicate(current.data) {
			if previous == nil {
				// head case
				if predicate(current.data) {
					l.SetHead(current.next)
				} else {
					previous = current
				}
			} else {
				// tail case
				if predicate(current.data) {
					previous.next = current.next
				} else {
					previous = current
				}
			}
		}
		current = current.next
	}
}

func (l *LinkedList[T]) PrintItems() string {
	res := make([]string, 0)
	current := l.head
	for current != nil {
		res = append(res, fmt.Sprint(current.data))
		current = current.next
		if current != nil {
			res = append(res, " -> ")
		} else {
			res = append(res, " -> nil")
		}
	}
	return strings.Join(res, "")
}

func (l *LinkedList[T]) PrintItemsInReverse() string {
	res := []string{" -> nil"}
	current := l.head
	for current != nil {
		// add to front of res
		var val string
		if current.next != nil {
			val = fmt.Sprint(" -> ", current.data)
		} else {
			val = fmt.Sprint(current.data)
		}
		res = append([]string{val}, res[:]...)
		current = current.next
	}
	// return string value of res
	return strings.Join(res, "")
}

func (l *LinkedList[T]) ReadLastItem() (T, error) {
	var value T
	if l.head == nil {
		return value, errors.New("list is empty")
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	return current.data, nil
}
