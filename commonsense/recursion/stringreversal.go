package recursion

import "errors"

var ErrEmptyInput = errors.New("minimum input of one word required")

func StringReversal(word string) (string, error) {
	if len(word) < 1 {
		return "", ErrEmptyInput
	}
	res := stringReversalRec([]rune(word))
	return string(res), nil
}

func stringReversalRec(array []rune) string {
	if len(array) == 1 {
		return string(array[0])
	}
	return stringReversalRec(array[1:]) + string(array[0])
}
