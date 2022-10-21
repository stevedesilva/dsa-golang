package stack

import (
	"errors"
	"fmt"
)

var (
	ErrEmpty = errors.New("empty stack")
)

type AllowedFunc interface {
	Push(string)
	Pop() (string, error)
	Peek() (string, error)
	Print() string
	Size() int
}

func New(data ...string) AllowedFunc {
	d := make([]string, 0, len(data))
	d = append(d, data...)
	return &stack{
		data: d,
	}
}

type stack struct {
	data []string
}

func (s *stack) Print() string {
	return fmt.Sprintf("%v", s.data)
}

func (s *stack) Push(s2 string) {
	s.data = append(s.data, s2)
}

func (s *stack) Pop() (string, error) {
	if s.Size() < 1 {
		return "", ErrEmpty
	}
	return s.data[len(s.data)-1], nil
}

func (s *stack) Peek() (string, error) {
	if s.Size() < 1 {
		return "", ErrEmpty
	}
	return s.data[len(s.data)-1], nil
}

func (s *stack) Size() int {
	return len(s.data)
}
