package tree

type testData struct {
	id int
}

func NewData(id int) *testData {
	return &testData{
		id: id,
	}
}

func (t *testData) GetID() int {
	return t.id
}

func buildTreeWithoutDuplicateIDs() *Tree {
	node1 := NewNode(NewData(1))
	node2 := NewNode(NewData(2))
	node3 := NewNode(NewData(3))
	node4 := NewNode(NewData(4))
	node5 := NewNode(NewData(5))
	node6 := NewNode(NewData(6))
	node7 := NewNode(NewData(7))
	node8 := NewNode(NewData(8))

	node1.AddChildren(node2, node3)
	node2.AddChildren(node4)
	node3.AddChildren(node5, node6, node7)
	node7.AddChildren(node8)

	return New(node1)
}

func buildTreeWithDuplicateIDs() *Tree {
	node1 := NewNode(NewData(1))
	node2 := NewNode(NewData(2))
	node3 := NewNode(NewData(3))
	node4 := NewNode(NewData(4))
	node5 := NewNode(NewData(5))
	node6 := NewNode(NewData(6))
	node7 := NewNode(NewData(7))
	node8 := NewNode(NewData(8))

	node7_duplicate := NewNode(NewData(7))
	node8_duplicate := NewNode(NewData(8))

	node3.AddChildren(node6, node8)
	node6.AddChildren(node2, node5, node7)
	node5.AddChildren(node4, node8_duplicate, node1)
	node1.AddChildren(node7_duplicate)

	return New(node3)
}

func buildRoot() Node {
	data := NewData(1)

	return NewNode(data)
}
