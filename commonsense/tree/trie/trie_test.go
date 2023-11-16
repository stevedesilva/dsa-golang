package trie

import (
	"fmt"
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
	assert.Nil(t, err)
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

func TestTrie_Search(t1 *testing.T) {
	tests := []struct {
		name    string
		root    *Node
		word    string
		want    *Node
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "word found",
			root: &Node{
				children: map[rune]*Node{
					't': {
						children: map[rune]*Node{
							'e': {
								children: map[rune]*Node{
									's': {
										children: map[rune]*Node{
											't': {
												children: map[rune]*Node{
													'*': nil,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			word: "test",
			want: &Node{
				children: map[rune]*Node{
					'*': nil,
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "word found",
			root: &Node{
				children: map[rune]*Node{
					's': {
						children: map[rune]*Node{
							'e': {
								children: map[rune]*Node{
									'l': {
										children: map[rune]*Node{
											'l': {
												children: map[rune]*Node{
													'*': nil,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			word:    "test",
			want:    newNode(),
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trie{
				Root: tt.root,
			}
			got, err := t.Search(tt.word)
			if !tt.wantErr(t1, err, fmt.Sprintf("Search(%v)", tt.word)) {
				return
			}
			if err == nil {
				assert.Equalf(t1, tt.want, got, "Search(%v)", tt.word)
			}
		})
	}
}

func TestTrie_PrintAllWords(t1 *testing.T) {

	tests := []struct {
		name    string
		root    *Node
		want    []string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "error on nil root",
			root:    nil,
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name:    "error on empty root",
			root:    newNode(),
			want:    []string{},
			wantErr: assert.NoError,
		},
		{
			name:    "print word",
			root:    testTrie("test"),
			want:    []string{"test"},
			wantErr: assert.NoError,
		},
		/*		{
					name:    "print words with same prefix",
					root:    testTrie("test", "tester", "testing", "testify"),
					want:    []string{"test", "tester", "testing", "testify"},
					wantErr: assert.NoError,
				},
				{
					name:    "print words with different prefix",
					root:    testTrie("can", "banner", "testing", "testify", "cat", "dog", "canada", "candy"),
					want:    []string{"can", "banner", "testing", "testify", "cat", "dog", "canada", "candy"},
					wantErr: assert.NoError,
				},*/
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trie{
				Root: tt.root,
			}
			got, err := t.PrintAllWords()
			if !tt.wantErr(t1, err, fmt.Sprintf("PrintAllWords()")) {
				return
			}
			assert.Equalf(t1, tt.want, got, "PrintAllWords()")
		})
	}
}

func testTrie(words ...string) *Node {
	trie := NewTrie()
	for _, word := range words {
		_ = trie.Insert(word)
	}
	return trie.Root
}
