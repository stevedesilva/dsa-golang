package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAllPossibleAnagrams(t *testing.T) {

	tests := []struct {
		name string
		word string
		want []string
	}{
		{
			name: "a",
			word: "a",
			want: []string{"a"},
		},
		{
			name: "ab",
			word: "ab",
			want: []string{"ab", "ba"},
		},
		{
			name: "abc",
			word: "abc",
			want: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FindAllPossibleAnagrams(tt.word), "FindAllPossibleAnagrams(%v)", tt.word)
		})
	}
}
