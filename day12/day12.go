package day12

import (
	"adventofcode/utils"
	"adventofcode/utils/graph"
	"adventofcode/utils/queue"
	"fmt"
	"log"
	"math"
	"unicode"
)

type Area struct {
	nodes        []*graph.Node
	distances    []int
	startNodeIdx int
	endNodeIdx   int
}

func NewArea() *Area {
	// Initialize with a some length 10
	nodes := make([]*graph.Node, 0, 10)
	distances := make([]int, 0)
	return &Area{nodes, distances, -1, -1}
}

func (a *Area) AddNode(node *graph.Node) {
	a.nodes = append(a.nodes, node)
}

func (a *Area) FindIndex(node *graph.Node) int {
	for idx, item := range a.nodes {
		if item == node {
			return idx
		}
	}
	return -1
}

func Day12(part2 bool) {
	contents := utils.GetFileContents("./day12/input")
	area := NewArea()

	// Iterate counts and create nodes for each letter
	for i := 0; i < len(contents); i++ {
		for j := 0; j < len(contents[i]); j++ {

			ch := rune(contents[i][j])
			elevation := unicode.ToUpper(ch)

			// Check if node is the start or end
			if ch == 83 {
				area.startNodeIdx = len(area.nodes)
				elevation = 'A'
			} else if ch == 69 {
				area.endNodeIdx = len(area.nodes)
				elevation = 'Z'
			}

			area.AddNode(graph.New(elevation))
		}
	}

	width := len(contents[0])
	height := len(contents)

	// Iterate contents and build the graph
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {

			currentIdx := i*width + j
			currentNode := area.nodes[currentIdx]

			checkIdxs := [4]int{
				currentIdx - 1,
				currentIdx + 1,
				currentIdx - width,
				currentIdx + width,
			}

			// Check surrounding nodes and add as children if possible
			for _, idx := range checkIdxs {
				if idx < 0 || idx >= len(area.nodes) {
					continue
				}

				node := area.nodes[idx]

				// Calculate elevations and check if step is possible
				currentElevation := int(currentNode.Label() - 'A')
				nodeElevation := int(node.Label() - 'A')

				if nodeElevation-currentElevation <= 1 {
					currentNode.AddChild(node)
				}
			}
		}
	}

	startNodesIdxs := make([]int, 0)

	if part2 {
		for idx, node := range area.nodes {
			if node.Label() == 'A' {
				startNodesIdxs = append(startNodesIdxs, idx)
			}
		}
	} else {
		startNodesIdxs = append(startNodesIdxs, area.startNodeIdx)
	}

	lowestDist := math.MaxInt32

	for _, nodeIdx := range startNodesIdxs {

		// Reinitialize distances with length nodes
		area.distances = make([]int, len(area.nodes))
		for i := 0; i < len(area.distances); i++ {

			area.distances[i] = math.MaxInt32
			for j := 0; j < len(startNodesIdxs); j++ {
				if i == startNodesIdxs[j] {
					area.distances[i] = 0
					break
				}
			}
		}

		// Queue of the nodes to visit next
		queue := queue.New[*graph.Node]()
		// parent := make([]*graph.Node, len(area.nodes))
		visited := make(map[int]bool)

		// Add start node to the queue
		queue.Add(area.nodes[nodeIdx])
		visited[nodeIdx] = true

		// BFS transversal algorithm
		for {
			if queue.Empty() {
				// If the queue has no items, all nodes has been visited
				break
			}

			// Get node from the queue
			currentNode := queue.Get()
			currentNodeIdx := area.FindIndex(currentNode)

			for _, child := range currentNode.Children() {
				childIdx := area.FindIndex(child)
				if childIdx == -1 {
					log.Fatal("Could not find the node in the area nodes.")
				}

				if _, ok := visited[childIdx]; !ok {
					// Add child node to the queue
					visited[childIdx] = true
					area.distances[childIdx] = area.distances[currentNodeIdx] + 1
					queue.Add(child)
				}
			}
		}

		if lowestDist > area.distances[area.endNodeIdx] {
			lowestDist = area.distances[area.endNodeIdx]
		}
	}
	fmt.Println("Result:", lowestDist)
}
