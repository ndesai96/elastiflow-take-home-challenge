# tree

The tree package is designed to be a standalone package that implements a generic N-ary tree.


## Tree

### Traversal

The `Traversal()` method performs a breadth-first search (BFS) traversal of the tree and returns a channel of `NodeTraversal`. Each `NodeTraversal` contains a node and its level in the tree.

### `NodeTraversal` Struct

Represents a node during traversal.

## Node

### `Node` Interface

The `Node` interface defines the basic operations supported by a node in the tree.

### `withID` Interface

The `withID` interface specifies a `GetID()` method which must be implemented by any struct intended to be used as data in a node. This interface enables the tree to maintain a unique identifier for each node's data while also allowing a node hold generic struct data.

### `node` Struct

The `node` struct represents a single node in the tree. It contains the following fields:

- `data`: The data associated with the node, implementing the withID interface.
- `children`: A slice containing the child nodes of the current node.