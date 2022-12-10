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
		// 1 = 1
		// 2 = 11, 2
		// 3 = 111, 12, 21, 3
		// 4 = 1111, 112, 121, 211, 22, 31, 13
		// 5 = 11111, 1112, 1121, 1211, 2111, 221, 212, 122, 311, 131, 113, 32, 23
		// 6 = 111111, 11112, 11121, 11211, 12111, 21111, 2211,2121,2112,222,231,213, 3111,1311,1131,1113, 321,312, 33
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
		{
			name:   "step 5",
			number: 5,
			want:   13,
		},
		{
			name:   "step 5",
			number: 6,
			want:   19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CountNumberOfSteps(tt.number), "CountNumberOfSteps(%v)", tt.number)
		})
	}
}
