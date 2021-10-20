package section14

import (
	"reflect"
	"testing"
)

func TestCreateSpiral(t *testing.T) {

	tests := []struct {
		name string
		size int
		want [][]int
	}{
		{
			name: "1",
			size: 1,
			want: [][]int{
				{1},
			},
		},
		{
			name: "2",
			size: 2,
			want: [][]int{
				{1, 2},
				{4, 3},
			},
		},
		{
			name: "3",
			size: 3,
			want: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			name: "4",
			size: 4,
			want: [][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
		},
		{
			name: "5",
			size: 5,
			want: [][]int{
				{1, 2, 3, 4, 5},
				{16, 17, 18, 19, 6},
				{15, 24, 25, 20, 7},
				{14, 23, 22, 21, 8},
				{13, 12, 11, 10, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateSpiral(tt.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSpiral() = %v, want %v", got, tt.want)
			}
		})
	}
}
