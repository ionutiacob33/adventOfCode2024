package days

import (
	"aoc2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func Day02(part int) {
	input := utils.ReadInput("inputs/day02.txt")

	switch part {
	case 1:
		Part1Day02(input)
	case 2:
		Part2Day02(input)
	default:
		fmt.Println("Invalid part specified (use 1 or 2).")
	}
}

func Part1Day02(input string) {
	fmt.Println("=== Day 2, Part 1 ===")

	lines := strings.Split(strings.TrimSpace(input), "\n")
	safeReports := 0

	for _, line := range lines {
		stringNums := strings.Fields(line)
		nums := make([]int, len(stringNums))

		for i, stringNum := range stringNums {
			num, _ := strconv.Atoi(stringNum)
			nums[i] = num
		}

		isSafe := isRowSafe(nums)

		if isSafe {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}

func Part2Day02(input string) {
	fmt.Println("=== Day 2, Part 2 ===")

	lines := strings.Split(strings.TrimSpace(input), "\n")
	safeReports := 0

	for _, line := range lines {
		stringNums := strings.Fields(line)
		nums := make([]int, len(stringNums))

		for i, stringNum := range stringNums {
			num, _ := strconv.Atoi(stringNum)
			nums[i] = num
		}

		isSafe := false
		isSafe = isRowSafe(nums)

		if !isSafe {
			for i := 0; i < len(nums); i++ {
				modifiedArray := []int{}
				modifiedArray = append(modifiedArray, nums[:i]...)
				modifiedArray = append(modifiedArray, nums[i+1:]...)
				if isRowSafe(modifiedArray) {
					safeReports++
					break
				}
			}
		}

		if isSafe {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}

func isRowSafe(nums []int) bool {
	isIncreasing := true
	isSafe := true

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]

		if diff == 0 {
			isSafe = false
			break
		}

		if diff > 3 || diff < -3 {
			isSafe = false
			break
		}

		if i == 1 && diff < 0 {
			isIncreasing = false
		} else if (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
			isSafe = false
			break
		}
	}

	return isSafe
}
