package days

import (
	"aoc2024/utils"
	"fmt"
)

func Day10(part int) {
	input := utils.ReadInput("inputs/day10.txt")

	switch part {
	case 1:
		Part1Day10(input)
	case 2:
		Part2Day10(input)
	default:
		fmt.Println("Invalid part specified (use 1 or 2).")
	}
}

type point struct {
	x int
	y int
}

var score map[point]int
var visited map[point][]point

func findScore(grid [][]rune, copyGrid [][]rune, originI int, originJ int, currentI int, currentJ int, previous rune) {
	if currentI < 0 || currentJ < 0 || currentI >= len(grid) || currentJ >= len(grid[0]) {
		return
	}

	if grid[currentI][currentJ] == '9' && grid[currentI][currentJ] == previous+1 {
		origin := point{originI, originJ}
		peak := point{currentI, currentJ}
		for i := 0; i < len(visited[origin]); i++ {
			if visited[origin][i] == peak {
				return
			}
		}
		score[origin]++
		visited[origin] = append(visited[origin], peak)
		return
	}

	if !(grid[currentI][currentJ] == previous+1) {
		return
	}

	findScore(grid, copyGrid, originI, originJ, currentI-1, currentJ, grid[currentI][currentJ])
	findScore(grid, copyGrid, originI, originJ, currentI, currentJ+1, grid[currentI][currentJ])
	findScore(grid, copyGrid, originI, originJ, currentI+1, currentJ, grid[currentI][currentJ])
	findScore(grid, copyGrid, originI, originJ, currentI, currentJ-1, grid[currentI][currentJ])
}

func findRating(grid [][]rune, copyGrid [][]rune, originI int, originJ int, currentI int, currentJ int, previous rune) {
	if currentI < 0 || currentJ < 0 || currentI >= len(grid) || currentJ >= len(grid[0]) {
		return
	}

	if grid[currentI][currentJ] == '9' && grid[currentI][currentJ] == previous+1 {
		origin := point{originI, originJ}
		score[origin]++
		return
	}

	if !(grid[currentI][currentJ] == previous+1) {
		return
	}

	findRating(grid, copyGrid, originI, originJ, currentI-1, currentJ, grid[currentI][currentJ])
	findRating(grid, copyGrid, originI, originJ, currentI, currentJ+1, grid[currentI][currentJ])
	findRating(grid, copyGrid, originI, originJ, currentI+1, currentJ, grid[currentI][currentJ])
	findRating(grid, copyGrid, originI, originJ, currentI, currentJ-1, grid[currentI][currentJ])
}

func Part1Day10(input string) {
	fmt.Println("=== Day 10, Part 1 ===")

	grid := utils.Read2dRuneArray(input)
	copyGrid := utils.Read2dRuneArray(input)
	score = make(map[point]int)
	visited = map[point][]point{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '0' {
				findScore(grid, copyGrid, i, j, i, j, '0'-1)
			}
		}
	}

	total := 0
	for _, val := range score {
		total += val
	}

	fmt.Println(total)
}

func Part2Day10(input string) {
	fmt.Println("=== Day 10, Part 2 ===")

	grid := utils.Read2dRuneArray(input)
	copyGrid := utils.Read2dRuneArray(input)
	score = make(map[point]int)
	visited = map[point][]point{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '0' {
				findRating(grid, copyGrid, i, j, i, j, '0'-1)
			}
		}
	}

	total := 0
	for _, val := range score {
		total += val
	}

	fmt.Println(total)
}
