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

// Constructor for Node
func NewNode(parent *Node, size int, is_file bool, name string) *Node {
	var children []*Node
	node := Node{parent, children, size, is_file, name}
	return &node
}

// Insert a new node to the slice
func (n *Node) AddNode(node *Node) {
	if n.is_file {
		log.Fatal("cannot add to a file")
	}

	n.children = append(n.children, node)
}

// Recursively transverse children nodes and
// calculate the sizes of directories and files
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

// Recursively transverse children nodes and count
// directory sizes below a certain limit
func (n *Node) CountDirSize(limit int) int {
	var size int = 0
	for _, node := range n.children {
		if node.is_file {
			continue
		}

		if node.size < limit {
			size += node.size
		}

		size += node.CountDirSize(limit)
	}
	return size
}

func (n *Node) Size() int {
	return n.size
}

func (n *Node) Name() string {
	return n.name
}

func (n *Node) Parent() *Node {
	return n.parent
}

// Check if a node has a child with specific name
func (n *Node) ExistChild(name string) (*Node, error) {
	for _, node := range n.children {
		if node.name == name {
			return node, nil
		}
	}
	return nil, errors.New("no child by name: " + name)
}

// Recursively transvese folders and look for the smallest folder
// that is larger than the specified limit
func (n *Node) FindSmallestDirAboveLimit(smallestNode *Node, limit int) *Node {
	for _, node := range n.children {
		if node.is_file {
			continue
		}

		if node.size > limit {
			if smallestNode == nil {
				smallestNode = node
			} else if node.size < smallestNode.size {
				smallestNode = node
			}

			smallestNode = node.FindSmallestDirAboveLimit(smallestNode, limit)
		}
	}
	return smallestNode
}
