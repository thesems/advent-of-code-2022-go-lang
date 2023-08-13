package day18

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Cube struct {
	x, y, z    int
	external   bool
	neighbours []*Cube
}

func CalculateSurfaceArea(cubes []*Cube) int {
	area := 0
	for _, cube1 := range cubes {
		surfaces := 6
		for _, cube2 := range cubes {
			dist := utils.Abs(cube2.x-cube1.x) + utils.Abs(cube2.y-cube1.y) + utils.Abs(cube2.z-cube1.z)

			if dist == 1 {
				surfaces -= 1
				cube1.neighbours = append(cube1.neighbours, cube2)
			}
		}
		area += surfaces
	}
	return area
}

func markCubeAsExternal(cubes []*Cube) {
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

	for _, cube := range cubes {
		if cube.x == minX || cube.x == maxX || cube.y == minY || cube.y == maxY || cube.z == minZ || cube.z == maxZ {
			cube.external = true
		}
	}
}

func CalculateNeighbourSurfaces(cubes []*Cube) int {
    startCube := cubes[0]
	area := 0
	queue := []*Cube{startCube}
	visited := map[*Cube]struct{}{}

	for len(queue) > 0 {
		cube := queue[0]
		queue = queue[1:]

		if _, ok := visited[cube]; ok {
			continue
		}

		area += 6 - len(cube.neighbours)
		queue = append(queue, cube.neighbours...)
		visited[cube] = struct{}{}
	}
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

		neighbours := make([]*Cube, 0)
		cubes = append(cubes, &Cube{x, y, z, false, neighbours})
	}

	// fmt.Print(cubes)

	result := CalculateSurfaceArea(cubes)
	fmt.Println("Result part 1:", result)

    result = CalculateNeighbourSurfaces(cubes)
	fmt.Println("Result part 2:", result)
}
