package stack

import "fmt"

type AllowedFunc interface {
	Push(string)
	Pop() string
	Peek() string
	Print() string
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
	panic("implement me")
}

func (s *stack) Pop() string {
	panic("implement me")
}

func (s *stack) Peek() string {
	panic("implement me")
}
