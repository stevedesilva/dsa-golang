package stack

import (
	"errors"
	"strings"
)

var ErrMinimumInputNotMet = errors.New("minimum input not met")

func reverse(data string) (string, error) {
	if len(data) == 0 {
		return "", ErrMinimumInputNotMet
	}
	s := stack[rune]{}
	for _, d := range data {
		s.Push(d)
	}
	sb := strings.Builder{}
	for i := 0; s.Size() > 0; i++ {
		r, err := s.Pop()
		if err == nil {
			sb.WriteRune(r)
		}
	}
	return sb.String(), nil
}
