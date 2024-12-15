package main

import (
	"aoc2024/days"
	"fmt"
	"os"
)

func main() {
	solutions := map[int]func(part int){
		1:  days.Day01,
		2:  days.Day02,
		3:  days.Day03,
		4:  days.Day04,
		5:  days.Day05,
		6:  days.Day06,
		7:  days.Day07,
		8:  days.Day08,
		9:  days.Day09,
		10: days.Day10,
		11: days.Day11,
		// TODO
	}

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <day> <part>")
		return
	}

	var day, part int
	fmt.Sscanf(os.Args[1], "%d", &day)
	fmt.Sscanf(os.Args[2], "%d", &part)

	if solution, exists := solutions[day]; exists {
		solution(part)
	} else {
		fmt.Printf("Solution for day %d not found\n", day)
	}
}
