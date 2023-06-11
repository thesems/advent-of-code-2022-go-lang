package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var moves = make(map[string]map[string]string)
var points_per_shape = make(map[string]int)
var points_per_outcome = make(map[int]int)
var winner = make(map[string]map[string]int)

func initVariables() {
	moves["A"] = make(map[string]string)
	moves["A"]["X"] = "Z"
	moves["A"]["Y"] = "X"
	moves["A"]["Z"] = "Y"

	moves["B"] = make(map[string]string)
	moves["B"]["X"] = "X"
	moves["B"]["Y"] = "Y"
	moves["B"]["Z"] = "Z"

	moves["C"] = make(map[string]string)
	moves["C"]["X"] = "Y"
	moves["C"]["Y"] = "Z"
	moves["C"]["Z"] = "X"

	points_per_shape["X"] = 1
	points_per_shape["Y"] = 2
	points_per_shape["Z"] = 3

	points_per_outcome[-1] = 3
	points_per_outcome[1] = 6
	points_per_outcome[0] = 0

	winner["A"] = make(map[string]int)
	winner["B"] = make(map[string]int)
	winner["C"] = make(map[string]int)

	winner["A"]["X"] = -1
	winner["A"]["Y"] = 1
	winner["A"]["Z"] = 0
	winner["B"]["X"] = 0
	winner["B"]["Y"] = -1
	winner["B"]["Z"] = 1
	winner["C"]["X"] = 1
	winner["C"]["Y"] = 0
	winner["C"]["Z"] = -1
}

func getWinner(hand1 string, hand2 string) int {
	return winner[hand1][hand2]
}

func getNextMove(hand string, symbol string) string {
	return moves[hand][symbol]
}

func Day2() {
	initVariables()

	file, err := os.Open("day2/day2-input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	points := 0

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")

		next_move := getNextMove(tokens[0], tokens[1])
		winner := getWinner(tokens[0], next_move)

		points += points_per_outcome[winner]
		points += points_per_shape[next_move]
	}

	fmt.Println(points)
}
