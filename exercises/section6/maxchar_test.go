package section4

import "testing"

func TestReverse(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  string
	}{
		//{name: "1", input: "1", want: "1"},
		//{name: "12", input: "12", want: "1"},
		//{name: "122", input: "122", want: "2"},
		{name: "1 2 2 3", input: "1 2 2 3 ", want: "2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMax(tt.input); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
