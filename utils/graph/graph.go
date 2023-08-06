package graph

type Node struct {
	Children []*Node
	Label    string
	Weight   int
}

func New(label string, weight int) *Node {
	var children []*Node
	return &Node{children, label, weight}
}

func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}
