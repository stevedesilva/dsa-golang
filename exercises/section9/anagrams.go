package section9

import (
	"reflect"
	"regexp"
	"strings"
)

func AnagramChecker(inWordA, inWordB string) bool {
	if eitherWordIsEmpty(inWordA, inWordB) {
		return false
	}

	regexp, err := regexp.Compile("\\W")
	if err != nil {
		// to return error
		return false
	}

	wordA := regexp.ReplaceAllString(strings.ToLower(inWordA), "")
	wordB := regexp.ReplaceAllString(strings.ToLower(inWordB), "")
	if len(wordA) != len(wordB) {
		return false
	}

	wordAMap := createWordMap(wordA)
	wordBMap := createWordMap(wordB)

	return reflect.DeepEqual(wordAMap, wordBMap)
}

func createWordMap(word string) map[rune]int {
	wordMap := make(map[rune]int)
	for _, char := range word {
		if _, ok := wordMap[char]; ok {
			wordMap[char] = wordMap[char] + 1
		} else {
			wordMap[char] = 1
		}
	}
	return wordMap
}

func eitherWordIsEmpty(inWordA string, inWordB string) bool {
	return len(inWordA) <= 0 || len(inWordB) <= 0
}
