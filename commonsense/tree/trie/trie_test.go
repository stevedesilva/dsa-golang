package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie_InsertEmptyWord_Error(t *testing.T) {
	// create test to insert a single word into a trie
	trie := NewTrie()
	err := trie.Insert("")
	assert.NotNil(t, err)
}

func TestTrie_InsertSingleLetter(t *testing.T) {
	// create test to insert a single word into a trie
	trie := NewTrie()
	err := trie.Insert("a")
	assert.Nil(t, err)
	notFound := trie.Root.children['z']
	assert.Nil(t, notFound)
	actual := trie.Root.children['a']
	assert.NotNil(t, actual)
	actualChild := trie.Root.children['*']
	assert.Nil(t, actualChild)
}
func TestTrie_SearchWord(t *testing.T) {
	// create test to insert a single word into a trie
	trie := NewTrie()
	err := trie.Insert("word")
	assert.Nil(t, err)
	res, err := trie.Search("word")
	assert.NotNil(t, res)
	assert.Nil(t, err)
}

func TestTrie_InsertWord(t *testing.T) {
	// create test to insert a single word into a trie
	trie := NewTrie()
	err := trie.Insert("test")
	assert.NotNil(t, err)
	res := trie.Root.children['t']
	assert.NotNil(t, res)
	res = res.children['e']
	assert.NotNil(t, res)
	res = res.children['s']
	assert.NotNil(t, res)
	res = res.children['t']
	assert.NotNil(t, res)
	res = res.children['*']
	assert.Nil(t, res)
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

// create  test for search where word is found
func TestFunctionalTest_Search(t *testing.T) {

}
