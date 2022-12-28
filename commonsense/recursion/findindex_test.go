package recursion

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findFirstIndex(t *testing.T) {
	type args struct {
		word   string
		letter rune
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "cannot find single letter a",
			args: args{
				word:   "xxx",
				letter: 'a',
			},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name: "cannot find single letter f",
			args: args{
				word:   "xxdsdsewgwjowx",
				letter: 'f',
			},
			want:    0,
			wantErr: assert.Error,
		},

		{
			name: "find single letter a",
			args: args{
				word:   "abc",
				letter: 'a',
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "find single letter b",
			args: args{
				word:   "abc",
				letter: 'b',
			},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name: "find single letter c",
			args: args{
				word:   "abc",
				letter: 'c',
			},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name: "find single letter d",
			args: args{
				word:   "abcdefghijk",
				letter: 'd',
			},
			want:    3,
			wantErr: assert.NoError,
		},
		{
			name: "find single letter k",
			args: args{
				word:   "abcdefghijk",
				letter: 'k',
			},
			want:    10,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindFirstIndex(tt.args.word, tt.args.letter)
			if !tt.wantErr(t, err, fmt.Sprintf("FindFirstIndex(%v, %v)", tt.args.word, tt.args.letter)) {
				return
			}
			assert.Equalf(t, tt.want, got, "FindFirstIndex(%v, %v)", tt.args.word, tt.args.letter)
		})
	}
}
