package binarysearch

type TreeNode[T comparable] struct {
	data  T
	left  *TreeNode[T]
	right *TreeNode[T]
}

func (t *TreeNode[T]) Search(value T) bool {
	return t.search(value, t)
}

func (t *TreeNode[T]) search(value T, node *TreeNode[T]) bool {
	// base
	if node == nil {
		return false
	}
	if value == node.data {
		return true
	}

	return t.search(value, node.left) || t.search(value, node.right)
}
