package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CheckDuplicateIDs(t *testing.T) {
	tests := map[string]struct {
		tree          *Tree
		expectedNode  Node
		expectedLevel int
	}{
		"tree_without_duplicate_ids": {
			tree:          buildTreeWithoutDuplicateIDs(),
			expectedNode:  nil,
			expectedLevel: 0,
		},
		"tree_with_duplicate_ids": {
			tree:          buildTreeWithDuplicateIDs(),
			expectedNode:  NewNode(NewData(8)),
			expectedLevel: 3,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			node, level := CheckDuplicateIDs(tt.tree)

			assert.Equal(t, tt.expectedNode, node)
			assert.Equal(t, tt.expectedLevel, level)
		})
	}
}
