package day10

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Day10() {
	contents := utils.GetFileContents("day10/input")

	x := 1
	cycle := 0
	signalStrength := 0
	lineSize := 40
	lineCnt := 0

	for _, line := range contents {
		tokens := strings.Split(line, " ")
		op := tokens[0]

		doCycles := 1
		if op == "addx" {
			doCycles = 2
		}

		for i := 0; i < doCycles; i++ {
			if lineCnt == x || lineCnt-1 == x || lineCnt+1 == x {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			if lineCnt+1 == lineSize {
				fmt.Println()
				lineCnt = 0
			} else {
				lineCnt += 1
			}

			cycle += 1
			if cycle == 20 || ((cycle+20)%40 == 0 && cycle <= 220) {
				signalStrength += cycle * x
			}

			if op == "addx" && i == 1 {
				num, err := strconv.Atoi(tokens[1])
				if err != nil {
					log.Fatal("could not convert string to number")
				}
				x += num
			}
		}

	}

	fmt.Println("")
	fmt.Println("Result part 1: ", signalStrength)
	fmt.Println("Tip: squint to see the letters easier!")
}
