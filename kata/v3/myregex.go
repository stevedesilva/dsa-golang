package v3

import (
	"fmt"
	"regexp"
)

func MatchMyRegex(pattern, word string) bool {
	var exp *regexp.Regexp = regexp.MustCompile(pattern)
	var result string = exp.FindString(word)
	fmt.Println(result)
	return true
}
