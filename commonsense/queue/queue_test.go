package queue

import (
	"testing"
)

func Test_queue_Print(t *testing.T) {
	queue := New[string]("a", "b", "c")
	got := queue.Print()
	want := "[a b c]"
	if want != got {
		t.Errorf("want %s got %s", want, got)
	}
}

func Test_queue_Peek(t *testing.T) {
	queue := New[string]("a", "b", "c")
	got, _ := queue.Read()
	want := "c"
	if want != got {
		t.Errorf("want %s got %s", want, got)
	}
}

func Test_queue_PeekEmpty(t *testing.T) {
	queue := New[string]()
	if l := queue.Size(); l != 0 {
		t.Errorf("want 0 got %d", l)
	}
}

func Test_queue_Push(t *testing.T) {
	queue := New[string]()
	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c")
	if l := queue.Size(); l != 3 {
		t.Errorf("want 0 got %d", l)
	}
	want := "[a b c]"
	if got := queue.Print(); want != got {
		t.Errorf("want %s got %s", want, got)
	}
}

func Test_queue_Pop(t *testing.T) {
	queue := New[string]("a", "b", "c")
	queue.Dequeue()
	queue.Dequeue()
	if l := queue.Size(); l != 1 {
		t.Errorf("want 0 got %d", l)
	}
	want := "[a]"
	if got := queue.Print(); want != got {
		t.Errorf("want %s got %s", want, got)
	}
}

func Test_queue_Size(t *testing.T) {
	queue := New[string]("a", "b", "c")
	if l := queue.Size(); l != 3 {
		t.Errorf("want 0 got %d", l)
	}
}
