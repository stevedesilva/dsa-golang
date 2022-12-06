package recursion

import "errors"

var ErrEmptyInput = errors.New("minimum input of one word required")

func StringReversal(word string) (string, error) {
	if len(word) < 1 {
		return "", ErrEmptyInput
	}
	return stringReversalRec([]rune(word)), nil
}

func stringReversalRec(array []rune) string {
	if len(array) == 0 {
		return ""
	}
	return stringReversalRec(array[1:]) + string(array[0])
}
