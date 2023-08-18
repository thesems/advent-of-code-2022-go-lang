package day20

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"strconv"
)

func normalizeIdx(array []item, index int) int {
	size := len(array) - 1

	idx := index % size
	if idx < 0 {
		idx = idx + size
	}
	return idx
}

type item struct {
	index int
	value int
}

func IndexOf(array []item, value int) int {
	for idx, num := range array {
		if num.value == value {
			return idx
		}
	}
	log.Fatalf("could not find value %d\n", value)
	return 0
}

func InsertAt(items []item, index int, value item) []item {
	items = append(items[:index+1], items[index:]...)
	items[index] = value
	return items
}

func FindByIndex(items []item, index int) int {
	for i := 0; i < len(items); i++ {
		if items[i].index == index {
			return i
		}
	}
	return -1
}

func Day20(part2 bool) {
	contents := utils.GetFileContents("day20/input")
	items := make([]item, len(contents))

	for idx, line := range contents {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal("NaN")
		}

		items[idx] = item{idx, num}
	}

	if !part2 {
		solvePart1(items)
	} else {
		solvePart2(items)
	}
}

func solvePart1(items []item) {
	for index := range items {

		// find index of item
		idx := FindByIndex(items, index)
		item := items[idx]

		// newIdx := (currentIndex + moves) % (len(items) - 1)
		newIdx := idx + item.value
		newIdx = normalizeIdx(items, newIdx)

		// remove item on index
		items = utils.RemoveAt(items, idx)

		// reinsert item at new index
		items = InsertAt(items, newIdx, item)
	}

	nullIdx := IndexOf(items, 0)

	sum := 0
	sum += items[(nullIdx+1000)%len(items)].value
	sum += items[(nullIdx+2000)%len(items)].value
	sum += items[(nullIdx+3000)%len(items)].value

	fmt.Println("Result part 1:", sum)
}

func solvePart2(items []item) {
	for i := 0; i < len(items); i++ {
		items[i].value = items[i].value * 811589153
	}

	for i := 0; i < 10; i++ {
		for initialIndex := range items {

			// find index of item
			currentIndex := FindByIndex(items, initialIndex)
			currentEntry := items[currentIndex]

			// newIdx := (currentIndex + moves) % (len(items) - 1)
			newIdx := currentIndex + currentEntry.value
			newIdx = normalizeIdx(items, newIdx)

			// remove item on index
			items = utils.RemoveAt(items, currentIndex)

			// reinsert item at new index
			items = InsertAt(items, newIdx, currentEntry)
		}
	}

	nullIdx := IndexOf(items, 0)

	sum := 0
	sum += items[(nullIdx+1000)%len(items)].value
	sum += items[(nullIdx+2000)%len(items)].value
	sum += items[(nullIdx+3000)%len(items)].value

	fmt.Println("Result part 2:", sum)
}
