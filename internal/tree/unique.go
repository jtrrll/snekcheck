// TODO: package comment
package tree

import (
	"iter"
)

// A node in a tree with unique children.
type UniqueNode[T comparable] map[T]UniqueNode[T]

// Adds a path to the tree by connecting any intermediate nodes.
func (tree UniqueNode[T]) AddPath(path []T) UniqueNode[T] {
	node := tree
	for _, part := range path {
		if _, exists := node[part]; !exists {
			node[part] = make(map[T]UniqueNode[T])
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

// Iterates over the full paths of a tree, always yielding parent nodes before their children.
func (tree UniqueNode[T]) All() iter.Seq[[]T] {
	return func(yield func(path []T) bool) {
		for id, node := range tree {
			path := []T{id}
			if !yield(path) {
				return
			}
			for childPath := range node.All() {
				if !yield(append(path, childPath...)) {
					return
				}
			}
		}
	}
}
