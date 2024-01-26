package vertex

import "fmt"

type Vertex[T any] struct {
	value            T
	adjacentVertices []*Vertex[T]
}

func NewVertex[T any](value T) *Vertex[T] {
	return &Vertex[T]{
		value: value,
	}
}

func (v *Vertex[T]) AddAdjacentVertex(adjacentVertex *Vertex[T]) {
	// if slice contains adjacentVertex, return
	for _, vertex := range v.adjacentVertices {
		if vertex == adjacentVertex {
			return
		}
	}
	// add to this vertex slice
	v.adjacentVertices = append(v.adjacentVertices, adjacentVertex)

	// add to calling vertex slice
	adjacentVertex.AddAdjacentVertex(v)
}

func (v *Vertex[T]) DfsTraverse(vertex *Vertex[T]) {
	visited := make(map[*Vertex[T]]bool, 0)
	dfsTraverse(vertex, visited)
}

func dfsTraverse[T any](vertex *Vertex[T], visited map[*Vertex[T]]bool) {
	// already visited
	visited[vertex] = true
	// print
	fmt.Printf("%v\n", vertex.value)
	// loop over adjacent vertx
	for _, v := range vertex.adjacentVertices {
		if _, ok := visited[v]; ok {
			continue
		} else {
			dfsTraverse(v, visited)
		}
	}
}

// dfs
func Dfs() {

}
