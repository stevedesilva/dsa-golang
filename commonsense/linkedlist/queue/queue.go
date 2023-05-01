package queue

import (
	"errors"
	"fmt"
	"github.com/stevedesilva/dsa-golang.git/commonsense/linkedlist/doublylinkedlist"
)

var (
	ErrEmpty = errors.New("empty queue")
)

type AllowedQueueFunc[T any] interface {
	Enqueue(value T)
	Dequeue() (T, error)
	Read() (T, error)
	Size() int
	Print() string
}

//
//func New[T any](data ...T) AllowedQueueFunc[T] {
//	d := make([]T, 0, len(data))
//	d = append(d, data...)
//	res := queue[T]{
//		data: d,
//	}
//	return &res
//}

type queue[T comparable] struct {
	data *doublylinkedlist.DoublyLinkedList[T]
}

func New[T comparable]() AllowedQueueFunc[T] {
	res := queue[T]{
		data: doublylinkedlist.NewDoublyLinkedList[T](),
	}
	return &res
}

func (q *queue[T]) Enqueue(value T) {
	q.data.AddAtEnd(value)
}

func (q *queue[T]) Dequeue() (T, error) {
	return q.data.ReadFromEnd()
}

func (q *queue[T]) Read() (T, error) {
	var res T
	//if len(q.data) < 1 {
	//	return res, ErrEmpty
	//}
	//return q.data[0], nil
	return res, nil
}

func (q *queue[T]) Size() int {
	//return len(q.data)
	return 0
}

func (q *queue[T]) Print() string {
	return fmt.Sprintf("%v", q.data)
}
