package hashmap_test

import (
	"reflect"
	"testing"

	. "github.com/stevedesilva/dsa-golang.git/commonsense/hashmap"
)

func TestNew(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	type res struct {
		val []int
		err error
	}
	tests := []struct {
		name string
		args args
		want res
	}{
		{
			name: "error on minimum not met",
			args: args{
				a: nil,
				b: nil,
			},
			want: res{nil, ErrInvalidInput},
		},
		{
			name: "single a",
			args: args{
				a: []int{4, 2},
				b: nil,
			},
			want: res{[]int{4, 2}, nil},
		},
		{
			name: "single b",
			args: args{
				a: nil,
				b: []int{1, 2},
			},
			want: res{[]int{1, 2}, nil},
		},
		{
			name: "join 1",
			args: args{
				a: []int{1, 2},
				b: []int{1, 3},
			},
			want: res{[]int{1}, nil},
		},
		{
			name: "join 2",
			args: args{
				a: []int{1, 2},
				b: []int{2, 3},
			},
			want: res{[]int{2}, nil},
		},
		{
			name: "join 3",
			args: args{
				a: []int{1, 3},
				b: []int{2, 3, 3, 3},
			},
			want: res{[]int{3}, nil},
		},
		{
			name: "join 3",
			args: args{
				a: []int{1, 3, 6, 9},
				b: []int{2, 3, 9, 32},
			},
			want: res{[]int{3, 9}, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sub := New(tt.args.a, tt.args.b)
			got, error := sub.GetIntersection()
			if error != tt.want.err {
				t.Errorf("GetIntersection() error = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got, tt.want.val) {
				t.Errorf("GetIntersection() val = %v, want %v", got, tt.want)
			}
		})
	}
}
