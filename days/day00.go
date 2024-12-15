package days

import (
	"aoc2024/utils"
	"fmt"
)

func Day00(part int) {
	input := utils.ReadInput("inputs/test.txt")

	switch part {
	case 1:
		Part1Day00(input)
	case 2:
		Part2Day00(input)
	default:
		fmt.Println("Invalid part specified (use 1 or 2).")
	}
}

func Part1Day00(input string) {
	fmt.Println("=== Day 00, Part 1 ===")
	fmt.Println(input)
}

func Part2Day00(input string) {
	fmt.Println("=== Day 00, Part 2 ===")
	fmt.Println(input)
}
