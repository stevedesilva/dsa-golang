package node

type Node[T comparable] struct {
	Data T
	Next *Node[T]
}

//type Node[T any] struct {
//	data T
//	next *Node[T]
//}
