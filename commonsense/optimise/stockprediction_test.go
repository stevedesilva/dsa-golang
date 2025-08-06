package optimise

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindGreatestProfit(t *testing.T) {
	type args struct {
		prices []float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Test with normal prices",
			args: args{
				prices: []float64{10.0, 7.0, 5.0, 8.0, 11.0, 9.0},
			},
			want:    6.0, // Buy at 5.0 and sell at 11.0
			wantErr: assert.NoError,
		},
		{
			name: "Test with prices in descending order",
			args: args{
				prices: []float64{10.0, 9.0, 8.0, 7.0, 6.0},
			},
			want:    0.0, // No profit possible
			wantErr: assert.NoError,
		},
		{
			name: "Test with prices in ascending order",
			args: args{
				prices: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			},
			want:    4.0, // Buy at 1.0 and sell at 5.0
			wantErr: assert.NoError,
		},
		{
			name: "Test with prices in ascending order with low number at the end",
			args: args{
				prices: []float64{2.0, 3.0, 4.0, 5.0, 1.0, 2.0, 3.0},
			},
			want:    3.0, // Buy at 2.0 and sell at 5.0
			wantErr: assert.NoError,
		},
		{
			name: "Test with prices in ascending order with low number at the end",
			args: args{
				prices: []float64{2.0, 3.0, 4.0, 5.0, 1.0, 2.0, 3.0, 10.0},
			},
			want:    9.0, // Buy at 1.0 and sell at 10.0
			wantErr: assert.NoError,
		},

		{
			name: "Test with prices in ascending order with low number at the end",
			args: args{
				prices: []float64{2.0, 3.0, 4.0, 5.0, 1.0, 2.0, 3.0, 5.0},
			},
			want:    4.0, // Buy at 1.0 and sell at 5.0
			wantErr: assert.NoError,
		},
		{
			name: "Test with single price",
			args: args{
				prices: []float64{10.0},
			},
			want:    0.0, // Not enough prices to calculate profit
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindGreatestProfit(tt.args.prices)
			if !tt.wantErr(t, err, fmt.Sprintf("FindGreatestProfit(%v)", tt.args.prices)) {
				return
			}
			assert.Equalf(t, tt.want, got, "FindGreatestProfit(%v)", tt.args.prices)
		})
	}
}
