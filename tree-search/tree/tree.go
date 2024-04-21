package tree

import (
	"github.com/golang-collections/collections/queue"
)

type Tree struct {
	root Node
}

func NewTree(root Node) *Tree {
	return &Tree{
		root: root,
	}
}

type NodeTraversal struct {
	node  Node
	level int
}

func (t *Tree) Traverse() <-chan NodeTraversal {
	ch := make(chan NodeTraversal)

	nodeQueue := queue.New()
	nodeQueue.Enqueue(NodeTraversal{node: t.root, level: 0})

	go func() {
		for nodeQueue.Len() > 0 {
			if nodeTraversal, ok := nodeQueue.Dequeue().(NodeTraversal); ok {
				ch <- nodeTraversal
				for _, n := range nodeTraversal.node.GetChildren() {
					nodeQueue.Enqueue(NodeTraversal{node: n, level: nodeTraversal.level + 1})
				}
			}
		}

		close(ch)
	}()

	return ch
}
