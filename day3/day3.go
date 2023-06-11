package day3

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
	"unicode"
)

// short strings: brutorce o(n^2) 4-5x faster than o(2n)
// go testing lib

func findCommonCharFromTwoStringsBruteForce(comp1 string, comp2 string) (rune, error) {
	defer duration(track("bruteforce 2"))
	for _, char1 := range comp1 {
		for _, char2 := range comp2 {
			if char1 == char2 {
				result := char1
				if unicode.IsLower(char1) {
					result = char1 - 96
				} else {
					result = char1 - 38
				}
				return result, nil
			}
		}
	}
	return 0, errors.New("no common character")
}

func findCommonCharFromThreeStringsBruteForce(comp1 string, comp2 string, comp3 string) (rune, error) {
	defer duration(track("bruteforce 3"))
	for _, char1 := range comp1 {
		for _, char2 := range comp2 {
			for _, char3 := range comp3 {
				if char1 == char2 && char2 == char3 {
					result := char1
					if unicode.IsLower(char1) {
						result = char1 - 96
					} else {
						result = char1 - 38
					}
					return result, nil
				}
			}
		}
	}
	return 0, errors.New("no common character")
}

func findCommonCharFromTwoStrings(comp1 string, comp2 string) (rune, error) {
	helper := make(map[int]bool)
	defer duration(track("optimized 2"))
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
	return 0, errors.New("no common character")
}

func findCommonCharFromThreeStrings(comp1 string, comp2 string, comp3 string) (rune, error) {
	helper := make(map[int]bool)
	defer duration(track("optimized 3"))
	for _, ch := range comp1 {
		helper[int(ch)] = true
	}
	for _, ch := range comp2 {
		_, ok := helper[int(ch)]

		if !ok {
			helper[int(ch)] = false
		}
	}
	for _, ch := range comp3 {
		res, ok := helper[int(ch)]

		if ok && res {
			result := ch
			if unicode.IsLower(ch) {
				result = ch - 96
			} else {
				result = ch - 38
			}
			return result, nil
		}
	}
	return 0, errors.New("no common character")
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
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

		// findCommonCharFromTwoStringsBruteForce(comp1, comp2)
		common_char, err := findCommonCharFromTwoStrings(comp1, comp2)
		if err != nil {
			log.Fatal("not common characters")
		}

		total_priorities += int(common_char)

		if len(group) == 3 {
			findCommonCharFromThreeStrings(
				group[0],
				group[1],
				group[2],
			)
			common_char, err := findCommonCharFromThreeStringsBruteForce(
				group[0],
				group[1],
				group[2],
			)

			if err != nil {
				log.Fatal("No common character among three groups.")
			}

			badges_priorities += int(common_char)
			group = group[:0]
		}
	}

	fmt.Println("Part 1: ", total_priorities)
	fmt.Println("Part 1: ", badges_priorities)
}
