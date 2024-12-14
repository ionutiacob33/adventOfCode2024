package days

import (
	"aoc2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func Day07(part int) {
	input := utils.ReadInput("inputs/day07.txt")

	switch part {
	case 1:
		Part1Day07(input)
	case 2:
		Part2Day07(input)
	default:
		fmt.Println("Invalid part specified (use 1 or 2).")
	}
}

func Part1Day07(input string) {
	fmt.Println("=== Day 7, Part 1 ===")
	lines := strings.Split(input, "\n")

	output := 0

	var evaluate func(target int, numbers []int, index int, currentResult int) bool
	evaluate = func(target int, numbers []int, index int, currentResult int) bool {
		if index >= len(numbers) {
			return target == currentResult
		}

		currentNumber := numbers[index]

		newResult := currentResult + currentNumber
		if newResult <= target {
			if evaluate(target, numbers, index+1, newResult) {
				return true
			}
		}

		newResult = currentResult * currentNumber
		if newResult <= target {
			return evaluate(target, numbers, index+1, newResult)
		}

		return false
	}

	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], ":")
		target, _ := strconv.Atoi(parts[0])

		numberStrings := strings.Fields(parts[1])
		numbers := make([]int, len(numberStrings))
		for i, numString := range numberStrings {
			numbers[i], _ = strconv.Atoi(numString)
		}

		if evaluate(target, numbers, 0, 0) {
			output += target
		}
	}

	fmt.Println(output)
}

func Part2Day07(input string) {
	fmt.Println("=== Day 7, Part 2 ===")
	lines := strings.Split(input, "\n")

	output := 0

	var evaluate func(target int, numbers []int, index int, currentResult int) bool
	evaluate = func(target int, numbers []int, index int, currentResult int) bool {
		if index >= len(numbers) {
			return target == currentResult
		}

		currentNumber := numbers[index]

		newResult := currentResult + currentNumber
		if newResult <= target {
			if evaluate(target, numbers, index+1, newResult) {
				return true
			}
		}

		newResult = currentResult * currentNumber
		if newResult <= target {
			if evaluate(target, numbers, index+1, newResult) {
				return true
			}
		}

		newResult, _ = strconv.Atoi(fmt.Sprintf("%d%d", currentResult, currentNumber))
		if newResult <= target {
			return evaluate(target, numbers, index+1, newResult)
		}

		return false
	}

	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], ":")
		target, _ := strconv.Atoi(parts[0])

		numberStrings := strings.Fields(parts[1])
		numbers := make([]int, len(numberStrings))
		for i, numString := range numberStrings {
			numbers[i], _ = strconv.Atoi(numString)
		}

		if evaluate(target, numbers, 0, 0) {
			output += target
		}
	}

	fmt.Println(output)
}
