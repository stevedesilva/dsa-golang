package optimise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxNum(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Test with empty array",
			args:    args{array: []int{}},
			want:    0,
			wantErr: true,
		},
		{
			name:    "Test with single element",
			args:    args{array: []int{10}},
			want:    10,
			wantErr: false,
		},
		{
			name:    "Test with positive and negative numbers",
			args:    args{array: []int{1, -2, 3, 4, -5, 8}},
			want:    8,
			wantErr: false,
		},
		{
			name:    "Test with all negative numbers",
			args:    args{array: []int{-1, -2, -3, -4, -5}},
			want:    -1,
			wantErr: false,
		},
		{
			name:    "Test with all positive numbers",
			args:    args{array: []int{1, 2, 3, 4, 5}},
			want:    5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxNum(tt.args.array)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equalf(t, tt.want, got, "MaxNum(%v)", tt.args.array)
		})
	}
}
