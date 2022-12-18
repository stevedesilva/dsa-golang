package recursion

import "strings"

func FindAllPossibleAnagrams(word string) []string {
	// create a collection to hold all the anagrams
	res := make([]string, 0)

	// return char if len == 1
	if len(word) == 1 {
		return append(res, word)
	}

	// get all substring anagrams word e.g. abc - note, string([]rune{...}) converts []runes to a string
	wordArray := []rune(word)
	// e.g. a
	firstChr := string(wordArray[:1])
	// e.g. bc
	subWord := string(wordArray[1:])
	substringAnagrams := FindAllPossibleAnagrams(subWord)
	// iterate over each substring bc, cb
	for _, subAnagram := range substringAnagrams {

		newSubStringAnagramLen := len(subAnagram) + 1
		for i := 0; i < newSubStringAnagramLen; i++ {
			subRes := make([]string, 0, newSubStringAnagramLen)
			start := subAnagram[:i]
			end := subAnagram[i:]

			subRes = append(subRes, start, firstChr, end)
			joined := strings.Join(subRes, "")
			res = append(res, joined)
		}
	}

	return res
}
