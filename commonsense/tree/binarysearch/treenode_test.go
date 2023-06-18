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

	leftChild := TreeNode[int]{5, nil, nil}
	rightChild := TreeNode[int]{7, nil, nil}
	left := TreeNode[int]{6, &leftChild, &rightChild}
	right := TreeNode[int]{13, nil, nil}
	root := TreeNode[int]{12, &left, &right}
	assert.Equal(t, 12, root.Search(12).data)
	assert.Equal(t, 13, root.Search(13).data)
	assert.Equal(t, 5, root.Search(5).data)
	assert.Equal(t, 6, root.Search(6).data)
	assert.Equal(t, 7, root.Search(7).data)
	assert.Nil(t, root.Search(19))
}
