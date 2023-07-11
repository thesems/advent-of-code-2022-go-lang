package day8

import (
	"adventofcode/utils"
	"fmt"
)

func Day8() {
	contents := utils.GetFileContents("day8/input")
	width := len(contents[0])
	height := len(contents)

	var grid [][]int = make([][]int, height)
	for i, line := range contents {
		grid[i] = make([]int, width)
		for j, ch := range line {
			grid[i][j] = int(ch)
		}
	}

	visible := 0
	maxScore := 0

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			// externally visible trees
			if i == 0 || j == 0 || i == height-1 || j == width-1 {
				visible += 1
				continue
			}

			visibleFromLeft := true
			leftVisibleCount := 0
			for n := j - 1; n >= 0; n-- {
				if visibleFromLeft = visibleFromLeft && grid[i][j] > grid[i][n]; !visibleFromLeft {
					leftVisibleCount += 1
					break
				} else {
					leftVisibleCount += 1
				}
			}

			visibleFromRight := true
			rightVisibleCount := 0
			for n := j + 1; n < width; n++ {
				if visibleFromRight = visibleFromRight && grid[i][j] > grid[i][n]; !visibleFromRight {
					rightVisibleCount += 1
					break
				} else {
					rightVisibleCount += 1
				}
			}

			visibleFromUp := true
			upVisibleCount := 0
			for n := i - 1; n >= 0; n-- {
				if visibleFromUp = visibleFromUp && grid[i][j] > grid[n][j]; !visibleFromUp {
					upVisibleCount += 1
					break
				} else {
					upVisibleCount += 1
				}
			}

			visibleFromDown := true
			downVisibleCount := 0
			for n := i + 1; n < height; n++ {
				if visibleFromDown = visibleFromDown && grid[i][j] > grid[n][j]; !visibleFromDown {
					downVisibleCount += 1
					break
				} else {
					downVisibleCount += 1
				}
			}

			// Check if the tree is visible from any direction, and if so
			// increase the count of visible trees
			if visibleFromUp || visibleFromLeft || visibleFromDown || visibleFromRight {
				visible += 1
			}

			// Obtain the scenic score
			score := leftVisibleCount * rightVisibleCount * upVisibleCount * downVisibleCount
			if score > maxScore {
				maxScore = score
			}
		}
	}

	fmt.Println("Result part 1: ", visible)
	fmt.Println("Result part 2: ", maxScore)
}
