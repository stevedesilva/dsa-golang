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
