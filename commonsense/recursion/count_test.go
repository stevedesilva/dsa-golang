package recursion

import (
	"reflect"
	"testing"
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
			if got := countTo(tt.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
