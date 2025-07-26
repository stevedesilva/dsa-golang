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
				numbers: []int{1, 2, 4, 5, 6},
			},
			want:    3,
			wantErr: assert.NoError,
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
