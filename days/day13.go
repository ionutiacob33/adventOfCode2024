package days

import (
	"aoc2024/utils"
	"fmt"
)

func Day13(part int) {
	input := utils.ReadInput("inputs/test.txt")

	switch part {
	case 1:
		Part1Day13(input)
	case 2:
		Part2Day13(input)
	default:
		fmt.Println("Invalid part specified (use 1 or 2).")
	}
}

func Part1Day13(input string) {
	fmt.Println("=== Day 13, Part 1 ===")
	fmt.Println(input)
}

func Part2Day13(input string) {
	fmt.Println("=== Day 13, Part 2 ===")
	fmt.Println(input)
}
