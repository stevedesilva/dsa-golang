package binarysearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTreeNode(t *testing.T) {
	left := TreeNode[int]{1, nil, nil}
	right := TreeNode[int]{3, nil, nil}
	root := TreeNode[int]{2, &left, &right}
	assert.Equal(t, root.data, 2)
	assert.Equal(t, root.right.data, 3)
	assert.Equal(t, root.left.data, 1)
}
