package doublylinkedlist

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
	d.size = d.size + 1
}

func (d *DoublyLinkedList[T]) AddByIndex(index int, value T) error {
	return nil
}

//func (d *DoublyLinkedList[T]) RemoveFromFront() T {
//	return nil
//}
//
//
//func (d *DoublyLinkedList[T]) DeleteItems(predicate func(T) bool) {
//
//}
//
//func (d *DoublyLinkedList[T]) Delete(index int) error {
//	return nil
//}
//
//func (d *DoublyLinkedList[T]) Search(value T) (int, error) {
//	return -1, nil
//}
//
//func (d *DoublyLinkedList[T]) Read(index int) (T, error) {
//	return nil, nil
//}
