package optimise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortBodyTemperatures(t *testing.T) {
	type args struct {
		temp []float64
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr bool
	}{
		{
			name:    "Test with valid temperatures 1 element",
			args:    args{temp: []float64{97.5}},
			want:    []float64{97.5},
			wantErr: false,
		},
		{
			name:    "Test with valid temperatures 2 elements",
			args:    args{temp: []float64{99.0, 97.5}},
			want:    []float64{97.5, 99.0},
			wantErr: false,
		},
		{
			name:    "Test with valid temperatures 2 elements wrong order",
			args:    args{temp: []float64{97.5, 99.0}},
			want:    []float64{97.5, 99.0},
			wantErr: false,
		},
		{
			name:    "Test with valid temperatures 3 elements",
			args:    args{temp: []float64{99.0, 97.5, 98.6}},
			want:    []float64{97.5, 98.6, 99.0},
			wantErr: false,
		},
		{
			name:    "Test with valid temperatures 3 elements",
			args:    args{temp: []float64{99.0, 98.6, 97.5}},
			want:    []float64{97.5, 98.6, 99.0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SortBodyTemperatures(tt.args.temp)
			if tt.wantErr {
				assert.Error(t, err, "Expected an error but got none")
			} else {
				assert.NoError(t, err, "Did not expect an error but got one")
				assert.Equalf(t, tt.want, got, "SortBodyTemperatures(%v)", tt.args.temp)
			}
		})
	}
}
