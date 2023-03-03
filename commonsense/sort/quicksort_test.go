package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuicksort(t *testing.T) {
	type data struct {
		name  string
		input []int
		want  []int
	}
	tests := []data{
		{
			name:  "empty test",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "1 test",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "test in order",
			input: []int{1, 2},
			want:  []int{1, 2},
		},
		{
			name:  "test out of order",
			input: []int{2, 1},
			want:  []int{1, 2},
		},
		{
			name:  "3 test in order",
			input: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
		{
			name:  "3 test out of order",
			input: []int{3, 2, 1},
			want:  []int{1, 2, 3},
		},
		{
			name:  "4 test out of order",
			input: []int{3, 2, 4, 1},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "5 test out of order",
			input: []int{5, 3, 2, 4, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "6 test out of order",
			input: []int{5, 3, 2, 6, 4, 1},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:  "6 - worst case test out of order",
			input: []int{6, 5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQuickSort(tt.input)
			assert.Equalf(t, tt.want, q.Quicksort(), "Quicksort(%v)", tt.want)
		})
	}
}
