package section4

func Palindrome(input string) bool {
	inputAsRune := []rune(input)
	for i, j := 0, len(inputAsRune)-1; i < j; i, j = i+1, j-1 {
		if inputAsRune[i] != inputAsRune[j] {
			return false
		}
	}
	return true
}
