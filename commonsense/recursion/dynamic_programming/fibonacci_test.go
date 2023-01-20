package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciWithMemoization(t *testing.T) {

	tests := []struct {
		name   string
		number int
		want   int
	}{
		// {"0,0","1,1","2,1","3,2","4,3","5,5","6,8","7,13","8,21","9,34","10,55", "11,89"}
		{name: "0,0", number: 0, want: 0},
		{name: "1,1", number: 1, want: 1},
		{name: "2,1", number: 2, want: 1},
		{name: "3,2", number: 3, want: 2},
		{name: "4,3", number: 4, want: 3},
		{name: "5,5", number: 5, want: 5},
		{name: "6,8", number: 6, want: 8},
		{name: "7,13", number: 7, want: 13},
		{name: "8,21", number: 8, want: 21},
		{name: "9,34", number: 9, want: 34},
		{name: "10,55", number: 10, want: 55},
		{name: "10,55", number: 11, want: 89},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Fibonacci(tt.number), "Fibonacci(%v)", tt.number)
		})
	}
}

func TestFibonacciNoRecursive(t *testing.T) {

	tests := []struct {
		name   string
		number int
		want   int
	}{
		// {"0,0","1,1","2,1","3,2","4,3","5,5","6,8","7,13","8,21","9,34","10,55", "11,89"}
		{name: "0,0", number: 0, want: 0},
		{name: "1,1", number: 1, want: 1},
		{name: "2,1", number: 2, want: 1},
		{name: "3,2", number: 3, want: 2},
		{name: "4,3", number: 4, want: 3},
		{name: "5,5", number: 5, want: 5},
		{name: "6,8", number: 6, want: 8},
		{name: "7,13", number: 7, want: 13},
		{name: "8,21", number: 8, want: 21},
		{name: "9,34", number: 9, want: 34},
		{name: "10,55", number: 10, want: 55},
		{name: "10,55", number: 11, want: 89},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FibonacciNoRecursive(tt.number), "FibonacciNoRecursive(%v)", tt.number)
		})
	}
}
func TestFibonacciNoRecursiveNoAdditionalSpace(t *testing.T) {

	tests := []struct {
		name   string
		number int
		want   int
	}{
		// {"0,0","1,1","2,1","3,2","4,3","5,5","6,8","7,13","8,21","9,34","10,55", "11,89"}
		{name: "0,0", number: 0, want: 0},
		{name: "1,1", number: 1, want: 1},
		{name: "2,1", number: 2, want: 1},
		{name: "3,2", number: 3, want: 2},
		{name: "4,3", number: 4, want: 3},
		{name: "5,5", number: 5, want: 5},
		{name: "6,8", number: 6, want: 8},
		{name: "7,13", number: 7, want: 13},
		{name: "8,21", number: 8, want: 21},
		{name: "9,34", number: 9, want: 34},
		{name: "10,55", number: 10, want: 55},
		{name: "10,55", number: 11, want: 89},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FibonacciNoRecursiveAlt(tt.number), "FibonacciNoRecursiveAlt(%v)", tt.number)
		})
	}
}
