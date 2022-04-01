package generics_test

import (
	"testing"

	. "github.com/stevedesilva/dsa-golang.git/exercises/section20/generics"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	n := New[string]("Test 1")
	assert.Equal(t, "Test 1", n.Value)
	assert.Nil(t, nil, n.Next)

	n2 := NewWithNext[string]("Test 2", n)
	assert.Equal(t, "Test 2", n2.Value)
	assert.Equal(t, "Test 1", n2.Next.Value)

}
