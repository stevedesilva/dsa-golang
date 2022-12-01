package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SumArrayStartingAtZeroIndex(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  int
	}{
		{
			name:  "1",
			array: []int{1},
			want:  1,
		},
		{
			name:  "1 - 2",
			array: []int{1, 2},
			want:  3,
		},
		{
			name:  "1 - 3",
			array: []int{1, 2, 3},
			want:  6,
		},
		{
			name:  "1 - 4",
			array: []int{1, 2, 3, 4},
			want:  10,
		},
		{
			name:  "1 - 5",
			array: []int{1, 2, 3, 4, 5},
			want:  15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SumArrayStartingAtZeroIndex(tt.array), "SumArrayStartingAtZeroIndex(%v)", tt.array)
		})
	}
}

func Test_SumArrayStartingAtIndexEnd(t *testing.T) {

	tests := []struct {
		name  string
		array []int
		want  int
	}{
		{
			name:  "1",
			array: []int{1},
			want:  1,
		},
		{
			name:  "1 - 2",
			array: []int{1, 2},
			want:  3,
		},

		{
			name:  "1 - 3",
			array: []int{1, 2, 3},
			want:  6,
		},
		{
			name:  "1 - 4",
			array: []int{1, 2, 3, 4},
			want:  10,
		},
		{
			name:  "1 - 5",
			array: []int{1, 2, 3, 4, 5},
			want:  15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SumArrayStartingAtIndexEnd(tt.array), "SumArrayStartingAtIndexEnd(%v)", tt.array)
		})
	}
}
