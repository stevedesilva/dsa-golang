package section9

import "testing"

func TestAnagramChecker(t *testing.T) {
	type args struct {
		wordA string
		wordB string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty returns false",
			args: args{
				wordA: "",
				wordB: "",
			},
			want: false,
		},
		{
			name: "remove none chars",
			args: args{
				wordA: "abc!&$!",
				wordB: "$abc!!&&",
			},
			want: true,
		},
		{
			name: "length not equal (abc, abcd)",
			args: args{
				wordA: "abc",
				wordB: "abcd",
			},
			want: false,
		},
		{
			name: "length not equal (abc, CbA)",
			args: args{
				wordA: "abc",
				wordB: "CbA",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnagramChecker(tt.args.wordA, tt.args.wordB); got != tt.want {
				t.Errorf("AnagramChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
