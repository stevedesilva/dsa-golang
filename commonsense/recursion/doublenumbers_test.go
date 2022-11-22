package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_doubleNumbersInArray(t *testing.T) {

	tests := []struct {
		name  string
		array []int
		want  []int
	}{
		{
			name:  "1-3",
			array: []int{1, 2, 3},
			want:  []int{2, 4, 6},
		},
		{
			name:  "1-4",
			array: []int{1, 2, 3, 4},
			want:  []int{2, 4, 6, 8},
		},
		{
			name:  "1-5",
			array: []int{1, 2, 3, 4, 5},
			want:  []int{2, 4, 6, 8, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, doubleNumbersInArray(tt.array), "doubleNumbersInArray(%v)", tt.array)
		})
	}
}
