package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode_NewNode(t *testing.T) {
	tests := map[string]struct {
		data withID
	}{
		"happy": {
			data: &testData{id: 1},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {

			n := NewNode(tt.data)

			assert.IsType(t, &node{}, n)
			assert.Equal(t, tt.data, n.GetData())
			assert.Equal(t, tt.data.GetID(), n.GetID())
			assert.Empty(t, n.GetChildren())
		})
	}
}

func TestNode_AddChildren(t *testing.T) {
	node2 := NewNode(&testData{id: 2})
	node3 := NewNode(&testData{id: 3})

	tests := map[string]struct {
		node             Node
		children         []Node
		expectedChildren []Node
	}{
		"one child": {
			node:             NewNode(&testData{id: 1}),
			children:         []Node{node2},
			expectedChildren: []Node{node2},
		},
		"multiple children": {
			node:             NewNode(&testData{id: 1}),
			children:         []Node{node2, node3},
			expectedChildren: []Node{node2, node3},
		},
		"nil children": {
			node:             NewNode(&testData{id: 1}),
			children:         []Node{},
			expectedChildren: []Node{},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tt.node.AddChildren(tt.children...)

			assert.Equal(t, tt.expectedChildren, tt.node.GetChildren())
		})
	}
}

type testData struct {
	id int
}

func (t *testData) GetID() int {
	return t.id
}
