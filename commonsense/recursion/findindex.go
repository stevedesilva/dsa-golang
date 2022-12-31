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

func FindFirstIndexAlt(word string, letter rune) (int, error) {
	return findFirstIndexAlt([]rune(word), letter)
}

func findFirstIndexAlt(word []rune, letter rune) (int, error) {
	if len(word) == 0 {
		return 0, errors.New("not found")
	}
	if word[0] == letter {
		return 0, nil
	}
	val, err := findFirstIndexAlt(word[1:], letter)
	if err == nil {
		val += 1
	}
	return val, err
}
