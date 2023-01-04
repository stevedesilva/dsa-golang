package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
70 35 15  5  1
35 20 10  4  1
15 10  6  3  1
5   4  3  2  1
1   1  1  1  1
*/
func Test_findAllUniquePaths(t *testing.T) {
	type args struct {
		row    int
		column int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1,1:1", args: args{1, 1}, want: 1},
		{name: "1,2:1", args: args{1, 2}, want: 1},
		{name: "1,3:1", args: args{1, 3}, want: 1},
		{name: "1,4:1", args: args{1, 4}, want: 1},
		{name: "1,5:1", args: args{1, 5}, want: 1},
		{name: "1,6:1", args: args{1, 6}, want: 1},
		{name: "1,7:1", args: args{1, 7}, want: 1},

		{name: "2,1:1", args: args{2, 1}, want: 1},
		{name: "2,2:1", args: args{2, 2}, want: 2},
		{name: "2,3:1", args: args{2, 3}, want: 3},
		{name: "2,4:1", args: args{2, 4}, want: 4},
		{name: "2,5:1", args: args{2, 5}, want: 5},

		{name: "3,1:1", args: args{3, 1}, want: 1},
		{name: "3,2:1", args: args{3, 2}, want: 3},
		{name: "3,3:1", args: args{3, 3}, want: 6},
		{name: "3,4:1", args: args{3, 4}, want: 10},
		{name: "3,5:1", args: args{3, 5}, want: 15},

		{name: "4,1:1", args: args{4, 1}, want: 1},
		{name: "4,2:1", args: args{4, 2}, want: 4},
		{name: "4,3:1", args: args{4, 3}, want: 10},
		{name: "4,4:1", args: args{4, 4}, want: 20},
		{name: "4,5:1", args: args{4, 5}, want: 35},

		{name: "5,1:1", args: args{5, 1}, want: 1},
		{name: "5,2:1", args: args{5, 2}, want: 5},
		{name: "5,3:1", args: args{5, 3}, want: 15},
		{name: "5,4:1", args: args{5, 4}, want: 35},
		{name: "5,5:1", args: args{5, 5}, want: 70},

		{name: "5,1:1", args: args{6, 1}, want: 1},

		{name: "5,1:1", args: args{7, 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findAllUniquePaths(tt.args.row, tt.args.column), "findAllUniquePaths(%v, %v)", tt.args.row, tt.args.column)
		})
	}
}
