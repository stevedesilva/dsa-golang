package applications

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name    string
	numbers []int
	want    *int
	wantErr assert.ErrorAssertionFunc
}{
	{
		name:    "test error",
		numbers: []int{},
		want:    nil,
		wantErr: assert.Error,
	},
	{
		name:    "test 1",
		numbers: []int{1, 2, 4, 5, 6, 6, 4, 10, 5, 2, 1, 23},
		want:    intPtr(23),
		wantErr: assert.NoError,
	},
	{
		name:    "test 2",
		numbers: []int{1, 2, 4, 5, 6, 6, 4, 10, 5, 2, 1},
		want:    intPtr(10),
		wantErr: assert.NoError,
	},
	{
		name:    "test 3",
		numbers: []int{1, 2, 4, 50, 6, 6, 4, 10, 5, 2, 1, 23},
		want:    intPtr(50),
		wantErr: assert.NoError,
	},
}

func TestLargestValueN(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LargestValueN(tt.numbers)
			if !tt.wantErr(t, err, fmt.Sprintf("LargestValueN(%v)", tt.numbers)) {
				return
			}
			assert.Equalf(t, tt.want, got, "LargestValueN(%v)", tt.numbers)
		})
	}
}

func TestLargestValueN2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LargestValueN2(tt.numbers)
			if !tt.wantErr(t, err, fmt.Sprintf("LargestValueN2(%v)", tt.numbers)) {
				return
			}
			assert.Equalf(t, tt.want, got, "LargestValueN2(%v)", tt.numbers)
		})
	}
}

func TestLargestValueNLogN(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LargestValueNLogN(tt.numbers)
			if !tt.wantErr(t, err, fmt.Sprintf("LargestValueNLogN(%v)", tt.numbers)) {
				return
			}
			assert.Equalf(t, tt.want, got, "LargestValueNLogN(%v)", tt.numbers)
		})
	}
}
