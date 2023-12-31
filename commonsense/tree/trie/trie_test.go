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
			name:    "print single word",
			root:    testTrieRoot("test"),
			want:    []string{"test"},
			wantErr: assert.NoError,
		},
		{
			name:    "print words with different prefix",
			root:    testTrieRoot("test", "can", "man", "testing", "testify", "cat"),
			want:    []string{"test", "can", "man", "testing", "testify", "cat"},
			wantErr: assert.NoError,
		},
		{
			name:    "print words with same prefix",
			root:    testTrieRoot("test", "tester", "testing", "testify"),
			want:    []string{"test", "tester", "testing", "testify"},
			wantErr: assert.NoError,
		},
		{
			name:    "print words with different prefix",
			root:    testTrieRoot("can", "banner", "testing", "testify", "cat", "dog", "canada", "candy"),
			want:    []string{"can", "banner", "testing", "testify", "cat", "dog", "canada", "candy"},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trie{
				Root: tt.root,
			}

			got, err := t.CollectAllWords()
			if !tt.wantErr(t1, err, fmt.Sprintf("CollectAllWords()")) {
				return
			}
			assert.ElementsMatchf(t1, tt.want, got, "CollectAllWords()")
		})
	}
}

func TestTrie_shouldCollectKeysIsTrie(t *testing.T) {
	trie := createTestTrie("word", "worker")
	res, err := trie.CollectAllKeys()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	expected := []rune{'w', 'o', 'r', 'd', '*', 'k', 'e', 'r', '*'}
	assert.ElementsMatchf(t, expected, res, "CollectAllKeys()")
}

func TestTrie_shouldAutoCompletePrefixEmpty(t *testing.T) {
	trie := NewTrie()
	res, err := trie.AutoComplete("book")
	assert.Nil(t, res)
	assert.NotNil(t, err)
}

func TestTrie_shouldAutoCompletePrefixNotFound(t *testing.T) {
	trie := createTestTrie("word", "worker")
	res, err := trie.AutoComplete("book")
	assert.Nil(t, res)
	assert.NotNil(t, err)
}

func TestTrie_AutoComplete(t1 *testing.T) {
	trie := createTestTrie("word", "worker", "starter", "cube", "candle", "cat", "canter")
	got, err := trie.AutoComplete("wor")
	assert.NoError(t1, err)
	want := []string{"word", "worker"}
	assert.ElementsMatchf(t1, want, got, "AutoComplete()")
}

/*
 @Test
    public void shouldAutoCorrectWordNotFound() {
        Trie t = new Trie();
        final List<String> words = List.of("word", "worker", "starter", "cube", "candle", "cat", "canter");
        words.forEach(t::insert);
        final String result = t.autoCorrect("wors");
        MatcherAssert.assertThat(result, Matchers.equalTo("word"));
    }
*/

func TestTrie_AutoCorrectWordNotFound(t1 *testing.T) {
	trie := createTestTrie("word", "worker", "starter", "cube", "candle", "cat", "canter")
	_, err := trie.AutoCorrect("wors")
	assert.NotNil(t1, err)
}

/*
@Test
    public void shouldAutoCorrectWordFound() {
        Trie t = new Trie();
        final List<String> words = List.of("word", "worker", "starter", "cube", "candle", "cat", "canter");
        words.forEach(t::insert);
        MatcherAssert.assertThat(t.autoCorrect("word"), Matchers.equalTo("word"));
        MatcherAssert.assertThat(t.autoCorrect("cube"), Matchers.equalTo("cube"));
    }
*/

func TestTrie_AutoCorrectWordFound(t1 *testing.T) {
	trie := createTestTrie("word", "worker", "starter", "cube", "candle", "cat", "canter")
	got, err := trie.AutoCorrect("word")
	assert.Nil(t1, err)
	want := "word"
	assert.Equal(t1, want, got)
	got, err = trie.AutoCorrect("cube")
	assert.Nil(t1, err)
	want = "cube"
	assert.Equal(t1, want, got)
}

/*
@Test

	public void shouldThrowExceptionWhenAutoCorrectWordEmpty() {
	    Trie t = new Trie();
	    final List<String> words = List.of("word", "worker", "starter", "cube", "candle", "cat", "canter");
	    words.forEach(t::insert);
	    assertThrows(IllegalArgumentException.class, () -> t.autoCorrect(""));
	}
*/
func TestTrie_shouldReturnErrorWhenAutoCompleteWordEmpty(t *testing.T) {
	trie := createTestTrie("word", "worker", "starter", "cube", "candle", "cat", "canter")
	_, err := trie.AutoCorrect("")
	assert.NotNil(t, err)
}

func testTrieRoot(words ...string) *Node {
	trie := NewTrie()
	for _, word := range words {
		_ = trie.Insert(word)
	}
	return trie.Root
}

func createTestTrie(words ...string) *Trie {
	trie := NewTrie()
	for _, word := range words {
		_ = trie.Insert(word)
	}
	return trie
}
