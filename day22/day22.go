// Solution translated from hyper-neutrino's python solution.
// https://github.com/hyper-neutrino/advent-of-code

package day22

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func Day22() {
	contents := utils.GetFileContents("day22/input")
	sgrid := make([]string, 0)
	done := false
	sequence := ""
	width := 0

	for _, line := range contents {
		if line == "" {
			done = true
		}
		if done {
			sequence = line
		} else {
			sgrid = append(sgrid, line)
			width = utils.Max(len(line), width)
		}
	}

	r := 0
	c := 0
	dr := 0
	dc := 1

	grid := make([]string, len(sgrid))
	for idx, line := range sgrid {
		acc := ""
		for i := 0; i < width-len(line); i++ {
			acc += " "
		}
		grid[idx] = line + acc
	}

	for grid[r][c] != '.' {
		c += 1
	}

	rg := regexp.MustCompile(`(\d+)([RL]?)`)
	tokens := rg.FindAllStringSubmatch(sequence, -1)

	{
		for _, token := range tokens {
			x, err := strconv.Atoi(token[1])
			if err != nil {
				log.Fatal("NaN")
			}
			y := token[2]

			for i := 0; i < x; i++ {
				nr := r
				nc := c
				for {
					nr = (nr + dr + len(grid)) % len(grid)
					nc = (nc + dc + len(grid[0])) % len(grid[0])
					if grid[nr][nc] != ' ' {
						break
					}
				}
				if grid[nr][nc] == '#' {
					break
				}
				r = nr
				c = nc
			}
			if y == "R" {
				dr, dc = dc, -dr
			} else if y == "L" {
				dr, dc = -dc, dr
			}
		}

		k := 0
		if dr == 0 {
			if dc == 1 {
				k = 0
			} else {
				k = 2
			}
		} else {
			if dr == 1 {
				k = 1
			} else {
				k = 3
			}
		}

        fmt.Println("Results part 1", 1000*(r+1) + 4*(c+1) + k)
	}

	r = 0
	c = 0
	dr = 0
	dc = 1

	for grid[r][c] != '.' {
		c += 1
	}

	{
		for _, token := range tokens {
			x, err := strconv.Atoi(token[1])
			if err != nil {
				log.Fatal("NaN")
			}
			y := token[2]

			for i := 0; i < x; i++ {
                cdr := dr
                cdc := dc
                nr := r + dr
                nc := c + dc

                if nr < 0 && nc >= 50 && nc < 100 && dr == -1 {
                    dr, dc = 0, 1
                    nr, nc = nc + 100, 0
                } else if nc < 0 && nr >= 150 && nr < 200 && dc == -1 {
                    dr, dc = 1, 0
                    nr, nc = 0, nr - 100
                } else if nr < 0 && nc >= 100 && nc < 150 && dr == -1 {
                    nr, nc = 199, nc - 100
                } else if nr >= 200 && nc >= 0 && nc < 50 && dr == 1 {
                    nr, nc = 0, nc + 100
                } else if nc >= 150 && nr >= 0 && nr < 50 && dc == 1 {
                    dc = -1
                    nr, nc = 149 - nr, 99
                } else if nc == 100 && nr >= 100 && nr < 150 && dc == 1 {
                    dc = -1
                    nr, nc = 149 - nr, 149
                } else if nr == 50 && nc >= 100 && nc < 150 && dr == 1 {
                    dr, dc = 0, -1
                    nr, nc = nc - 50, 99
                } else if nc == 100 && nr >= 50 && nr < 100 && dc == 1 {
                    dr, dc = -1, 0
                    nr, nc = 49, nr + 50
                } else if nr == 150 && nc >= 50 && nc < 100 && dr == 1 {
                    dr, dc = 0, -1
                    nr, nc = nc + 100, 49
                } else if nc == 50 && nr >= 150 && nr < 200 && dc == 1 {
                    dr, dc = -1, 0
                    nr, nc = 149, nr - 100
                } else if nr == 99 && nc >= 0 && nc < 50 && dr == -1 {
                    dr, dc = 0, 1
                    nr, nc = nc + 50, 50
                } else if nc == 49 && nr >= 50 && nr < 100 && dc == -1 {
                    dr, dc = 1, 0
                    nr, nc = 100, nr - 50
                } else if nc == 49 && nr >= 0 && nr < 50 && dc == -1 {
                    dc = 1
                    nr, nc = 149 - nr, 0
                } else if nc < 0 && nr >= 100 && nr < 150 && dc == -1 {
                    dc = 1
                    nr, nc = 149 - nr, 50
                }

                if grid[nr][nc] == '#' {
                    dr = cdr
                    dc = cdc
                    break
                }
                r = nr
                c = nc
			}

			if y == "R" {
				dr, dc = dc, -dr
			} else if y == "L" {
				dr, dc = -dc, dr
			}
		}

		k := 0
		if dr == 0 {
			if dc == 1 {
				k = 0
			} else {
				k = 2
			}
		} else {
			if dr == 1 {
				k = 1
			} else {
				k = 3
			}
		}

        fmt.Println("Results part 2:", 1000*(r+1) + 4*(c+1) + k)
	}
}
