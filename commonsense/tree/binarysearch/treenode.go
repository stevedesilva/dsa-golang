package binarysearch

type TreeNode[T comparable] struct {
	data  T
	left  *TreeNode[T]
	right *TreeNode[T]
}

func (t *TreeNode[T]) Search(value T) bool {
	return false

}
