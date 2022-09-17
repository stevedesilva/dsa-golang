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

func TestFindTheProduct1(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindTheProduct(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindTheProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindTheProduct() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindTheProductFromTwoArrays(t *testing.T) {
	type args struct {
		inputA []int
		inputB []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "empty", args: args{[]int{}, []int{}}, want: []int{}},
		{name: "empty - 1", args: args{[]int{}, []int{1}}, want: []int{}},
		{name: "1 - empty", args: args{[]int{1}, []int{}}, want: []int{}},
		{name: "1-1", args: args{[]int{1}, []int{1}}, want: []int{1}},
		{name: "1,2-1", args: args{[]int{1, 2}, []int{1}}, want: []int{1, 2}},
		{name: "1,2-1,2", args: args{[]int{1, 2}, []int{1, 2}}, want: []int{1, 2, 2, 4}},
		{name: "1,2-1,2", args: args{[]int{1, 2, 3}, []int{1, 2, 3}}, want: []int{1, 2, 3, 2, 4, 6, 3, 6, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindTheProductFromTwoArrays(tt.args.inputA, tt.args.inputB)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindTheProductFromTwoArrays() got = %v, want %v", got, tt.want)
			}
		})
	}
}
