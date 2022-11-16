package recursion_test

import (
	"reflect"
	"testing"

	"github.com/stevedesilva/dsa-golang.git/commonsense/recursion"
)

func Test_countTo(t *testing.T) {
	tests := []struct {
		name   string
		number int
		want   []int
	}{
		{
			name:   "2",
			number: 2,
			want:   []int{2, 1},
		},
		{
			name:   "3",
			number: 3,
			want:   []int{3, 2, 1},
		},
		{
			name:   "4",
			number: 4,
			want:   []int{4, 3, 2, 1},
		},
		{
			name:   "5",
			number: 5,
			want:   []int{5, 4, 3, 2, 1},
		},
		{
			name:   "10",
			number: 10,
			want:   []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recursion.CountTo(tt.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countEven(t *testing.T) {
	type data struct {
		low, high int
	}
	tests := []struct {
		name string
		data data
		want []int
	}{
		{
			name: "even 0 to 10",
			data: data{0, 10},
			want: []int{0, 2, 4, 6, 8, 10},
		},
		{
			name: "even 0 to 20",
			data: data{0, 20},
			want: []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recursion.CountEven(tt.data.low, tt.data.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
