package doublylinkedlist

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedList_AddToFront(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")

	list.AddToFront("y")
	assert.Equal(t, 3, list.Size())
	assert.Equal(t, "y", list.Head().data)
	assert.Equal(t, "a", list.Head().next.data)
	assert.Equal(t, "b", list.Head().next.next.data)
	assert.Equal(t, "b", list.Tail().data)

	list.AddToFront("z")
	assert.Equal(t, 4, list.Size())
	assert.Equal(t, "z", list.Head().data)
	assert.Equal(t, "y", list.Head().next.data)
	assert.Equal(t, "a", list.Head().next.next.data)
	assert.Equal(t, "b", list.Head().next.next.next.data)
	assert.Equal(t, "b", list.Tail().data)
}

func TestDoublyLinkedList_AddAtEnd(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.AddAtEnd(1)
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 1, list.Head().data)
	assert.Equal(t, 1, list.Tail().data)

	list.AddAtEnd(2)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 1, list.Head().data)
	assert.Equal(t, 2, list.Tail().data)

	assert.Equal(t, 2, list.Head().next.data)
}

func TestDoublyLinkedList_AddByIndexFront(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")

	list.AddByIndex(0, "z")
	assert.Equal(t, 4, list.Size())
	assert.Equal(t, "z", list.Head().data)
	assert.Equal(t, "c", list.Tail().data)
	assert.Equal(t, "a", list.Head().next.data)
	assert.Equal(t, "b", list.Head().next.next.data)
	assert.Equal(t, "c", list.Head().next.next.next.data)
}

func TestDoublyLinkedList_AddByIndexMid(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")

	list.AddByIndex(1, "z")
	assert.Equal(t, 4, list.Size())
	assert.Equal(t, "a", list.Head().data)
	assert.Equal(t, "c", list.Tail().data)
	assert.Equal(t, "z", list.Head().next.data)
	assert.Equal(t, "b", list.Head().next.next.data)
	//assert.Equal(t, "c", list.Head().next.next.next.data)
}

func TestDoublyLinkedList_AddByIndexMidEnd(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")

	list.AddByIndex(2, "z")
	assert.Equal(t, 4, list.Size())
	assert.Equal(t, "a", list.Head().data)
	assert.Equal(t, "c", list.Tail().data)
	assert.Equal(t, "b", list.Head().next.data)
	assert.Equal(t, "z", list.Head().next.next.data)
	assert.Equal(t, "c", list.Head().next.next.next.data)
}

func TestDoublyLinkedList_AddByIndexEnd(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")

	list.AddByIndex(3, "z")
	assert.Equal(t, 4, list.Size())
	assert.Equal(t, "a", list.Head().data)
	assert.Equal(t, "z", list.Tail().data)
	assert.Equal(t, "b", list.Head().next.data)
	assert.Equal(t, "c", list.Head().next.next.data)
	assert.Equal(t, "z", list.Head().next.next.next.data)
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

func TestLinkedList_RemoveFromFrontSingle(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")

	value, err := list.RemoveFromFront()
	assert.Nil(t, err)
	assert.Equal(t, 0, list.Size())
	assert.Equal(t, "a", value)
}

func TestLinkedList_RemoveFromFrontByIndex(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")
	value, err := list.RemoveByIndex(0)
	assert.Nil(t, err)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "a", value)

	value, err = list.RemoveByIndex(0)
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, "b", value)

	value, err = list.RemoveByIndex(0)
	assert.Equal(t, 0, list.Size())
	assert.Equal(t, "c", value)
}

func TestLinkedList_RemoveFromFront(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")
	value, err := list.RemoveFromFront()
	assert.Nil(t, err)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "a", value)

	value, err = list.RemoveFromFront()
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, "b", value)

	value, err = list.RemoveFromFront()
	assert.Equal(t, 0, list.Size())
	assert.Equal(t, "c", value)
}

func TestLinkedList_RemoveFromEnd(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")
	value, err := list.RemoveFromEnd()
	assert.Nil(t, err)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "c", value)

	value, err = list.RemoveFromEnd()
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, "b", value)

	value, err = list.RemoveByIndex(0)
	assert.Equal(t, 0, list.Size())
	assert.Equal(t, "a", value)
}

func TestLinkedList_RemoveFromEndByIndex(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")
	value, err := list.RemoveByIndex(2)
	assert.Nil(t, err)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "c", value)

	value, err = list.RemoveByIndex(1)
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, "b", value)

	value, err = list.RemoveByIndex(0)
	assert.Equal(t, 0, list.Size())
	assert.Equal(t, "a", value)
}

func TestLinkedList_RemoveFromMiddleByIndex(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")
	list.AddAtEnd("d")

	value, err := list.RemoveByIndex(2)
	assert.Nil(t, err)
	assert.Equal(t, 3, list.Size())
	assert.Equal(t, "c", value)

	value, err = list.RemoveByIndex(1)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "b", value)

	value, err = list.RemoveByIndex(1)
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, "d", value)

	value, err = list.RemoveByIndex(0)
	assert.Equal(t, 0, list.Size())
	assert.Equal(t, "a", value)
}

func TestLinkedList_DeleteAllListHead(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	value := list.DeleteFromFront()
	assert.Nil(t, value)
	assert.Equal(t, 1, list.Size())

	front, err := list.ReadFromFront()
	assert.Nil(t, err)

	assert.Equal(t, "b", front)
	value = list.DeleteFromFront()
	assert.Nil(t, value)
	assert.Equal(t, 0, list.Size())
}

func TestLinkedList_DeleteHead(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")
	value := list.DeleteFromFront()
	assert.Nil(t, value)
	assert.Equal(t, 2, list.Size())

	front, err := list.ReadFromFront()
	assert.Nil(t, err)

	assert.Equal(t, "b", front)
	value = list.DeleteFromFront()
	assert.Nil(t, value)
	assert.Equal(t, 1, list.Size())

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
	assert.Equal(t, 2, list.Size())

	front, err := list.ReadFromEnd()
	assert.Nil(t, err)

	assert.Equal(t, "b", front)
	value = list.DeleteFromEnd()
	assert.Nil(t, value)
	assert.Equal(t, 1, list.Size())

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
	assert.Equal(t, 3, list.Size())
}

func TestLinkedList_DeleteFrontByIndex(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")
	err := list.DeleteByIndex(0)
	assert.Nil(t, err)
	assert.Equal(t, 2, list.Size())

	assert.Equal(t, "b", list.Head().data)
	err = list.DeleteByIndex(0)
	assert.Equal(t, 1, list.Size())

	assert.Nil(t, err)
	assert.Equal(t, "c", list.Head().data)
}

func TestLinkedList_DeleteMiddleByIndex(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")
	err := list.DeleteByIndex(1)
	assert.Nil(t, err)
	assert.Equal(t, "a", list.Head().data)
	assert.Equal(t, "c", list.Head().next.data)
	assert.Equal(t, "c", list.Tail().data)
	assert.Equal(t, 2, list.Size())
}

func TestLinkedList_DeleteEndByIndex(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("a")
	list.AddAtEnd("b")
	list.AddAtEnd("c")
	err := list.DeleteByIndex(2)
	assert.Nil(t, err)
	assert.Equal(t, "a", list.Head().data)
	assert.Equal(t, "b", list.Tail().data)
	assert.Equal(t, "b", list.Head().next.data)
	assert.Equal(t, 2, list.Size())
}

func TestLinkedList_DeleteItemsHead(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.AddAtEnd("1 this is a password:1 that needs obfuscating")
	list.AddAtEnd("2 a")
	list.AddAtEnd("3 b")

	hasPassword := func(data string) bool {
		pattern := regexp.MustCompile("\\bpassword:[\\w]+\\b")
		return pattern.MatchString(data)
	}
	list.DeleteItems(hasPassword)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "2 a", list.Head().data)
	assert.Equal(t, "3 b", list.Tail().data)

	first, _ := list.ReadByIndex(0)
	assert.Equal(t, first, "2 a")
	second, _ := list.ReadByIndex(1)
	assert.Equal(t, second, "3 b")
}

func TestLinkedList_DeleteItemsEnd(t *testing.T) {
	list := NewDoublyLinkedList[string]()

	list.AddAtEnd("1 a")
	list.AddAtEnd("2 b")
	list.AddAtEnd("3 this is a password:1 that needs obfuscating")

	hasPassword := func(data string) bool {
		pattern := regexp.MustCompile("\\bpassword:[\\w]+\\b")
		return pattern.MatchString(data)
	}
	list.DeleteItems(hasPassword)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "1 a", list.Head().data)
	assert.Equal(t, "2 b", list.Tail().data)

	first, _ := list.ReadByIndex(0)
	assert.Equal(t, first, "1 a")
	second, _ := list.ReadByIndex(1)
	assert.Equal(t, second, "2 b")
}

func TestLinkedList_DeleteItemsMid(t *testing.T) {
	list := NewDoublyLinkedList[string]()

	list.AddAtEnd("1 a")
	list.AddAtEnd("2 this is a password:1 that needs obfuscating")
	list.AddAtEnd("3 this is a password:1 that needs obfuscating")
	list.AddAtEnd("4 b")

	hasPassword := func(data string) bool {
		pattern := regexp.MustCompile("\\bpassword:[\\w]+\\b")
		return pattern.MatchString(data)
	}
	list.DeleteItems(hasPassword)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "1 a", list.Head().data)
	assert.Equal(t, "4 b", list.Tail().data)

	first, _ := list.ReadByIndex(0)
	assert.Equal(t, first, "1 a")
	second, _ := list.ReadByIndex(1)
	assert.Equal(t, second, "4 b")
}

func TestLinkedList_DeleteItems(t *testing.T) {
	list := NewDoublyLinkedList[string]()

	list.AddAtEnd("1 this is a password:1 that needs obfuscating")
	list.AddAtEnd("2 this is a password:12 that needs obfuscating")
	list.AddAtEnd("3 this is a password:123 that needs obfuscating")
	list.AddAtEnd("4 this is a password:12cae that needs obfuscating")
	list.AddAtEnd("5 this is a password:some_password that needs obfuscating")
	list.AddAtEnd("6 a")
	list.AddAtEnd("7 b")
	list.AddAtEnd("8 this is a password:12cae that needs obfuscating")
	list.AddAtEnd("9 c")
	list.AddAtEnd("10 this is a password:some_password that needs obfuscating")

	hasPassword := func(data string) bool {
		pattern := regexp.MustCompile("\\bpassword:[\\w]+\\b")
		return pattern.MatchString(data)
	}
	list.DeleteItems(hasPassword)
	assert.Equal(t, 3, list.Size())
	assert.Equal(t, "6 a", list.Head().data)
	assert.Equal(t, "9 c", list.Tail().data)

	first, _ := list.ReadByIndex(0)
	assert.Equal(t, first, "6 a")
	second, _ := list.ReadByIndex(1)
	assert.Equal(t, second, "7 b")
	third, _ := list.ReadByIndex(2)
	assert.Equal(t, third, "9 c")
}
