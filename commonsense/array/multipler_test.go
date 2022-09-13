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
		{name: "", input: []int{}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := FindTheProduct(tt.input); !reflect.DeepEqual(got, tt.want) && err != nil {
				t.Errorf("FindTheProduct() = %v err = %v, want %v, nil", got, err, tt.want)
			}
		})
	}
}

func TestFindTheProduct_ErrorsWhenInputLengthLessThanTwo(t *testing.T) {
	product, err := FindTheProduct([]int{})
	if product != nil {
		t.Errorf("FindTheProduct() = %v should be nil", product)
	}

	if err != ErrInvalidInput {
		t.Errorf("FindTheProduct() err = %v want %v", err, ErrInvalidInput)
	}
}
