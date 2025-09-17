package optimise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Test with positive integers",
			args:    args{nums: []int{3, 4, 5, 2}},
			want:    20,
			wantErr: false,
		},
		{
			name:    "Test with larger integers",
			args:    args{nums: []int{1, 5, 4, 5}},
			want:    25,
			wantErr: false,
		},
		{
			name:    "Test with negative and positive integers 1",
			args:    args{nums: []int{5, -1, -2}},
			want:    2,
			wantErr: false,
		},
		{
			name:    "Test with negative and positive integers 2",
			args:    args{nums: []int{-10, 1, 7, -9, 1}},
			want:    90,
			wantErr: false,
		},
		{
			name:    "Test with negative and positive integers 3",
			args:    args{nums: []int{-5, 2, 1}},
			want:    2,
			wantErr: false,
		},
		{
			name:    "Test with negative and positive integers 4",
			args:    args{nums: []int{3, -1, 1}},
			want:    3,
			wantErr: false,
		},
		{
			name:    "Test with insufficient numbers",
			args:    args{nums: []int{1}},
			want:    0,
			wantErr: true,
		},
		{
			name:    "Test with empty list",
			args:    args{nums: []int{}},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxProduct(tt.args.nums)
			if tt.wantErr {
				assert.Error(t, err, "Expected an error but got none")
			} else {
				assert.NoError(t, err, "Did not expect an error but got one")
				assert.Equalf(t, tt.want, got, "MaxProduct(%v)", tt.args.nums)
			}
		})
	}
}
