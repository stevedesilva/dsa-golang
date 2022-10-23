package stack_test

import (
	"testing"

	. "github.com/stevedesilva/dsa-golang.git/commonsense/stack"
)

func Test_stack_Print(t *testing.T) {
	stack := New[string]("a", "b", "c")
	got := stack.Print()
	want := "[a b c]"
	if want != got {
		t.Errorf("want %s got %s", want, got)
	}
}

func Test_stack_Peek(t *testing.T) {
	stack := New[string]("a", "b", "c")
	got, _ := stack.Peek()
	want := "c"
	if want != got {
		t.Errorf("want %s got %s", want, got)
	}
}

func Test_stack_PeekEmpty(t *testing.T) {
	stack := New[string]()
	if l := stack.Size(); l != 0 {
		t.Errorf("want 0 got %d", l)
	}
}

func Test_stack_Push(t *testing.T) {
	stack := New[string]()
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	if l := stack.Size(); l != 3 {
		t.Errorf("want 0 got %d", l)
	}
	want := "[a b c]"
	if got := stack.Print(); want != got {
		t.Errorf("want %s got %s", want, got)
	}
}

func Test_stack_Pop(t *testing.T) {
	stack := New[string]("a", "b", "c")
	stack.Pop()
	stack.Pop()
	if l := stack.Size(); l != 1 {
		t.Errorf("want 0 got %d", l)
	}
	want := "[a]"
	if got := stack.Print(); want != got {
		t.Errorf("want %s got %s", want, got)
	}
}

func Test_stack_Size(t *testing.T) {
	stack := New[string]("a", "b", "c")
	if l := stack.Size(); l != 3 {
		t.Errorf("want 0 got %d", l)
	}
}
