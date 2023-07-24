package binarysearch

import "golang.org/x/exp/constraints"

type TreeNode[T constraints.Ordered] struct {
	data  *T
	left  *TreeNode[T]
	right *TreeNode[T]
}

func (t *TreeNode[T]) toPointer(val T) *T {
	return &val
}

func (t *TreeNode[T]) Search(value T) *TreeNode[T] {
	return t.search(t.toPointer(value), t)
}

func (t *TreeNode[T]) search(value *T, node *TreeNode[T]) *TreeNode[T] {
	if node == nil || value == node.data {
		return node
	}
	if *value <= *node.data {
		return t.search(value, node.left)
	} else {
		return t.search(value, node.right)
	}
}

func (t *TreeNode[T]) Insert(value T) *TreeNode[T] {
	return t.insert(t.toPointer(value), t)
}

func (t *TreeNode[T]) insert(value *T, node *TreeNode[T]) *TreeNode[T] {
	if node == nil || node.data == nil || node.data == value {
		if node == nil {
			node = &TreeNode[T]{value, nil, nil}
		} else if node.data == nil {
			node.data = value
		}
		return node

	} else if *value < *node.data {
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
func (t *TreeNode[T]) Delete(value T, node *TreeNode[T]) *TreeNode[T] {
	return t.delete(t.toPointer(value), node)
}
func (t *TreeNode[T]) delete(value *T, node *TreeNode[T]) *TreeNode[T] {
	if node == nil {
		return nil
	} else if *value < *node.data {
		node.left = t.delete(value, node.left)
		return node
	} else if *value > *node.data {
		node.right = t.delete(value, node.right)
		return node
	} else if node.data == value {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		} else {
			node.right = t.lift(node.right, node)
			return node
		}
	}

	// if node to delete has no children, just delete it
	// if node to delete has one child, replace it with its child
	// if node to delete has two children, replace it with its successor
	// successor is the smallest node in the right subtree
	return node
}

func (t *TreeNode[T]) lift(node *TreeNode[T], nodeToDelete *TreeNode[T]) *TreeNode[T] {
	if node.left != nil {
		node.left = t.lift(node.left, nodeToDelete)
		return node
	} else {
		nodeToDelete.data = node.data
		return node.right
	}
}
