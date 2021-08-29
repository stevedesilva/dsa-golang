package section7

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {

	tests := []struct {
		name string
		args int
		want []string
	}{
		{name: "1", args: 1, want: []string{"1"}},
		{name: "1 2", args: 2, want: []string{"1", "2"}},
		{name: "1 2 Fizz", args: 3, want: []string{"1", "2", "Fizz"}},
		{name: "1 2 Fizz 4", args: 4, want: []string{"1", "2", "Fizz", "4"}},
		{name: "1 2 Fizz 4 5", args: 5, want: []string{"1", "2", "Fizz", "4", "Buzz"}},
		{name: "1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz", args: 15, want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzz(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
