package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasIntersection(t *testing.T) {
	type input struct {
		a1 []int
		a2 []int
	}
	tests := []struct {
		name string
		args input
		want []int
	}{
		{
			name: "{1}",
			args: input{[]int{1}, []int{1}},
			want: []int{1},
		},
		{
			name: "{1, 2}",
			args: input{[]int{1, 2}, []int{2, 2}},
			want: []int{2},
		},
		{
			name: "{1, 2, 3}",
			args: input{[]int{1, 2, 3}, []int{1, 2, 3}},
			want: []int{1, 2, 3},
		},
		{
			name: "{1, 2, 3, 4}",
			args: input{[]int{1, 2, 3, 4}, []int{4, 2, 1, 3}},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "{1, 2, 3, 4, 5}",
			args: input{[]int{1, 2, 3, 4, 5}, []int{4, 2, 1, 3, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasIntersection(tt.args.a1, tt.args.a2)
			assert.Equal(t, tt.want, got)
		})
	}
}
