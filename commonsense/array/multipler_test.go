package array

import (
	"reflect"
	"testing"
)

func TestFindTheProduct(t *testing.T) {

	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "1,2", input: []int{1, 2}, want: []int{2}},
		{name: "1,2,3", input: []int{1, 2, 3}, want: []int{2, 3, 6}},
		{name: "1,2,3,4", input: []int{1, 2, 3, 4}, want: []int{2, 3, 4, 6, 8, 12}},
		{name: "1,2,3,4,5", input: []int{1, 2, 3, 4, 5}, want: []int{2, 3, 4, 5, 6, 8, 10, 12, 15, 20}},
		{name: "1,2,3,4,5,6", input: []int{1, 2, 3, 4, 5, 6}, want: []int{2, 3, 4, 5, 6, 6, 8, 10, 12, 12, 15, 18, 20, 24, 30}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := FindTheProduct(tt.input); !reflect.DeepEqual(got, tt.want) && err == nil {
				t.Errorf("FindTheProduct() = %v err = %v, want %v, nil", got, err, tt.want)
			}
		})
	}
}

func TestFindTheProduct_ErrorsWhenInputLengthLessThanTwo(t *testing.T) {

	tests := []struct {
		name  string
		input []int
		want  error
	}{
		{name: "Empty", input: []int{}, want: ErrInvalidInput},
		{name: "Empty", input: []int{1}, want: ErrInvalidInput},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product, err := FindTheProduct(tt.input)
			if product != nil {
				t.Errorf("FindTheProduct() = %v should be nil", product)
			}

			if err != tt.want {
				t.Errorf("FindTheProduct() err = %v want %v", err, tt.want)
			}
		})
	}

}
