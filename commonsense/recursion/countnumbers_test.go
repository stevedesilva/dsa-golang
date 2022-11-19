package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
IntegerElement.of(1),
                IntegerElement.of(2),
                IntegerElement.of(3),
                ArrayElement.of(IntegerElement.of(4), IntegerElement.of(5), IntegerElement.of(6)),
                IntegerElement.of(7),
                ArrayElement.of(IntegerElement.of(8),
                        ArrayElement.of(IntegerElement.of(9), IntegerElement.of(10), IntegerElement.of(11),
                                ArrayElement.of(IntegerElement.of(12), IntegerElement.of(13), IntegerElement.of(14))
                        )
                ),
                ArrayElement.of(IntegerElement.of(15), IntegerElement.of(16), IntegerElement.of(17), IntegerElement.of(18), IntegerElement.of(19),
                        ArrayElement.of(IntegerElement.of(20), IntegerElement.of(21), IntegerElement.of(22),
                                ArrayElement.of(IntegerElement.of(23), IntegerElement.of(24), IntegerElement.of(25),
                                        ArrayElement.of(
                                                IntegerElement.of(26), IntegerElement.of(27), IntegerElement.of(29)
                                        )
                                ), IntegerElement.of(30), IntegerElement.of(31)
                        ), IntegerElement.of(32)
                ), IntegerElement.of(33)
*/

func TestGetNumbersFromSlice(t *testing.T) {
	tests := []struct {
		name  string
		array []any
		want  []int
	}{
		{
			name:  "1",
			array: []any{1},
			want:  []int{1},
		},
		{
			name:  "1 to 3",
			array: []any{1, 2, 3},
			want:  []int{1, 2, 3},
		},
		{
			name:  "1 to 14",
			array: []any{1, 2, 3, []any{4, 5, 6}, 7, []any{8, []any{9, 10, 11, []any{12, 13, 14}}}},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
		},
		{
			name:  "1 to 33",
			array: []any{1, 2, 3, []any{4, 5, 6}, 7, []any{8, []any{9, 10, 11, []any{12, 13, 14}}}, []any{15, 16, 17, 18, 19, []any{20, 21, 22, []any{23, 24, 25, []any{26, 27, 28, 29}, 30, 31}, 32}, 33}},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetNumbersFromSlice(tt.array), "GetNumbersFromSlice(%v)", tt.array)
		})
	}
}
