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
			name: "already sorted 1,2,3,4,5",
			data: []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "sorted 1",
			data: []int{1},
			want: []int{1},
		},
		{
			name: "sorted 1,2",
			data: []int{2, 1},
			want: []int{1, 2},
		},
		{
			name: "sorted 1,2,3",
			data: []int{3, 2, 1},
			want: []int{1, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BubbleSort{
				data: tt.data,
			}

			want := []int{1, 2, 3, 4, 5}
			b.Sort()
			assert.Equal(t, want, b.data)
		})
	}
}
