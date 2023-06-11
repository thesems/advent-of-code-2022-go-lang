package day3

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"unicode"
)

func findCommonCharFromTwoStrings(comp1 string, comp2 string) (rune, error) {
	helper := make(map[int]bool)
	for _, ch := range comp1 {
		helper[int(ch)] = true
	}
	for _, ch := range comp2 {
		_, ok := helper[int(ch)]

		if ok {
			result := ch
			if unicode.IsLower(ch) {
				result = ch - 96
			} else {
				result = ch - 38
			}
			return result, nil
		}
	}
	return 0, errors.New("no common characters found")
}

func findCommonCharFromThreeStrings(comp1 string, comp2 string, comp3 string) (rune, error) {
	helper1 := make(map[int]bool)
	helper2 := make(map[int]bool)

	for _, ch := range comp1 {
		helper1[int(ch)] = true
	}
	for _, ch := range comp2 {
		helper2[int(ch)] = true
	}
	for _, ch := range comp3 {
		_, is_in_comp1 := helper1[int(ch)]
		_, is_in_comp2 := helper2[int(ch)]

		if is_in_comp1 && is_in_comp2 {
			result := ch
			if unicode.IsLower(ch) {
				result = ch - 96
			} else {
				result = ch - 38
			}
			return result, nil
		}
	}
	return 0, errors.New("no common characters found")
}

func Day3() {
	file, err := os.Open("day3/day3-input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	total_priorities := 0
	badges_priorities := 0
	group := make([]string, 0, 3)

	for scanner.Scan() {
		line := scanner.Text()
		group = append(group, line)

		half_size := len(line) / 2
		comp1 := line[:half_size]
		comp2 := line[half_size:]

		common_char, err := findCommonCharFromTwoStrings(comp1, comp2)
		if err != nil {
			log.Fatal("No common characters found!")
		}

		total_priorities += int(common_char)

		if len(group) == 3 {
			common_char, err := findCommonCharFromThreeStrings(
				group[0],
				group[1],
				group[2],
			)

			if err != nil {
				log.Fatal("No common characters found!")
			}

			badges_priorities += int(common_char)
			group = group[:0]
		}
	}

	fmt.Println("Part 1: ", total_priorities)
	fmt.Println("Part 2: ", badges_priorities)
}
