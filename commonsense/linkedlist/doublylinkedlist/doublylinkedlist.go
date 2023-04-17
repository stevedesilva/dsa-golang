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
		next:     nil,
		previous: nil,
	}
	if d.head == nil {
		d.head = node
		d.tail = nil
	} else {
		curr := d.head
		for curr.next != nil {
			curr = curr.next
		}
		// end of list
		node.previous = curr
		curr.next = node
		d.tail = node
	}
	d.size = d.size + 1
}

//func (d *DoublyLinkedList[T]) RemoveFromFront() T {
//	return nil
//}
//
//func (d *DoublyLinkedList[T]) AddByIndex(index int, value T) error {
//	return nil
//}
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
