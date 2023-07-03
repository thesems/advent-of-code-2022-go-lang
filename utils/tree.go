package utils

import (
	"errors"
	"log"
)

type Node struct {
	parent   *Node
	children []*Node
	size     int
	is_file  bool
	name     string
}

func NewNode(parent *Node, size int, is_file bool, name string) *Node {
	var children []*Node
	node := Node{parent, children, size, is_file, name}
	return &node
}

func (n *Node) AddNode(node *Node) {
	if n.is_file {
		log.Fatal("cannot add to a file")
	}

	n.children = append(n.children, node)
}

func (n *Node) CalculateSize() {
	for _, node := range n.children {
		if node.is_file {
			n.size += node.size
		} else {
			node.CalculateSize()
			n.size += node.size
		}
	}
}

func (n *Node) CountDirSize(sizeLimit int) int {
	var size int = 0
	for _, node := range n.children {
		if node.is_file {
			continue
		}

		if node.size < sizeLimit {
			size += node.size
		}

		size += node.CountDirSize(sizeLimit)
	}
	return size
}

func (n *Node) Size() int {
	return n.size
}

func (n *Node) Parent() *Node {
	return n.parent
}

func (n *Node) ExistChild(name string) (*Node, error) {
	for _, node := range n.children {
		if node.name == name {
			return node, nil
		}
	}
	return nil, errors.New("no child by name: " + name)
}