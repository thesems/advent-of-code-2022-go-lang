package day6

import (
	"adventofcode/utils"
	"fmt"
)

func Day6() {
	data := utils.GetFileContents("day6/input")[0]

	queue_part1 := utils.CreateEvictingQueue(4)
	queue_part2 := utils.CreateEvictingQueue(14)

	for i := 0; i < len(data); i++ {
		queue_part1.Add(int(data[i]))

		if queue_part1.Distinct() && i > 3 {
			fmt.Println("Result part 1: ", i+1)
			break
		}
	}

	for i := 0; i < len(data); i++ {
		queue_part2.Add(int(data[i]))
		if queue_part2.Distinct() && i > 3 {
			fmt.Println("Result part 2: ", i+1)
			break
		}
	}
}
