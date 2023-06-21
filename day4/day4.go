package day4

import (
	"adventofcode/utils"
	"log"
	"strconv"
	"strings"
)

func Day4() {
	contents := utils.GetFileContents("day4/input")

	counter := 0
	for _, line := range contents {
		pairs := strings.Split(line, ",")
		pair1 := strings.Split(pairs[0], "-")
		pair2 := strings.Split(pairs[1], "-")

		val1, err1 := strconv.Atoi(pair1[0])
		val2, err2 := strconv.Atoi(pair1[1])
		val3, err3 := strconv.Atoi(pair2[0])
		val4, err4 := strconv.Atoi(pair2[1])

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			log.Fatal("failed to convert string to integer")
		}

		if val1 >= val3 && val2 <= val4 {
			counter += 1
		} else if val3 >= val1 && val4 <= val2 {
			counter += 1
		}
	}

	log.Println(counter)
}
