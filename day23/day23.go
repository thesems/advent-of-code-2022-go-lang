package day23

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"math"
)

type coord struct {
	x, y int
}

func (c *coord) getNeighbours() []coord {
	coords := []coord{
		{c.x, c.y - 1},
		{c.x, c.y + 1},
		{c.x - 1, c.y},
		{c.x + 1, c.y},
	}
	return coords
}

func (c *coord) getAllAdjacentFields() []coord {
    fields := c.getNeighbours()
    fields = append(fields, []coord{{c.x-1, c.y-1}, {c.x+1, c.y-1}, {c.x-1, c.y+1}, {c.x+1, c.y+1}}...)
    return fields
}

func (c *coord) getAdjacentFields(direction int) []coord {
    switch direction {
        case 0:
            return []coord{{c.x-1, c.y-1}, {c.x+1, c.y-1}}
        case 1:
            return []coord{{c.x-1, c.y+1}, {c.x+1, c.y+1}}
        case 2:
            return []coord{{c.x-1, c.y-1}, {c.x-1, c.y+1}}
        case 3:
            return []coord{{c.x+1, c.y-1}, {c.x+1, c.y+1}}
    }
    
    log.Fatal("should not be here. choose a direction 0-3!")
    return []coord{}
}

func Bounds(grid map[coord]struct{}) (int, int, int, int) {

    minX, minY := math.MaxInt16, math.MaxInt16
    maxX, maxY := math.MinInt16, math.MinInt16

    for currCoord := range grid {
        if currCoord.x < minX {
            minX = currCoord.x 
        }
        if currCoord.x > maxX {
            maxX = currCoord.x
        }
        if currCoord.y < minY {
            minY = currCoord.y 
        }
        if currCoord.y > maxY {
            maxY = currCoord.y
        }
    }

    return minX, maxX, minY, maxY
}

func PrintGrid(grid map[coord]struct{}) {

    minX, maxX, minY, maxY := Bounds(grid)

    for y := minY - 1; y < maxY + 2; y++ {
        for x := minX - 1; x < maxX + 2; x++ {
            _, ok := grid[coord{x,y}]
            if ok {
                fmt.Print("#") 
            } else {
                fmt.Print(".") 
            }
        }
        fmt.Println()
    }
}

func Day23(part2 bool) {
	contents := utils.GetFileContents("day23/input")

	grid := make(map[coord]struct{})

	for y, line := range contents {
		for x, point := range line {
			if point == '#' {
				grid[coord{x, y}] = struct{}{}
			}
		}
	}

    fmt.Println("== Initial State ==")
    PrintGrid(grid)
    fmt.Println()

	round := 0
	directionOrder := 0
	for round != 10 || part2 {
		// destination, sources
		proposedMoves := make(map[coord][]coord)
        noMoves := true

		for currCoord := range grid {
			neighbours := currCoord.getNeighbours()

            checked := 0
			for i := directionOrder; ; i++ {
				if i == len(neighbours) {
					i = 0
				}

                // Check if there are any other elves around it
                adjacentFields := currCoord.getAllAdjacentFields()
                alone := true
                for _, field := range adjacentFields {
                    _, ok := grid[field]
                    if ok {
                        alone = false
                        break
                    }
                }

                if alone {
                    break
                }

                noMoves = false

                // Get direct neighbour for the specific  direction
				neighbour := neighbours[i]

                // Get the neighbour's adjacent fields
                diagonals := currCoord.getAdjacentFields(i)

				_, ok := grid[neighbour]
                _, ok2 := grid[diagonals[0]]
                _, ok3 := grid[diagonals[1]]

				if !ok && !ok2 && !ok3 {
					// grid field is free. Check direction and direction adjecent fields!
					_, ok := proposedMoves[neighbour]
					if !ok {
						proposedMoves[neighbour] = make([]coord, 0)
					}
					proposedMoves[neighbour] = append(proposedMoves[neighbour], currCoord)
					break
				}

                checked += 1
                if checked == 4 {
                    break
                }
			}
		}

		for dest, sources := range proposedMoves {
			if len(sources) > 1 {
				continue
			}

			delete(grid, sources[0])
			grid[dest] = struct{}{}
		}

		round += 1
		directionOrder += 1
		if directionOrder == 4 {
			directionOrder = 0
		}

        // fmt.Println()
        fmt.Printf("== End of Round %d ==\n", round)
        // PrintGrid(grid)

        if noMoves {
            fmt.Println("Results part 2:", round)
            break
        }
	}

	freeTiles := make(map[coord]struct{})
    minX, maxX, minY, maxY := Bounds(grid)

    for y := minY; y <= maxY; y++ {
        for x := minX; x <= maxX; x++ {
            currCoord := coord{x,y}
            _, ok := grid[currCoord]
            if !ok && x >= minX && x <= maxX && y >= minY && y <= maxY {
                freeTiles[currCoord] = struct{}{}
            }
        }
    }

	fmt.Println("Results part 1:", len(freeTiles))
}
