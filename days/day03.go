package days

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Day03(part int) {
	input := utils.ReadInput("inputs/day03.txt")

	switch part {
	case 1:
		Part1Day03(input)
	case 2:
		Part2Day03(input)
	default:
		fmt.Println("Invalid part specified (use 1 or 2).")
	}
}

func Part1Day03(input string) {
	fmt.Println("=== Day 3, Part 1 ===")

	r, err := regexp.Compile(`mul\(\d+,\d+\)`)
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, matches := range r.FindAllString(input, -1) {
		sum += calculateValidMul(matches)
	}

	fmt.Println(sum)
}

func Part2Day03(input string) {
	fmt.Println("=== Day 3, Part 2 ===")

	r, err := regexp.Compile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	isMulAllowed := true
	for _, matches := range r.FindAllString(input, -1) {
		if matches == "don't()" {
			isMulAllowed = false
			continue
		}
		if matches == "do()" {
			isMulAllowed = true
			continue
		}
		if isMulAllowed {
			sum += calculateValidMul(matches)
		} else {
			continue
		}
	}

	fmt.Println(sum)
}

func calculateValidMul(input string) int {
	paranSplit := strings.Split(input, "(")
	commaSplit := strings.Split(paranSplit[1], ",")

	lastNumString := commaSplit[1][:len(commaSplit[1])-1]
	firstNum, _ := strconv.Atoi(commaSplit[0])
	lastNum, _ := strconv.Atoi(lastNumString)

	return firstNum * lastNum
}
