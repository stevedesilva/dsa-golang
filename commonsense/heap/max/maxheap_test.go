package max_test

import (
	"testing"

	"github.com/stevedesilva/dsa-golang.git/commonsense/heap/max"
	"github.com/stretchr/testify/assert"
)

// get root node
func TestGetRootNode(t *testing.T) {
	heap := max.New[int]()
	root, err := heap.Root()
	assert.Nil(t, root)
	assert.NotNil(t, err)
}

// insert node

// delete node
