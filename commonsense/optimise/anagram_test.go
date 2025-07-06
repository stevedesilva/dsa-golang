package optimise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_areAnagrams(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test with anagrams",
			args: args{s1: "listen", s2: "silent"},
			want: true,
		},
		{
			name: "Test with non-anagrams",
			args: args{s1: "hello", s2: "world"},
			want: false,
		},
		{
			name: "Test with empty strings",
			args: args{s1: "", s2: ""},
			want: true,
		},
		{
			name: "Test with one empty string",
			args: args{s1: "test", s2: ""},
			want: false,
		},
		{
			name: "Test with different lengths",
			args: args{s1: "abc", s2: "ab"},
			want: false,
		},
		{
			name: "Test with same characters but different counts",
			args: args{s1: "aabbcc", s2: "abc"},
			want: false,
		},
		{
			name: "Test with anagrams with different cases",
			args: args{s1: "Listen", s2: "Silent"},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, areAnagrams(tt.args.s1, tt.args.s2), "areAnagrams(%v, %v)", tt.args.s1, tt.args.s2)
		})
	}
}
