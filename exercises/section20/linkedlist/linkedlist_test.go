package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLinkedList(t *testing.T) {

}

func TestNew(t *testing.T) {
	n := NewNode[string]("Hi")
	assert.Equal(t, "Hi", n.Value)
	assert.Nil(t, nil, n.Next)

	n2 := NewWithNext[string]("There", n)
	assert.Equal(t, "There", n2.Value)
	assert.Equal(t, "Hi", n2.Next.Value)

}
