package binarysearch

type TreeNode[T comparable] struct {
	data  T
	left  *TreeNode[T]
	right *TreeNode[T]
}
