package search_test

import (
	"testing"

	"github.com/stevedesilva/dsa-golang.git/commonsense/search"
	"github.com/stretchr/testify/assert"
)

func TestWhenNotFoundReturnsErrorEmpty(t *testing.T) {
	_, err := search.FindInSortedArray(8, []int{})
	assert.Equal(t, search.NotFound, err)
}
func TestWhenNotFoundReturnsError(t *testing.T) {
	sortedArray := []int{1, 2, 4, 5, 7, 9}
	_, err := search.FindInSortedArray(8, sortedArray)
	assert.Equal(t, search.NotFound, err)
}

func TestWhenFoundReturnsIndexOneItem(t *testing.T) {
	sortedArray := []int{7}
	res, err := search.FindInSortedArray(7, sortedArray)
	assert.Nil(t, err)
	indexOfFoundItem := 0
	assert.Equal(t, indexOfFoundItem, res)
}
func TestWhenFoundReturnsIndexTwoItem(t *testing.T) {
	sortedArray := []int{1, 7}
	res, err := search.FindInSortedArray(7, sortedArray)
	assert.Nil(t, err)
	indexOfFoundItem := 1
	assert.Equal(t, indexOfFoundItem, res)
}
func TestWhenFoundReturnsIndexTwoItemReverse(t *testing.T) {
	sortedArray := []int{7, 1}
	res, err := search.FindInSortedArray(7, sortedArray)
	assert.Nil(t, err)
	indexOfFoundItem := 0
	assert.Equal(t, indexOfFoundItem, res)
}

func TestWhenFoundReturnsIndex(t *testing.T) {
	sortedArray := []int{1, 2, 4, 5, 7, 9}
	res, err := search.FindInSortedArray(7, sortedArray)
	assert.Nil(t, err)
	indexOfFoundItem := 4
	assert.Equal(t, indexOfFoundItem, res)
}
