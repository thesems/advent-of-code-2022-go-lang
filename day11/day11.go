package day11

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Operation interface {
	do(value int) int
	print(worryLevel int)
}

type AddOperation struct {
	cofactor int
}

func (op *AddOperation) do(value int) int {
	return op.cofactor + value
}
func (op *AddOperation) print(worryLevel int) {
	fmt.Printf("    Worry level increases by %d to %d.\n", op.cofactor, worryLevel)
}

type MultiplyOperation struct {
	cofactor int
}

func (op *MultiplyOperation) do(value int) int {
	return op.cofactor * value
}
func (op *MultiplyOperation) print(worryLevel int) {
	fmt.Printf("    Worry level is multiplied by %d to %d.\n", op.cofactor, worryLevel)
}

type SquareOperation struct{}

func (op *SquareOperation) do(value int) int {
	return value * value
}
func (op *SquareOperation) print(worryLevel int) {
	fmt.Printf("    Worry level is multiplied by itself to %d.\n", worryLevel)
}

type Monkey struct {
	items                []int
	operation            Operation
	testDivision         int
	testMonkeyIdxIfTrue  int
	testMonkeyIdxIfFalse int
	inspectionCount      int
}

func (m *Monkey) getCommonDivisor() int {
	return 0
}

func (m *Monkey) doOperation() {
	m.items[0] = m.operation.do(m.items[0])
}
func (m *Monkey) checkDivisionBy() bool {
	return m.items[0]%m.testDivision == 0
}
func (m *Monkey) trisectWorry() {
	m.items[0] = m.items[0] / 3
}
func (m *Monkey) throwItem() {
	m.items = utils.DeleteElement(m.items, 0)
}

func parseMonkeys(contents []string) []Monkey {
	monkeys := make([]Monkey, 0)
	reNum := regexp.MustCompile("[0-9]+")
	reOpVar := regexp.MustCompile(`old [\*|+] ([0-9]+|old)`)

	for i := 0; i < len(contents); i++ {
		line := strings.TrimSpace(contents[i])

		if strings.HasPrefix(line, "Monkey") {
			_, err := strconv.Atoi(reNum.FindAllString(line, -1)[0])
			if err != nil {
				log.Fatal("could not convert string to int")
			}

			line = strings.TrimSpace(contents[i+1])
			line = line[strings.Index(line, ":")+1:]
			items := strings.Split(line, ",")
			monkeyItems := make([]int, 0, len(items))
			for _, item := range items {
				item = strings.TrimSpace(item)
				itemInt, err := strconv.Atoi(item)
				if err != nil {
					log.Fatal("could not convert string to int")
				}

				monkeyItems = append(monkeyItems, itemInt)
			}

			line = contents[i+2]
			tokens := strings.Split(reOpVar.FindAllString(line, -1)[0], " ")
			operationSign := tokens[1]
			operationVarString := tokens[2]
			operationVar, _ := strconv.Atoi(operationVarString)

			var operation Operation
			if operationSign == "+" {
				operation = &AddOperation{cofactor: operationVar}
			} else if operationSign == "*" {
				operation = &MultiplyOperation{cofactor: operationVar}
			}

			if operationVarString == "old" {
				operation = &SquareOperation{}
			}

			line = contents[i+3]
			testDivision, err := strconv.Atoi(reNum.FindAllString(line, -1)[0])
			if err != nil {
				log.Fatal("could not convert string to int")
			}

			line = contents[i+4]
			testMonkeyIdxIfTrue, err := strconv.Atoi(reNum.FindAllString(line, -1)[0])
			if err != nil {
				log.Fatal("could not convert string to int")
			}

			line = contents[i+5]
			testMonkeyIdxIfFalse, err := strconv.Atoi(reNum.FindAllString(line, -1)[0])
			if err != nil {
				log.Fatal("could not convert string to int")
			}

			monkey := Monkey{monkeyItems, operation, testDivision, testMonkeyIdxIfTrue, testMonkeyIdxIfFalse, 0}

			i = i + 6
			monkeys = append(monkeys, monkey)
		}
	}
	return monkeys
}

func Day11() {
	contents := utils.GetFileContents("day11/example1")
	monkeys := parseMonkeys(contents)

	gameLoop := true
	round := 0
	partTwo := true

	divisorProduct := 1
	for i := 0; i < len(monkeys); i++ {
		divisorProduct *= monkeys[i].testDivision
	}

	for gameLoop {
		for i := 0; i < len(monkeys); i++ {
			if !partTwo {
				fmt.Printf("Monkey %d:\n", i)
			}
			// Iterate all items of the current monkey
			for 0 < len(monkeys[i].items) {
				if !partTwo {
					fmt.Printf("  Monkey inspects an item with a worry level of %d.\n", monkeys[i].items[0])
				}

				monkeys[i].items[0] = monkeys[i].items[0] % divisorProduct

				// Do operation (add or multiply or square)
				monkeys[i].doOperation()

				if !partTwo {
					monkeys[i].operation.print(monkeys[i].items[0])

					// Reduce worry levels by 3 and round it to nearest integer
					monkeys[i].trisectWorry()
					fmt.Printf("    Monkey gets bored with item. Worry level is divided by 3 to %d.\n", monkeys[i].items[0])
				}

				// Monkey tests your worry level
				test := monkeys[i].checkDivisionBy()
				nextMonkeyIdx := 0
				if test {
					nextMonkeyIdx = monkeys[i].testMonkeyIdxIfTrue
					if !partTwo {
						fmt.Printf("    Current worry level is divisible by %d.\n", monkeys[i].testDivision)
					}
				} else {
					nextMonkeyIdx = monkeys[i].testMonkeyIdxIfFalse
					if !partTwo {
						fmt.Printf("    Current worry level is not divisible by %d.\n", monkeys[i].testDivision)
					}
				}

				// Add item to the target monkey's items
				monkeys[nextMonkeyIdx].items = append(monkeys[nextMonkeyIdx].items, monkeys[i].items[0])

				// Remove item from the current monkey
				monkeys[i].throwItem()

				if !partTwo {
					fmt.Printf("    Item with worry level %d is thrown to monkey %d.\n", monkeys[i].items[0], nextMonkeyIdx)
				}

				monkeys[i].inspectionCount += 1
			}
		}

		round += 1

		if !partTwo {
			fmt.Println()
			for i := 0; i < len(monkeys); i++ {
				fmt.Printf("Monkey %d: %v\n", i, monkeys[i].items)
			}
			fmt.Println()
		} else if round%1000 == 0 || round == 1 || round == 20 {
			fmt.Printf("== After round %d ==\n", round)
			for i := 0; i < len(monkeys); i++ {
				fmt.Printf("Monkey %d inspected items %d times.\n", i, monkeys[i].inspectionCount)
			}
			fmt.Println()
		}

		// For part 1, exit game loop after 20 rounds
		// For part 2, exit game loop after 10000 rounds
		if round == 20 && !partTwo {
			break
		} else if round == 10000 && partTwo {
			break
		}
	}

	// Keep highest inspection counts by sorting the array
	// and comparing with the index 0 and replacing if greater
	mostInspections := [2]int{0, 0}
	for i := 0; i < len(monkeys); i++ {
		if monkeys[i].inspectionCount > mostInspections[0] {
			mostInspections[0] = monkeys[i].inspectionCount
		}
		sort.Ints(mostInspections[:])
	}

	fmt.Println(mostInspections)

	monkeyBusiness := mostInspections[0] * mostInspections[1]
	fmt.Println("Result:", monkeyBusiness)
}
