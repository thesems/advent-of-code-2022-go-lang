package graph

import (
	"math"
	"sort"
)

type Node struct {
	Children []*Node
	Weights  []int
	Label    string
	Weight   int
}

func New(label string, weight int) *Node {
	var children []*Node
	return &Node{children, []int{}, label, weight}
}

func (n *Node) AddChild(child *Node, weight int) {
	n.Children = append(n.Children, child)
	n.Weights = append(n.Weights, weight)
}

type DistanceMap map[*Node]map[*Node]int

func FloydWarshall(nodes []*Node) DistanceMap {
	distanceMap := make(DistanceMap)

	// Initialize all paths to a large number
	for _, node1 := range nodes {
		distanceMap[node1] = make(map[*Node]int)
		for _, node2 := range nodes {
			distanceMap[node1][node2] = math.MaxInt32
		}
	}

	// Initialize known distances with edges
	// Set self-distances to 0
	for _, node := range nodes {
		for j, child := range node.Children {
			distanceMap[node][child] = node.Weights[j]
		}

		distanceMap[node][node] = 0
	}

	for k := 0; k < len(nodes); k++ {
		for i := 0; i < len(nodes); i++ {
			for j := 0; j < len(nodes); j++ {
				distIj := distanceMap[nodes[i]][nodes[j]]
				distIk := distanceMap[nodes[i]][nodes[k]]
				distKj := distanceMap[nodes[k]][nodes[j]]

				if distIj > distIk+distKj {
					distanceMap[nodes[i]][nodes[j]] = distIk + distKj
				}
			}
		}
	}

	return distanceMap
}


func (startNode *Node) Dijkstra(nodes []*Node) map[*Node]int {
	// Contains all shortest paths lengths from startNode
	tight := make(map[*Node]bool)
	distances := make(map[*Node]int)
	distances[startNode] = 0

	for _, node := range nodes {
		tight[startNode] = false

		if node == startNode {
			continue
		}
		distances[node] = math.MaxInt32
	}

	previous := make(map[*Node]*Node)
	previous[startNode] = nil

	queue := make([]*Node, 0)
	queue = append(queue, startNode)

	// Iterate over queue until all nodes are marked as tight
	for len(queue) > 0 {
		// Pop off first element and reduce the queue
		node := queue[0]
		queue = queue[1:]

		// Add all children to the queue
		// Set the distances of the children
	OUTER:
		for idx, child := range node.Children {
			if tight[child] {
				continue
			}

			newDist := distances[node] + node.Weights[idx]
			if newDist < distances[child] {
				distances[child] = newDist
			}

			for _, n := range queue {
				if n.Label == child.Label {
					continue OUTER
				}
			}

			queue = append(queue, child)
		}

		sort.Slice(queue, func(i, j int) bool {
			return distances[queue[i]] < distances[queue[j]]
		})

		tight[node] = true
	}

	return distances
}
