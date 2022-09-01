package array

import "testing"

func TestAverageCelsuisReader(t *testing.T) {
	type args struct {
		fahrenheitReadings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "20",
			args: args{
				fahrenheitReadings: []int{20},
			},
			want: -7,
		},
		{name: "32",
			args: args{
				fahrenheitReadings: []int{32},
			},
			want: 0,
		},
		{name: "40",
			args: args{
				fahrenheitReadings: []int{40},
			},
			want: 4,
		},
		{name: "90",
			args: args{
				fahrenheitReadings: []int{90},
			},
			want: 32,
		},
		{name: "100",
			args: args{
				fahrenheitReadings: []int{100},
			},
			want: 38,
		},
		{name: "10c,20c,32c,40c,90c,100c",
			args: args{
				fahrenheitReadings: []int{10, 20, 32, 40, 90, 100},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AverageCelsuisReader(tt.args.fahrenheitReadings); got != tt.want {
				t.Errorf("AverageCelsuisReader() = %v, want %v", got, tt.want)
			}
		})
	}
}
