package v2

import "regexp"

func isMatched(regex, value string) bool {
	compile, err := regexp.Compile(regex)
	if err != nil {
		return false
	}

	return compile.MatchString(value)
}
