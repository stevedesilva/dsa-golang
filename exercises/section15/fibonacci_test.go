package section15

import "testing"

func TestGenerateFibSeries(t *testing.T) {
	type args struct {
		entry int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "fib(1) ==  1", args: args{1}, want: 1},
		{name: "fib(2) ==  1", args: args{2}, want: 1},
		{name: "fib(3) ==  2", args: args{3}, want: 2},
		{name: "fib(4) ==  3", args: args{4}, want: 3},
		{name: "fib(5) ==  5", args: args{5}, want: 5},
		{name: "fib(6) ==  8", args: args{6}, want: 8},
		{name: "fib(7) == 13", args: args{7}, want: 13},
		{name: "fib(8) == 21", args: args{8}, want: 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateFibSeriesIterative(tt.args.entry); got != tt.want {
				t.Errorf("GenerateFibSeriesIterative() = %v, want %v", got, tt.want)
			}
			if got := GenerateFibSeriesIterative2(tt.args.entry); got != tt.want {
				t.Errorf("GenerateFibSeriesIterative2() = %v, want %v", got, tt.want)
			}
			if got := GenerateFibSeriesRecursive(tt.args.entry); got != tt.want {
				t.Errorf("GenerateFibSeriesRecursive() = %v, want %v", got, tt.want)
			}
			if got := GenerateFibSeriesRecursive2(tt.args.entry); got != tt.want {
				t.Errorf("GenerateFibSeriesRecursive2() = %v, want %v", got, tt.want)
			}
		})
	}
}
