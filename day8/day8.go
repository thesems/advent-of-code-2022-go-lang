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
			if i == 0 || j == 0 || i == height-1 || j == width-1 {
				visible += 1
				continue
			}

			leftSmaller := true
			leftScore := 0
			for n := j - 1; n >= 0; n-- {
				if leftSmaller = leftSmaller && grid[i][j] > grid[i][n]; !leftSmaller {
					leftScore += 1
					break
				} else {
					leftScore += 1
				}
			}

			rightSmaller := true
			rightScore := 0
			for n := j + 1; n < width; n++ {
				if rightSmaller = rightSmaller && grid[i][j] > grid[i][n]; !rightSmaller {
					rightScore += 1
					break
				} else {
					rightScore += 1
				}
			}

			upSmaller := true
			upScore := 0
			for n := i - 1; n >= 0; n-- {
				if upSmaller = upSmaller && grid[i][j] > grid[n][j]; !upSmaller {
					upScore += 1
					break
				} else {
					upScore += 1
				}
			}

			downSmaller := true
			downScore := 0
			for n := i + 1; n < height; n++ {
				if downSmaller = downSmaller && grid[i][j] > grid[n][j]; !downSmaller {
					downScore += 1
					break
				} else {
					downScore += 1
				}
			}
			if upSmaller || leftSmaller || downSmaller || rightSmaller {
				visible += 1
			}

			score := leftScore * rightScore * upScore * downScore
			if score > maxScore {
				maxScore = score
			}
		}
	}

	fmt.Println("Result part 1: ", visible)
	fmt.Println("Result part 2: ", maxScore)
}
