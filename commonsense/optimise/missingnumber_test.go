package optimise

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Test with missing number in the middle",
			args: args{
				numbers: []int{0, 1, 2, 4, 5, 6},
			},
			want:    3,
			wantErr: assert.NoError,
		},
		{
			name: "Test with missing number in the middle",
			args: args{
				numbers: []int{0, 1, 2, 3, 5, 6},
			},
			want:    4,
			wantErr: assert.NoError,
		},
		// no missing number
		{
			name: "Test with no missing number",
			args: args{
				numbers: []int{0, 1, 2, 3, 4, 5, 6},
			},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name: "Test with empty array",
			args: args{
				numbers: []int{},
			},
			want:    0,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Find(tt.args.numbers)
			if !tt.wantErr(t, err, fmt.Sprintf("Find(%v)", tt.args.numbers)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Find(%v)", tt.args.numbers)
		})
	}
}
