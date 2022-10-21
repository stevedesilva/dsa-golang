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
func Test_stack_PeekEmpty(t *testing.T) {
	stack := New()
	if l := stack.Size(); l != 0 {
		t.Errorf("want 0 got %d", l)
	}
}
