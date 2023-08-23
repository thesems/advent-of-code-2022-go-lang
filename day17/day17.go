package day17

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"math"
	"strings"
)

type Coord struct {
	x, y int
}

type Rock struct {
	coords     []Coord
	minX, maxX int
}

func NewRock(line string) Rock {
	line = strings.Replace(line, "\n", "", -1)
	line = strings.Replace(line, " ", "", -1)

	coords := make([]Coord, 0)

	height := (len(line) / 7) - 1
	for idx, ch := range line {
		if ch != '#' {
			continue
		}

		y := height - idx/7
		x := idx % 7
		coords = append(coords, Coord{x, y})
	}

	rock := Rock{coords, 0, 0}
	rock.HorizontalBounds()
	return rock
}

func CopyRocks(rocks []Rock) []Rock {
	newRocks := make([]Rock, len(rocks))

	for i := 0; i < len(rocks); i++ {
		newRocks[i].coords = make([]Coord, len(rocks[i].coords))
		copy(newRocks[i].coords, rocks[i].coords)
		newRocks[i].maxX = rocks[i].maxX
		newRocks[i].minX = rocks[i].minX
	}

	return newRocks
}

func (r *Rock) HorizontalBounds() {
	r.minX = math.MaxInt32
	r.maxX = 0

	for _, coord := range r.coords {
		if coord.x < r.minX {
			r.minX = coord.x
		}
		if coord.x > r.maxX {
			r.maxX = coord.x
		}
	}
}

func (r *Rock) shiftCoords(dx int, dy int) {
	r.HorizontalBounds()

	if r.minX+dx < 0 || r.maxX+dx > 6 {
		return
	}

	for i := 0; i < len(r.coords); i++ {
		r.coords[i].x += dx
		r.coords[i].y += dy
	}

	r.HorizontalBounds()
}

func (r *Rock) ResetRock(rock Rock) {
	copy(r.coords, rock.coords)
	r.minX = rock.minX
	r.maxX = rock.maxX
}

func (r *Rock) DoesCollide(topLine []Coord, dx int, dy int) bool {
	for j := 0; j < len(r.coords); j++ {
		coord := r.coords[j]
		coord.x += dx
		coord.y += dy

		for i := 0; i < len(topLine); i++ {
			if topLine[i] == coord {
				return true
			}
		}
	}
	return false
}

func Day17(part2 bool) {
	contents := utils.GetFileContents("day17/example")
	jets := make([]int, 0)

	for _, ch := range contents[0] {
		jet := 0
		switch ch {
		case '<':
			jet = -1
		case '>':
			jet = 1
		default:
			log.Fatal("should not happen!")
		}
		jets = append(jets, jet)
	}

	rocks := []Rock{
		NewRock("..####."),
		NewRock(`...#...
                 ..###..
                 ...#...`),
		NewRock(`....#..
                 ....#..
                 ..###..`),
		NewRock(`..#....
                 ..#....
                 ..#....
                 ..#....`),
		NewRock(`..##...
                 ..##...`),
	}

	initRocks := CopyRocks(rocks)

	lastCoordsLimit := 100
	lastCoords := NewRock("#######").coords

	rockIdx := 0
	alt := 0
	jetsIdx := 0
	totalRocks := 0

	rocks[rockIdx].shiftCoords(0, 4)

	for {
		landed := false

		if alt == 0 {
			// move rock by gas jets
			if !rocks[rockIdx].DoesCollide(lastCoords, jets[jetsIdx], 0) {
				rocks[rockIdx].shiftCoords(jets[jetsIdx], 0)
			}
			jetsIdx = (jetsIdx + 1) % len(jets)
			alt = 1
		} else {
			// move down
			if !rocks[rockIdx].DoesCollide(lastCoords, 0, -1) {
				rocks[rockIdx].shiftCoords(0, -1)
			} else {
				landed = true
			}
			alt = 0
		}

		if rocks[rockIdx].coords[0].y < 0 {
			log.Fatal("fell through the floor")
		}

		if landed {
			// Clean topLine and assign new maxY x coords
			lastCoords = append(lastCoords, rocks[rockIdx].coords...)

			idx := 0
			if len(lastCoords) > lastCoordsLimit {
				idx = len(lastCoords) - lastCoordsLimit
			}

			lastCoords = lastCoords[idx:]

			maxY := 0
			for _, coord := range lastCoords {
				if coord.y > maxY {
					maxY = coord.y
				}
			}

			totalRocks += 1
			if (totalRocks == 2022 && !part2) || (totalRocks == 1000000000000 && part2) {
				fmt.Println("Results part 1:", maxY)
				return
			}

            if totalRocks % 1000000 == 0 {
                fmt.Println(totalRocks)
            }

			// reset rock to initial position
			rocks[rockIdx].ResetRock(initRocks[rockIdx])

			rockIdx = (rockIdx + 1) % len(rocks)

			rocks[rockIdx].shiftCoords(0, maxY+4)
		}

	}
}
