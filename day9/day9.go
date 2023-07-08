package day9

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func PrintField(rope [][]int) {

	head := rope[0]
	tail := rope[len(rope)-1]

	minWidth := utils.Min(head[0], tail[0])
	maxWidth := utils.Max(head[0], tail[0])

	minWidth = utils.Min(minWidth, 0)
	maxWidth = utils.Max(maxWidth, 6)

	minHeight := utils.Min(head[1], tail[1])
	maxHeight := utils.Max(head[1], tail[1])

	minHeight = utils.Min(minHeight, 0)
	maxHeight = utils.Max(maxHeight, 5)

	fmt.Println("")
	for i := maxHeight; i >= minHeight; i-- {
		for j := minWidth; j <= maxWidth; j++ {
			ch := "."

			if j == 0 && i == 0 {
				ch = "s"
			}
			if j == tail[0] && i == tail[1] {
				ch = "T"
			}

			if len(rope) > 2 {
				for n := len(rope) - 2; n > 0; n-- {
					if j == rope[n][0] && i == rope[n][1] {
						ch = strconv.Itoa(n)
					}
				}
			}
			if j == head[0] && i == head[1] {
				ch = "H"
			}

			fmt.Print(ch)
		}
		fmt.Println("")
	}
}

type Point struct {
	x int
	y int
}

func Day9() {
	contents := utils.GetFileContents("day9/input")
	enableLogs := false

	// Adjust this parameter for different rope sizes:
	// part 1 = 2, part 2 = 10
	rope := [10][]int{}
	ropeSize := len(rope)
	for i := 0; i < ropeSize; i++ {
		rope[i] = []int{0, 0}
	}

	visited := utils.Set[Point]{}
	visited[Point{0, 0}] = struct{}{}

	if enableLogs {
		fmt.Println("== Initial State ==")
		PrintField(rope[:])
	}

	for _, line := range contents {
		tokens := strings.Split(line, " ")
		direction := tokens[0]
		steps, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatal("could not convert to int")
		}

		if enableLogs {
			fmt.Println()
			fmt.Println("==", direction, steps, "==")
		}

		for i := 0; i < steps; i++ {
			if direction == "R" {
				rope[0][0] += 1
			} else if direction == "L" {
				rope[0][0] -= 1
			} else if direction == "U" {
				rope[0][1] += 1
			} else if direction == "D" {
				rope[0][1] -= 1
			} else {
				log.Fatal("invalid direction")
			}

			for j := 0; j < ropeSize-1; j++ {
				head := rope[j]
				tail := rope[j+1]

				diffX := head[0] - tail[0]
				diffY := head[1] - tail[1]

				if utils.Abs(diffX) > 1 || utils.Abs(diffY) > 1 {
					if diffX == 0 {
						tail[1] += diffY / 2
					} else if diffY == 0 {
						tail[0] += diffX / 2
					} else {
						if diffX > 0 {
							tail[0] += 1
						} else {
							tail[0] -= 1
						}
						if diffY > 0 {
							tail[1] += 1
						} else {
							tail[1] -= 1
						}
					}
				}
			}
			visited[Point{rope[ropeSize-1][0], rope[ropeSize-1][1]}] = struct{}{}

			if enableLogs {
				PrintField(rope[:])
			}
		}
	}

	fmt.Println("Result:", len(visited))
}
