package recursion

import "testing"

func Test_getFactorial(t *testing.T) {

	tests := []struct {
		name string
		num  int
		want int
	}{
		{
			"1", 1, 1,
		},
		{
			"2", 2, 2,
		},
		{
			"3", 3, 6,
		},
		{
			"4", 4, 24,
		},
		{
			"5", 5, 120,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFactorial(tt.num); got != tt.want {
				t.Errorf("getFactorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
