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

	assert.Equal(t, list.Head().Data, "a")
	assert.Equal(t, list.Head().Next.Data, "b")
	assert.Equal(t, list.Head().Next.Next.Data, "c")
	assert.Equal(t, list.Head().Next.Next.Next.Data, "d")
	assert.Equal(t, list.Head().Next.Next.Next.Next.Data, "e")
	assert.Equal(t, list.Head().Next.Next.Next.Next.Next.Data, "f")
}

func TestLinkedList_AddByIndexToHead(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.AddByIndex(0, "a")
	assert.Equal(t, list.Head().Data, "a")

	list.AddByIndex(0, "b")
	assert.Equal(t, list.Head().Data, "b")

	list.AddByIndex(0, "c")
	assert.Equal(t, list.Head().Data, "c")
}

func TestLinkedList_AddByIndexToMiddle(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")

	err := list.AddByIndex(1, "d")
	assert.Nil(t, err)
	assert.Equal(t, list.Head().Next.Data, "d")

	err = list.AddByIndex(1, "e")
	assert.Nil(t, err)
	assert.Equal(t, list.Head().Next.Data, "e")
	assert.Equal(t, list.Head().Next.Next.Data, "d")
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
	assert.Equal(t, list.Head().Next.Next.Next.Data, "d")
}

func TestLinkedList_Read(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")

	read, err := list.Read(0)
	assert.Nil(t, err)
	assert.Equal(t, "a", read)
	read, err = list.Read(1)
	assert.Nil(t, err)
	assert.Equal(t, "b", read)
	read, err = list.Read(2)
	assert.Nil(t, err)
	assert.Equal(t, "c", read)
}

func TestLinkedList_Search(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	idx, err := list.Search("a")
	assert.Nil(t, err)
	assert.Equal(t, 0, idx)
	idx, err = list.Search("b")
	assert.Nil(t, err)
	assert.Equal(t, 1, idx)
	idx, err = list.Search("c")
	assert.Nil(t, err)
	assert.Equal(t, 2, idx)
}

func TestLinkedList_SearchErrorWhenIndexNotFound(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	idx, err := list.Search("z")
	assert.NotNil(t, err)
	assert.Equal(t, -1, idx)
}

func TestLinkedList_DeleteHead(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	err := list.Delete(0)
	assert.Nil(t, err)
	assert.Equal(t, "b", list.Head().Data)
	err = list.Delete(0)
	assert.Nil(t, err)
	assert.Equal(t, "c", list.Head().Data)
}

func TestLinkedList_DeleteMiddle(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	err := list.Delete(1)
	assert.Nil(t, err)
	assert.Equal(t, "a", list.Head().Data)
	assert.Equal(t, "c", list.Head().Next.Data)
}

func TestLinkedList_DeleteEnd(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	err := list.Delete(2)
	assert.Nil(t, err)
	assert.Equal(t, "a", list.Head().Data)
	assert.Equal(t, "b", list.Head().Next.Data)
	err = list.Delete(0)
	assert.Nil(t, err)
	assert.Equal(t, "b", list.Head().Data)
}

func TestLinkedList_DeleteErrorWhenIndexNotFound(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	err := list.Delete(10)
	assert.NotNil(t, err)
}
