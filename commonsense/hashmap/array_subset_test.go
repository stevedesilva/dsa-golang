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
