package doublylinkedlist

import "errors"

type Node[T comparable] struct {
	data     T
	next     *Node[T]
	previous *Node[T]
}

type DoublyLinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewDoublyLinkedList[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (d *DoublyLinkedList[T]) AddToFront(value T) {
	node := &Node[T]{
		data:     value,
		previous: nil,
		next:     nil,
	}
	if d.head == nil {
		d.head = node
		d.tail = node
	} else {
		// new node points to previous tail
		node.next = d.head
		// update current tail
		d.head.previous = node
		// make new node the new head
		d.head = node
	}
	d.size++
}

func (d *DoublyLinkedList[T]) AddAtEnd(value T) {
	node := &Node[T]{
		data:     value,
		previous: nil,
		next:     nil,
	}
	if d.head == nil {
		d.head = node
		d.tail = node
	} else {
		// new node points to previous tail
		node.previous = d.tail
		// update current tail
		d.tail.next = node
		// make new node the new tail
		d.tail = node

	}
	d.size++
}

func (d *DoublyLinkedList[T]) AddByIndex(index int, value T) error {
	// a - b - c
	node := &Node[T]{
		data:     value,
		previous: nil,
		next:     nil,
	}
	if index == 0 {
		// head needs updating
		node.next = d.head
		d.head = node
		d.size++
		return nil
	}
	if index == d.size {
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

func (d *DoublyLinkedList[T]) ReadFromFront() (T, error) {
	return nil, nil
}

func (d *DoublyLinkedList[T]) ReadFromEnd() (T, error) {
	return nil, nil
}

func (d *DoublyLinkedList[T]) ReadByIndex(index int) (T, error) {
	return nil, nil
}

//func (d *DoublyLinkedList[T]) DeleteFromFront() T {
//	return nil
//}

//func (d *DoublyLinkedList[T]) DeleteFromEnd() T {
//	return nil
//}

//func (d *DoublyLinkedList[T]) Delete(index int) error {
//	return nil
//}

//
//
//func (d *DoublyLinkedList[T]) DeleteItems(predicate func(T) bool) {
//
//}
//

//
//func (d *DoublyLinkedList[T]) Search(value T) (int, error) {
//	return -1, nil
//}
