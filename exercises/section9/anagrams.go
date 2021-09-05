package section9

import (
	"reflect"
	"regexp"
	"strings"
)

func AnagramChecker(inWordA, inWordB string) bool {
	// if empty return
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

	wordAMap := make(map[rune]int)
	for _, char := range []rune(wordA) {
		if _, ok := wordAMap[char]; ok {
			wordAMap[char] = wordAMap[char] + 1
		} else {
			wordAMap[char] = 1
		}
	}
	wordBMap := make(map[rune]int)
	for _, char := range []rune(wordB) {
		if _, ok := wordBMap[char]; ok {
			wordBMap[char] = wordBMap[char] + 1
		} else {
			wordBMap[char] = 1
		}
	}

	return reflect.DeepEqual(wordAMap, wordBMap)
}

func eitherWordIsEmpty(inWordA string, inWordB string) bool {
	return len(inWordA) <= 0 || len(inWordB) <= 0
}
