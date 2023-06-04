package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func Day1() {
	file, err := os.Open("days/day1-input.txt")

	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewScanner(file)
	reader.Split(bufio.ScanLines)

	maxCalories := []int{0, 0, 0}
	currentCalories := 0

	for reader.Scan() {
		line := reader.Text()

		if line == "" {

			sort.Ints(maxCalories)
			if currentCalories > maxCalories[0] {
				maxCalories[0] = currentCalories
			}
			currentCalories = 0
			continue
		}

		calories, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal("File has improper structure on line: ", line)
			log.Fatal("Error: ", err)
		}

		currentCalories += calories
	}

	fmt.Println(maxCalories[0] + maxCalories[1] + maxCalories[2])
}
