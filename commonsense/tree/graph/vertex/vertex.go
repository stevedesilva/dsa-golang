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

func (v *Vertex[T]) AddAdjacentVertex(adjacentVertex *Vertex[T]) {
	v.adjacentVertices = append(v.adjacentVertices, adjacentVertex)
}
