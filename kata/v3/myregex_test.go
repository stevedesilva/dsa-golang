package v3

import "testing"

func TestMatchMyRegex(t *testing.T) {
	type args struct {
		pattern string
		word    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "MatchMyRegex(abc) = true",
			args: args{
				pattern: "abc",
				word:    "aaa abbbb abcd abc",
			},
			want: true,
		},
		{
			name: "MatchMyRegex(abcd) = true",
			args: args{
				pattern: "^[(abc)|(def)].*$",
				word:    "aaa abbbb abcd abc",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchMyRegex(tt.args.pattern, tt.args.word); got != tt.want {
				t.Errorf("MatchMyRegex() = %v, want %v", got, tt.want)
			}
		})
	}
}
