package day5

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type Cargo struct {
	stacks            []utils.Stack
	line_with_numbers int
	contents          []string
}

func (c *Cargo) Initialize(contents []string) {
	c.contents = contents

	// Find the line index of stack numbers
	for line_idx, line := range c.contents {
		if line[1] == '1' {
			c.line_with_numbers = line_idx
			break
		}
	}

	// Iterating the stack numberings by character
	// If character is a digit, append an additional stack
	for _, num := range contents[c.line_with_numbers] {
		if unicode.IsDigit(num) {
			c.stacks = append(c.stacks, utils.Stack{})
		}
	}
}

func (c *Cargo) AddCrates() {
	// Iterate stacks and add to stacks
	for i := c.line_with_numbers - 1; i >= 0; i-- {
		stack_idx := 0
		for idx, ch := range c.contents[i] {
			if idx%4 == 0 && idx != 0 {
				stack_idx++
			}
			if unicode.IsLetter(ch) {
				c.stacks[stack_idx].Push(ch)
			}
		}
	}
}

func (c *Cargo) PrintTop() {
	for i := 0; i < len(c.stacks); i++ {
		ch, err := c.stacks[i].GetTop()
		if err != nil {
			continue
		}
		fmt.Printf("%c", ch)
	}
}

func (c *Cargo) Move(part2 bool) {
	// Iterate moves
	for i := c.line_with_numbers + 2; i < len(c.contents); i++ {
		tokens := strings.Split(c.contents[i], " ")
		move_cnt, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatal("cannot convert string to int")
		}

		from_stack, err := strconv.Atoi(tokens[3])
		if err != nil {
			log.Fatal("cannot convert string to int")
		}
		from_stack -= 1

		to_stack, err := strconv.Atoi(tokens[5])
		if err != nil {
			log.Fatal("cannot convert string to int")
		}
		to_stack -= 1

		if !part2 {
			for j := 0; j < move_cnt; j++ {
				item, err := c.stacks[from_stack].Pop()
				if err != nil {
					log.Fatal("cannot pop an empty stack")
				}

				c.stacks[to_stack].Push(item)
			}
		} else {
			items := []rune{}
			for j := 0; j < move_cnt; j++ {
				item, err := c.stacks[from_stack].Pop()
				if err != nil {
					log.Fatal("cannot pop an empty stack")
				}
				items = append(items, item)
			}

			for j := len(items) - 1; j >= 0; j-- {
				c.stacks[to_stack].Push(items[j])
			}
		}
	}
}

func Day5() {
	contents := utils.GetFileContents("day5/input")

	cargo := Cargo{}
	cargo.Initialize(contents)
	cargo.AddCrates()
	cargo.Move(false)

	fmt.Print("Solution part 1: ")
	cargo.PrintTop()

	cargo.Initialize(contents)
	cargo.AddCrates()
	cargo.Move(true)

	fmt.Println()
	fmt.Print("Solution part 2: ")
	cargo.PrintTop()
}
