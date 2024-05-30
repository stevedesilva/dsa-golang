package comparator_test

import (
	"testing"

	"github.com/stevedesilva/dsa-golang.git/commonsense/tree/heap/min/comparator"

	"github.com/stretchr/testify/assert"
)

var testFunc = func(i, j int) bool {
	return i < j
}

func TestHeap_GetRootNode(t *testing.T) {
	heap := comparator.New[int](testFunc)
	root, err := heap.Root()
	assert.Nil(t, root)
	assert.NotNil(t, err)
}

func TestHeap_InsertSingleValue(t *testing.T) {
	heap := comparator.New[int](testFunc)
	err := heap.Insert(1)
	assert.Nil(t, err)
	root, err := heap.Root()
	assert.Nil(t, err)
	assert.NotNil(t, root)
	assert.Equal(t, 1, *root)
}

func TestHeap_Insert(t *testing.T) {
	heap := comparator.New[int](testFunc)
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
	assert.Equal(t, 1, *root)

	err = heap.Insert(10)
	assert.Nil(t, err)
	root, err = heap.Root()
	assert.Nil(t, err)
	assert.Equal(t, 1, *root)

	err = heap.Insert(2)
	assert.Nil(t, err)
	root, err = heap.Root()
	assert.Nil(t, err)
	assert.Equal(t, 1, *root)

	err = heap.Insert(12)
	assert.Nil(t, err)
	root, err = heap.Root()
	assert.Nil(t, err)
	assert.Equal(t, 1, *root)

	err = heap.Insert(6)
	assert.Nil(t, err)
	root, err = heap.Root()
	assert.Nil(t, err)
	assert.Equal(t, 1, *root)
}

func TestHeap_DeleteEmptyHeap(t *testing.T) {
	minHeap := comparator.New[int](testFunc)
	v, err := minHeap.Delete()
	assert.Nil(t, v)
	assert.NotNil(t, err)
}

func TestHeap_DeleteSingleItem(t *testing.T) {
	minHeap := comparator.New[int](testFunc)
	_ = minHeap.Insert(9)

	v, err := minHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 9, *v)

	v, err = minHeap.Delete()
	assert.Nil(t, v)
	assert.NotNil(t, err)
}

func TestHeap_DeleteTwoItems(t *testing.T) {
	minHeap := comparator.New[int](testFunc)
	_ = minHeap.Insert(9)
	_ = minHeap.Insert(6)

	v, err := minHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 6, *v)

	v, err = minHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 9, *v)

	v, err = minHeap.Delete()
	assert.Nil(t, v)
	assert.NotNil(t, err)
}

func TestHeap_DeleteThreeItems(t *testing.T) {
	minHeap := comparator.New[int](testFunc)
	_ = minHeap.Insert(9)
	_ = minHeap.Insert(6)
	_ = minHeap.Insert(5)

	v, err := minHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 5, *v)

	v, err = minHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 6, *v)

	v, err = minHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 9, *v)

	v, err = minHeap.Delete()
	assert.Nil(t, v)
	assert.NotNil(t, err)
}

func TestHeap_Delete(t *testing.T) {
	minHeap := comparator.New[int](testFunc)
	_ = minHeap.Insert(9)
	_ = minHeap.Insert(6)
	_ = minHeap.Insert(5)
	_ = minHeap.Insert(1)
	_ = minHeap.Insert(2)
	_ = minHeap.Insert(10)
	_ = minHeap.Insert(3)
	_ = minHeap.Insert(7)
	_ = minHeap.Insert(8)
	_ = minHeap.Insert(4)
	root, err := minHeap.Root()

	assert.Equal(t, 1, *root)

	v, err := minHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 1, *v)

	v, err = minHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 2, *v)

	v, err = minHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 3, *v)

	v, err = minHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 4, *v)

	v, err = minHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 5, *v)

	v, err = minHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 6, *v)

	v, err = minHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 7, *v)

	v, err = minHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 8, *v)

	v, err = minHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 9, *v)

	v, err = minHeap.Delete()
	assert.Nil(t, err)
	assert.Equal(t, 10, *v)

	v, err = minHeap.Delete()
	assert.Nil(t, v)
	assert.NotNil(t, err)
}
