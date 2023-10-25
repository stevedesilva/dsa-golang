package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie_Insert(t *testing.T) {
	// create test to insert a single word into a trie
	trie := NewTrie()
	trie.Insert("a")
	assert.Equal(t, 1, len(trie.Root.children))
	assert.Equal(t, 1, len(trie.Root.children['a'].children))
	assert.Equal(t, 0, len(trie.Root.children['a'].children['a'].children))

}

//func TestTrie_Insert(t *testing.T) {
//	tests := []struct {
//		name string
//		word string
//		want []string
//	}{
//		{
//			name: "exact match",
//			word: "a",
//			want: []string{"a"},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			res := NewTrie()
//			res.Insert(tt.word)
//
//		})
//	}
//}
