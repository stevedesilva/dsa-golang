package array_test

import (
	"reflect"
	"testing"

	"github.com/stevedesilva/dsa-golang.git/commonsense/array"
)

func TestPasswordCracker(t *testing.T) {

	tests := []struct {
		name   string
		number int
		want   []string
	}{
		{
			name:   "empty",
			number: 0,
			want:   []string{},
		},
		{
			name:   "single combi",
			number: 1,
			want:   []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"},
		},
		{
			name:   "a b - 2",
			number: 1,
			want:   []string{"aa", "ab", "ba", "bb"},
		},
		{
			name:   "a b - 3",
			number: 1,
			want:   []string{"aaa", "aab", "aba", "abb", "baa", "bab", "bba", "bbb"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := array.PasswordCracker(tt.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PasswordCracker() = %v, want %v", got, tt.want)
			}
		})
	}
}
