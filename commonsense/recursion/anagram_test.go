package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAllPossibleAnagrams(t *testing.T) {
	// O(N!)
	tests := []struct {
		name string
		word string
		want []string
	}{
		{
			name: "a",
			word: "a",
			want: []string{"a"},
		},
		{
			name: "ab",
			word: "ab",
			want: []string{"ab", "ba"},
		},
		{
			name: "abc",
			word: "abc",
			want: []string{"abc", "bac", "bca", "acb", "cab", "cba"},
		},
		{
			name: "abcd",
			word: "abcd",
			want: []string{
				"abcd", "bacd", "bcad", "bcda",
				"acbd", "cabd", "cbad", "cbda",
				"acdb", "cadb", "cdab", "cdba",
				"abdc", "badc", "bdac", "bdca",
				"adbc", "dabc", "dbac", "dbca",
				"adcb", "dacb", "dcab", "dcba"},
		},
		{
			name: "abcde",
			word: "abcde",
			want: []string{
				"abcde", "bacde", "bcade", "bcdae", "bcdea",
				"acbde", "cabde", "cbade", "cbdae", "cbdea",
				"acdbe", "cadbe", "cdabe", "cdbae", "cdbea",
				"acdeb", "cadeb", "cdaeb", "cdeab", "cdeba",
				"abdce", "badce", "bdace", "bdcae", "bdcea",
				"adbce", "dabce", "dbace", "dbcae", "dbcea",
				"adcbe", "dacbe", "dcabe", "dcbae", "dcbea",
				"adceb", "daceb", "dcaeb", "dceab", "dceba",
				"abdec", "badec", "bdaec", "bdeac", "bdeca",
				"adbec", "dabec", "dbaec", "dbeac", "dbeca", //50
				"adebc", "daebc", "deabc", "debac", "debca",
				"adecb", "daecb", "deacb", "decab", "decba",
				"abced", "baced", "bcaed", "bcead", "bceda",
				"acbed", "cabed", "cbaed", "cbead", "cbeda",
				"acebd", "caebd", "ceabd", "cebad", "cebda",
				"acedb", "caedb", "ceadb", "cedab", "cedba",
				"abecd", "baecd", "beacd", "becad", "becda",
				"aebcd", "eabcd", "ebacd", "ebcad", "ebcda",
				"aecbd", "eacbd", "ecabd", "ecbad", "ecbda",
				"aecdb", "eacdb", "ecadb", "ecdab", "ecdba", //100
				"abedc", "baedc", "beadc", "bedac", "bedca",
				"aebdc", "eabdc", "ebadc", "ebdac", "ebdca",
				"aedbc", "eadbc", "edabc", "edbac", "edbca",
				"aedcb", "eadcb", "edacb", "edcab", "edcba",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FindAllPossibleAnagrams(tt.word), "FindAllPossibleAnagrams(%v)", tt.word)
		})
	}
}
