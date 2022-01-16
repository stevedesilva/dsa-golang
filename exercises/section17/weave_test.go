package section17

import (
	"testing"

	"github.com/stevedesilva/dsa-golang.git/exercises/section16"
	"github.com/stretchr/testify/assert"
)

func TestQueue_ShouldRemoveFromQueue(t *testing.T) {

	want := []section16.Any{"Hi", 1, "To", 2, "You", 3}
	q1 := section16.New()
	q1.Add("Hi")
	q1.Add("To")
	q1.Add("You")

	q2 := section16.New()
	q2.Add(1)
	q2.Add(2)
	q2.Add(3)

	q, _ := Weave(*q1, *q2)

	assert.Equal(t, want, q.Values)

	v, _ := q.Remove()
	assert.Equal(t, "Hi", v)

	v, _ = q.Remove()
	assert.Equal(t, 1, v)

	v, _ = q.Remove()
	assert.Equal(t, "To", v)

	v, _ = q.Remove()
	assert.Equal(t, 2, v)

	v, _ = q.Remove()
	assert.Equal(t, "You", v)

	v, _ = q.Remove()
	assert.Equal(t, 3, v)

}
