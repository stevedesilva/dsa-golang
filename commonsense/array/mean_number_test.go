package array

import "testing"

func TestFindMeanNumber(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{array: []int{1}},
			want: 1,
		},
		{
			name: "10,10",
			args: args{array: []int{1}},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMeanForEvenNumbers(tt.args.array); got != tt.want {
				t.Errorf("FindMeanForEvenNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
