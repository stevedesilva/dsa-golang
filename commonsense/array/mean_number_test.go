package array_test

import (
	"testing"

	. "github.com/stevedesilva/dsa-golang.git/commonsense/array"
)

func TestFindMeanNumber(t *testing.T) {
	type args struct {
		array []int
	}
	type want struct {
		res float64
		err error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "1",
			args: args{array: []int{1}},
			want: want{1, ErrNotFound},
		},
		{
			name: "10,10",
			args: args{array: []int{10, 10}},
			want: want{10, ErrNotFound},
		},
		{
			name: "5,10",
			args: args{array: []int{5, 10}},
			want: want{10, ErrNotFound},
		},
		{
			name: "8,8,8",
			args: args{array: []int{8, 8, 8}},
			want: want{8, ErrNotFound},
		},
		{
			name: "6,8,20,12",
			args: args{array: []int{6, 8, 20, 12}},
			want: want{11.5, ErrNotFound},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := FindMeanForEvenNumbers(tt.args.array); got != tt.want.res && err != tt.want.err {
				t.Errorf("FindMeanForEvenNumbers() = %v, want result %v with error %v ", got, tt.want.res, tt.want.err)
			}
		})
	}
}
