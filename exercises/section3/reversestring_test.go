package section3

import "testing"

func TestReverse(t *testing.T) {
	type data struct {
		name  string
		input string
		want  string
	}
	/*
		func TestReverse(t *testing.T) {
		    assert.Equal(t, Reverse(""), "")
		    assert.Equal(t, Reverse("X"), "X")
		    assert.Equal(t, Reverse("b\u0301"), "b\u0301")
		    assert.Equal(t, Reverse("😎⚽"), "⚽😎")
		    assert.Equal(t, Reverse("Les Mise\u0301rables"), "selbare\u0301siM seL")
		    assert.Equal(t, Reverse("ab\u0301cde"), "edcb\u0301a")
		    assert.Equal(t, Reverse("This `\xc5` is an invalid UTF8 character"), "retcarahc 8FTU dilavni na si `�` sihT")
		    assert.Equal(t, Reverse("The quick bròwn 狐 jumped over the lazy 犬"), "犬 yzal eht revo depmuj 狐 nwòrb kciuq ehT")
		}
	*/
	tests := []data{
		{name: "empty", input: "", want: ""},
		{name: "a -> a", input: "a", want: "a"},
		{name: "ab -> ba", input: "ab", want: "ba"},
		{name: "abc -> cba", input: "abc", want: "cba"},
		{name: "123 -> 321", input: "123", want: "321"},
		{name: "abcd -> dcba", input: "abcd", want: "dcba"},
		{name: "abcde -> edcba", input: "abcde", want: "edcba"},
		//{name: "b\u0301 -> b\u0301", input: "b\u0301", want: "b\u0301"},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := ReverseInlineOptimal(v.input); got != v.want {
				t.Errorf("want %s, got %s", v.want, got)
			}
			if got := ReverseWithResultArray(v.input); got != v.want {
				t.Errorf("want %s, got %s", v.want, got)
			}
		})
	}
}
