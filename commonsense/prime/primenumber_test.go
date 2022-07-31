package prime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {

	data := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "0",
			input:    0,
			expected: false,
		},
		{
			name:     "1",
			input:    0,
			expected: false,
		},

		{
			name:     "2",
			input:    2,
			expected: true,
		},
		{
			name:     "3",
			input:    3,
			expected: true,
		},
		{
			name:     "4",
			input:    4,
			expected: false,
		},
		{
			name:     "5",
			input:    5,
			expected: true,
		},
		{
			name:     "6",
			input:    6,
			expected: false,
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			res := IsPrime(d.input)
			assert.Equal(t, d.expected, res, d.name)
		})
	}

}
