package binarysearch

import "golang.org/x/exp/constraints"

type TreeNode[T constraints.Ordered] struct {
	data  T
	left  *TreeNode[T]
	right *TreeNode[T]
}

func (t *TreeNode[T]) Search(value T) *TreeNode[T] {
	return t.search(value, t)
}

func (t *TreeNode[T]) search(value T, node *TreeNode[T]) *TreeNode[T] {
	if node == nil || value == node.data {
		return node
	}
	if value <= node.data {
		return t.search(value, node.left)
	} else {
		return t.search(value, node.right)
	}
}

func (t *TreeNode[T]) Insert(value T) *TreeNode[T] {
	return t.insert(value, t)
}

func (t *TreeNode[T]) insert(value T, node *TreeNode[T]) *TreeNode[T] {
	if node == nil || node.data == value {
		return node
	} else if value < node.data {
		if node.left == nil {
			node.left = &TreeNode[T]{value, nil, nil}
		}
		return t.insert(value, node.left)
	} else {
		if node.right == nil {
			node.right = &TreeNode[T]{value, nil, nil}
		}
		return t.insert(value, node.right)
	}
}

func (t *TreeNode[T]) Delete(i int) {
	// if node to delete has no children, just delete it
	// if node to delete has one child, replace it with its child
	// if node to delete has two children, replace it with its successor
	// successor is the smallest node in the right subtree
	panic("implement me")
}
