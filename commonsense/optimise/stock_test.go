package optimise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStockPredictor(t *testing.T) {
	tests := []struct {
		name string
		args []float64
		want bool
	}{
		{
			name: "Test with positive",
			args: []float64{22, 25, 18, 19.6, 17, 16, 20.5},
			want: true,
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
