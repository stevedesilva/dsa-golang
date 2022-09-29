package array

import (
	"reflect"
	"testing"
)

func Test_mergeTwoArrays(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{
			name:    "empty arrays",
			args:    args{a: []int{}, b: []int{}},
			wantRes: []int{},
		},
		{
			name:    "empty array b",
			args:    args{a: []int{1}, b: []int{}},
			wantRes: []int{1},
		},
		{
			name:    "empty array a",
			args:    args{a: []int{}, b: []int{1}},
			wantRes: []int{1},
		},
		{
			name:    "merge a and b single",
			args:    args{a: []int{1}, b: []int{2}},
			wantRes: []int{1, 2},
		},
		{
			name:    "merge a and b end on a",
			args:    args{a: []int{1, 3}, b: []int{2}},
			wantRes: []int{1, 2, 3},
		},
		{
			name:    "merge a and b end on b",
			args:    args{a: []int{1}, b: []int{2, 3}},
			wantRes: []int{1, 2, 3},
		},
		{
			name:    "merge a then b",
			args:    args{a: []int{1, 2, 3}, b: []int{4, 5, 6}},
			wantRes: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:    "merge b then a single",
			args:    args{a: []int{2}, b: []int{1}},
			wantRes: []int{1, 2},
		},
		{
			name:    "merge b then a",
			args:    args{a: []int{4, 5, 6}, b: []int{1, 2, 3}},
			wantRes: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:    "merge random odd",
			args:    args{a: []int{2, 4}, b: []int{1, 3, 5}},
			wantRes: []int{1, 2, 3, 4, 5},
		},

		{
			name:    "merge random even",
			args:    args{a: []int{2, 3, 4, 11, 12}, b: []int{1, 2, 3, 5, 10}},
			wantRes: []int{1, 2, 2, 3, 3, 4, 5, 10, 11, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := mergeTwoArrays(tt.args.a, tt.args.b); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("mergeTwoArrays() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
