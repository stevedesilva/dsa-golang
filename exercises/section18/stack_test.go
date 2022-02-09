package section18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldPushToStack(t *testing.T) {
	s := New()
	s.Push("a")
	s.Push("b")
	s.Push("c")

	assert.Equal(t, 3, s.Size())
}

func TestShouldPeekFromStack(t *testing.T) {
	s := New()
	s.Push("a")
	s.Push("b")
	s.Push("c")

	assert.Equal(t, 3, s.Size())
	r, _ := s.Peek()
	assert.Equal(t, "c", r)

	s.Pop()
	r, _ = s.Peek()
	assert.Equal(t, "b", r)

	s.Pop()
	r, _ = s.Peek()
	assert.Equal(t, "a", r)

	s.Pop()
	assert.Equal(t, 0, s.Size())

}

func TestShouldPopFromStack(t *testing.T) {
	s := New()
	s.Push("a")
	s.Push("b")
	s.Push("c")

	assert.Equal(t, 3, s.Size())
	r, _ := s.Pop()
	assert.Equal(t, "c", r)
	r, _ = s.Pop()
	assert.Equal(t, "b", r)
	r, _ = s.Pop()
	assert.Equal(t, "a", r)
	assert.Equal(t, 0, s.Size())
}

func TestShouldNotPopFromEmptyStack(t *testing.T) {
	s := New()

	_, err := s.Pop()
	if err == nil {
		t.Errorf("should throw error")
	}
	assert.Equal(t, EmptyStackErr, err)
}

func TestShouldNotPeekFromEmptyStack(t *testing.T) {
	s := New()
	_, err := s.Peek()
	if err == nil {
		t.Errorf("should throw error")
	}
	assert.Equal(t, EmptyStackErr, err)
}
