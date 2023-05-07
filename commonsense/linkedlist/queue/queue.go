package queue

import (
	"errors"
	"fmt"
	"github.com/stevedesilva/dsa-golang.git/commonsense/linkedlist/doublylinkedlist"
)

var (
	ErrEmpty = errors.New("empty queue")
)

//type AllowedQueueFunc[T any] interface {
//	Enqueue(value T)
//	Dequeue() (T, error)
//	Read() (T, error)
//	Size() int
//	Print() string
//}

//
//func New[T any](data ...T) AllowedQueueFunc[T] {
//	d := make([]T, 0, len(data))
//	d = append(d, data...)
//	res := queue[T]{
//		data: d,
//	}
//	return &res
//}

type Queue[T comparable] struct {
	data *doublylinkedlist.DoublyLinkedList[T]
}

func New[T comparable]() *Queue[T] {
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
