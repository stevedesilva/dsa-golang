package optimise

import (
	"reflect"
	"testing"
)

func TestSwapToMakeEqual(t *testing.T) {
	type args struct {
		a, b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "'4,2,3,1', '4,2,5,3,6', '3,4'",
			args: args{[]int{4, 2, 3, 1}, []int{4, 2, 5, 3, 6}},
			want: []int{3, 4},
		},
		{
			name: "'4,2,5,3,6', '4,2,3,1', '4,3'",
			args: args{[]int{4, 2, 5, 3, 6}, []int{4, 2, 3, 1}},
			want: []int{4, 3},
		},
		{
			name: "'5,3,2,9,1', '1,12,5', '2,0'",
			args: args{[]int{5, 3, 2, 9, 1}, []int{1, 12, 5}},
			want: []int{2, 0},
		},
		{
			name: "'1,12,5','5,3,2,9,1',  '0,2'",
			args: args{[]int{1, 12, 5}, []int{5, 3, 2, 9, 1}},
			want: []int{0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SwapToMakeEqual(tt.args.a, tt.args.a)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}
