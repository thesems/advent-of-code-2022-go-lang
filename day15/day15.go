package day15

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

const (
	NORMAL = 0
	SENSOR = 1
	BEACON = 2
)

type Point struct {
	x             int
	y             int
	ptype         int
	ClosestBeacon *Point
}

type Cave struct {
	InvalidPoints utils.Set[Point]
	Points        []*Point
	MinX          int
	MaxX          int
	MinY          int
	MaxY          int
}

func (c *Cave) PointExists(x int, y int, ptype int) bool {
	for _, point := range c.Points {
		if point.x == x && point.y == y && point.ptype == ptype {
			return true
		}
	}
	return false
}

func (c *Cave) Print() {
	fmt.Print("   ")
	for j := c.MinX; j <= c.MaxX; j++ {
		if j%5 == 0 && j >= 10 {
			fmt.Printf("%d", j/10)
		} else {
			fmt.Print(" ")
		}
	}

	fmt.Println()
	fmt.Print("   ")
	for j := c.MinX; j <= c.MaxX; j++ {
		if j%5 == 0 {
			fmt.Printf("%d", j%10)
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	for i := c.MinY; i <= c.MaxY; i++ {
		fmt.Printf("%02d ", i)
		for j := c.MinX; j <= c.MaxX; j++ {
			if c.PointExists(j, i, SENSOR) {
				fmt.Print("S")
			} else if c.PointExists(j, i, BEACON) {
				fmt.Print("B")
			} else if _, ok := (c.InvalidPoints[Point{j, i, NORMAL, nil}]); ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func manhattanDistance(point1 *Point, point2 *Point) int {
	return utils.Abs(point1.x-point2.x) + utils.Abs(point1.y-point2.y)
}

func Day15() {
	contents := utils.GetFileContents("day15/input")

	// rowToCount := 10
	// outOfBounds := 20

	rowToCount := 2000000
	outOfBounds := 4000000

	cave := Cave{utils.Set[Point]{}, make([]*Point, 0), math.MaxInt32, 0, math.MaxInt32, 0}

	// Extract all points and construct a path
	for _, line := range contents {

		tokens := strings.Split(line, " ")
		words := []string{}

		for _, token := range tokens {
			if strings.Contains(token, "=") {
				words = append(words, token)
			}
		}

		x := 0
		y := 0

		for i, word := range words {
			word = strings.ReplaceAll(word, ",", "")
			word = strings.ReplaceAll(word, ":", "")
			tokens := strings.Split(word, "=")

			number, err := strconv.Atoi(tokens[1])
			if err != nil {
				log.Fatal("NaN")
			}

			switch i {
			case 0:
				x = number
			case 1:
				y = number
				if !cave.PointExists(x, y, SENSOR) {
					cave.Points = append(cave.Points, &Point{x, y, SENSOR, nil})

					if x < cave.MinX {
						cave.MinX = x
					}
					if x > cave.MaxX {
						cave.MaxX = x
					}
					if y < cave.MinY {
						cave.MinY = y
					}
					if y > cave.MaxY {
						cave.MaxY = y
					}
				}
			case 2:
				x = number
			case 3:
				y = number
				if !cave.PointExists(x, y, BEACON) {
					cave.Points = append(cave.Points, &Point{x, y, BEACON, nil})

					if x < cave.MinX {
						cave.MinX = x
					}
					if x > cave.MaxX {
						cave.MaxX = x
					}
					if y < cave.MinY {
						cave.MinY = y
					}
					if y > cave.MaxY {
						cave.MaxY = y
					}
				}
			}
		}
	}

	// cave.Print()
	// Iterate points (sensors and beacons)

	for _, point := range cave.Points {
		closestBeaconDist := math.MaxInt
		// Iterate sensors and find closest beacon
		for _, otherPoint := range cave.Points {
			// Ignore if it is not a beacon
			if otherPoint.ptype != BEACON {
				continue
			}

			// Calculate manhanttan distance between sensor and beacon
			dist := manhattanDistance(point, otherPoint)
			if dist < closestBeaconDist {
				closestBeaconDist = dist
				point.ClosestBeacon = otherPoint
			}
		}
	}
	for _, point := range cave.Points {
		// Determine fields without any other beacons inside them
		dist := manhattanDistance(point, point.ClosestBeacon)

		// Iterate points above the sensor
		for y := point.y; y >= point.y-dist; y-- {
			if y != rowToCount {
				continue
			}
			for x := point.x - dist + (point.y - y); x < point.x+dist-(point.y-y)+1; x++ {
				cave.InvalidPoints[Point{x, y, NORMAL, nil}] = struct{}{}
			}
		}

		// Iterate points below the sensor
		for y := point.y + 1; y <= point.y+dist; y++ {
			if y != rowToCount {
				continue
			}
			for x := point.x - dist + (y - point.y); x < point.x+dist-(y-point.y)+1; x++ {
				cave.InvalidPoints[Point{x, y, NORMAL, nil}] = struct{}{}
			}
		}
	}
	count := 0

OUTER:
	for point := range cave.InvalidPoints {
		if point.y != rowToCount {
			continue
		}
		for _, otherPoint := range cave.Points {
			if otherPoint.x == point.x && otherPoint.y == point.y && otherPoint.ptype == BEACON {
				continue OUTER
			}
		}
		count++
	}

	fmt.Println("Results part 1:", count)
	// cave.Print()

	for _, point := range cave.Points {
		dist := manhattanDistance(point, point.ClosestBeacon)

		if FindBeacon(&cave, point, point.ClosestBeacon, dist, outOfBounds) {
			break
		}
	}
}

func FindBeacon(cave *Cave, sensor *Point, closestBeacon *Point, distance int, outOfBounds int) bool {
	width := 0
OUTER_SIDE:
	for y := sensor.y - distance - 1; y <= sensor.y+distance+1; y++ {
		xRight := sensor.x + width
		xLeft := sensor.x - width

		foundDistressRight := true
		foundDistressLeft := true

		if y < 0 || y > outOfBounds {
			continue
		}
		if xLeft < 0 || xLeft > outOfBounds {
			foundDistressLeft = false
		}
		if xRight < 0 || xRight > outOfBounds {
			foundDistressRight = false
		}

		for _, sensor2 := range cave.Points {
			sensor2BeaconDist := manhattanDistance(sensor2, sensor2.ClosestBeacon)

			newDist := manhattanDistance(sensor2, &Point{xRight, y, NORMAL, nil})
			if newDist <= sensor2BeaconDist {
				foundDistressRight = false
			}

			newDist = manhattanDistance(sensor2, &Point{xLeft, y, NORMAL, nil})
			if newDist <= sensor2BeaconDist {
				foundDistressLeft = false
			}
		}

		if !foundDistressRight && !foundDistressLeft {
			if y >= sensor.y {
				width--
			} else {
				width++
			}
			continue OUTER_SIDE
		}

		if foundDistressRight && !cave.PointExists(xRight, y, BEACON) {
			fmt.Println("X,Y=", xRight, y)
			fmt.Println("Right side Results part 2:", xRight*4000000+y)
			return true
		} else if foundDistressLeft && !cave.PointExists(xLeft, y, BEACON) {
			fmt.Println("X,Y=", xLeft, y)
			fmt.Println("Left side Results part 2:", xLeft*4000000+y)
			return true
		}
	}

	return false
}
