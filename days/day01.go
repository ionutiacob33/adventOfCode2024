package days

import (
	"aoc2024/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day01(part int) {
	input := utils.ReadInput("inputs/day01.txt")

	switch part {
	case 1:
		Part1Day01(input)
	case 2:
		Part2Day01(input)
	default:
		fmt.Println("Invalid part specified (use 1 or 2).")
	}
}

func Part1Day01(input string) {
	fmt.Println("=== Day 1, Part 1 ===")
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var left []int
	var right []int

	for _, line := range lines {
		nums := strings.Fields(line)

		leftNum, _ := strconv.Atoi(nums[0])
		rightNum, _ := strconv.Atoi(nums[1])

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0

	for i := 0; i < len(left); i++ {
		distance := left[i] - right[i]

		if distance < 0 {
			distance *= -1
		}

		totalDistance += distance
	}

	fmt.Println(totalDistance)
}

func Part2Day01(input string) {
	fmt.Println("=== Day 1, Part 2 ===")
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var left []int
	right := map[int]int{}

	for _, line := range lines {
		nums := strings.Fields(line)

		leftNum, _ := strconv.Atoi(nums[0])
		rightNum, _ := strconv.Atoi(nums[1])

		left = append(left, leftNum)
		value, exists := right[rightNum]
		if exists {
			right[rightNum] = value + 1
		} else {
			right[rightNum] = 1
		}
	}

	totalDistance := 0

	for _, num := range left {
		value, exists := right[num]
		if exists {
			totalDistance += num * value
		}
	}

	fmt.Println(totalDistance)
}
