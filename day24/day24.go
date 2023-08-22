// Efficient python solution translated from hyper-neutrino.
// URL: https://github.com/hyper-neutrino/advent-of-code
// Contains multiple optmizations such as: memoization, LCM to cut possible states, shifting player instead of blizzards.

package day24

import (
	"adventofcode/utils"
	"adventofcode/utils/deque"
	"fmt"
	"log"
)

const (
	NORTH = '^'
	EAST  = '<'
	SOUTH = 'v'
	WEST  = '>'
)

type Coord struct {
	x, y int
}

func (c *Coord) getNeighbours() []Coord {
	coords := []Coord{
		{c.x, c.y - 1},
		{c.x, c.y + 1},
		{c.x - 1, c.y},
		{c.x + 1, c.y},
	}
	return coords
}

type State struct {
	time  int
	coord Coord
	stage int
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func Mod(a, b int) int {
	remainder := a % b
	if remainder < 0 {
		remainder += b
	}
	return remainder
}

func Day24() {
	contents := utils.GetFileContents("day24/input")
	blizzards := make(map[rune][]Coord)

	maxX := len(contents[0][1:]) - 1
	maxY := len(contents[1:]) - 1

	for y, line := range contents[1:] {
		for x, ch := range line[1:] {
			switch ch {
			case NORTH, EAST, SOUTH, WEST:
				coord := Coord{x, y}
				blizz, ok := blizzards[ch]
				if !ok {
					blizzards[ch] = make([]Coord, 0)
				}
				blizz = append(blizz, coord)
				blizzards[ch] = blizz
			}
		}
	}

	solve1(blizzards, maxX, maxY)
    solve2(blizzards, maxX, maxY)
}

func solve1(blizzards map[rune][]Coord, maxX int, maxY int) {
	queue := deque.New[State]()
	queue.Push(State{0, Coord{0, -1}, 0})

	seen := make(map[State]struct{})
	target := Coord{maxX - 1, maxY}

	lcm := (maxY * maxX) / GCD(maxY, maxX)

	for !queue.IsEmpty() {
		state, err := queue.Front()
		if err != nil {
			return
		}

		// fmt.Println(state)

		time := state.time + 1

		neighbours := state.coord.getNeighbours()
		neighbours = append(neighbours, state.coord)

		for _, newCoord := range neighbours {
			if newCoord == target {
				fmt.Println("Results part 1:", time)
				return
			}

			if (newCoord.x < 0 || newCoord.y < 0 || newCoord.x >= maxX || newCoord.y >= maxY) && (newCoord != Coord{0, -1}) {
				continue
			}

			fail := false

			if (newCoord != Coord{0, -1}) {
				blizzDirs := [][]int{
					{NORTH, 0, -1}, {EAST, -1, 0}, {SOUTH, 0, 1}, {WEST, 1, 0},
				}
				for _, dir := range blizzDirs {
					position := Coord{Mod(newCoord.x-dir[1]*time, maxX), Mod((newCoord.y - dir[2]*time), maxY)}

					dirBlizzards, ok := blizzards[rune(dir[0])]
					if !ok {
						log.Fatal("should not happen!")
					}

					for _, blizz := range dirBlizzards {
						if blizz.x == position.x && blizz.y == position.y {
							fail = true
							continue
						}
					}
				}
			}

			if !fail {
				key := State{time % lcm, newCoord, 0}
				_, ok := seen[key]
				if ok {
					continue
				}
				seen[key] = struct{}{}
				queue.Push(State{time, newCoord, 0})
			}
		}
	}
}

func solve2(blizzards map[rune][]Coord, maxX int, maxY int) {
	queue := deque.New[State]()
	queue.Push(State{0, Coord{0, -1}, 0})

	seen := make(map[State]struct{})
	target := []Coord{{maxX - 1, maxY}, {0, -1}}

	lcm := (maxY * maxX) / GCD(maxY, maxX)

	for !queue.IsEmpty() {
		state, err := queue.Front()
		if err != nil {
			return
		}

		// fmt.Println(state)

		time := state.time + 1

		neighbours := state.coord.getNeighbours()
		neighbours = append(neighbours, state.coord)

		for _, newCoord := range neighbours {
            nstage := state.stage

			if newCoord == target[state.stage % 2] {
                if state.stage == 2 {
				    fmt.Println("Results part 2:", time)
				    return
                }
                nstage += 1
			}

            foundTarget := false
            for _, t := range target {
                if newCoord == t {
                    foundTarget = true
                    break
                }
            }

			if (newCoord.x < 0 || newCoord.y < 0 || newCoord.x >= maxX || newCoord.y >= maxY) && !foundTarget {
				continue
			}

			fail := false

			if !foundTarget {
				blizzDirs := [][]int{
					{NORTH, 0, -1}, {EAST, -1, 0}, {SOUTH, 0, 1}, {WEST, 1, 0},
				}
				for _, dir := range blizzDirs {
					position := Coord{Mod(newCoord.x-dir[1]*time, maxX), Mod((newCoord.y - dir[2]*time), maxY)}

					dirBlizzards, ok := blizzards[rune(dir[0])]
					if !ok {
						log.Fatal("should not happen!")
					}

					for _, blizz := range dirBlizzards {
						if blizz.x == position.x && blizz.y == position.y {
							fail = true
							continue
						}
					}
				}
			}

			if !fail {
				key := State{time % lcm, newCoord, nstage}
				_, ok := seen[key]
				if ok {
					continue
				}
				seen[key] = struct{}{}
				queue.Push(State{time, newCoord, nstage})
			}
		}
	}
}
