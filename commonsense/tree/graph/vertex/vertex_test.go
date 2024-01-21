package vertex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVertex_AddAdjacentVertex(t *testing.T) {
	v1 := NewVertex(1)
	v2 := NewVertex(2)

	v1.AddAdjacentVertex(v2)
	// v1 slice contains v2 testify assert
	assert.Equal(t, v2, v1.adjacentVertices[0])
	assert.Equal(t, v1, v2.adjacentVertices[0])

}
