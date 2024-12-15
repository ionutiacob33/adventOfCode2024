package days

import (
	"aoc2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func Day11(part int) {
	input := utils.ReadInput("inputs/day11.txt")

	switch part {
	case 1:
		Part1Day11(input)
	case 2:
		Part2Day11(input)
	default:
		fmt.Println("Invalid part specified (use 1 or 2).")
	}
}

func digits(num int) int {
	digits := 0
	for num > 0 {
		digits++
		num /= 10
	}
	return digits
}

func splitNumber(n int) (int, int) {
	numStr := strconv.Itoa(n)

	length := len(numStr)

	leftStr := numStr[:length/2]
	rightStr := numStr[length/2:]

	left, _ := strconv.Atoi(leftStr)
	right, _ := strconv.Atoi(rightStr)

	return left, right
}

func Part1Day11(input string) {
	fmt.Println("=== Day 11, Part 1 ===")

	blinks := 25
	stringNums := strings.Fields(input)
	nums := make([]int, len(stringNums))
	numsAfterBlink := []int{}

	for i, stringNum := range stringNums {
		num, _ := strconv.Atoi(stringNum)
		nums[i] = num
	}

	for range blinks {
		for i := 0; i < len(nums); i++ {
			if nums[i] == 0 {
				numsAfterBlink = append(numsAfterBlink, 1)
				continue
			}

			num := nums[i]
			numOfDigits := digits(nums[i])
			if numOfDigits%2 == 0 {
				left, right := splitNumber(num)

				numsAfterBlink = append(numsAfterBlink, left)
				numsAfterBlink = append(numsAfterBlink, right)
			} else {
				numsAfterBlink = append(numsAfterBlink, nums[i]*2024)
			}
		}
		// fmt.Println(numsAfterBlink)
		nums = append([]int{}, numsAfterBlink...)
		numsAfterBlink = []int{}
	}

	fmt.Println(len(nums))
}

func Part2Day11(input string) {
	fmt.Println("=== Day 11, Part 2 ===")

	blinks := 75
	stringNums := strings.Fields(input)
	stones := map[int]int{}

	for _, stringNum := range stringNums {
		num, _ := strconv.Atoi(stringNum)
		stones[num]++
	}

	for range blinks {
		newStones := map[int]int{}

		for stone, count := range stones {
			if stone == 0 {
				newStones[1] += count
				continue
			}

			if len(strconv.Itoa(stone))%2 == 0 {
				left, right := splitNumber(stone)

				newStones[left] += count
				newStones[right] += count
			} else {
				newStones[stone*2024] += count
			}
		}
		// fmt.Println(numsAfterBlink)
		stones = newStones
	}

	output := 0
	for _, count := range stones {
		output += count
	}

	fmt.Println(output)
}
