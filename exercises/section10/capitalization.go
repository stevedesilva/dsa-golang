package section10

import (
	"fmt"
	"strings"
)

func CapitalizeSentence(sentence string) string {
	split := strings.Split(sentence, " ")
	result := make([]string, 0, len(split))

	for _, w := range split {
		word := []rune(w)
		head := strings.ToUpper(string(word[:1]))
		tail := string(word[1:])
		result = append(result, fmt.Sprintf("%s%s", head, tail))
	}
	return strings.Join(result, " ")

}
