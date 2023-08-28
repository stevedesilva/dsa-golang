package v2

import "testing"

func Test_isMatched(t *testing.T) {
	type args struct {
		regex string
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"match p1",
			args{"^steve$", "developer steve is a developer"},
			false,
		},
		{
			"match p2",
			args{"^[\\d].*[a-z]$", "111117878 Steve is a developer"},
			true,
		},
		{
			"match p3",
			args{"steven", "steve"},
			false,
		},
		{
			"match p4",
			args{"steve", "steve"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatched(tt.args.regex, tt.args.value); got != tt.want {
				t.Errorf("isMatched() = %v, want %v", got, tt.want)
			}
		})
	}
}
