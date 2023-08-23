package day17

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"math"
	"sort"
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

func (r *Rock) DoesCollide(lastRocks map[int][]Coord, dx int, dy int) bool {
	for j := 0; j < len(r.coords); j++ {
		coord := r.coords[j]
		coord.x += dx
		coord.y += dy

		coords := lastRocks[coord.x]
		for i := 0; i < len(coords); i++ {
			if coords[i] == coord {
				return true
			}
		}
	}
	return false
}

func Day17(part2 bool) {
	contents := utils.GetFileContents("day17/input")
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

	// naive(jets, rocks)
	optimized(jets, rocks, part2)
}

// func naive(jets []int, rocks []Rock) {
// 	initRocks := CopyRocks(rocks)
//
// 	lastCoordsLimit := 100
// 	lastCoords := NewRock("#######").coords
//
// 	rockIdx := 0
// 	alt := 0
// 	jetsIdx := 0
// 	totalRocks := 0
//
// 	rocks[rockIdx].shiftCoords(0, 4)
//
// 	for {
// 		landed := false
//
// 		if alt == 0 {
// 			// move rock by gas jets
// 			if !rocks[rockIdx].DoesCollide(lastCoords, jets[jetsIdx], 0) {
// 				rocks[rockIdx].shiftCoords(jets[jetsIdx], 0)
// 			}
// 			jetsIdx = (jetsIdx + 1) % len(jets)
// 			alt = 1
// 		} else {
// 			// move down
// 			if !rocks[rockIdx].DoesCollide(lastCoords, 0, -1) {
// 				rocks[rockIdx].shiftCoords(0, -1)
// 			} else {
// 				landed = true
// 			}
// 			alt = 0
// 		}
//
// 		if rocks[rockIdx].coords[0].y < 0 {
// 			log.Fatal("fell through the floor")
// 		}
//
// 		if landed {
// 			// Clean topLine and assign new maxY x coords
// 			lastCoords = append(lastCoords, rocks[rockIdx].coords...)
//
// 			idx := 0
// 			if len(lastCoords) > lastCoordsLimit {
// 				idx = len(lastCoords) - lastCoordsLimit
// 			}
//
// 			lastCoords = lastCoords[idx:]
//
// 			maxY := 0
// 			for _, coord := range lastCoords {
// 				if coord.y > maxY {
// 					maxY = coord.y
// 				}
// 			}
//
// 			totalRocks += 1
// 			if totalRocks == 2022 {
// 				fmt.Println("Results part 1:", maxY)
// 				return
// 			}
//
// 			// reset rock to initial position
// 			rocks[rockIdx].ResetRock(initRocks[rockIdx])
//
// 			rockIdx = (rockIdx + 1) % len(rocks)
//
// 			rocks[rockIdx].shiftCoords(0, maxY+4)
// 		}
//
// 	}
// }

type State struct {
	coords  map[int][]Coord
	rockIdx int
	jetsIdx int
}

func (s *State) toString() string {
    allCoords := make([]Coord, 0)

    for _, coords := range s.coords {
        for _, coord := range coords {
            allCoords = append(allCoords, coord)
        }
    }

	sort.Slice(allCoords, func(i, j int) bool {
		return allCoords[i].x < allCoords[j].x
	})

	maxY := 0
	for _, coord := range allCoords {
		if coord.y > maxY {
			maxY = coord.y
		}
	}

	coords := ""
	for _, coord := range allCoords {
		coords += fmt.Sprintf("%d,%d,", coord.x, maxY-coord.y)
	}

	str := fmt.Sprintf("%s,%d,%d", coords[:len(coords)-1], s.rockIdx, s.jetsIdx)
	return str
}

func optimized(jets []int, rocks []Rock, part2 bool) {
	initRocks := CopyRocks(rocks)

	lastCoordsLimit := 25
	lastCoords := map[int][]Coord{
		0: {{0, 0}},
		1: {{1, 0}},
		2: {{2, 0}},
		3: {{3, 0}},
		4: {{4, 0}},
		5: {{5, 0}},
		6: {{6, 0}},
	}

	seen := make(map[string][2]int)

	rockIdx := 0
	jetsIdx := 0
	totalRocks := 0

	rocks[rockIdx].shiftCoords(0, 4)

	trillion := 1000000000000
	offset := 0

	for {
		// move rock by gas jets
		if !rocks[rockIdx].DoesCollide(lastCoords, jets[jetsIdx], 0) {
			rocks[rockIdx].shiftCoords(jets[jetsIdx], 0)
		}
		jetsIdx = (jetsIdx + 1) % len(jets)

		// move down
		if !rocks[rockIdx].DoesCollide(lastCoords, 0, -1) {
			rocks[rockIdx].shiftCoords(0, -1)
		} else {

			// Add rocks to lastRocks
			for _, coord := range rocks[rockIdx].coords {
				value, ok := lastCoords[coord.x]
				if !ok {
					log.Fatal("should not happen!")
				}
				lastCoords[coord.x] = append(value, coord)

				if len(lastCoords[coord.x]) > lastCoordsLimit {
					idx := len(lastCoords[coord.x]) - lastCoordsLimit
					lastCoords[coord.x] = lastCoords[coord.x][idx:]
				}
			}

			maxY := 0
			for _, coords := range lastCoords {
				for _, coord := range coords {
					if coord.y > maxY {
						maxY = coord.y
					}
				}
			}

			totalRocks += 1
			if (!part2 && totalRocks == 2022) || totalRocks == trillion {
				fmt.Println("Results part 1:", maxY+offset)
				return
			}

			// reset rock to initial position
			rocks[rockIdx].ResetRock(initRocks[rockIdx])

			// advance rock idx and shift coordinate to start position
			rockIdx = (rockIdx + 1) % len(rocks)
			rocks[rockIdx].shiftCoords(0, maxY+4)

			state := State{lastCoords, rockIdx, jetsIdx}
			key := state.toString()
			value, ok := seen[key]
			if ok {
				rem := trillion - totalRocks
				rep := rem / (totalRocks - value[0])
				offset = rep * (maxY - value[1])
				totalRocks += rep * (totalRocks - value[0])
				clear(seen)
			}

			seen[key] = [2]int{totalRocks, maxY}
		}

		if rocks[rockIdx].coords[0].y < 0 {
			log.Fatal("fell through the floor")
		}
	}
}
