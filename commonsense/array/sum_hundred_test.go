package array

import "testing"

func TestOppositeValuesSumToHundredReturnErrorWhenMinimumInputNotMet(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "Empty",
			args: args{},
			want: ErrMinimumInputNotMet,
		},
		{
			name: "One ",
			args: args{numbers: []int{50}},
			want: ErrMinimumInputNotMet,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := OppositeValuesSumToHundred(tt.args.numbers); got != false && err != tt.want {
				t.Errorf("OppositeValuesSumToHundred() = %v %v, want false %v", got, err, tt.want)
			}
		})
	}
}

func TestOppositeValuesSumToHundred(t *testing.T) {
	type args struct {
		numbers []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Two  ",
			args: args{numbers: []int{50, 50}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := OppositeValuesSumToHundred(tt.args.numbers); got != tt.want {
				t.Errorf("OppositeValuesSumToHundred() = %v, want %v", got, tt.want)
			}
		})
	}
}
