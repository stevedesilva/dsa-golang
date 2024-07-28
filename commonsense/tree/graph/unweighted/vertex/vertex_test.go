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
