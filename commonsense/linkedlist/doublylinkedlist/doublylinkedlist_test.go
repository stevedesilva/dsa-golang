package doublylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedList_AddAtEnd(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.AddAtEnd(1)
	assert.Equal(t, 1, list.size)
	assert.Equal(t, 1, list.head.data)
	assert.Equal(t, 1, list.tail.data)

	list.AddAtEnd(2)
	assert.Equal(t, 2, list.size)
	assert.Equal(t, 1, list.head.data)
	assert.Equal(t, 2, list.tail.data)

	assert.Equal(t, 2, list.head.next.data)
}

func TestDoublyLinkedList_AddByIndexFront(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")

	list.AddByIndex(0, "z")
	assert.Equal(t, 4, list.size)
	assert.Equal(t, "z", list.head.data)
	assert.Equal(t, "c", list.tail.data)
	assert.Equal(t, "a", list.head.next.data)
	assert.Equal(t, "b", list.head.next.next.data)
	assert.Equal(t, "c", list.head.next.next.next.data)
}

func TestDoublyLinkedList_AddByIndexMid(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")

	list.AddByIndex(1, "z")
	assert.Equal(t, 4, list.size)
	assert.Equal(t, "a", list.head.data)
	assert.Equal(t, "c", list.tail.data)
	assert.Equal(t, "z", list.head.next.data)
	assert.Equal(t, "b", list.head.next.next.data)
	//assert.Equal(t, "c", list.head.next.next.next.data)
}

func TestDoublyLinkedList_AddByIndexMidEnd(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")

	list.AddByIndex(2, "z")
	assert.Equal(t, 4, list.size)
	assert.Equal(t, "a", list.head.data)
	assert.Equal(t, "c", list.tail.data)
	assert.Equal(t, "b", list.head.next.data)
	assert.Equal(t, "z", list.head.next.next.data)
	assert.Equal(t, "c", list.head.next.next.next.data)
}

func TestDoublyLinkedList_AddByIndexEnd(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")

	list.AddByIndex(3, "z")
	assert.Equal(t, 4, list.size)
	assert.Equal(t, "a", list.head.data)
	assert.Equal(t, "z", list.tail.data)
	assert.Equal(t, "b", list.head.next.data)
	assert.Equal(t, "c", list.head.next.next.data)
	assert.Equal(t, "z", list.head.next.next.next.data)
}

func TestLinkedList_AddByIndexReturnErrorWhenIndexOutOfBounds(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")

	err := list.AddByIndex(12, "d")
	assert.NotNil(t, err)
}
