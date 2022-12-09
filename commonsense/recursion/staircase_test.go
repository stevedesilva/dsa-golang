package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countNumberOfSteps(t *testing.T) {

	tests := []struct {
		name   string
		number int
		want   int
	}{
		{
			name:   "step 1",
			number: 1,
			want:   1,
		},
		{
			name:   "step 2",
			number: 2,
			want:   2,
		},
		{
			name:   "step 3",
			number: 3,
			want:   4,
		},
		{
			name:   "step 4",
			number: 4,
			want:   7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countNumberOfSteps(tt.number), "countNumberOfSteps(%v)", tt.number)
		})
	}
}
