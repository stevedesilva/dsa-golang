package array_test

import (
	"reflect"
	"testing"

	"github.com/stevedesilva/dsa-golang.git/commonsense/array"
)

func TestPasswordCracker(t *testing.T) {
	type args struct {
		charSet    []rune
		desiredLen int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty",
			args: args{},
			want: []string{},
		},
		{
			name: "single combination",
			args: args{
				charSet:    []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'},
				desiredLen: 1,
			},
			want: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"},
		},
		{
			name: "a b - len 2",
			args: args{
				charSet:    []rune{'a', 'b'},
				desiredLen: 2,
			},
			want: []string{"aa", "ab", "ba", "bb"},
		},
		{
			name: "a b - 3",
			args: args{
				charSet:    []rune{'a', 'b'},
				desiredLen: 3,
			},
			want: []string{"aaa", "aab", "aba", "abb", "baa", "bab", "bba", "bbb"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := array.PasswordCracker(tt.args.charSet, tt.args.desiredLen); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PasswordCracker() = %v, want %v", got, tt.want)
			}
		})
	}
}
