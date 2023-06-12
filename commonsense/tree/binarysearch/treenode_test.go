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

func TestTreeNodeSearch(t *testing.T) {
	left := TreeNode[int]{1, nil, nil}
	right := TreeNode[int]{3, nil, nil}
	root := TreeNode[int]{2, &left, &right}
	assert.Equal(t, 1, root.Search(1).data)
	assert.Equal(t, 2, root.Search(2).data)
	assert.Equal(t, 3, root.Search(3).data)
	assert.Nil(t, root.Search(13))
}
