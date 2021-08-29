package section4

import "testing"

func TestPalindrome(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "a", input: "a", want: true},
		{name: "", input: "", want: true},
		{name: "ab", input: "ab", want: false},
		{name: "abc", input: "abc", want: false},
		{name: "aba", input: "aba", want: true},
		{name: "abba", input: "abba", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Palindrome(tt.input); got != tt.want {
				t.Errorf("Palindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
