// DFS approach mostly translated from hyper-neutrino's python solution.
// https://github.com/hyper-neutrino/advent-of-code

package day19

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

type Blueprint struct {
	id      int
	recipes []RobotRecipe
	maxCost []int
}

type RobotRecipe struct {
	ingredients []int
}

type State struct {
	minute    int
	inventory []int
	robots    []int
}

func (s *State) toString() string {
	inventory := strings.Trim(strings.Replace(fmt.Sprint(s.inventory), " ", ",", -1), "[]")
	robots := strings.Trim(strings.Replace(fmt.Sprint(s.robots), " ", ",", -1), "[]")

	key := fmt.Sprintf("%d,%s,%s", s.minute, inventory, robots)
	return key
}

func dfs(blueprint Blueprint, seen map[string]int, state State) int {
	if state.minute == 0 {
		return state.inventory[3]
	}

	value, ok := seen[state.toString()]
	if ok {
		return value
	}

	maxValue := state.inventory[3] + state.robots[3]*state.minute

	for i, recipe := range blueprint.recipes {
		if i != 3 && state.robots[i] >= blueprint.maxCost[i] {
			continue
		}

		waitTime := 0
		built := true
		for j, cost := range recipe.ingredients {
            if cost == 0 {
                continue
            }
			if state.robots[j] == 0 {
				built = false
				break
			}

            waitTime = utils.Max(waitTime, int(math.Ceil(float64(cost - state.inventory[j]) / float64(state.robots[j]))))
		}

		if built {
            remainingTime := state.minute - waitTime - 1
            if remainingTime <= 0 {
                continue 
            }

            robots := make([]int, len(state.robots))
            copy(robots, state.robots)

			inventory := make([]int, len(state.inventory))
            for i := 0; i < len(state.inventory); i++ {
                inventory[i] = state.inventory[i] + state.robots[i]*(waitTime+1)
            }
			for j, cost := range recipe.ingredients {
				inventory[j] -= cost
			}

            robots[i] += 1

            for j := range inventory {
                if j == 3 {
                    continue
                }
                inventory[j] = utils.Min(inventory[j], blueprint.maxCost[j] * remainingTime)
            }

            maxValue = utils.Max(maxValue, dfs(blueprint, seen, State{remainingTime, inventory, robots}))
		}
	}

    seen[state.toString()] = maxValue
	return maxValue
}

func Day19() {
	contents := utils.GetFileContents("day19/input")
	blueprints := make([]Blueprint, 0)

	for _, line := range contents {
		tokens := strings.Split(line, " ")

		tokens[1] = strings.ReplaceAll(tokens[1], ":", "")
		id, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatal("NaN")
		}

		oreRobotCost, err := strconv.Atoi(tokens[6])
		if err != nil {
			log.Fatal("NaN")
		}

		clayRobotCostOre, err := strconv.Atoi(tokens[12])
		if err != nil {
			log.Fatal("NaN")
		}

		obsidianRobotCostOre, err := strconv.Atoi(tokens[18])
		if err != nil {
			log.Fatal("NaN")
		}

		obsidianRobotCostClay, err := strconv.Atoi(tokens[21])
		if err != nil {
			log.Fatal("NaN")
		}

		geodeRobotCostOre, err := strconv.Atoi(tokens[27])
		if err != nil {
			log.Fatal("NaN")
		}

		geodeRobotCostObsidian, err := strconv.Atoi(tokens[30])
		if err != nil {
			log.Fatal("NaN")
		}

		recipes := []RobotRecipe{
			{[]int{oreRobotCost, 0, 0, 0}},
			{[]int{clayRobotCostOre, 0, 0, 0}},
			{[]int{obsidianRobotCostOre, obsidianRobotCostClay, 0, 0}},
			{[]int{geodeRobotCostOre, 0, geodeRobotCostObsidian, 0}},
		}

		maxOreCost := make([]int, 4)

		for _, r := range recipes {
			for i, cost := range r.ingredients {
				maxOreCost[i] = utils.Max(maxOreCost[i], cost)
			}
		}

		blueprints = append(blueprints, Blueprint{id, recipes, maxOreCost})
	}

	start := time.Now()

	sum := 0
	for _, bp := range blueprints {
		seen := make(map[string]int)
		inventory := make([]int, len(bp.recipes))
		robots := make([]int, len(bp.recipes))
        robots[0] = 1

		maxGeodes := dfs(bp, seen, State{
			24, inventory, robots,
		})
		sum += maxGeodes * bp.id
	}

	fmt.Println("Results part 1:", sum)

	elapsed := time.Since(start)
	fmt.Println(elapsed)

    // part 2

	start = time.Now()

    result := 1
    maxSize := utils.Min(len(blueprints), 3)
    for _, bp := range blueprints[:maxSize] {
		seen := make(map[string]int)
		inventory := make([]int, len(bp.recipes))
		robots := make([]int, len(bp.recipes))
        robots[0] = 1

		result *= dfs(bp, seen, State{
			32, inventory, robots,
		})
	}

	fmt.Println("Results part 2:", result)

	elapsed = time.Since(start)
	fmt.Println(elapsed)
}

