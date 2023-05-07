package queue

import (
	"errors"
	"fmt"
	"github.com/stevedesilva/dsa-golang.git/commonsense/linkedlist/doublylinkedlist"
)

var (
	ErrEmpty = errors.New("empty queue")
)

type AllowedQueueFunc[T comparable] interface {
	Enqueue(value T)
	Dequeue() (T, error)
	Read() (T, error)
	Size() int
	Print() string
}

type Queue[T comparable] struct {
	data *doublylinkedlist.DoublyLinkedList[T]
}

func New[T comparable]() AllowedQueueFunc[T] {
	res := Queue[T]{
		data: doublylinkedlist.NewDoublyLinkedList[T](),
	}
	return &res
}
func NewQueueAsStruct[T comparable]() *Queue[T] {
	res := Queue[T]{
		data: doublylinkedlist.NewDoublyLinkedList[T](),
	}
	return &res
}

func (q *Queue[T]) Read() (T, error) {
	return q.data.ReadFromFront()
}

func (q *Queue[T]) Enqueue(value T) {
	q.data.AddAtEnd(value)
}

func (q *Queue[T]) Dequeue() (T, error) {
	return q.data.RemoveFromFront()
}

func (q *Queue[T]) Size() int {
	return q.data.Size()
}

func (q *Queue[T]) Print() string {
	data := q.data.PrintItems()
	return fmt.Sprintf("%v", data)
}
