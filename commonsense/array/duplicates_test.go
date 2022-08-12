package array

import "testing"

func TestHasDuplicates(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want bool
	}{
		{
			name: "no dup 1",
			args: []int{1},
			want: false,
		},
		{
			name: "no dup 1,2",
			args: []int{1, 2},
			want: false,
		},
		{
			name: "no dup 1,2,3",
			args: []int{1, 2, 3},
			want: false,
		},
		{
			name: "has dup 1 1,1,5,5,6,7",
			args: []int{1, 1, 1, 5, 5, 6, 7},
			want: true,
		},
		{
			name: "has dup 1,2,2",
			args: []int{1, 2, 2},
			want: true,
		},
		{
			name: "has dup 1,2,3, 2, 1",
			args: []int{1, 2, 3, 2, 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasDuplicates(tt.args); got != tt.want {
				t.Errorf("TestHasDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
