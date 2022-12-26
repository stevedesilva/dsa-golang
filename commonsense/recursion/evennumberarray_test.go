package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvenNumbersFromArray(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  []int
	}{
		{
			name:  "odd  digit",
			array: []int{1},
			want:  []int{},
		},
		{
			name:  "single even digit",
			array: []int{2},
			want:  []int{2},
		},
		{
			name:  "odd  digits",
			array: []int{1, 3, 5, 9, 11},
			want:  []int{},
		},
		{
			name:  "even digits",
			array: []int{2, 4, 6, 8, 10},
			want:  []int{2, 4, 6, 8, 10},
		},
		{
			name:  "mixed  digits 1",
			array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:  []int{2, 4, 6, 8, 10},
		},
		{
			name:  "mixed  digits 2",
			array: []int{1, 3, 5, 9, 11, 2, 4, 6, 8, 10},
			want:  []int{2, 4, 6, 8, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, EvenNumbersFromArray(tt.array), "EvenNumbersFromArray(%v)", tt.array)
		})
	}
}
