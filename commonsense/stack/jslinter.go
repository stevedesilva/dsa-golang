package stack

import "errors"

var (
	ErrMissingClosingBrace = errors.New("opening brace does not have a corresponding closing brace")
	ErrMissingOpeningBrace = errors.New("closing brace doe not have a corresponding opening brace")
	ErrBraceMismatch       = errors.New("brace ordering is not correct")
	ErrInvalidInput        = errors.New("minimum input data required")
)

type Linter interface {
	Validate() (bool, error)
}

func NewLinter[T AllowedStackTypes](data []T) Linter {
	// convert string to slice

	return &braceLinter[T]{
		data: data,
	}
}

type braceLinter[T AllowedStackTypes] struct {
	data []T
}

func (b *braceLinter[T]) Validate() (bool, error) {
	// move on if character isn't type of brace
	//for i, v := range b.data {
	//
	//}
	return true, nil
}
