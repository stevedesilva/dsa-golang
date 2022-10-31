package queue

import "fmt"

type AllowedFunc[T any] interface {
	Enqueue(value T)
	Dequeue() (T, error)
	Read() (T, error)
	Size() int
	Print() string
}

func New[T any](data ...T) AllowedFunc[T] {
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
	//TODO implement me
	panic("implement me")
}

func (q *queue[T]) Dequeue() (T, error) {
	//TODO implement me
	panic("implement me")
}

func (q *queue[T]) Read() (T, error) {
	//TODO implement me
	panic("implement me")
}

func (q *queue[T]) Size() int {
	//TODO implement me
	panic("implement me")
}

func (s *queue[T]) Print() string {
	return fmt.Sprintf("%v", s.data)
}
