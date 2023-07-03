package day7

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Day7() {
	contents := utils.GetFileContents("day7/input")

	rootNode := utils.NewNode(nil, 0, false, "/")
	var activeNode *utils.Node = rootNode

	for idx, line := range contents {
		if idx == 0 {
			// skip cd /
			continue
		}

		tokens := strings.Split(line, " ")

		// command
		if tokens[0] == "$" {
			if tokens[1] == "cd" {
				if tokens[2] == ".." {
					activeNode = activeNode.Parent()
				} else {
					child, err := activeNode.ExistChild(tokens[2])
					if err != nil {
						newNode := utils.NewNode(activeNode, 0, false, tokens[2])
						activeNode.AddNode(newNode)
						activeNode = newNode
					} else {
						activeNode = child
					}
				}
			} else {
				// ls - no action necessary
			}
		} else if tokens[0] == "dir" {
			_, err := activeNode.ExistChild(tokens[1])

			if err != nil {
				newNode := utils.NewNode(activeNode, 0, false, tokens[1])
				activeNode.AddNode(newNode)
			}
		} else {
			// file
			_, err := activeNode.ExistChild(tokens[1])

			if err != nil {
				size, err := strconv.Atoi(tokens[0])
				if err != nil {
					log.Fatal("expected int size, but got " + tokens[0])
				}
				newNode := utils.NewNode(activeNode, size, true, tokens[1])
				activeNode.AddNode(newNode)
			}
		}
	}

	rootNode.CalculateSize()
	fmt.Println("Result Part 1: ", rootNode.CountDirSize(100000))

	sizeNeeded := 30000000 - (70000000 - rootNode.Size())
	var deleteDir *utils.Node = nil
	deleteDir = rootNode.FindSmallestDirAboveLimit(deleteDir, sizeNeeded)
	if deleteDir == nil {
		log.Fatal("no dir found to delete")
	}
	fmt.Println("Result Part 2: ", deleteDir.Name(), deleteDir.Size())
}
