package section18

import "errors"

type Any interface{}

var (
	EmptyStackErr = errors.New("empty stack")
)

type Stack interface {
	Peek() (Any, error)
	Pop() (Any, error)
	Push(Any)
	Size() int
}

func New() Stack {
	return &StackImpl{
		stack: []Any{},
	}
}

type StackImpl struct {
	stack []Any
}

func (s *StackImpl) Pop() (Any, error) {
	if len(s.stack) <= 0 {
		return nil, EmptyStackErr
	}
	r := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return r, nil
}

func (s *StackImpl) Peek() (Any, error) {
	if len(s.stack) <= 0 {
		return nil, EmptyStackErr
	}
	return s.stack[len(s.stack)-1], nil
}

func (s *StackImpl) Push(any Any) {
	s.stack = append(s.stack, any)
}
func (s *StackImpl) Size() int {
	return len(s.stack)
}
