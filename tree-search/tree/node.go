package tree

// withID requires an ID getter method
type withID interface {
	GetID() int
}

type Node interface {
	withID

	GetData() withID
	AddChildren(children ...Node)
	GetChildren() []Node
}

// A node in a tree contains any struct type data
// as long as the struct has an ID getter method
type node struct {
	data     withID
	children []Node
}

func NewNode(data withID) Node {
	return &node{
		data:     data,
		children: []Node{},
	}
}

func (n *node) GetData() withID {
	return n.data
}

// Exposes the data's underlying ID getter method
func (n *node) GetID() int {
	return n.data.GetID()
}

func (n *node) AddChildren(children ...Node) {
	if children != nil {
		n.children = append(n.children, children...)
	}
}

func (n *node) GetChildren() []Node {
	return n.children
}
