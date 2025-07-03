package optimise

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIncreasingTriplet(t *testing.T) {
	type args struct {
		prices []float64
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "Test with increasing triplet",
			args:    args{prices: []float64{1.0, 2.0, 3.0, 4.0, 5.0}},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name:    "Test with no increasing triplet",
			args:    args{prices: []float64{5.0, 4.0, 3.0, 2.0, 1.0}},
			want:    false,
			wantErr: assert.NoError,
		},
		{
			name:    "Test with mixed values",
			args:    args{prices: []float64{1.0, 5.0, 2.0, 4.0, 3.0}},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name:    "Test with less than three elements",
			args:    args{prices: []float64{1.0, 2.0}},
			want:    false,
			wantErr: assert.Error, // Should expect an error, not false
		},
		{
			name:    "Test with nil slice",
			args:    args{prices: nil},
			want:    false,
			wantErr: assert.Error,
		},
		{
			name:    "Test with empty slice",
			args:    args{prices: []float64{}},
			want:    false,
			wantErr: assert.Error,
		},
		{
			name:    "Test with exactly three elements - increasing",
			args:    args{prices: []float64{1.0, 2.0, 3.0}},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name:    "Test with exactly three elements - decreasing",
			args:    args{prices: []float64{3.0, 2.0, 1.0}},
			want:    false,
			wantErr: assert.NoError,
		},
		{
			name:    "Test with non-consecutive triplet",
			args:    args{prices: []float64{3.0, 1.0, 2.0, 4.0}},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name:    "Test with duplicate values",
			args:    args{prices: []float64{1.0, 1.0, 2.0, 2.0, 3.0}},
			want:    true,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsIncreasingTriplet(tt.args.prices)
			if !tt.wantErr(t, err, fmt.Sprintf("IsIncreasingTriplet(%v)", tt.args.prices)) {
				return
			}
			assert.Equalf(t, tt.want, got, "IsIncreasingTriplet(%v)", tt.args.prices)
		})
	}
}
