package day9

import (
	"adventofcode/utils"
	"log"
	"strconv"
	"strings"
)

func Day9() {
	contents := utils.GetFileContents("day9/simple-input")

	for _, line := range contents {
		tokens := strings.Split(line, " ")
		direction := tokens[0]
		steps, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatal("could not convert to int")
		}

		for i := 0; i < steps; i++ {

		}
	}
}
