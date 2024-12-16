package optimise

import (
	"reflect"
	"testing"
)

func TestSumSwap(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SumSwap(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("SumSwap() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SumSwap() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

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
