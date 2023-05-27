package doublylinkedlist

import (
	"errors"
	"fmt"
	"strings"
)

type AllowedDoublyLinkedListFunc[T comparable] interface {
	AddAtEnd(value T)
	AddByIndex(index int, value T) error
	AddToFront(value T)
	DeleteByIndex(index int) error
	DeleteFromEnd() error
	DeleteFromFront() error
	DeleteItems(func(T) bool)
	PrintItems() string
	PrintItemsInReverse() string
	ReadByIndex(index int) (T, error)
	ReadFromEnd() (T, error)
	ReadFromFront() (T, error)
	RemoveByIndex(index int) (T, error)
	RemoveFromEnd() (T, error)
	RemoveFromFront() (T, error)
	Search(value T) (int, error)
	Size() int
	Head() *node[T]
	Tail() *node[T]
}

func NewDoublyLinkedList[T comparable]() AllowedDoublyLinkedListFunc[T] {
	return &doublyLinkedList[T]{}
}

type node[T comparable] struct {
	data     T
	next     *node[T]
	previous *node[T]
}

type doublyLinkedList[T comparable] struct {
	head *node[T]
	tail *node[T]
	size int
}

func (d *doublyLinkedList[T]) Head() *node[T] {
	return d.head
}
func (d *doublyLinkedList[T]) Tail() *node[T] {
	return d.tail
}

func (d *doublyLinkedList[T]) AddByIndex(index int, value T) error {
	node := &node[T]{
		data:     value,
		previous: nil,
		next:     nil,
	}
	if index == 0 {
		if d.head == nil {
			d.tail = node
		}
		node.next = d.head
		d.head = node
		d.size++
		return nil
	}

	if index == d.size {
		// last element
		node.previous = d.tail
		d.tail.next = node
		d.tail = node

		d.size++
		return nil
	}

	count := 0
	curr := d.head

	for count < index-1 {
		count++
		curr = curr.next
		if curr == nil {
			return errors.New("index not found")
		}
	}
	node.previous = curr
	node.next = curr.next
	curr.next = node
	d.size++
	return nil
}

func (d *doublyLinkedList[T]) AddToFront(value T) {
	err := d.AddByIndex(0, value)
	if err != nil {
		return
	}
}

func (d *doublyLinkedList[T]) AddAtEnd(value T) {
	err := d.AddByIndex(d.size, value)
	if err != nil {
		return
	}
}

func (d *doublyLinkedList[T]) ReadFromFront() (T, error) {
	if d.head != nil {
		return d.head.data, nil
	}
	var v T
	return v, errors.New("not found")
}

func (d *doublyLinkedList[T]) ReadFromEnd() (T, error) {
	if d.tail != nil {
		return d.tail.data, nil
	}
	var v T
	return v, errors.New("not found")
}

func (d *doublyLinkedList[T]) Search(value T) (int, error) {
	curr := d.head
	index := 0
	for curr != nil {
		if curr.data == value {
			return index, nil
		}
		curr = curr.next
		index++
	}
	return -1, errors.New("not found")
}

func (d *doublyLinkedList[T]) ReadByIndex(index int) (T, error) {
	var v T
	if index >= d.size {
		return v, errors.New("index not found")
	}
	count := 0
	curr := d.head
	for count < index {
		curr = curr.next
		count++
	}
	return curr.data, nil
}

func (d *doublyLinkedList[T]) RemoveByIndex(index int) (T, error) {
	var value T
	if index >= d.size {
		return value, errors.New("index not found")
	}

	if index == 0 {
		value = d.head.data
		d.head = d.head.next
		d.size--
		return value, nil
	}

	curr := d.head
	count := 0
	for count < index {
		count++
		curr = curr.next
	}
	if curr.next == nil {
		value = curr.data
		d.tail = curr.previous
		curr.previous.next = nil
	} else {
		value = curr.data
		curr.previous.next = curr.next
		curr.next.previous = curr.previous
	}
	d.size--
	return value, nil
}

func (d *doublyLinkedList[T]) RemoveFromFront() (T, error) {
	return d.RemoveByIndex(0)
}

func (d *doublyLinkedList[T]) RemoveFromEnd() (T, error) {
	return d.RemoveByIndex(lastIndex(d.size))
}

func (d *doublyLinkedList[T]) DeleteByIndex(index int) error {
	_, err := d.RemoveByIndex(index)
	return err
}

func (d *doublyLinkedList[T]) DeleteFromFront() error {
	return d.DeleteByIndex(0)
}

func (d *doublyLinkedList[T]) DeleteFromEnd() error {
	return d.DeleteByIndex(lastIndex(d.size))
}

func lastIndex(size int) int {
	var index = size - 1
	if index < 0 {
		index = 0
	}
	return index
}

func (d *doublyLinkedList[T]) DeleteItems(predicate func(T) bool) {
	curr := d.head
	for curr != nil {
		if predicate(curr.data) {
			if curr.previous == nil {
				// head
				d.head = curr.next
				curr.next.previous = nil
			} else if curr.next == nil {
				// tail
				d.tail = curr.previous
				curr.previous.next = nil
			} else {
				curr.previous.next = curr.next
				curr.next.previous = curr.previous
			}
			d.size--
		}
		curr = curr.next
	}
}

func (d *doublyLinkedList[T]) PrintItems() string {
	res := make([]string, 0, d.size)
	curr := d.head
	for curr != nil {
		elem := fmt.Sprintf("%v", curr.data)
		res = append(res, elem)
		curr = curr.next
	}
	return strings.Join(res, ",")
}

func (d *doublyLinkedList[T]) PrintItemsInReverse() string {
	res := make([]string, 0, d.size)
	curr := d.tail
	for curr != nil {
		elem := fmt.Sprintf("%v", curr.data)
		res = append(res, elem)
		curr = curr.previous
	}
	return strings.Join(res, ",")
}

func (d *doublyLinkedList[T]) Size() int {
	return d.size
}
