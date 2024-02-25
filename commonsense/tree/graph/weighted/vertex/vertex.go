package vertex

type Vertx[T comparable] struct {
	Value    T
	Adjacent map[*Vertx[T]]int
}

func New[T comparable](value T) *Vertx[T] {
	return &Vertx[T]{
		Value:    value,
		Adjacent: make(map[*Vertx[T]]int),
	}
}

func (v *Vertx[T]) addAdjacentVertx(vertx *Vertx[T], value int) {
	v.Adjacent[vertx] = value
}
