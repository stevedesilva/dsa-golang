package optimise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSequence(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test with consecutive numbers",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Test with non-consecutive numbers",
			args: args{nums: []int{10, 5, 12, 3, 55, 30, 4, 11, 2}},
			want: []int{2, 3, 4, 5},
		},
		{
			name: "Test with negative and positive numbers",
			args: args{nums: []int{-1, 0, 1, 2, -2, -3, 3}},
			want: []int{-3, -2, -1, 0, 1, 2, 3},
		},
		{
			name: "Test with duplicates",
			args: args{nums: []int{1, 2, 2, 3, 4, 4, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Test with single element",
			args: args{nums: []int{42}},
			want: []int{42},
		},
		{
			name: "Test with empty array",
			args: args{nums: []int{}},
			want: []int{},
		},
		{
			name: "Test with multiple sequences",
			args: args{nums: []int{100, 4, 200, 1, 3, 2}},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "Test with all identical elements",
			args: args{nums: []int{7, 7, 7, 7}},
			want: []int{7},
		},
		{
			name: "Test with large range",
			args: args{nums: []int{0, -1, 1, 2, -2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: []int{-2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, LongestSequence(tt.args.nums), "LongestSequence(%v)", tt.args.nums)
		})
	}
}
