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

func (d *DoublyLinkedList[T]) AddByIndex(index int, value T) error {
	node := &Node[T]{
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

func (d *DoublyLinkedList[T]) AddToFront(value T) {
	d.AddByIndex(0, value)
}

func (d *DoublyLinkedList[T]) AddAtEnd(value T) {
	d.AddByIndex(d.size, value)
	//node := &Node[T]{
	//	data:     value,
	//	previous: nil,
	//	next:     nil,
	//}
	//if d.head == nil {
	//	d.head = node
	//	d.tail = node
	//} else {
	//	// new node points to previous tail
	//	node.previous = d.tail
	//	// update current tail
	//	d.tail.next = node
	//	// make new node the new tail
	//	d.tail = node
	//
	//}
	//d.size++
}

func (d *DoublyLinkedList[T]) ReadFromFront() (T, error) {
	if d.head != nil {
		return d.head.data, nil
	}
	var v T
	return v, errors.New("not found")
}

func (d *DoublyLinkedList[T]) ReadFromEnd() (T, error) {
	//return d.ReadByIndex(d.size - 1)
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

func (d *DoublyLinkedList[T]) RemoveByIndex(index int) (T, error) {
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

func (d *DoublyLinkedList[T]) RemoveFromFront() (T, error) {
	return d.RemoveByIndex(0)
}

func (d *DoublyLinkedList[T]) RemoveFromEnd() (T, error) {
	return d.RemoveByIndex(lastIndex(d.size))
}

func (d *DoublyLinkedList[T]) DeleteByIndex(index int) error {
	_, err := d.RemoveByIndex(index)
	return err
}

func (d *DoublyLinkedList[T]) DeleteFromFront() error {
	return d.DeleteByIndex(0)
}

func (d *DoublyLinkedList[T]) DeleteFromEnd() error {
	return d.DeleteByIndex(lastIndex(d.size))
}

func lastIndex(size int) int {
	var index = size - 1
	if index < 0 {
		index = 0
	}
	return index
}

func (d *DoublyLinkedList[T]) DeleteItems(predicate func(T) bool) {
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
