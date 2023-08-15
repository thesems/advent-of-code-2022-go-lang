package day18

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	UNKNOWN = 0
	AIR     = 1
	DROPLET = 2
)

type Cube struct {
	x, y, z    int
}

type World struct {
	minX, maxX int
	minY, maxY int
	minZ, maxZ int
	cubes      []*Cube
}

func (c *Cube) GetNeighbours() []*Cube {
	return []*Cube{
		{c.x + 1, c.y, c.z},
		{c.x, c.y + 1, c.z},
		{c.x, c.y, c.z + 1},
		{c.x - 1, c.y, c.z},
		{c.x, c.y - 1, c.z},
		{c.x, c.y, c.z - 1},
	}
}

func (w *World) Exists(cube *Cube) bool {
	for _, c := range w.cubes {
		for cube.x == c.x && cube.y == c.y && cube.z == c.z {
			return true
		}
	}
	return false
}

func (w *World) InBounds(cube *Cube) bool {
	if cube.x > w.minX - 2 && cube.x < w.maxX + 2 && cube.y > w.minY - 2 && cube.y < w.maxY + 2 && cube.z > w.minZ - 2 && cube.z < w.maxZ + 2 {
		return true
	}
	return false
}

func (w *World) IsAir(cube *Cube) bool {
	for _, item := range w.cubes {
		if item.x == cube.x && item.y == cube.y && item.z == cube.z {
			return false
		}
	}
	return true
}

func CalculateSurfaceArea(cubes []*Cube) int {
	area := 0
	for _, cube1 := range cubes {
		surfaces := 6
		for _, cube2 := range cubes {
			dist := utils.Abs(cube2.x-cube1.x) + utils.Abs(cube2.y-cube1.y) + utils.Abs(cube2.z-cube1.z)

			if dist == 1 {
				surfaces -= 1
			}
		}
		area += surfaces
	}
	return area
}

func NewWorld(cubes []*Cube) *World {
	minX, minY, minZ := int(^uint(0)>>1), int(^uint(0)>>1), int(^uint(0)>>1) // maximum int values
	maxX, maxY, maxZ := -minX, -minY, -minZ                                  // minimum int values

	for _, cube := range cubes {
		minX = utils.Min(minX, cube.x)
		minY = utils.Min(minY, cube.y)
		minZ = utils.Min(minZ, cube.z)
		maxX = utils.Max(maxX, cube.x)
		maxY = utils.Max(maxY, cube.y)
		maxZ = utils.Max(maxZ, cube.z)
	}

	return &World{
		minX - 1, maxX + 1, minY - 1, maxY + 1, minZ - 1, maxZ + 1, cubes,
	}
}

func (w *World) CalculateExternalArea(cubes []*Cube) int {
	startCube := &Cube{w.minX, w.minY, w.minZ}
	queue := []*Cube{startCube}
	visited := map[Cube]struct{}{}
    area := 0
	
    for len(queue) > 0 {
		cube := queue[0]
		queue = queue[1:]

		if _, ok := visited[*cube]; ok {
			continue
		}

        if w.Exists(cube) || !w.InBounds(cube) {
            continue
        }

        // fmt.Printf("x,y,z = %d,%d,%d\n", cube.x, cube.y, cube.z)

        neighbours := cube.GetNeighbours()
        queue = append(queue, neighbours...)
		visited[*cube] = struct{}{}

        for _, c1 := range neighbours {
            if w.Exists(c1) {
                area += 1
            }
        }
	}

    fmt.Println(len(visited))
    fmt.Println(len(cubes))

	return area
}

func Day18(part2 bool) {
	contents := utils.GetFileContents("day18/input")

	cubes := []*Cube{}
	for _, line := range contents {
		tokens := strings.Split(line, ",")

		x, err := strconv.Atoi(tokens[0])
		if err != nil {
			log.Fatal("NaN")
		}

		y, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatal("NaN")
		}

		z, err := strconv.Atoi(tokens[2])
		if err != nil {
			log.Fatal("NaN")
		}

		cubes = append(cubes, &Cube{x, y, z})
	}

	// fmt.Print(cubes)
	world := NewWorld(cubes)

    result := CalculateSurfaceArea(cubes)
	fmt.Println("Result part 1:", result)

	result = world.CalculateExternalArea(cubes)
	fmt.Println("Result part 2:", result)
}
