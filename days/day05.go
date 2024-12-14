package days

import (
	"aoc2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func Day05(part int) {
	input := utils.ReadInput("inputs/day05.txt")

	switch part {
	case 1:
		Part1Day05(input)
	case 2:
		Part2Day05(input)
	default:
		fmt.Println("Invalid part specified (use 1 or 2).")
	}
}

func Part1Day05(input string) {
	fmt.Println("=== Day 5, Part 1 ===")
	lines := strings.Split(input, "\n")

	rules := make(map[int][]int)
	updates := [][]int{}
	output := 0

	ruleMode := true
	for _, line := range lines {
		if len(line) == 0 {
			ruleMode = false
			continue
		}

		if ruleMode {
			rule := strings.Split(line, "|")
			ruleX, _ := strconv.Atoi(rule[0])
			ruleY, _ := strconv.Atoi(rule[1])

			rules[ruleX] = append(rules[ruleX], ruleY)
		} else {
			currentUpdate := []int{}
			parts := strings.Split(line, ",")

			for _, part := range parts {
				num, _ := strconv.Atoi(part)
				currentUpdate = append(currentUpdate, num)
			}

			updates = append(updates, currentUpdate)
		}
	}

	for _, currentUpdate := range updates {
		isValid := true
		for i := 1; i < len(currentUpdate); i++ {
			currentNum := currentUpdate[i]
			for _, num := range currentUpdate[:i] {
				for _, ruleNum := range rules[currentNum] {
					if num == ruleNum {
						isValid = false
						break
					}
				}

				if !isValid {
					break
				}
			}
		}

		if isValid {
			middle := currentUpdate[len(currentUpdate)/2]
			output += middle
		}
	}

	fmt.Println(output)
}

func Part2Day05(input string) {
	fmt.Println("=== Day 5, Part 2 ===")
	lines := strings.Split(input, "\n")

	rules := make(map[int][]int)
	updates := [][]int{}
	output := 0

	ruleMode := true
	for _, line := range lines {
		if len(line) == 0 {
			ruleMode = false
			continue
		}

		if ruleMode {
			rule := strings.Split(line, "|")
			ruleX, _ := strconv.Atoi(rule[0])
			ruleY, _ := strconv.Atoi(rule[1])

			rules[ruleX] = append(rules[ruleX], ruleY)
		} else {
			currentUpdate := []int{}
			parts := strings.Split(line, ",")

			for _, part := range parts {
				num, _ := strconv.Atoi(part)
				currentUpdate = append(currentUpdate, num)
			}

			updates = append(updates, currentUpdate)
		}
	}

	for _, currentUpdate := range updates {
		isValid := true
		for i := 1; i < len(currentUpdate); i++ {
			currentNum := currentUpdate[i]
			for _, num := range currentUpdate[:i] {
				for _, ruleNum := range rules[currentNum] {
					if num == ruleNum {
						isValid = false
						break
					}
				}

				if !isValid {
					break
				}
			}
		}

		if !isValid {
			ordered := []int{}
			remaining := map[int]bool{}
			dependencies := make(map[int][]int)

			for _, num := range currentUpdate {
				remaining[num] = true

				for _, dep := range rules[num] {
					dependencies[dep] = append(dependencies[dep], num)
				}
			}

			for len(remaining) > 0 {
				for num := range remaining {
					if len(dependencies[num]) == 0 {
						ordered = append(ordered, num)
						delete(remaining, num)

						for key, val := range dependencies {
							newList := []int{}

							for _, n := range val {
								if n != num {
									newList = append(newList, n)
								}
							}

							dependencies[key] = newList
						}
					}
				}
			}

			middle := ordered[len(ordered)/2]
			output += middle
		}
	}

	fmt.Println(output)
}
