package day14

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Point []int

type Path struct {
	Points []*Point
}

func (p *Path) Print() {
	for _, point := range p.Points {
		fmt.Printf("%d,%d ", (*point)[0], (*point)[1])
	}
	fmt.Println()
}

const (
	Rock   = 1
	Air    = 2
	Sand   = 3
	Source = 4
)

type Cave struct {
	Grid    [][]int
	xOffset int
}

func (c *Cave) Print() {
	for i := 0; i < len(c.Grid); i++ {
		for j := 0; j < len(c.Grid[i]); j++ {
			pointType := c.Grid[i][j]

			switch pointType {
			case Rock:
				fmt.Print("#")
			case Air:
				fmt.Print(".")
			case Sand:
				fmt.Print("o")
			case Source:
				fmt.Print("+")
			default:
				log.Fatal("invalid point type")
			}
		}
		fmt.Println()
	}
}

func Day14() {
	contents := utils.GetFileContents("day14/input")

	paths := []Path{}

	// Extract all points and construct a path
	for _, line := range contents {

		points := strings.Split(line, " -> ")

		path := Path{}

		for _, point := range points {
			tokens := strings.Split(point, ",")
			x, err := strconv.Atoi(tokens[0])
			if err != nil {
				log.Fatal("NaN")
			}
			y, err := strconv.Atoi(tokens[1])
			if err != nil {
				log.Fatal("NaN")
			}

			path.Points = append(path.Points, &Point{x, y})
		}

		paths = append(paths, path)
	}

	// Iterate all points and find min./max. of X's, max. of Y's
	minX := math.MaxInt32
	maxX := 500
	maxY := 0

	for _, path := range paths {
		for _, point := range path.Points {
			if (*point)[0] > maxX {
				maxX = (*point)[0]
			}
			if (*point)[0] < minX {
				minX = (*point)[0]
			}
			if (*point)[1] > maxY {
				maxY = (*point)[1]
			}
		}
	}

	// Initialize a cave grid with air
	cave := Cave{xOffset: minX}
	cave.Grid = make([][]int, maxY+1)

	for i := 0; i <= maxY; i++ {
		width := maxX - cave.xOffset
		cave.Grid[i] = make([]int, width+1)

		for j := 0; j <= width; j++ {
			cave.Grid[i][j] = Air
		}
	}

	// Set cave grid source
	source := [2]int{500, 0}
	cave.Grid[source[1]][source[0]-cave.xOffset] = Source

	// Fill up cave grid with rocks
	for _, path := range paths {

		var lastPoint *Point = path.Points[0]

		// Iterate rest of the points
		for _, point := range path.Points {

			x := (*lastPoint)[0]
			y := (*lastPoint)[1]

			// Set first point
			cave.Grid[y][x-cave.xOffset] = Rock

			distX := (*point)[0] - (*lastPoint)[0]
			distY := (*point)[1] - (*lastPoint)[1]

			// Iterate all coordinates between two points
			// Only a single direction will change (X or Y)
			for x != (*point)[0] || y != (*point)[1] {

				if distX == 0 {
					y += 1 * (distY / utils.Abs(distY))
				} else {
					x += 1 * (distX / utils.Abs(distX))
				}

				cave.Grid[y][x-cave.xOffset] = Rock
			}

			lastPoint = point
		}
	}

	// cave.Print()

	// Simulate sand
	tempX := source[0] - cave.xOffset
	tempY := source[1]

	countSettled := 0

	for {
		settled := true
		if tempY+1 >= len(cave.Grid) {
			break
		} else if cave.Grid[tempY+1][tempX] == Air {
			tempY += 1
			settled = false
		} else if tempX-1 < 0 {
			break
		} else if cave.Grid[tempY+1][tempX-1] == Air {
			tempX -= 1
			settled = false
		} else if tempX+1 > len(cave.Grid[tempY]) {
			break
		} else if cave.Grid[tempY+1][tempX+1] == Air {
			tempX += 1
			settled = false
		}

		if settled {
			cave.Grid[tempY][tempX] = Sand
			countSettled += 1

			tempX = source[0] - cave.xOffset
			tempY = source[1]
		}

		// cave.Print()
		// fmt.Println()
	}

	// cave.Print()
	fmt.Println("Result part 1:", countSettled)
}
