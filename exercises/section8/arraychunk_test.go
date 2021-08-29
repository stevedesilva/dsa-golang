package section8_test

import (
	"reflect"
	"testing"

	"github.com/stevedesilva/dsa-golang.git/exercises/section8"
)

func TestExecute(t *testing.T) {
	type args struct {
		nums      []int
		chuckSize int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		//{name: "[],1 => []", args: args{nums: []int{}, chuckSize: 1}, want: [][]int{}},
		{name: "[1],1 => [[1]]", args: args{nums: []int{1}, chuckSize: 1}, want: [][]int{{1}}},
		//{name: "[1,2],1 => [[1],[2]]", args: args{nums: []int{1, 2}, chuckSize: 1}, want: [][]int{{1}, {2}}},
		//{name: "[1,2,3],1 => [[1],[2],[3]]", args: args{nums: []int{1, 2, 3}, chuckSize: 1}, want: [][]int{{1}, {2}, {3}}},
		//{name: "[1],2 => [[1]]", args: args{nums: []int{1}, chuckSize: 1}, want: [][]int{{1}}},
		//{name: "[1,2],2 => [[1,2]]", args: args{nums: []int{1, 2}, chuckSize: 2}, want: [][]int{{1, 2}}},
		//{name: "[1,2,3],2 => [[1,2],[3]", args: args{nums: []int{1, 2, 3}, chuckSize: 2}, want: [][]int{{1, 2}, {3}}},
		//{name: "[1,2,3,4,5],2 => [[1,2],[3,4],[5]", args: args{nums: []int{1, 2, 3, 4, 5}, chuckSize: 2}, want: [][]int{{1, 2}, {3, 4}, {5}}},
		//{name: "[1,2,3,4,5,6,7,8],3 => [[1,2,3],[4,5,6],[7,8]", args: args{nums: []int{1, 2, 3, 4, 5, 6, 7, 8}, chuckSize: 3}, want: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8}}},
		//{name: "[1,2,3,4,5],4 => [[1,2,3,4],[5]", args: args{nums: []int{1, 2, 3, 4, 5}, chuckSize: 4}, want: [][]int{{1, 2, 3, 4}, {5}}},
		//{name: "[1,2,3,4,5],10 => [[1,2,3,4,5]", args: args{nums: []int{1, 2, 3, 4, 5}, chuckSize: 10}, want: [][]int{{1, 2, 3, 4, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := section8.ChunkIntoSlices(tt.args.nums, tt.args.chuckSize)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
