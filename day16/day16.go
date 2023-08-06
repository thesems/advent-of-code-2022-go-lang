package day16

import (
	"adventofcode/utils"
	"adventofcode/utils/graph"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Exists(nodes []*graph.Node, label string) *graph.Node {
	// Try to find the node within slice nodes
	for i := 0; i < len(nodes); i++ {
		if nodes[i].Label == label {
			return nodes[i]
		}
	}
	return nil
}

func Day16() {
	contents := utils.GetFileContents("day16/example")

	nodes := make([]*graph.Node, 0)

	for _, line := range contents {
		line = strings.ReplaceAll(line, "=", " ")
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, ";", "")

		tokens := strings.Split(line, " ")
		label := tokens[1]
		flowRate, err := strconv.Atoi(tokens[5])
		if err != nil {
			log.Fatal("NaN")
		}

		node := Exists(nodes, label)
		if node == nil {
			// If no nodes are found, create a new one
			node = graph.New(label, flowRate)
			nodes = append(nodes, node)
		} else {
			node.Weight = flowRate
		}

		for i := 10; i < len(tokens); i++ {
			adjNode := Exists(nodes, tokens[i])
			if adjNode == nil {
				adjNode = graph.New(tokens[i], 0)
				nodes = append(nodes, adjNode)
			}
			node.AddChild(adjNode)
		}
	}
	fmt.Println()
}
