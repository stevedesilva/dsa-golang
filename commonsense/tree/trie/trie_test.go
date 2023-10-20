package trie

import (
	"reflect"
	"testing"
)

func TestNewTrie(t *testing.T) {
	tests := []struct {
		name string
		want *Trie
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTrie(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTrie() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_Insert(t1 *testing.T) {
	type fields struct {
		Root *node
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trie{
				Root: tt.fields.Root,
			}
			t.Insert(tt.args.word)
		})
	}
}
