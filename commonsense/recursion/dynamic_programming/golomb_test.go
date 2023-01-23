package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGolomb(t *testing.T) {

	tests := []struct {
		name   string
		number int
		want   int
	}{
		{
			name:   "1",
			number: 1,
			want:   1,
		},
		{
			name:   "2",
			number: 2,
			want:   2,
		},
		{
			name:   "3",
			number: 3,
			want:   2,
		},
		{
			name:   "4",
			number: 4,
			want:   3,
		},
		{
			name:   "5",
			number: 5,
			want:   3,
		},
		{
			name:   "10",
			number: 10,
			want:   5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Golomb(tt.number), "Golomb(%v)", tt.number)
		})
	}
}
