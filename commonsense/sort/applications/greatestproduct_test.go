package applications

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindGreatestProduct(t *testing.T) {
	type args struct {
		array       []int
		numberRange int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "error when minimum array length not met",
			args: args{
				array:       nil,
				numberRange: -1,
			},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name: "error when numberRange less then 1",
			args: args{
				array:       nil,
				numberRange: -1,
			},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name: "find product in 2",
			args: args{
				array:       []int{2, 1},
				numberRange: 1,
			},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name: "find product in 5",
			args: args{
				array:       []int{10, 5, 2, 1},
				numberRange: 3,
			},
			want:    17,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product, err := FindGreatestProduct(tt.args.array, tt.args.numberRange)
			if !tt.wantErr(t, err, fmt.Sprintf("FindGreatestProduct() = %v, want %v", product, tt.want)) {
				return
			}
			assert.Equalf(t, tt.want, product, "FindGreatestProduct(%v, %v)", tt.args.array, tt.args.numberRange)
		})
	}
}
