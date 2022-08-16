package sort_test

import (
	"testing"

	"github.com/stevedesilva/dsa-golang.git/commonsense/sort"
	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	type data struct {
		name  string
		input []int
		want  []int
	}
	testData := []data{
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
			name:  "2 test in order",
			input: []int{1, 2},
			want:  []int{1, 2},
		},
		{
			name:  "2 test out of order",
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
	}
	for _, d := range testData {
		t.Run(d.name, func(t *testing.T) {
			res := sort.InsertionSort(d.input)
			assert.Equal(t, d.want, res)
		})
	}
}
