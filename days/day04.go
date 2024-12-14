package days

import (
	"aoc2024/utils"
	"fmt"
	"strings"
)

func Day04(part int) {
	input := utils.ReadInput("inputs/day04.txt")

	switch part {
	case 1:
		Part1Day04(input)
	case 2:
		Part2Day04(input)
	default:
		fmt.Println("Invalid part specified (use 1 or 2).")
	}
}

func Part1Day04(input string) {
	fmt.Println("=== Day 4, Part 1 ===")

	lines := strings.Split(input, "\n")
	grid := make([][]rune, len(lines))
	output := 0

	word := "XMAS"
	directions := [][2]int{
		{-1, 0},  // up
		{-1, 1},  // up-right
		{0, 1},   // right
		{1, 1},   // down-right
		{1, 0},   // down
		{1, -1},  // down-left
		{0, -1},  // left
		{-1, -1}, // up-left
	}

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			for _, dir := range directions {
				rowDir := dir[0]
				colDir := dir[1]

				isXmas := true
				for charIndex := 0; charIndex < len(word); charIndex++ {
					offsetRow := row + (rowDir * charIndex)
					offsetCol := col + (colDir * charIndex)

					if offsetRow < 0 || offsetRow >= len(grid) || offsetCol < 0 || offsetCol >= len(grid[row]) {
						isXmas = false
						break
					}

					if grid[offsetRow][offsetCol] != rune(word[charIndex]) {
						isXmas = false
						break
					} else {
						fmt.Printf("%d %d %s\n", offsetRow, offsetCol, string(word[charIndex]))
					}
				}

				if isXmas {
					fmt.Println("found")
					output++
				}
			}
		}
	}

	// fmt.Println(grid)
	fmt.Println(output)
}

func Part2Day04(input string) {
	fmt.Println("=== Day 4, Part 2 ===")

	lines := strings.Split(input, "\n")
	grid := make([][]rune, len(lines))
	output := 0

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[row])-1; col++ {
			if grid[row][col] == 'A' {
				if (grid[row-1][col-1] == 'M' && grid[row+1][col+1] == 'S' && grid[row+1][col-1] == 'M' && grid[row-1][col+1] == 'S') ||
					(grid[row-1][col-1] == 'M' && grid[row+1][col+1] == 'S' && grid[row+1][col-1] == 'S' && grid[row-1][col+1] == 'M') ||
					(grid[row-1][col-1] == 'S' && grid[row+1][col+1] == 'M' && grid[row+1][col-1] == 'S' && grid[row-1][col+1] == 'M') ||
					(grid[row-1][col-1] == 'S' && grid[row+1][col+1] == 'M' && grid[row+1][col-1] == 'M' && grid[row-1][col+1] == 'S') {
					output++
				}
			}
		}
	}

	fmt.Println(output)
}
