package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTriangularNumber(t *testing.T) {
	tests := []struct {
		name                string
		positionToCalculate int
		want                int
	}{
		//{"1,1","2,3","3,6","4,10","5,15","6,21","7,28","8,36","9,45","10,55"}
		{
			name:                "1,1",
			positionToCalculate: 1,
			want:                1,
		},
		{
			name:                "2,3",
			positionToCalculate: 2,
			want:                3,
		},
		{
			name:                "3,6",
			positionToCalculate: 3,
			want:                6,
		},
		{
			name:                "4,10",
			positionToCalculate: 4,
			want:                10,
		},
		{
			name:                "5,15",
			positionToCalculate: 5,
			want:                15,
		},
		{
			name:                "6,21",
			positionToCalculate: 6,
			want:                21,
		},
		{
			name:                "7,28",
			positionToCalculate: 7,
			want:                28,
		},
		{
			name:                "8,36",
			positionToCalculate: 8,
			want:                36,
		},
		{
			name:                "9,45",
			positionToCalculate: 9,
			want:                45,
		},
		{
			name:                "10,55",
			positionToCalculate: 10,
			want:                55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CalculateTriangularNumber(tt.positionToCalculate), "CalculateTriangularNumber(%v)", tt.positionToCalculate)
		})
	}
}
