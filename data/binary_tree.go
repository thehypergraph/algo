package data

type Node struct {
	data  interface{}
	left  *Node
	right *Node
}

func (n *Node) New() *Node {
	return &Node{data: nil, left: nil, right: nil}
}
