package recursion

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	word   string
	letter rune
}

var testData = []struct {
	name    string
	args    args
	want    int
	wantErr assert.ErrorAssertionFunc
}{
	{
		name: "find single letter a",
		args: args{
			word:   "a",
			letter: 'a',
		},
		want:    0,
		wantErr: assert.NoError,
	},
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
	{
		name: "find single letter x",
		args: args{
			word:   "abcdefghijklmnopqrstuvwxyz",
			letter: 'x',
		},
		want:    23,
		wantErr: assert.NoError,
	},
}

func Test_findFirstIndex(t *testing.T) {
	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindFirstIndex(tt.args.word, tt.args.letter)
			if !tt.wantErr(t, err, fmt.Sprintf("FindFirstIndex(%v, %v)", tt.args.word, tt.args.letter)) {
				return
			}
			assert.Equalf(t, tt.want, got, "FindFirstIndex(%v, %v)", tt.args.word, tt.args.letter)
		})
	}
}

func Test_findFirstIndexAlt(t *testing.T) {
	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindFirstIndexAlt(tt.args.word, tt.args.letter)
			if !tt.wantErr(t, err, fmt.Sprintf("Test_findFirstIndexAlt(%v, %v)", tt.args.word, tt.args.letter)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Test_findFirstIndexAlt(%v, %v)", tt.args.word, tt.args.letter)
		})
	}
}
