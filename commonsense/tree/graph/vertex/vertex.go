package vertex

type Vertex[T any] struct {
	value            T
	adjacentVertices []*Vertex[T]
}

func NewVertex[T any](value T) *Vertex[T] {
	return &Vertex[T]{
		value: value,
	}
}

// TODO find something more efficient than linear search
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