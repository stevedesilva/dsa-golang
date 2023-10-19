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

func TestHeap_InsertSingleValue(t *testing.T) {
	heap := max.New[int]()
	err := heap.Insert(1)
	assert.Nil(t, err)
	root, err := heap.Root()
	assert.Nil(t, err)
	assert.NotNil(t, root)
	assert.Equal(t, 1, *root)
}

/*
 //0,1,2,3,4,5,6,7,8,9,10
        //           0
        //      /          \
        //     1            2
        //    /  \        /  \
        //   3    4      5    6
        //  / \  / \    / \   / \
        // 7  8  9  10 11 12 12 14
        MaxHeap<Integer> maxHeap = new MaxHeap<>();
        maxHeap.insert(9);
        maxHeap.insert(6);
        maxHeap.insert(5);
        maxHeap.insert(1);
        maxHeap.insert(2);
        maxHeap.insert(10);
        maxHeap.insert(3);
        maxHeap.insert(7);
        maxHeap.insert(8);
        maxHeap.insert(4);
        MatcherAssert.assertThat( MaxHeap.getParentNode(1), Matchers.equalTo(0));
        MatcherAssert.assertThat( MaxHeap.getParentNode(2), Matchers.equalTo(0));
        MatcherAssert.assertThat( MaxHeap.getParentNode(3), Matchers.equalTo(1));
        MatcherAssert.assertThat( MaxHeap.getParentNode(4), Matchers.equalTo(1));
        MatcherAssert.assertThat( MaxHeap.getParentNode(5), Matchers.equalTo(2));
        MatcherAssert.assertThat( MaxHeap.getParentNode(6), Matchers.equalTo(2));
        MatcherAssert.assertThat( MaxHeap.getParentNode(7), Matchers.equalTo(3));
        MatcherAssert.assertThat( MaxHeap.getParentNode(8), Matchers.equalTo(3));
        MatcherAssert.assertThat( MaxHeap.getParentNode(11), Matchers.equalTo(5));
        MatcherAssert.assertThat( MaxHeap.getParentNode(14), Matchers.equalTo(6));
*/

// insert node
func TestHeap_Insert(t *testing.T) {
	heap := max.New[int]()
	err := heap.Insert(1)
	assert.Nil(t, err)
	root, err := heap.Root()
	assert.Nil(t, err)
	assert.NotNil(t, root)
	assert.Equal(t, 1, *root)

	err = heap.Insert(4)
	assert.Nil(t, err)
	root, err = heap.Root()
	assert.Nil(t, err)
	assert.Equal(t, 4, *root)

	err = heap.Insert(10)
	assert.Nil(t, err)
	root, err = heap.Root()
	assert.Nil(t, err)
	assert.Equal(t, 10, *root)

	err = heap.Insert(2)
	assert.Nil(t, err)
	root, err = heap.Root()
	assert.Nil(t, err)
	assert.Equal(t, 10, *root)

	err = heap.Insert(12)
	assert.Nil(t, err)
	root, err = heap.Root()
	assert.Nil(t, err)
	assert.Equal(t, 12, *root)

	err = heap.Insert(6)
	assert.Nil(t, err)
	root, err = heap.Root()
	assert.Nil(t, err)
	assert.Equal(t, 12, *root)
}

func TestHeap_DeleteEmptyHeap(t *testing.T) {
	maxHeap := max.New[int]()
	v, err := maxHeap.Delete()
	assert.Nil(t, v)
	assert.NotNil(t, err)
}

func TestHeap_DeleteSingleItem(t *testing.T) {
	maxHeap := max.New[int]()
	_ = maxHeap.Insert(9)

	v, err := maxHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 9, *v)

	v, err = maxHeap.Delete()
	assert.Nil(t, v)
	assert.NotNil(t, err)
}

func TestHeap_DeleteTwoItems(t *testing.T) {
	maxHeap := max.New[int]()
	_ = maxHeap.Insert(9)
	_ = maxHeap.Insert(6)

	v, err := maxHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 9, *v)

	v, err = maxHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 6, *v)

	v, err = maxHeap.Delete()
	assert.Nil(t, v)
	assert.NotNil(t, err)
}

func TestHeap_DeleteThreeItems(t *testing.T) {
	maxHeap := max.New[int]()
	_ = maxHeap.Insert(9)
	_ = maxHeap.Insert(6)
	_ = maxHeap.Insert(5)

	v, err := maxHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 9, *v)

	v, err = maxHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 6, *v)

	v, err = maxHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 5, *v)

	v, err = maxHeap.Delete()
	assert.Nil(t, v)
	assert.NotNil(t, err)
}

func TestHeap_Delete(t *testing.T) {
	maxHeap := max.New[int]()
	_ = maxHeap.Insert(9)
	_ = maxHeap.Insert(6)
	_ = maxHeap.Insert(5)
	_ = maxHeap.Insert(1)
	_ = maxHeap.Insert(2)
	_ = maxHeap.Insert(10)
	_ = maxHeap.Insert(3)
	_ = maxHeap.Insert(7)
	_ = maxHeap.Insert(8)
	_ = maxHeap.Insert(4)
	root, err := maxHeap.Root()

	assert.Equal(t, 10, *root)

	v, err := maxHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 10, *v)

	v, err = maxHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 9, *v)

	v, err = maxHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 8, *v)

	v, err = maxHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 7, *v)

	v, err = maxHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 6, *v)

	v, err = maxHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 5, *v)

	v, err = maxHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 4, *v)

	v, err = maxHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 3, *v)

	v, err = maxHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 2, *v)

	v, err = maxHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 1, *v)

	v, err = maxHeap.Delete()
	assert.Nil(t, v)
	assert.NotNil(t, err)
}
