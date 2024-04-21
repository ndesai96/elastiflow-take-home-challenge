package tree

func CheckDuplicateIDs(tree *Tree) (Node, int) {
	seen := map[int]bool{}

	for nodeTraversal := range tree.Traverse() {
		node := nodeTraversal.GetNode()

		if _, exists := seen[node.GetID()]; exists {
			return node, nodeTraversal.GetLevel()
		} else {
			seen[node.GetID()] = true
		}
	}

	return nil, 0
}
