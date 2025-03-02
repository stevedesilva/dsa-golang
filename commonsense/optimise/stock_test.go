package optimise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStockPredictor(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "Test with positive and negative numbers",
			args: []int{1, -2, 3, 4, -5, 8},
			want: []int{1, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StockPredictor(tt.args)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
