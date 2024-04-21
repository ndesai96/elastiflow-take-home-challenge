# Solution

`main.go` provides an example usecase of the `CheckDuplicateIDs` method. This example implements a tree structure from the `tree` package to build a hierarchical graph that represents a company's organizational structure. Once tree representation is built the `CheckDuplicateIDs` function is called to return the shallowest duplicate employee (if any) in said graph.

The example considers two scenarios:
1. An employee graph with at least one instance of employee duplication
2. An employee graph with no duplicated employees

For the first example, `CheckDuplicateIDs` returns the duplicated employee closest to the top of the tree (shallowest) along with the level of the duplicatation. Since the second example has no duplication, no employee is returned.

Under the hood, `CheckDuplicateIDs` traverses the provided tree using a breadth-first search (BFS) algorithm and stores each node's ID in a map. The first node ID found to already exist in this map is considered to be the shallowest duplicated node.

Since `CheckDuplicateIDs` uses a BFS algorithm and each node and node-to-node connection is visited once (worst case), the time complexity of this function is O(V+E) where V is the number of vertices (nodes) and E is the number of edges in the tree. Furthermore, the space complexity is O(V) since `CheckDuplicateIDs` uses a map to store all seen nodes. In the worst case where all nodes are the children of the root node, `CheckDuplicateIDs` will store all nodes in the seen map.

Check out the tree package [README](tree/README.md) to better understand how the Tree type itself is built.