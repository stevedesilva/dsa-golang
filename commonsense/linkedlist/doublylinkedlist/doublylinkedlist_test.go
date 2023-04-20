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
