package section19

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_ShouldAddToExistingQueue(t *testing.T) {
	q := New()
	q.Add(1)
	q.Add(2)
	q.Add(3)
	assert.Equal(t, 3, q.Size())
}

func TestQueue_ShouldRemoveOneItemFromQueue(t *testing.T) {
	q := New()
	q.Add(1)
	assert.Equal(t, 1, q.Size())
	i, _ := q.Remove()
	assert.Equal(t, 1, i)

}
func TestQueue_ShouldRemoveTwoItemsFromQueue(t *testing.T) {
	q := New()
	q.Add(1)
	q.Add(2)
	assert.Equal(t, 2, q.Size())
	i, _ := q.Remove()
	assert.Equal(t, 1, i)
	i, _ = q.Remove()
	assert.Equal(t, 2, i)

}

func TestQueue_ShouldRemoveNItemsFromQueue(t *testing.T) {
	q := New()
	q.Add(1)
	q.Add(2)
	q.Add(3)
	q.Add(4)
	assert.Equal(t, 4, q.Size())
	i, _ := q.Remove()
	assert.Equal(t, 1, i)
	i, _ = q.Remove()
	assert.Equal(t, 2, i)
	i, _ = q.Remove()
	assert.Equal(t, 3, i)
	i, _ = q.Remove()
	assert.Equal(t, 4, i)
}

func TestQueue_ShouldAddAndRemoveItemsFromQueue(t *testing.T) {
	// given
	q := New()
	q.Add(1)
	q.Add(2)

	// when adding and removing multiple items
	assert.Equal(t, 2, q.Size())
	i, _ := q.Remove()
	assert.Equal(t, 1, i)
	i, _ = q.Remove()
	assert.Equal(t, 2, i)

	q.Add(3)
	q.Add(4)
	assert.Equal(t, 2, q.Size())
	i, _ = q.Remove()
	assert.Equal(t, 3, i)
	i, _ = q.Remove()
	assert.Equal(t, 4, i)
}

func TestQueue_ShouldPeekFromQueue(t *testing.T) {
	q := New()
	q.Add(1)
	q.Add(2)
	q.Add(3)
	q.Add(4)
	assert.Equal(t, 4, q.Size())
	i, _ := q.Peek()
	assert.Equal(t, 1, i)
}

func TestQueue_ShouldNotPeekFromEmptyQueue(t *testing.T) {
	q := New()
	assert.Equal(t, 0, q.Size())
	_, err := q.Peek()

	assert.True(t, errors.Is(err, ErrQueue))
}

func TestQueue_ShouldNotRemoveFromEmptyQueue(t *testing.T) {
	q := New()
	assert.Equal(t, 0, q.Size())
	_, err := q.Remove()

	assert.True(t, errors.Is(err, ErrQueue))
}
