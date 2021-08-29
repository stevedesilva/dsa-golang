package section4

import "testing"

func TestReverse(t *testing.T) {

	tests := []struct {
		name  string
		input int
		want  int
	}{
		{name: "0", input: 0, want: 0},
		{name: "1", input: 1, want: 1},
		{name: "12", input: 12, want: 21},
		{name: "15", input: 15, want: 51},
		{name: "981", input: 981, want: 189},
		{name: "500", input: 500, want: 5},
		{name: "-15", input: -15, want: -51},
		{name: "-90", input: -90, want: -9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.input); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
