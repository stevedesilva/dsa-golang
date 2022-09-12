package array

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		word string
		want bool
	}{
		{
			name: "₼aRa₼",
			word: "₼aRa₼",
			want: true,
		},
		{
			name: "a",
			word: "a",
			want: true,
		},
		{
			name: "ADa",
			word: "ADa",
			want: true,
		},
		{
			name: "KayAk",
			word: "KayAk",
			want: true,
		},
		{
			name: "kayak",
			word: "kayak",
			want: true,
		},
		{
			name: "Mom",
			word: "Mom",
			want: true,
		},
		{
			name: "NoMatch",
			word: "NoMatch",
			want: false,
		},
		{
			name: "Rampage",
			word: "Rampage",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.word); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
