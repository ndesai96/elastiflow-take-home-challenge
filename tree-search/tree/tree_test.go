package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode_New(t *testing.T) {
	tests := map[string]struct {
		root Node
	}{
		"happy": {
			root: buildRoot(),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {

			tree := New(tt.root)

			assert.IsType(t, &Tree{}, tree)
			assert.Equal(t, tt.root.GetID(), tree.root.GetID())
		})
	}
}

func TestNode_Traverse(t *testing.T) {
	tests := map[string]struct {
		tree           *Tree
		expectedIDs    []int
		expectedLevels []int
	}{
		"tree_without_duplicate_ids": {
			tree:           buildTreeWithoutDuplicateIDs(),
			expectedIDs:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			expectedLevels: []int{0, 1, 1, 2, 2, 2, 2, 3},
		},
		"tree_with_duplicate_ids": {
			tree:           buildTreeWithDuplicateIDs(),
			expectedIDs:    []int{3, 6, 8, 2, 5, 7, 4, 8, 1, 7},
			expectedLevels: []int{0, 1, 1, 2, 2, 2, 3, 3, 3, 4},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			idx := 0
			for n := range tt.tree.Traverse() {
				assert.Equal(t, tt.expectedIDs[idx], n.node.GetID())
				assert.Equal(t, tt.expectedLevels[idx], n.level)
				idx++
			}
		})
	}
}
