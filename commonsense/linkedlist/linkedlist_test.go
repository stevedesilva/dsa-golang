package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList_Add(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	list.Add("d")
	list.Add("e")
	list.Add("f")

	assert.Equal(t, list.Head.Data, "a")
	assert.Equal(t, list.Head.Next.Data, "b")
	assert.Equal(t, list.Head.Next.Next.Data, "c")
	assert.Equal(t, list.Head.Next.Next.Next.Data, "d")
	assert.Equal(t, list.Head.Next.Next.Next.Next.Data, "e")
	assert.Equal(t, list.Head.Next.Next.Next.Next.Next.Data, "f")
}

func TestLinkedList_AddByIndexToHead(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.AddByIndex(0, "a")
	assert.Equal(t, list.Head.Data, "a")

	list.AddByIndex(0, "b")
	assert.Equal(t, list.Head.Data, "b")

	list.AddByIndex(0, "c")
	assert.Equal(t, list.Head.Data, "c")
}

func TestLinkedList_AddByIndexToMiddle(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")

	err := list.AddByIndex(1, "d")
	assert.Nil(t, err)
	assert.Equal(t, list.Head.Next.Data, "d")

	err = list.AddByIndex(1, "e")
	assert.Nil(t, err)
	assert.Equal(t, list.Head.Next.Data, "e")
	assert.Equal(t, list.Head.Next.Next.Data, "d")
}

func TestLinkedList_AddByIndexReturnErrorWhenIndexOutOfBounds(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")

	err := list.AddByIndex(12, "d")
	assert.NotNil(t, err)
}
func TestLinkedList_AddByIndexToEndOfList(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")

	err := list.AddByIndex(3, "d")
	assert.Nil(t, err)
	assert.Equal(t, list.Head.Next.Next.Next.Data, "d")
}

func TestLinkedList_Delete(t *testing.T) {

}

func TestLinkedList_Read(t *testing.T) {

}

func TestLinkedList_Search(t *testing.T) {

}
