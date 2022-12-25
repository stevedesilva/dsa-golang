package recursion

import (
	"errors"
	"fmt"
)

func CountLetter(word string, letter rune) (int, error) {
	if len(word) < 1 {
		return 0, errors.New("minimum word length not met")
	}
	res := countLetter([]rune(word), letter)
	return res, nil
}

func countLetter(word []rune, letter rune) int {
	if len(word) == 0 {
		return 0
	}
	if word[0] == letter {
		// add 1 to the letters
		return countLetter(word[1:], letter) + 1
	} else {
		return countLetter(word[1:], letter)
	}
}

func CountLetterInArray(words []string) int {
	fmt.Printf("word %s \n", words)
	if len(words) <= 0 {
		return 0
	}
	if len(words) == 1 {
		return len([]rune(words[0]))
	}
	other := words[1:]
	return len([]rune(words[0])) + CountLetterInArray(other)
}
