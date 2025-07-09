package optimise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortCharactersInAnyOrder(t *testing.T) {
	type args struct {
		characters []rune
	}
	tests := []struct {
		name string
		args args
		want map[rune]int
	}{
		{
			name: "Test with mixed characters",
			args: args{characters: []rune{'a', 'b', 'a', 'c', 'b', 'a'}},
			want: map[rune]int{'a': 3, 'b': 2, 'c': 1},
		},
		{
			name: "Test with single character",
			args: args{characters: []rune{'x'}},
			want: map[rune]int{'x': 1},
		},
		{
			name: "Test with repeated characters",
			args: args{characters: []rune{'z', 'z', 'z'}},
			want: map[rune]int{'z': 3},
		},
		{
			name: "Test with empty input",
			args: args{characters: []rune{}},
			want: map[rune]int{},
		},
		{
			name: "Test with special characters",
			args: args{characters: []rune{'!', '@', '@', '#'}},
			want: map[rune]int{'!': 1, '@': 2, '#': 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SortCharactersInAnyOrder(tt.args.characters)

			// Convert the result to a map for comparison
			gotMap := make(map[rune]int)
			for _, char := range got {
				gotMap[char]++
			}

			assert.Equal(t, tt.want, gotMap, "SortCharactersInAnyOrder(%v)", tt.args.characters)
		})
	}
}
