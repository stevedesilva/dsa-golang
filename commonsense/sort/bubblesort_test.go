package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort_Sort(t *testing.T) {

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
			data: []int{3, 5, 2, 1, 4},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BubbleSort{
				data: tt.data,
			}
			b.Sort()
			assert.Equal(t, tt.want, b.data)
		})
	}
}
