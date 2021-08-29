package section4

import "strings"

func FindMax(input string) string {
	// remove spaces
	input = strings.ReplaceAll(input, " ", "")
	inputMap := make(map[rune]int)
	var (
		maxChar rune = -1
		maxVal       = -1
	)
	// to map
	for _, v := range input {
		if _, ok := inputMap[v]; ok {
			inputMap[v] = inputMap[v] + 1
		} else {
			inputMap[v] = 1
		}

		// update max
		if val := inputMap[v]; val > maxVal {
			maxVal = val
			maxChar = v
		}
	}

	return string(maxChar)
}
