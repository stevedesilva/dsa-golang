package stack

import (
	"errors"
	"fmt"
)

var (
	ErrEmpty = errors.New("empty stack")
)

type stackable interface {
	string | int
}

type AllowedFunc[T stackable] interface {
	Push(T)
	Pop() (T, error)
	Peek() (T, error)
	Print() string
	Size() int
}

func New[T stackable](data ...T) AllowedFunc[T] {
	d := make([]T, 0, len(data))
	d = append(d, data...)
	res := stack[T]{
		data: d,
	}
	return &res
}

type stack[T stackable] struct {
	data []T
}

func (s *stack[T]) Print() string {
	return fmt.Sprintf("%v", s.data)
}

func (s *stack[T]) Push(s2 T) {
	s.data = append(s.data, s2)
}

func (s *stack[T]) Pop() (T, error) {
	if s.Size() < 1 {
		return nil, ErrEmpty
	}
	i := len(s.data) - 1
	item := s.data[i]
	s.data = s.data[:i]
	return item, nil
}

func (s *stack[T]) Peek() (T, error) {
	if s.Size() < 1 {
		return nil, ErrEmpty
	}
	return s.data[len(s.data)-1], nil
}

func (s *stack[T]) Size() int {
	return len(s.data)
}
