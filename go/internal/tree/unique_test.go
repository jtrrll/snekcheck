package tree_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"snekcheck/internal/tree"
)

func TestUniqueTree(t *testing.T) {
	t.Parallel()
	t.Run("NewUniqueTree()", func(t *testing.T) {
		t.Parallel()
		t.Run("creates a tree", func(t *testing.T) {
			t.Parallel()
			assert.NotNil(t, tree.NewUniqueTree[string]())
		})
	})
	t.Run("AddPath()", func(t *testing.T) {
		t.Parallel()
		t.Run("creates a full path of nodes", func(t *testing.T) {
			t.Parallel()
			testCases := [][]string{
				{"a"},
				{"a", "b", "c"},
				{"x", "y", "z"},
				{"x", "y", "x"},
			}
			uniqueTree := tree.NewUniqueTree[string]()
			require.NotNil(t, uniqueTree)
			for _, input := range testCases {
				assert.NotNil(t, uniqueTree.AddPath(input))
				node := uniqueTree
				var exists bool
				for _, part := range input {
					node, exists = node[part]
					assert.True(t, exists)
				}
			}
		})
	})
	t.Run("FindByPath()", func(t *testing.T) {
		t.Parallel()
		t.Run("finds the node at the end of a path", func(t *testing.T) {
			t.Parallel()
			testCases := [][]string{
				{"a"},
				{"a", "b", "c"},
				{"x", "y", "z"},
				{"x", "y", "x"},
			}
			uniqueTree := tree.NewUniqueTree[string]()
			require.NotNil(t, uniqueTree)
			for _, input := range testCases {
				node := uniqueTree.AddPath(input)
				require.NotNil(t, node)
				assert.Equal(t, node, uniqueTree.FindByPath(input))
			}
		})
		t.Run("returns nil for an invalid path", func(t *testing.T) {
			t.Parallel()
			testCases := [][]string{
				{"a"},
				{"a", "b", "c"},
				{"x", "y", "z"},
				{"x", "y", "x"},
			}
			uniqueTree := tree.NewUniqueTree[string]()
			require.NotNil(t, uniqueTree)
			for _, input := range testCases {
				assert.Nil(t, uniqueTree.FindByPath(input))
			}
		})
	})
	t.Run("IterPreOrder()", func(t *testing.T) {
		t.Parallel()
		t.Run("yields every node", func(t *testing.T) {
			t.Parallel()
			testCases := [][]string{
				{"a", "b", "c"},
				{"a", "d", "e"},
				{"d", "f", "g"},
			}
			seen := make(map[string]bool)
			uniqueTree := tree.NewUniqueTree[string]()
			require.NotNil(t, uniqueTree)
			for _, input := range testCases {
				require.NotNil(t, uniqueTree.AddPath(input))
				for _, part := range input {
					seen[part] = false
				}
			}
			for path := range uniqueTree.IterPreOrder() {
				seen[path[len(path)-1]] = true
			}
			for _, v := range seen {
				assert.True(t, v)
			}
		})
		t.Run("yields parent nodes before their children", func(t *testing.T) {
			t.Parallel()
			testCases := [][]string{
				{"a", "b", "c"},
				{"a", "d", "e"},
				{"d", "f", "g"},
			}
			seen := make(map[string]bool)
			uniqueTree := tree.NewUniqueTree[string]()
			require.NotNil(t, uniqueTree)
			for _, input := range testCases {
				require.NotNil(t, uniqueTree.AddPath(input))
				for _, part := range input {
					seen[part] = false
				}
			}
			for path := range uniqueTree.IterPreOrder() {
				seen[path[len(path)-1]] = true
				if len(path) > 1 {
					assert.True(t, seen[path[len(path)-2]])
				}
			}
			for _, v := range seen {
				assert.True(t, v)
			}
		})
		t.Run("can break out of iteration", func(t *testing.T) {
			t.Parallel()
			testCases := [][]string{
				{"a", "b", "c"},
				{"a", "d", "e"},
				{"d", "f", "g"},
			}
			uniqueTree := tree.NewUniqueTree[string]()
			require.NotNil(t, uniqueTree)
			for _, input := range testCases {
				require.NotNil(t, uniqueTree.AddPath(input))
			}
			iterations := 0
			for range uniqueTree.IterPreOrder() {
				if iterations > 2 {
					break
				}
				iterations++
			}
		})
	})
}
