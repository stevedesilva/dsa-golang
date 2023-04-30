package doublylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedList_AddToFront(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")

	list.AddToFront("y")
	assert.Equal(t, 3, list.size)
	assert.Equal(t, "y", list.head.data)
	assert.Equal(t, "a", list.head.next.data)
	assert.Equal(t, "b", list.head.next.next.data)
	assert.Equal(t, "b", list.tail.data)

	list.AddToFront("z")
	assert.Equal(t, 4, list.size)
	assert.Equal(t, "z", list.head.data)
	assert.Equal(t, "y", list.head.next.data)
	assert.Equal(t, "a", list.head.next.next.data)
	assert.Equal(t, "b", list.head.next.next.next.data)
	assert.Equal(t, "b", list.tail.data)
}

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

func TestDoublyLinkedList_ReadFromFront(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")

	res, err := list.ReadFromFront()
	assert.Nil(t, err)
	assert.Equal(t, "a", res)
}

func TestDoublyLinkedList_ReadFromFrontShouldErrorEmptyList(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	_, err := list.ReadFromFront()
	assert.NotNil(t, err)
}

func TestDoublyLinkedList_ReadFromEnd(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")

	res, err := list.ReadFromEnd()
	assert.Nil(t, err)
	assert.Equal(t, "c", res)
}

func TestDoublyLinkedList_ReadFromEndShouldErrorEmptyList(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	_, err := list.ReadFromEnd()
	assert.NotNil(t, err)
}

func TestDoublyLinkedList_ReadByIndex(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")

	res, err := list.ReadByIndex(0)
	assert.Nil(t, err)
	assert.Equal(t, "a", res)
	res, err = list.ReadByIndex(1)
	assert.Nil(t, err)
	assert.Equal(t, "b", res)
	res, err = list.ReadByIndex(2)
	assert.Nil(t, err)
	assert.Equal(t, "c", res)
}

func TestDoublyLinkedList_ReadByIndexShouldErrorWhenEmptyList(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	_, err := list.ReadByIndex(0)
	assert.NotNil(t, err)
}

func TestDoublyLinkedList_ReadByIndexShouldErrorWhenIndexNotFound(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")

	_, err := list.ReadByIndex(3)
	assert.NotNil(t, err)
}

func TestLinkedList_Search(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")

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
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")

	idx, err := list.Search("z")
	assert.NotNil(t, err)
	assert.Equal(t, -1, idx)
}

func TestLinkedList_DeleteHead(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")
	value := list.DeleteFromFront()
	assert.Nil(t, value)

	front, err := list.ReadFromFront()
	assert.Nil(t, err)

	assert.Equal(t, "b", front)
	value = list.DeleteFromFront()
	assert.Nil(t, value)

	front, err = list.ReadFromFront()
	assert.Equal(t, "c", front)
}

func TestLinkedList_DeleteTail(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")
	value := list.DeleteFromEnd()
	assert.Nil(t, value)

	front, err := list.ReadFromEnd()
	assert.Nil(t, err)

	assert.Equal(t, "b", front)
	value = list.DeleteFromEnd()
	assert.Nil(t, value)

	front, err = list.ReadFromEnd()
	assert.Equal(t, "a", front)
}

func TestLinkedList_DeleteErrorWhenIndexNotFound(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")
	err := list.DeleteByIndex(3)
	assert.NotNil(t, err)
}

func TestLinkedList_DeleteFrontByIndex(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")
	err := list.DeleteByIndex(0)
	assert.Nil(t, err)
	assert.Equal(t, "b", list.head.data)
	err = list.DeleteByIndex(0)
	assert.Nil(t, err)
	assert.Equal(t, "c", list.head.data)
}

func TestLinkedList_DeleteMiddleByIndex(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")
	err := list.DeleteByIndex(1)
	assert.Nil(t, err)
	assert.Equal(t, "a", list.head.data)
	assert.Equal(t, "c", list.head.next.data)
	assert.Equal(t, "c", list.tail.data)
}

func TestLinkedList_DeleteEndByIndex(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")
	err := list.DeleteByIndex(2)
	assert.Nil(t, err)
	assert.Equal(t, "a", list.head.data)
	assert.Equal(t, "b", list.tail.data)
	assert.Equal(t, "b", list.head.next.data)
	err = list.DeleteByIndex(0)
	assert.Nil(t, err)
	assert.Equal(t, "b", list.head.data)
}

//func TestLinkedList_DeleteItems(t *testing.T) {
//	list := NewDoublyLinkedList[string]()
//
//	list.Add("this is a password:1 that needs obfuscating")
//	list.Add("this is a password:12 that needs obfuscating")
//	list.Add("this is a password:123 that needs obfuscating")
//	list.Add("this is a password:12cae that needs obfuscating")
//	list.Add("this is a password:some_password that needs obfuscating")
//	list.Add("a")
//	list.Add("b")
//
//	predicate := func(data string) bool {
//		pattern := regexp.MustCompile("\\bpassword:[\\w]+\\b")
//		return pattern.MatchString(data)
//	}
//	list.DeleteItems(predicate)
//	first, _ := list.Read(0)
//	assert.Equal(t, first, "a")
//	second, _ := list.Read(1)
//	assert.Equal(t, second, "b")
//}
