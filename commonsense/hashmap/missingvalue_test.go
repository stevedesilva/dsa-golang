package hashmap

import (
	"reflect"
	"testing"
)

func TestMissingValue(t *testing.T) {
	type data struct {
		name  string
		input string
		want  string
	}
	tests := []data{
		{
			name:  "missing f",
			input: "the quick brown _ox jumps over the lazy dog",
			want:  "f",
		},
		{
			name:  "no missing character",
			input: "the quick brown fox jumps over the lazy dog",
			want:  "",
		},
		{
			name:  "missing t",
			input: "the quick brown fox jumps over the la_y dog",
			want:  "z",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findFirstMissingCharacter(tt.input)
			if res != tt.want {
				t.Errorf("findFirstMissingCharacter() want %s =, got %s", tt.want, res)
			}
		})
	}
}

func TestFindMissingCharacters(t *testing.T) {
	type data struct {
		name  string
		input string
		want  []string
	}
	tests := []data{
		{
			name:  "missing f",
			input: "the quick brown _ox jumps over the lazy dog",
			want:  []string{"f"},
		},
		{
			name:  "no missing character",
			input: "the quick brown fox jumps over the lazy dog",
			want:  []string{},
		},
		{
			name:  "missing t",
			input: "the quick brown fox jumps over the la_y dog",
			want:  []string{"z"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findMissingCharacters(tt.input)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("findMissingCharacters() want %s =, got %s", tt.want, res)
			}
		})
	}
}
