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
		// zero indexed 0 1 2 3 4 5 6  7  8  9 10
		// value        0 1 1 2 3 5 8 13 21 34 54
		//{name: "fib(0) ==  0", args: args{0}, want: 0},
		{name: "fib(1) ==  1", args: args{1}, want: 1},
		{name: "fib(2) ==  1", args: args{2}, want: 1},
		{name: "fib(3) ==  2", args: args{3}, want: 2},
		{name: "fib(4) ==  3", args: args{4}, want: 3},
		{name: "fib(5) ==  5", args: args{5}, want: 5},
		{name: "fib(6) ==  8", args: args{6}, want: 8},
		{name: "fib(7) == 13", args: args{7}, want: 13},
		{name: "fib(8) == 21", args: args{8}, want: 21},
		{name: "fib(8) == 34", args: args{9}, want: 34},
		{name: "fib(8) == 55", args: args{10}, want: 55},
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
			if got := FibRecursiveNaive(tt.args.entry); got != tt.want {
				t.Errorf("FibRecursiveNaive() = %v, want %v", got, tt.want)
			}
			if got := FibRecursiveMemoizedV1(tt.args.entry); got != tt.want {
				t.Errorf("FibRecursiveMemoizedV1() = %v, want %v", got, tt.want)
			}
			if got := FibRecursiveMemoizedV2(tt.args.entry); got != tt.want {
				t.Errorf("FibRecursiveMemoizedV2() = %v, want %v", got, tt.want)
			}
		})
	}

}
