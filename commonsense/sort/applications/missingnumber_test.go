package applications

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMissingNumber(t *testing.T) {

	tests := []struct {
		name    string
		numbers []int
		want    *int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "minimum value not found",
			numbers: []int{},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name:    "value not found error",
			numbers: []int{1, 0, 3, 2},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name:    "missing 1",
			numbers: []int{2, 0},
			want:    IntPointer(1),
			wantErr: assert.NoError,
		},
		{
			name:    "missing 2",
			numbers: []int{3, 1, 0},
			want:    IntPointer(2),
			wantErr: assert.NoError,
		},
		{
			name:    "missing 5",
			numbers: []int{8, 7, 2, 4, 3, 6, 1, 0},
			want:    IntPointer(5),
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindMissingNumber(tt.numbers)
			if !tt.wantErr(t, err, fmt.Sprintf("FindMissingNumber(%v)", tt.numbers)) {
				return
			}
			assert.Equalf(t, tt.want, got, "FindMissingNumber(%v)", tt.numbers)
		})
	}
}
