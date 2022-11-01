package queue

import (
	"errors"
	"fmt"
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

func New[T any](data ...T) AllowedQueueFunc[T] {
	d := make([]T, 0, len(data))
	d = append(d, data...)
	res := queue[T]{
		data: d,
	}
	return &res
}

type queue[T any] struct {
	data []T
}

func (q *queue[T]) Enqueue(value T) {
	q.data = append(q.data, value)
}

func (q *queue[T]) Dequeue() (T, error) {
	var res T
	if len(q.data) < 1 {
		return res, ErrEmpty
	}
	res = q.data[0]
	q.data = q.data[1:]
	return res, nil
}

func (q *queue[T]) Read() (T, error) {
	var res T
	if len(q.data) < 1 {
		return res, ErrEmpty
	}
	return q.data[0], nil
}

func (q *queue[T]) Size() int {
	return len(q.data)
}

func (q *queue[T]) Print() string {
	return fmt.Sprintf("%v", q.data)
}
