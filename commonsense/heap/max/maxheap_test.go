package max_test

import (
	"testing"

	"github.com/stevedesilva/dsa-golang.git/commonsense/heap/max"
	"github.com/stretchr/testify/assert"
)

// get root node
func TestHeap_GetRootNode(t *testing.T) {
	heap := max.New[int]()
	root, err := heap.Root()
	assert.Nil(t, root)
	assert.NotNil(t, err)
}

// insert node
func TestHeap_Insert(t *testing.T) {
	heap := max.New[int]()
	err := heap.Insert(1)
	assert.Nil(t, err)
	root, err := heap.Root()
	assert.Nil(t, err)
	assert.NotNil(t, root)
	assert.Equal(t, 1, root)

	err = heap.Insert(4)
	assert.Nil(t, err)
	root, err = heap.Root()
	assert.Nil(t, err)
	assert.NotNil(t, root)
	assert.Equal(t, 4, root)

	err = heap.Insert(10)
	assert.Nil(t, err)
	root, err = heap.Root()
	assert.Nil(t, err)
	assert.NotNil(t, root)
	assert.Equal(t, 10, root)

	err = heap.Insert(2)
	assert.Nil(t, err)
	root, err = heap.Root()
	assert.Nil(t, err)
	assert.NotNil(t, root)
	assert.Equal(t, 10, root)
}

// delete node
