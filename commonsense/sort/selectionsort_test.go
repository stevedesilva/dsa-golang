package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort_Sort(t *testing.T) {

	tests := []struct {
		name string
		data []int
		want []int
	}{
		{
			name: "sorted 1",
			data: []int{1},
			want: []int{1},
		},
		{
			name: "already sorted 1,2,3,4,5",
			data: []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "sorted 1,2",
			data: []int{2, 1},
			want: []int{1, 2},
		},
		{
			name: "sorted 1,2,3",
			data: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
		{
			name: "sorted 1,2,3,4",
			data: []int{3, 4, 2, 1},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "sorted 1,2,3,4,5",
			data: []int{4, 5, 3, 1, 2},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "sorted 1,2,3,4,5",
			data: []int{1, 2, 4, 5, 3},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "test first place",
			data: []int{1, 2, 4, 3},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SelectionSort{
				data: tt.data,
			}
			s.Sort()
			assert.Equal(t, tt.want, s.data)
		})
	}
}

func TestSelectionSort_SortOfficial(t *testing.T) {

	tests := []struct {
		name string
		data []int
		want []int
	}{
		{
			name: "sorted 1",
			data: []int{1},
			want: []int{1},
		},
		{
			name: "already sorted 1,2,3,4,5",
			data: []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "sorted 1,2",
			data: []int{2, 1},
			want: []int{1, 2},
		},
		{
			name: "sorted 1,2,3",
			data: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
		{
			name: "sorted 1,2,3,4",
			data: []int{3, 4, 2, 1},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "sorted 1,2,3,4,5",
			data: []int{4, 5, 3, 1, 2},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "sorted 1,2,3,4,5",
			data: []int{1, 2, 4, 5, 3},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "test first place",
			data: []int{1, 2, 4, 3},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SelectionSort{
				data: tt.data,
			}
			s.SortOfficial()
			assert.Equal(t, tt.want, s.data)
		})
	}
}
