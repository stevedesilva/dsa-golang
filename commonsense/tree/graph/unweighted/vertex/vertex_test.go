package vertex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVertx_Shortest_Path(t *testing.T) {
	// create vertices
	idris := NewVertex("idris")
	kamil := NewVertex("kamil")
	lina := NewVertex("lina")
	sasha := NewVertex("sasha")
	marco := NewVertex("marco")
	ken := NewVertex("ken")
	talia := NewVertex("talia")

	// add adjacent vertices
	idris.AddAdjacentVertex(kamil)
	idris.AddAdjacentVertex(talia)

	kamil.AddAdjacentVertex(idris)
	kamil.AddAdjacentVertex(lina)

	lina.AddAdjacentVertex(kamil)
	lina.AddAdjacentVertex(sasha)

	sasha.AddAdjacentVertex(lina)
	sasha.AddAdjacentVertex(marco)

	marco.AddAdjacentVertex(sasha)
	marco.AddAdjacentVertex(ken)

	ken.AddAdjacentVertex(marco)
	ken.AddAdjacentVertex(talia)

	talia.AddAdjacentVertex(ken)
	talia.AddAdjacentVertex(idris)

	// given a graph
	graph := NewVertex("")
	graph.shortestPath(idris, lina)

	// then we should get the shortest path
	expected := []string{"idris", "kamil", "lina"}
	assert.Equal(t, graph.shortestPath(idris, lina), expected)

}

func TestDfs(t *testing.T) {
	v1 := NewVertex[int](1)
	v2 := NewVertex[int](2)
	v3 := NewVertex[int](3)

	v1.AddAdjacentVertex(v2)
	v1.AddAdjacentVertex(v3)

	dfs, err := v1.DepthFirstSearch(1)
	assert.Nil(t, err)
	assert.Equal(t, dfs, v1)

	dfs, err = v2.DepthFirstSearch(1)
	assert.Nil(t, err)
	assert.Equal(t, dfs, v1)

	dfs, err = v3.DepthFirstSearch(1)
	assert.Nil(t, err)
	assert.Equal(t, dfs, v1)

	dfs, err = v1.DepthFirstSearch(2)
	assert.Nil(t, err)
	assert.Equal(t, dfs, v2)

	dfs, err = v2.DepthFirstSearch(2)
	assert.Nil(t, err)
	assert.Equal(t, dfs, v2)

	dfs, err = v3.DepthFirstSearch(2)
	assert.Nil(t, err)
	assert.Equal(t, dfs, v2)

	dfs, err = v1.DepthFirstSearch(3)
	assert.Nil(t, err)
	assert.Equal(t, dfs, v3)

	dfs, err = v2.DepthFirstSearch(3)
	assert.Nil(t, err)
	assert.Equal(t, dfs, v3)

	dfs, err = v3.DepthFirstSearch(3)
	assert.Nil(t, err)
	assert.Equal(t, dfs, v3)
}
