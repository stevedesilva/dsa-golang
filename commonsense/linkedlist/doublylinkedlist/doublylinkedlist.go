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
	if d.head != nil {
		return d.head.data, nil
	}
	var v T
	return v, errors.New("not found")
}

func (d *DoublyLinkedList[T]) ReadFromEnd() (T, error) {
	if d.tail != nil {
		return d.tail.data, nil
	}
	var v T
	return v, errors.New("not found")
}

func (d *DoublyLinkedList[T]) Search(value T) (int, error) {
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

func (d *DoublyLinkedList[T]) ReadByIndex(index int) (T, error) {
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

func (d *DoublyLinkedList[T]) DeleteFromFront() error {
	if d.head == nil {
		return errors.New("no items to delete")
	}
	d.head = d.head.next
	d.head.previous = nil
	return nil
}

func (d *DoublyLinkedList[T]) DeleteFromEnd() error {
	if d.tail == nil {
		return errors.New("no items to delete")
	}
	d.tail = d.tail.previous
	d.tail.next = nil
	return nil
}

func (d *DoublyLinkedList[T]) DeleteByIndex(index int) error {
	if index >= d.size {
		return errors.New("index not found")
	}
	if index == 0 {
		d.head = d.head.next
		d.head.previous = nil
		return nil
	}
	curr := d.head
	count := 0
	for count < index {
		count++
		curr = curr.next
		if curr == nil {
			return errors.New("index not found")
		}
	}

	curr.previous.next = curr.next
	curr.next.previous = curr.previous
	return nil
}

//func (d *DoublyLinkedList[T]) DeleteItems(predicate func(T) bool) {
//
//}
