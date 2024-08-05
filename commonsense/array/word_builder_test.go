package array

import (
	"reflect"
	"testing"
)

func TestBuildTwoLevelWordArray(t *testing.T) {

	tests := []struct {
		name string
		args []rune
		want []string
	}{
		{
			name: "a,b",
			args: []rune{'a', 'b'},
			want: []string{"ab", "ba"},
		},
		{
			name: "a,b,c",
			args: []rune{'a', 'b', 'c'},
			want: []string{"ab", "ac", "ba", "bc", "ca", "cb"},
		},
		{
			name: "a,b,c,d",
			args: []rune{'a', 'b', 'c', 'd'},
			want: []string{"ab", "ac", "ad", "ba", "bc", "bd", "ca", "cb", "cd", "da", "db", "dc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildTwoLevelWordArray(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildTwoLevelWordArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildThreeLevelWordArray(t *testing.T) {

	tests := []struct {
		name string
		args []rune
		want []string
	}{
		{
			name: "a,b,c",
			args: []rune{'a', 'b', 'c'},
			want: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{
			name: "a,b,c,d",
			args: []rune{'a', 'b', 'c', 'd'},
			want: []string{"abc", "abd", "acb", "acd", "adb", "adc", "bac", "bad", "bca", "bcd", "bda", "bdc", "cab", "cad", "cba", "cbd", "cda", "cdb", "dab", "dac", "dba", "dbc", "dca", "dcb"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildThreeLevelWordArray(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildThreeLevelWordArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildFourLevelWordArray(t *testing.T) {

	tests := []struct {
		name string
		args []rune
		want []string
	}{
		{
			name: "a,b,c",
			args: []rune{'a', 'b', 'c'},
			want: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{
			name: "a,b,c,d",
			args: []rune{'a', 'b', 'c', 'd'},
			want: []string{"abc", "abd", "acb", "acd", "adb", "adc", "bac", "bad", "bca", "bcd", "bda", "bdc", "cab", "cad", "cba", "cbd", "cda", "cdb", "dab", "dac", "dba", "dbc", "dca", "dcb"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildFourLevelWordArray(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildThreeLevelWordArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
