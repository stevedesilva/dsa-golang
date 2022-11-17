package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumNumberBetweenRangeTailRecursion(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 to 2",
			args: args{
				1, 2,
			},
			want: 3,
		},
		{
			name: "1 to 3",
			args: args{
				1, 3,
			},
			want: 6,
		},
		{
			name: "1 to 4",
			args: args{
				1, 4,
			},
			want: 10,
		},
		{
			name: "1 to 5",
			args: args{
				1, 5,
			},
			want: 15,
		},
		{
			name: "1 to 6",
			args: args{
				1, 6,
			},
			want: 21,
		},
		{
			name: "1 to 7",
			args: args{
				1, 7,
			},
			want: 28,
		},
		{
			name: "1 to 8",
			args: args{
				1, 8,
			},
			want: 36,
		},
		{
			name: "1 to 9",
			args: args{
				1, 9,
			},
			want: 45,
		},
		{
			name: "1 to 10",
			args: args{
				1, 10,
			},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SumNumberBetweenRange(tt.args.start, tt.args.end), "SumNumberBetweenRange(%v, %v)", tt.args.start, tt.args.end)
		})
	}
}

func TestSumNumberBetweenRangeNoTailRec(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 to 2",
			args: args{
				1, 2,
			},
			want: 3,
		},
		{
			name: "1 to 3",
			args: args{
				1, 3,
			},
			want: 6,
		},
		{
			name: "1 to 4",
			args: args{
				1, 4,
			},
			want: 10,
		},
		{
			name: "1 to 5",
			args: args{
				1, 5,
			},
			want: 15,
		},
		{
			name: "1 to 6",
			args: args{
				1, 6,
			},
			want: 21,
		},
		{
			name: "1 to 7",
			args: args{
				1, 7,
			},
			want: 28,
		},
		{
			name: "1 to 8",
			args: args{
				1, 8,
			},
			want: 36,
		},
		{
			name: "1 to 9",
			args: args{
				1, 9,
			},
			want: 45,
		},
		{
			name: "1 to 10",
			args: args{
				1, 10,
			},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SumNumberBetweenRangeNoTailRec(tt.args.start, tt.args.end), "SumNumberBetweenRange(%v, %v)", tt.args.start, tt.args.end)
		})
	}
}
