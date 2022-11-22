package counternumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestGetNumbersRestricted(t *testing.T) {
	tests := []struct {
		name  string
		array Element
		want  []int
	}{
		{
			name: "1",
			array: Element{
				ArrValue: ArrayElement{
					Elements: []Element{{IntValue: 1}},
				},
			},
			want: []int{1},
		},
		{
			name: "1 - 3",
			array: Element{
				ArrValue: ArrayElement{
					Elements: []Element{{IntValue: 1}, {IntValue: 2}, {IntValue: 3}},
				},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "1 - 7",
			array: Element{
				ArrValue: ArrayElement{[]Element{{IntValue: 1}, {IntValue: 2}, {IntValue: 3}, {
					ArrValue: ArrayElement{
						[]Element{{IntValue: 4}, {IntValue: 5}, {IntValue: 6}},
					}}, {IntValue: 7}},
				},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "1 - 14",
			array: Element{
				ArrValue: ArrayElement{[]Element{
					{IntValue: 1}, {IntValue: 2}, {IntValue: 3}, {ArrValue: ArrayElement{
						[]Element{{IntValue: 4}, {IntValue: 5}, {IntValue: 6}},
					}},
					{IntValue: 7},
					{IntValue: 8}, {ArrValue: ArrayElement{[]Element{{IntValue: 9}, {IntValue: 10}, {IntValue: 11},
						{ArrValue: ArrayElement{[]Element{{IntValue: 12}, {IntValue: 13}, {IntValue: 14}}}}},
					}},
				}},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
		},
		{
			name: "1 - 33",
			array: Element{
				ArrValue: ArrayElement{[]Element{
					{IntValue: 1}, {IntValue: 2}, {IntValue: 3}, {ArrValue: ArrayElement{
						[]Element{{IntValue: 4}, {IntValue: 5}, {IntValue: 6}},
					}},
					{IntValue: 7},
					{IntValue: 8}, {ArrValue: ArrayElement{[]Element{{IntValue: 9}, {IntValue: 10}, {IntValue: 11},
						{ArrValue: ArrayElement{[]Element{{IntValue: 12}, {IntValue: 13}, {IntValue: 14}}}}},
					}},
					{ArrValue: ArrayElement{[]Element{{IntValue: 15}, {IntValue: 16}, {IntValue: 17}, {IntValue: 18}, {IntValue: 19},
						{ArrValue: ArrayElement{[]Element{{IntValue: 20}, {IntValue: 21}, {IntValue: 22},
							{ArrValue: ArrayElement{[]Element{{IntValue: 23}, {IntValue: 24}, {IntValue: 25},
								{ArrValue: ArrayElement{[]Element{{IntValue: 26}, {IntValue: 27}, {IntValue: 28}, {IntValue: 29}}}},
								{IntValue: 30}, {IntValue: 31}}}}, {IntValue: 32}},
						}}, {IntValue: 33}},
					}},
				},
				},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetNumbersRestricted(tt.array), "GetNumbersFromSlice(%v)", tt.array)
		})
	}
}

func TestGetNumbersRestrictedInf(t *testing.T) {
	tests := []struct {
		name  string
		array ElementI[[]any]
		want  []int
	}{
		{
			name: "1",
			//array: ArrayEl[[]any]{},
			//array: ArrayEl{
			//
			//		Elements: []Element{{IntValue: 1}},
			//	},
			//},
			want: []int{1},
		},
		//	{
		//		name: "1 - 3",
		//		array: Element{
		//			ArrValue: ArrayElement{
		//				Elements: []Element{{IntValue: 1}, {IntValue: 2}, {IntValue: 3}},
		//			},
		//		},
		//		want: []int{1, 2, 3},
		//	},
		//	{
		//		name: "1 - 7",
		//		array: Element{
		//			ArrValue: ArrayElement{[]Element{{IntValue: 1}, {IntValue: 2}, {IntValue: 3}, {
		//				ArrValue: ArrayElement{
		//					[]Element{{IntValue: 4}, {IntValue: 5}, {IntValue: 6}},
		//				}}, {IntValue: 7}},
		//			},
		//		},
		//		want: []int{1, 2, 3, 4, 5, 6, 7},
		//	},
		//	{
		//		name: "1 - 14",
		//		array: Element{
		//			ArrValue: ArrayElement{[]Element{
		//				{IntValue: 1}, {IntValue: 2}, {IntValue: 3}, {ArrValue: ArrayElement{
		//					[]Element{{IntValue: 4}, {IntValue: 5}, {IntValue: 6}},
		//				}},
		//				{IntValue: 7},
		//				{IntValue: 8}, {ArrValue: ArrayElement{[]Element{{IntValue: 9}, {IntValue: 10}, {IntValue: 11},
		//					{ArrValue: ArrayElement{[]Element{{IntValue: 12}, {IntValue: 13}, {IntValue: 14}}}}},
		//				}},
		//			}},
		//		},
		//		want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
		//	},
		//	{
		//		name: "1 - 33",
		//		array: Element{
		//			ArrValue: ArrayElement{[]Element{
		//				{IntValue: 1}, {IntValue: 2}, {IntValue: 3}, {ArrValue: ArrayElement{
		//					[]Element{{IntValue: 4}, {IntValue: 5}, {IntValue: 6}},
		//				}},
		//				{IntValue: 7},
		//				{IntValue: 8}, {ArrValue: ArrayElement{[]Element{{IntValue: 9}, {IntValue: 10}, {IntValue: 11},
		//					{ArrValue: ArrayElement{[]Element{{IntValue: 12}, {IntValue: 13}, {IntValue: 14}}}}},
		//				}},
		//				{ArrValue: ArrayElement{[]Element{{IntValue: 15}, {IntValue: 16}, {IntValue: 17}, {IntValue: 18}, {IntValue: 19},
		//					{ArrValue: ArrayElement{[]Element{{IntValue: 20}, {IntValue: 21}, {IntValue: 22},
		//						{ArrValue: ArrayElement{[]Element{{IntValue: 23}, {IntValue: 24}, {IntValue: 25},
		//							{ArrValue: ArrayElement{[]Element{{IntValue: 26}, {IntValue: 27}, {IntValue: 28}, {IntValue: 29}}}},
		//							{IntValue: 30}, {IntValue: 31}}}}, {IntValue: 32}},
		//					}}, {IntValue: 33}},
		//				}},
		//			},
		//			},
		//		},
		//		want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33},
		//	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetNumbersRestrictedInf(tt.array), "GetNumbersFromSlice(%v)", tt.array)
		})
	}
}
