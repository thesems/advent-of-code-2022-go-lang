package day21

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Monkey struct {
	monkey1 *Monkey
	monkey2 *Monkey
	name    string
	op      string
	num     int
}

func solve1(monkey *Monkey) int {
	if monkey.monkey1 != nil {
		solve1(monkey.monkey1)
	}
	if monkey.monkey2 != nil {
		solve1(monkey.monkey2)
	}

	if monkey.op == "" {
		return monkey.num
	}

	result := 0

	switch monkey.op {
	case "+":
		result = monkey.monkey1.num + monkey.monkey2.num
	case "-":
		result = monkey.monkey1.num - monkey.monkey2.num
	case "*":
		result = monkey.monkey1.num * monkey.monkey2.num
	case "/":
		result = monkey.monkey1.num / monkey.monkey2.num
	}

	monkey.num = result
	return result
}

func getNonHumanPath(monkey *Monkey) (int, bool) {
	if monkey.name == "humn" {
		return -1, true
	}

	if monkey.monkey1 != nil {
		_, found := getNonHumanPath(monkey.monkey1)
		if found {
			return 0, true
		}
	}

	if monkey.monkey2 != nil {
		_, found := getNonHumanPath(monkey.monkey2)
		if found {
			return 1, true
		}
	}

	return -1, false
}

func solve2(rootMonkey *Monkey, human *Monkey, desc bool) {
	pathChoice, _ := getNonHumanPath(rootMonkey)

	var humanPath, otherPath *Monkey
	if pathChoice == 0 {
		humanPath = rootMonkey.monkey1
		otherPath = rootMonkey.monkey2
	} else {
		humanPath = rootMonkey.monkey2
		otherPath = rootMonkey.monkey1
	}

	needValue := solve1(otherPath)

	min := 0
	max := int(math.Pow(2, 50))
	result := 0
	guess := 0

	for result != needValue {
		guess = (max + min) / 2
		human.num = guess
		result = solve1(humanPath)

		if result > needValue {
			if desc {
				min = guess
			} else {
				max = guess
			}
		} else if result < needValue {
			if desc {
				max = guess
			} else {
				min = guess
			}
		}

		if max-min == 1 {
			break
		}
	}
}

func Day21() {
	dataset := "input"
	contents := utils.GetFileContents("day21/" + dataset)
	monkeyToExpr := make(map[string]*Monkey)

	for _, line := range contents {
		tokens := strings.Split(line, ":")
		name := tokens[0]

		monkeyToExpr[name] = &Monkey{nil, nil, name, "", 0}
	}

	for _, line := range contents {
		line = strings.ReplaceAll(line, " ", "")
		tokens := strings.Split(line, ":")
		name := tokens[0]

		expr, ok := monkeyToExpr[name]
		if !ok {
			log.Fatal("should not happen!")
		}

		operations := []string{"+", "-", "*", "/"}
		operation := ""
		subtokens := make([]string, 0)

		for _, op := range operations {
			res := strings.Index(tokens[1], op)
			if res != -1 {
				subtokens = strings.Split(tokens[1], op)
				operation = op
				break
			}
		}

		if operation == "" {
			num, err := strconv.Atoi(tokens[1])
			if err != nil {
				log.Fatal("should not happen!")
			}
			expr.num = num
		} else {
			expr1, ok := monkeyToExpr[subtokens[0]]
			if !ok {
				log.Fatal("should not happen!")
			}

			expr2, ok := monkeyToExpr[subtokens[1]]
			if !ok {
				log.Fatal("should not happen!")
			}

			expr.monkey1 = expr1
			expr.monkey2 = expr2
			expr.op = operation
		}
	}

	rootMonkey, ok := monkeyToExpr["root"]
	if !ok {
		log.Fatal("should not happen! no root found!")
	}

	human, ok := monkeyToExpr["humn"]
	if !ok {
		log.Fatal("should not happen! no humn found!")
	}

	result := solve1(rootMonkey)
	fmt.Println("Results part 1:", result)

	human.num = 0

	solve2(rootMonkey, human, "example" != dataset)
	fmt.Println("Results part 2", human.num)
}
