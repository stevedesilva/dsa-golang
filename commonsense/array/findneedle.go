package array

import "errors"

var (
	ErrNeedleAndHaystackRequired = errors.New("needle and haystack cannot be empty")
)

func SearchHaystack(needle string, haystack string) (bool, error) {
	if len(needle) < 1 || len(haystack) < 1 {
		return false, ErrNeedleAndHaystackRequired
	}
	needleRunes := []rune(needle)
	haystackRunes := []rune(haystack)
	foundNeedle := false
	for i := 0; i < len(haystackRunes); i++ {
		for j := 0; j < len(needleRunes); j++ {
			foundNeedle = true
			if haystackRunes[i+j] != needleRunes[j] {
				// move to next letter
				foundNeedle = false
				break
			}
		}
		//
		if foundNeedle {
			break
		}
	}
	return foundNeedle, nil
}
