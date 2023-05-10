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

func New[T comparable]() AllowedQueueFunc[T] {
	res := queue[T]{
		data: doublylinkedlist.NewDoublyLinkedList[T](),
	}
	return &res
}

type queue[T comparable] struct {
	data doublylinkedlist.AllowedDoublyLinkedListFunc[T]
}

func (q *queue[T]) Read() (T, error) {
	return q.data.ReadFromFront()
}

func (q *queue[T]) Enqueue(value T) {
	q.data.AddAtEnd(value)
}

func (q *queue[T]) Dequeue() (T, error) {
	return q.data.RemoveFromFront()
}

func (q *queue[T]) Size() int {
	return q.data.Size()
}

func (q *queue[T]) Print() string {
	data := q.data.PrintItems()
	return fmt.Sprintf("%v", data)
}
