package array

import (
	"fmt"
	"unicode"
)

func IsPalindrome(word string) bool {
	wordAsRune := []rune(word)
	start := 0
	end := len(wordAsRune) - 1
	for start < end {
		a := wordAsRune[start]
		b := wordAsRune[end]

		fmt.Printf("1 - rune a: %v \trune b:%v\n", a, b)
		a = unicode.ToLower(a)
		b = unicode.ToLower(b)
		fmt.Printf("2 - rune a: %v \trune b:%v\n", a, b)
		if a != b {
			return false
		}
		start++
		end--
	}
	return true
}
