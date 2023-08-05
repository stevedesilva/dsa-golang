package binarysearch

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTreeNode(t *testing.T) {
	left := TreeNode[int]{toPointer(1), nil, nil}
	right := TreeNode[int]{toPointer(3), nil, nil}
	root := TreeNode[int]{toPointer(2), &left, &right}
	assert.Equal(t, *root.data, 2)
	assert.Equal(t, *root.right.data, 3)
	assert.Equal(t, *root.left.data, 1)
}

func toPointer(val int) *int {
	return &val
}
func TestTreeNodeSearch(t *testing.T) {
	leftChild := TreeNode[int]{toPointer(5), nil, nil}
	rightChild := TreeNode[int]{toPointer(7), nil, nil}
	left := TreeNode[int]{toPointer(6), &leftChild, &rightChild}
	right := TreeNode[int]{toPointer(13), nil, nil}
	root := TreeNode[int]{toPointer(12), &left, &right}
	assert.Equal(t, 12, *root.Search(12).data)
	assert.Equal(t, 13, *root.Search(13).data)
	assert.Equal(t, 5, *root.Search(5).data)
	assert.Equal(t, 6, *root.Search(6).data)
	assert.Equal(t, 7, *root.Search(7).data)
	assert.Nil(t, root.Search(19))
}

func TestTreeNodeInsertInOrder(t *testing.T) {
	//      1
	//     / \
	//    2   nil
	//    ..
	//   /
	//  7
	root := TreeNode[int]{toPointer(1), nil, nil}
	root.Insert(2)
	root.Insert(3)
	root.Insert(4)
	root.Insert(5)
	root.Insert(6)
	root.Insert(7)

	assert.Equal(t, 1, *root.Search(1).data)
	assert.Equal(t, 2, *root.Search(2).data)
	assert.Equal(t, 3, *root.Search(3).data)
	assert.Equal(t, 4, *root.Search(4).data)
	assert.Equal(t, 5, *root.Search(5).data)
	assert.Equal(t, 6, *root.Search(6).data)
	assert.Equal(t, 7, *root.Search(7).data)
}

func TestTreeNodeInsertRandom(t *testing.T) {
	//      4
	//     / \
	//    3   5
	//   / \  / \
	//  1  2  6  7
	root := TreeNode[int]{toPointer(4), nil, nil}
	root.Insert(5)
	root.Insert(3)
	root.Insert(6)
	root.Insert(1)
	root.Insert(2)
	root.Insert(7)

	assert.Equal(t, 1, *root.Search(1).data)
	assert.Equal(t, 2, *root.Search(2).data)
	assert.Equal(t, 3, *root.Search(3).data)
	assert.Equal(t, 4, *root.Search(4).data)
	assert.Equal(t, 5, *root.Search(5).data)
	assert.Equal(t, 6, *root.Search(6).data)
	assert.Equal(t, 7, *root.Search(7).data)
}

func TestTreeNode_DeleteNodeNotFound(t *testing.T) {
	root := TreeNode[int]{toPointer(4), nil, nil}
	root.Insert(5)
	root.Insert(3)
	root.Insert(6)
	root.Insert(1)
	root.Insert(2)
	root.Insert(7)

	root.Delete(10, &root)
	assert.Nil(t, root.Search(10))
}

func TestTreeNode_DeleteNodeWithNoChildren(t *testing.T) {
	//       4
	//     /   \
	//    2     6
	//   / \   / \
	//  1   3 5   7
	root := TreeNode[int]{toPointer(4), nil, nil}
	root.Insert(2)
	root.Insert(6)
	root.Insert(1)
	root.Insert(3)
	root.Insert(5)
	root.Insert(7)

	root.Delete(1, &root)
	assert.Nil(t, root.Search(1))
	root.Delete(3, &root)
	assert.Nil(t, root.Search(3))
	root.Delete(5, &root)
	assert.Nil(t, root.Search(5))
	root.Delete(7, &root)
	assert.Nil(t, root.Search(7))
}

func TestTreeNode_DeleteNodeWithOneChildRight(t *testing.T) {
	//       4
	//     /   \
	//    2     6
	//     \    / \
	//      3  5   7
	root := TreeNode[int]{toPointer(4), nil, nil}
	root.Insert(2)
	root.Insert(6)
	root.Insert(3)
	root.Insert(5)
	root.Insert(7)

	root.Delete(2, &root)
	assert.Nil(t, root.Search(2))
	assert.Equal(t, 3, *root.Search(3).data)
}

func TestTreeNode_DeleteNodeWithOneChildLeft(t *testing.T) {
	//       4
	//     /   \
	//    2     6
	//   /     / \
	//  1     5   7
	root := TreeNode[int]{toPointer(4), nil, nil}
	root.Insert(2)
	root.Insert(6)
	root.Insert(1)
	root.Insert(5)
	root.Insert(7)

	root.Delete(2, &root)
	assert.Nil(t, root.Search(2))
	assert.Equal(t, 1, *root.Search(1).data)

}

func TestTreeNode_DeleteNodeWithOneChild_LargerTree(t *testing.T) {
	//         50
	//      / 	    \
	//     /         \
	//   25          75
	//   / \        /  \
	// 10   33      56    89
	//  \    /  \    / \    / \
	//  11  30  40  52 61  82  95

	root := TreeNode[int]{toPointer(50), nil, nil}
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
	assert.Equal(t, 11, *root.Search(11).data)
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

	root := TreeNode[int]{toPointer(50), nil, nil}
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
	assert.Equal(t, 61, *root.Search(61).data)
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

	root := TreeNode[int]{toPointer(50), nil, nil}
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
	assert.Equal(t, 52, *root.Search(52).data)
	assert.Equal(t, 55, *root.Search(55).data)
}

func TestTreeNode_PrintTree(t *testing.T) {
	//          50
	//     /        \
	//   25         75
	//   / \        / \
	// 11  33      61  89
	//    /  \    /   / \
	//   30  40  52  82  95
	//            \
	//            55

	root := TreeNode[int]{toPointer(50), nil, nil}
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
	outputCapture := captureOutput(
		func() {
			root.Print()
		})
	expectedOutput := "11\n25\n30\n33\n40\n50\n52\n55\n61\n75\n82\n89\n95\n"
	assert.Equal(t, expectedOutput, outputCapture)
}

func captureOutput(f func()) string {
	// redirect stand output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	_ = old

	// call function
	f()

	// reset standard output
	w.Close()
	os.Stdout = old

	// read the capture output from the buffer
	output := make(chan string)
	go func() {
		var buf string
		b := make([]byte, 100)
		for {
			read, err := r.Read(b)
			if err != nil {
				break
			}
			buf += string(b[:read])
		}
		output <- buf
	}()
	return <-output
}
