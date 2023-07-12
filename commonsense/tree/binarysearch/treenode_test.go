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
	root := TreeNode[int]{1, nil, nil}
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
	root := TreeNode[int]{4, nil, nil}
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

func TestTreeNode_DeleteNodeNotFound(t *testing.T) {
	root := TreeNode[int]{4, nil, nil}
	root.Insert(5)
	root.Insert(3)
	root.Insert(6)
	root.Insert(1)
	root.Insert(2)
	root.Insert(7)

	root.Delete(10, &root)
	assert.Nil(t, root.Search(1))

}

func TestTreeNode_DeleteNodeWithNoChildren(t *testing.T) {
	//      4
	//     / \
	//    3   5
	//   / \  / \
	//  1  2  6  7
	root := TreeNode[int]{4, nil, nil}
	root.Insert(5)
	root.Insert(3)
	root.Insert(6)
	root.Insert(1)
	root.Insert(2)
	root.Insert(7)

	root.Delete(1, &root)
	assert.Nil(t, root.Search(1))
	root.Delete(2, &root)
	assert.Nil(t, root.Search(2))
	root.Delete(6, &root)
	assert.Nil(t, root.Search(2))
	root.Delete(7, &root)
	assert.Nil(t, root.Search(2))
}

func TestTreeNode_DeleteNodeWithOneChild(t *testing.T) {
	//         50
	//      / 	    \
	//     /         \
	//   25          75
	//   / \        /  \
	// 10   33      56    89
	//  \    /  \    / \    / \
	//  11  30  40  52 61  82  95

	root := TreeNode[int]{50, nil, nil}
	root.Insert(25)
	root.Insert(75)
	root.Insert(11)
	root.Insert(33)
	root.Insert(56)
	root.Insert(89)
	root.Insert(11)
	root.Insert(30)
	root.Insert(40)
	root.Insert(52)
	root.Insert(61)
	root.Insert(95)

	root.Delete(10, &root)
	assert.Nil(t, root.Search(10))
}

func TestTreeNode_DeleteNodeWithTwoChildren(t *testing.T) {
	//          50
	//     /         \
	//   25          75
	//   / \        /  \
	// 11  33      56    89
	//    /  \    / \    / \
	//   30  40  52 61  82  95

	root := TreeNode[int]{50, nil, nil}
	root.Insert(25)
	root.Insert(75)
	root.Insert(11)
	root.Insert(33)
	root.Insert(56)
	root.Insert(89)
	root.Insert(30)
	root.Insert(40)
	root.Insert(52)
	root.Insert(61)
	root.Insert(95)

	root.Delete(56, &root)
	assert.Nil(t, root.Search(56))
}

func TestTreeNode_DeleteNodeWhereSuccessorNodeHasRightChild(t *testing.T) {
	//          50
	//     /        \
	//   25         75
	//   / \        / \
	// 11  33      61  89
	//    /  \    /   / \
	//   30  40  52  82  95
	//            \
	//            55

	root := TreeNode[int]{50, nil, nil}
	root.Insert(25)
	root.Insert(75)
	root.Insert(11)
	root.Insert(33)
	root.Insert(61)
	root.Insert(89)
	root.Insert(30)
	root.Insert(40)
	root.Insert(52)
	root.Insert(82)
	root.Insert(95)
	root.Insert(55)

	root.Delete(50, &root)
	assert.Nil(t, root.Search(50))
}
