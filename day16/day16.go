package day16

import (
	"adventofcode/utils"
	"adventofcode/utils/graph"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func GetNodeIdx(nodes []*graph.Node, label string) (*graph.Node, int) {
	// Try to find the node within slice nodes
	for i := 0; i < len(nodes); i++ {
		if nodes[i].Label == label {
			return nodes[i], i
		}
	}
	return nil, 0
}

type MemoryTable struct {
	memories map[string]int
}

func (mt *MemoryTable) isMemorized(searchNode *graph.Node, minutes int, openValves []string, ele bool) (bool, int) {
	key := mt.encode(searchNode, minutes, openValves, ele)
	reward, ok := mt.memories[key]
	if ok {
		return true, reward
	}
	return false, 0
}

func (mt *MemoryTable) encode(searchNode *graph.Node, minutes int, openValves []string, ele bool) string {
	sort.Strings(openValves)
	key := fmt.Sprintf("%s-%d-%s-%d", searchNode.Label, minutes, strings.Join(openValves, "-"), ele)
	return key
}

func (mt *MemoryTable) memorize(searchNode *graph.Node, minutes int, openValves []string, reward int, ele bool) {
	key := mt.encode(searchNode, minutes, openValves, ele)
	mt.memories[key] = reward
}

func searchPath(memories *MemoryTable, distances graph.DistanceMap, valves []*graph.Node, openValves []*graph.Node, searchNode *graph.Node, minutes int, ele bool) int {
	if minutes <= 0 {
		// Ran out of minutes
		return 0
	}

	if len(openValves) == len(valves)-1 {
		// All valves are open
		return 0
	}

	openValvesCopy := make([]*graph.Node, len(openValves))
	copy(openValvesCopy, openValves)

	maxPressure := 0
	if ele {
		var openValvesStr []string
		for _, valve := range openValvesCopy {
			// sort them by label
			openValvesStr = append(openValvesStr, valve.Label)
			sort.Slice(openValvesStr, func(i, j int) bool {
				return openValvesStr[i] < openValvesStr[j]
			})
		}

		var aaValve *graph.Node
		for _, node := range valves {
			if node.Label == "AA" {
				aaValve = node
				break
			}
		}
		var memorized bool
		memorized, maxPressure = memories.isMemorized(aaValve, 26, openValvesStr, false)
		if !memorized {
			maxPressure = searchPath(memories, distances, valves, openValvesCopy, aaValve, 26, false)
			memories.memorize(aaValve, 26, openValvesStr, maxPressure, false)
		}
	}

	currentPressure := 0

	alreadyOpened := false
	for _, valve := range openValves {
		if searchNode.Label == valve.Label {
			alreadyOpened = true
			break
		}
	}
	if searchNode.Weight == 0 {
		alreadyOpened = true
	}

	if !alreadyOpened {
		openValvesCopy = append(openValvesCopy, searchNode)

		minutes -= 1
		currentPressure = minutes * searchNode.Weight
	}

	// Go deeper with opening the current valve
	destinations := distances[searchNode]
	for child, childWeight := range destinations {
		if child.Label == searchNode.Label {
			continue
		}

		var openValvesStr []string
		for _, valve := range openValvesCopy {
			// sort them by label
			openValvesStr = append(openValvesStr, valve.Label)
			sort.Slice(openValvesStr, func(i, j int) bool {
				return openValvesStr[i] < openValvesStr[j]
			})
		}

		memorized, pressure := memories.isMemorized(child, minutes-childWeight, openValvesStr, ele)

		if !memorized {
			pressure = searchPath(memories, distances, valves, openValvesCopy, child, minutes-childWeight, ele)
			memories.memorize(child, minutes-childWeight, openValvesStr, pressure, ele)
		}

		maxPressure = utils.Max(maxPressure, currentPressure+pressure)
	}

	return maxPressure
}

func compressGraph(nodes []*graph.Node) ([]*graph.Node, graph.DistanceMap) {
	flowingValves := make([]*graph.Node, 0)

	for _, node := range nodes {
		if node.Weight != 0 || node.Label == "AA" {
			flowingValves = append(flowingValves, node)
		}
	}

	distances := graph.FloydWarshall(nodes)

	for _, node1 := range nodes {
		for _, node2 := range nodes {
			if node2.Weight == 0 || distances[node1][node2] == 0 {
				// If the flow rate is 0 or is the same node (hence distance 0), remove it
				delete(distances[node1], node2)
			}
		}
	}

	return flowingValves, distances
}

func Day16(part2 bool) {
	contents := utils.GetFileContents("day16/input")
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

		node, _ := GetNodeIdx(nodes, label)
		if node == nil {
			// If no nodes are found, create a new one
			node = graph.New(label, flowRate)
			nodes = append(nodes, node)
		} else {
			node.Weight = flowRate
		}

		for i := 10; i < len(tokens); i++ {
			adjNode, _ := GetNodeIdx(nodes, tokens[i])
			if adjNode == nil {
				adjNode = graph.New(tokens[i], 0)
				nodes = append(nodes, adjNode)
			}
			node.AddChild(adjNode, 1)
			node.Weight = flowRate
		}
	}
	fmt.Println()

	valves, distances := compressGraph(nodes)

	// fmt.Println("Compressed graph:")
	// PrintGraph(valves)

	var openValves []*graph.Node

	memories := MemoryTable{}
	memories.memories = make(map[string]int)

	var aaValve *graph.Node
	for _, valve := range valves {
		if valve.Label == "AA" {
			aaValve = valve
			break
		}
	}

	// start := time.Now()
	minutes := 30
	if part2 {
		minutes = 26
	}
	reward := searchPath(&memories, distances, valves, openValves, aaValve, minutes, part2)
	fmt.Println("Result:", reward)
	// fmt.Printf("Time: %s", time.Since(start))
}
