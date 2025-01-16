// A collection of tree-like data structures.
package tree

import (
	"iter"
)

// A node in a tree with unique children.
type UniqueNode[T comparable] map[T]UniqueNode[T]

// Creates a new tree with unique children.
func NewUniqueTree[T comparable]() UniqueNode[T] {
	return make(map[T]UniqueNode[T])
}

// Adds a path to the tree by connecting any intermediate nodes.
func (tree UniqueNode[T]) AddPath(path []T) UniqueNode[T] {
	node := tree
	for _, part := range path {
		if _, exists := node[part]; !exists {
			node[part] = NewUniqueTree[T]()
		}
		node = node[part]
	}
	return node
}

// Attempts to find a node by its full path.
func (tree UniqueNode[T]) FindByPath(path []T) UniqueNode[T] {
	node := tree
	for _, part := range path {
		if _, exists := node[part]; !exists {
			return nil
		}
		node = node[part]
	}
	return node
}

// Iterates over the nodes of the tree by yielding root nodes before its children.
func (tree UniqueNode[T]) IterPreOrder() iter.Seq[[]T] {
	return func(yield func(path []T) bool) {
		for id, node := range tree {
			path := []T{id}
			if !yield(path) {
				return
			}
			for childPath := range node.IterPreOrder() {
				if !yield(append(path, childPath...)) {
					return
				}
			}
		}
	}
}
