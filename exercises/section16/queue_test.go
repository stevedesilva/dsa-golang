package section16_test

import (
	"testing"

	"github.com/stevedesilva/dsa-golang.git/exercises/section16"
	"github.com/stretchr/testify/assert"
)

func TestQueue_ShouldRemoveFromQueue(t *testing.T) {
	want := []section16.Any{1, 2, 3, 4}
	q := section16.New()
	q.Add(1)
	q.Add(2)
	q.Add(3)
	q.Add(4)

	assert.Equal(t, want, q.Values)
	v, _ := q.Remove()
	assert.Equal(t, 1, v)
	v, _ = q.Remove()
	assert.Equal(t, 2, v)
	v, _ = q.Remove()
	assert.Equal(t, 3, v)
	v, _ = q.Remove()
	assert.Equal(t, 4, v)

}
func TestQueue_ShouldPeekFromQueue(t *testing.T) {
	want := []section16.Any{1, 2, 3, 4}
	q := section16.New()
	q.Add(1)
	q.Add(2)
	q.Add(3)
	q.Add(4)

	assert.Equal(t, want, q.Values)
	v, _ := q.Peek()
	assert.Equal(t, 1, v)

	q.Remove()
	v, _ = q.Peek()
	assert.Equal(t, 2, v)

	q.Remove()
	v, _ = q.Peek()
	assert.Equal(t, 3, v)

	q.Remove()
	v, _ = q.Peek()
	assert.Equal(t, 4, v)

}

func TestQueue_ShouldAddToExistingQueue(t *testing.T) {
	want := []section16.Any{1, 2, 3, 4}
	q := section16.New()
	q.Add(1)
	q.Add(2)
	q.Add(3)
	q.Add(4)

	assert.Equal(t, want, q.Values)
}

func TestQueue_ShouldAddToEmptyQueue(t *testing.T) {
	want := make([]section16.Any, 0)
	q := section16.New()

	assert.Equal(t, want, q.Values)
}

func TestQueue_ShouldNotPeekFromEmptyQueue(t *testing.T) {
	q := section16.New()
	_, err := q.Peek()

	assert.Equal(t, section16.EmptyQueueErr, err)
}

func TestQueue_ShouldRemoveFromEmptyQueue(t *testing.T) {
	q := section16.New()
	_, err := q.Remove()

	assert.Equal(t, section16.EmptyQueueErr, err)
}
