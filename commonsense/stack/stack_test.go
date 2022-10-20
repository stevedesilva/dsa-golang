package stack

import (
	"testing"
)

func Test_stack_Peek(t *testing.T) {
	stack := New("a", "b", "c")
	got := stack.Print()
	want := "[a b c]"
	if want != got {
		t.Errorf("want %s got %s", want, got)
	}
}
