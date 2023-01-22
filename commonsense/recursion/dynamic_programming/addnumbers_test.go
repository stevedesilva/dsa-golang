package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUntilOneHundred(t *testing.T) {

	tests := []struct {
		name    string
		numbers []int
		want    int
	}{
		{
			name:    "1",
			numbers: []int{1},
			want:    1,
		},
		{
			name:    "1-3",
			numbers: []int{1, 2, 3},
			want:    6,
		},
		{
			name:    "mixture adding up to 100",
			numbers: []int{1, 2, 3, 4, 5, 10, 20, 40, 15},
			want:    100,
		},
		{
			name:    "1-99",
			numbers: []int{1, 98},
			want:    99,
		},
		{
			name:    "100",
			numbers: []int{100},
			want:    100,
		},
		{
			name:    "Over 110",
			numbers: []int{110},
			want:    0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, AddUntilOneHundred(tt.numbers), "AddUntilOneHundred(%v)", tt.numbers)
		})
	}
}
