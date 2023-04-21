package doublylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedList_AddAtEnd(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.AddAtEnd(1)
	assert.Equal(t, list.size, 1)
	assert.Equal(t, list.head.data, 1)
	assert.Equal(t, list.tail.data, 1)

	list.AddAtEnd(2)
	assert.Equal(t, list.size, 2)
	assert.Equal(t, list.head.data, 1)
	assert.Equal(t, list.tail.data, 2)

	assert.Equal(t, list.head.next.data, 2)

}

func TestDoublyLinkedList_AddByIndexFront(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")

	list.AddByIndex(0, "z")
	assert.Equal(t, list.size, 4)
	assert.Equal(t, list.head.data, "z")
	assert.Equal(t, list.head.next.data, "b")
	assert.Equal(t, list.head.next.next.data, "c")
	assert.Equal(t, list.tail.data, "c")
}

func TestDoublyLinkedList_AddByIndexMid(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")

	list.AddByIndex(1, "z")
	assert.Equal(t, list.size, 4)
	assert.Equal(t, list.head.data, "a")
	assert.Equal(t, list.head.next.data, "z")
	assert.Equal(t, list.tail.data, "c")
}

func TestDoublyLinkedList_AddByIndexEnd(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")

	list.AddByIndex(3, "z")
	assert.Equal(t, list.size, 4)
	assert.Equal(t, list.head.data, "a")
	assert.Equal(t, list.head.next.data, "b")
	assert.Equal(t, list.head.next.next.data, "c")
	assert.Equal(t, list.head.next.next.next.data, "z")
	assert.Equal(t, list.tail.data, "z")
}

func TestLinkedList_AddByIndexReturnErrorWhenIndexOutOfBounds(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")

	err := list.AddByIndex(12, "d")
	assert.NotNil(t, err)
}
