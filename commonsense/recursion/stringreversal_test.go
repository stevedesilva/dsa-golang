package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringReversal(t *testing.T) {

	tests := []struct {
		name    string
		word    string
		want    string
		wantErr error
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
		{
			name:    "empty",
			word:    "",
			want:    "",
			wantErr: ErrEmptyInput,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reversal, err := StringReversal(tt.word)
			assert.Equalf(t, tt.want, reversal, "StringReversal(%v)", tt.word)
			assert.Equalf(t, tt.wantErr, err, "StringReversal Error(%v)", tt.word)
		})
	}
}
