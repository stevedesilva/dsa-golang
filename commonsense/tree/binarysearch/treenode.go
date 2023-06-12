package binarysearch

import "golang.org/x/exp/constraints"

type TreeNode[T constraints.Ordered] struct {
	data  T
	left  *TreeNode[T]
	right *TreeNode[T]
}

func (t *TreeNode[T]) Search(value interface{}) *TreeNode[T] {
	return t.search(value, t)
}

func (t *TreeNode[T]) search(value T, node *TreeNode[T]) *TreeNode[T] {
	if node == nil || value == node.data {
		return node
	}
	if value <= node.data {
		return t.search(value, node.left)
	}
	return t.search(value, node.right)
}
