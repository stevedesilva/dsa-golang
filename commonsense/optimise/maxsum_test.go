package optimise

import "testing"

func TestMaxSum(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test with positive and negative numbers",
			args: args{array: []int{1, -2, 3, 4, -5, 8}},
			want: 10,
		},
		{
			name: "Test with all negative numbers",
			args: args{array: []int{-1, -2, -3, -4, -5}},
			want: 0,
		},
		{
			name: "Test with all positive numbers",
			args: args{array: []int{1, 2, 3, 4, 5}},
			want: 15,
		},
		{
			name: "Test with empty array",
			args: args{array: []int{}},
			want: 0,
		},
		{
			name: "Test with single element",
			args: args{array: []int{10}},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSum(tt.args.array); got != tt.want {
				t.Errorf("MaxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
