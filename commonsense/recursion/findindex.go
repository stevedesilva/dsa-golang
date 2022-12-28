package recursion

import "errors"

func FindFirstIndex(word string, letter rune) (int, error) {

	return findFirstIndex([]rune(word), letter, 0)
}

func findFirstIndex(word []rune, letter rune, index int) (int, error) {
	if len(word) == index {
		return 0, errors.New("not found")
	}
	if word[index] == letter {
		return index, nil
	}
	return findFirstIndex(word, letter, index+1)
}
