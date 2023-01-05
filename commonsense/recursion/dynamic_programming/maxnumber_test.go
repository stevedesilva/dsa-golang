package dynamic_programming

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindMaxNumber(t *testing.T) {

	tests := []struct {
		name    string
		numbers []int
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "empty",
			numbers: []int{},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name:    "1",
			numbers: []int{1},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name:    "1-2",
			numbers: []int{1, 2},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name:    "1-4",
			numbers: []int{1, 2, 3, 4},
			want:    4,
			wantErr: assert.NoError,
		},
		{
			name:    "random",
			numbers: []int{10, 2, 30, 4, 58, 21},
			want:    58,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindMaxNumber(tt.numbers)
			if !tt.wantErr(t, err, fmt.Sprintf("FindMaxNumber() = %v, want %v", got, tt.want)) {
				return
			}
			assert.Equalf(t, tt.want, got, "FindMaxNumber() = %v, want %v", got, tt.want)
		})
	}
}
