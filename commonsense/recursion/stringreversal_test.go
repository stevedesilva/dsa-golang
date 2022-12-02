package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringReversal(t *testing.T) {

	tests := []struct {
		name string
		word string
		want string
	}{
		{
			name: "a",
			word: "a",
			want: "a",
		},
		{
			name: "ab",
			word: "ab",
			want: "ba",
		},
		{
			name: "abc",
			word: "abc",
			want: "cba",
		},
		{
			name: "abcd",
			word: "abcd",
			want: "dcba",
		},
		{
			name: "abcde",
			word: "abcde",
			want: "edcba",
		},
		{
			name: "abcdef",
			word: "abcdef",
			want: "fedcba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, StringReversal(tt.word), "StringReversal(%v)", tt.word)
		})
	}
}
