package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode_New(t *testing.T) {
	tests := map[string]struct {
		root Node
	}{
		"tree_without_duplicate_ids": {
			root: buildTestRoot1(),
		},
		"tree_with_duplicate_ids": {
			root: buildTestRoot2(),
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
		root           Node
		expectedIDs    []int
		expectedLevels []int
	}{
		"tree_without_duplicate_ids": {
			root:           buildTestRoot1(),
			expectedIDs:    []int{1, 2, 3, 4, 5, 6},
			expectedLevels: []int{0, 1, 1, 2, 2, 2},
		},
		"tree_with_duplicate_ids": {
			root:           buildTestRoot2(),
			expectedIDs:    []int{3, 6, 2, 5, 3, 4, 1, 6},
			expectedLevels: []int{0, 1, 2, 2, 2, 3, 3, 4},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {

			tree := New(tt.root)

			idx := 0
			for n := range tree.Traverse() {
				assert.Equal(t, tt.expectedIDs[idx], n.node.GetID())
				assert.Equal(t, tt.expectedLevels[idx], n.level)
				idx++
			}
		})
	}
}

func buildTestRoot1() Node {
	node1 := NewNode(&testData{id: 1})
	node2 := NewNode(&testData{id: 2})
	node3 := NewNode(&testData{id: 3})
	node4 := NewNode(&testData{id: 4})
	node5 := NewNode(&testData{id: 5})
	node6 := NewNode(&testData{id: 6})

	node1.AddChildren(node2, node3)
	node2.AddChildren(node4)
	node3.AddChildren(node5, node6)

	return node1
}

func buildTestRoot2() Node {
	node1 := NewNode(&testData{id: 1})
	node2 := NewNode(&testData{id: 2})
	node3 := NewNode(&testData{id: 3})
	node4 := NewNode(&testData{id: 4})
	node5 := NewNode(&testData{id: 5})
	node6 := NewNode(&testData{id: 6})

	node3_duplicate := NewNode(&testData{id: 3})
	node6_duplicate := NewNode(&testData{id: 6})

	node3.AddChildren(node6)
	node6.AddChildren(node2, node5, node3_duplicate)
	node5.AddChildren(node4, node1)
	node1.AddChildren(node6_duplicate)

	return node3
}
