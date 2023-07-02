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

func TestTreeNodeInsertInOrder(t *testing.T) {
	//      1
	//     / \
	//    2   nil
	//    ..
	//   /
	//  7
	root := TreeNode[int]{}
	root.Insert(1)
	root.Insert(2)
	root.Insert(3)
	root.Insert(4)
	root.Insert(5)
	root.Insert(6)
	root.Insert(7)

	assert.Equal(t, 1, root.Search(1).data)
	assert.Equal(t, 2, root.Search(2).data)
	assert.Equal(t, 3, root.Search(3).data)
	assert.Equal(t, 4, root.Search(4).data)
	assert.Equal(t, 5, root.Search(5).data)
	assert.Equal(t, 6, root.Search(6).data)
	assert.Equal(t, 7, root.Search(7).data)
}

func TestTreeNodeInsertRandom(t *testing.T) {
	//      4
	//     / \
	//    3   5
	//   / \  / \
	//  1  2  6  7
	root := TreeNode[int]{}
	root.Insert(4)
	root.Insert(5)
	root.Insert(3)
	root.Insert(6)
	root.Insert(1)
	root.Insert(2)
	root.Insert(7)

	assert.Equal(t, 1, root.Search(1).data)
	assert.Equal(t, 2, root.Search(2).data)
	assert.Equal(t, 3, root.Search(3).data)
	assert.Equal(t, 4, root.Search(4).data)
	assert.Equal(t, 5, root.Search(5).data)
	assert.Equal(t, 6, root.Search(6).data)
	assert.Equal(t, 7, root.Search(7).data)
}

func TestTreeNode_DeleteNodeWithNoChildren(t *testing.T) {
	//      4
	//     / \
	//    3   5
	//   / \  / \
	//  1  2  6  7
	root := TreeNode[int]{}
	root.Insert(4)
	root.Insert(5)
	root.Insert(3)
	root.Insert(6)
	root.Insert(1)
	root.Insert(2)
	root.Insert(7)

	root.Delete(1)
	assert.Nil(t, root.Search(1))
	root.Delete(2)
	assert.Nil(t, root.Search(2))
	root.Delete(6)
	assert.Nil(t, root.Search(2))
	root.Delete(7)
	assert.Nil(t, root.Search(2))
}

func TestTreeNode_DeleteNodeWithOneChild(t *testing.T) {
	//      4
	//     / \
	//    3   5
	//   /    / \
	//  1    6   7
	//            \
	//             9
	root := TreeNode[int]{}
	root.Insert(4)
	root.Insert(5)
	root.Insert(3)
	root.Insert(6)
	root.Insert(1)
	root.Insert(7)
	root.Insert(9)

	root.Delete(3)
	assert.Nil(t, root.Search(3))

	root.Delete(7)
	assert.Nil(t, root.Search(7))
}

func TestTreeNode_DeleteRootNode(t *testing.T) {
	//      4
	//     / \
	//    3   5
	//   / \  / \
	//  1  2  6  7
	root := TreeNode[int]{}
	root.Insert(4)
	root.Insert(5)
	root.Insert(3)
	root.Insert(6)
	root.Insert(1)
	root.Insert(2)
	root.Insert(7)

	root.Delete(4)
	assert.Nil(t, root.Search(4))
}
