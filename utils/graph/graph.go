package graph

type Node struct {
	children []*Node
	label    rune
}

func New(label rune) *Node {
	var children []*Node
	return &Node{children, label}
}

func (n *Node) AddChild(child *Node) {
	n.children = append(n.children, child)
}

func (n *Node) Label() rune {
	return n.label
}

func (n *Node) Children() []*Node {
	return n.children
}
