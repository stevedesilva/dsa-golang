package recursion

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountLetter(t *testing.T) {
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
			name: "error minimum value word ",
			args: args{
				word:   "",
				letter: 'f',
			},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name: "count x's",
			args: args{
				word:   "abcxef",
				letter: 'x',
			},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name: "count a's",
			args: args{
				word:   "aabbccddeeffggaabbccddeeffgg",
				letter: 'a',
			},
			want:    4,
			wantErr: assert.NoError,
		},
		{
			name: "count f's",
			args: args{
				word:   "aabbccddeeffggaabbccddeeffgg",
				letter: 'f',
			},
			want:    4,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountLetter(tt.args.word, tt.args.letter)
			if !tt.wantErr(t, err, fmt.Sprintf("CountLetter(%v, %v)", tt.args.word, tt.args.letter)) {
				return
			}
			assert.Equalf(t, tt.want, got, "CountLetter(%v, %v)", tt.args.word, tt.args.letter)
		})
	}
}
