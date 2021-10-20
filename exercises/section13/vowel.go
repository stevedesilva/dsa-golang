package section13

import (
	"regexp"
	"strings"
)

type Vowel struct {
	Input string
}

func (v *Vowel) CalculateNumberOfVowelsIterativeUsingMap() int {
	vowels := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}

	// range over runes
	count := 0
	lowerCaseInput := strings.ToLower(v.Input)
	for _, v := range lowerCaseInput {
		if _, ok := vowels[string(v)]; ok {
			count++
		}
	}

	return count
}

func (v *Vowel) CalculateNumberOfVowelsIterativeUsingContains() int {
	vowels := "aeiou"
	// range over runes
	count := 0
	lowerCaseInput := strings.ToLower(v.Input)
	for _, v := range lowerCaseInput {
		if ok := strings.ContainsRune(vowels, v); ok {
			count++
		}
	}

	return count
}

func (v *Vowel) CalculateNumberOfVowelsRegex() int {
	rx, _ := regexp.Compile("(?i)[aeiou]") //same as regexp.Compile("a|e|i|o|u|A|E|I|O|U")

	allString := rx.FindAllString(v.Input, -1)
	return len(allString)
}
