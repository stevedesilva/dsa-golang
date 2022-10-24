package stack

import "errors"

var (
	ErrMissingClosingBrace = errors.New("opening brace does not have a corresponding closing brace")
	ErrMissingOpeningBrace = errors.New("closing brace does not have a corresponding opening brace")
	ErrBraceMismatch       = errors.New("brace ordering is not correct")
	ErrInvalidInput        = errors.New("minimum input data required")
)

type Linter interface {
	Validate() (bool, error)
}

func NewLinter(data string) Linter {
	// convert string to slice

	return &braceLinter{
		data: data,
	}
}

type braceLinter struct {
	data string
}

func (b *braceLinter) Validate() (bool, error) {

	openBraces := map[rune]bool{
		'{': true,
		'[': true,
		'(': true,
	}
	closingBraces := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	// move on if character isn't type of brace
	s := stack[rune]{}
	for _, ch := range b.data {
		if openBraces[ch] {
			s.Push(ch)
		} else if openingBraceMatch, found := closingBraces[ch]; found {
			val, err := s.Pop()
			if err != nil {
				// empty stack
				if errors.Is(err, ErrEmpty) {
					return false, ErrMissingOpeningBrace
				}
				// some unknown error
				return false, err
			}
			// check if ch matches
			if val != openingBraceMatch {
				return false, ErrBraceMismatch
			}
		}
	}

	// stack should be empty
	if s.Size() > 0 {
		return false, ErrMissingClosingBrace
	}
	return true, nil
}
