package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQuickSelect(t *testing.T) {
	type data struct {
		name  string
		input []int
		index int
		want  int
	}
	tests := []data{
		//{
		//	name:  "empty test",
		//	input: []int{},
		//	want:  0,
		//},
		{
			name:  "find index 0 single value",
			input: []int{1},
			index: 0,
			want:  1,
		},
		{
			name:  "find index 0 two sorted values",
			input: []int{1, 2},
			index: 0,
			want:  1,
		},
		{
			name:  "find index 1 two sorted values",
			input: []int{1, 2},
			index: 1,
			want:  2,
		},
		{
			name:  "find index 0 two unsorted values",
			input: []int{2, 1},
			index: 0,
			want:  1,
		},
		{
			name:  "find index 1 two unsorted values",
			input: []int{2, 1},
			index: 1,
			want:  2,
		},
		{
			name:  "find index 0 two sorted values",
			input: []int{1, 2, 3},
			index: 0,
			want:  1,
		},
		{
			name:  "find index 1 two sorted values",
			input: []int{1, 2, 3},
			index: 1,
			want:  2,
		},
		{
			name:  "find index 2 two sorted values",
			input: []int{1, 2, 3},
			index: 2,
			want:  3,
		},
		{
			name:  "find index 0 3 unsorted values",
			input: []int{3, 2, 1},
			index: 0,
			want:  1,
		},
		{
			name:  "find index 2 3 unsorted values",
			input: []int{3, 2, 1},
			index: 2,
			want:  3,
		},
		{
			name:  "find index 0  10 unsorted values",
			input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			index: 0,
			want:  1,
		},
		{
			name:  "find index 4  10 unsorted values",
			input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			index: 4,
			want:  5,
		},
		{
			name:  "find index 9  10 unsorted values",
			input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			index: 9,
			want:  10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQuickSelect(tt.input)
			assert.Equalf(t, tt.want, q.QuickSelect(tt.index), "QuickSelect(%v)", tt.want)
		})
	}
}
