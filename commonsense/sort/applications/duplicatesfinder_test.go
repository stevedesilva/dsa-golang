package applications

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasDuplicates(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want bool
	}{
		{
			name: "empty test",
			args: []int{},
			want: false,
		},
		{
			name: "single test",
			args: []int{1},
			want: false,
		},
		{
			name: "2 duplicates",
			args: []int{1, 1},
			want: true,
		},
		{
			name: "2 test in order",
			args: []int{1, 2, 1},
			want: true,
		},
		{
			name: "3 test in order",
			args: []int{1, 2, 3},
			want: false,
		},
		{
			name: "3 test in order duplicate",
			args: []int{1, 2, 3, 2},
			want: true,
		},
		{
			name: "2 test out of order",
			args: []int{2, 1},
			want: false,
		},
		{
			name: "2 test out of order",
			args: []int{2, 2},
			want: true,
		},
		{
			name: "3 test out of order",
			args: []int{3, 2, 1},
			want: false,
		},
		{
			name: "4 test out of order",
			args: []int{3, 2, 4, 1},
			want: false,
		},
		{
			name: "5 test out of order",
			args: []int{5, 3, 2, 4, 1},
			want: false,
		},
		{
			name: "6 test out of order",
			args: []int{5, 3, 2, 6, 4, 1},
			want: false,
		},
		{
			name: "6 - worst case test out of order",
			args: []int{6, 5, 4, 3, 2, 1},
			want: false,
		},
		{
			name: "6 - worst case test out of order duplicate",
			args: []int{6, 5, 4, 4, 6, 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, HasDuplicates(tt.args), "HasDuplicates(%v)", tt.args)
		})
	}
}
