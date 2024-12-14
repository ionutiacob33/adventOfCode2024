package days

import (
	"aoc2024/utils"
	"fmt"
)

type pos struct {
	x int
	y int
}

func Day08(part int) {
	input := utils.ReadInput("inputs/day08.txt")

	switch part {
	case 1:
		Part1Day08(input)
	case 2:
		Part2Day08(input)
	default:
		fmt.Println("Invalid part specified (use 1 or 2).")
	}
}

func Part1Day08(input string) {
	fmt.Println("=== Day 8, Part 1 ===")
	grid := utils.Read2dRuneArray(input)

	antenas := make(map[rune][]pos)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != '.' {
				antena := grid[i][j]
				antenas[antena] = append(antenas[antena], pos{i, j})
			}
		}
	}

	visitedPos := make(map[pos]bool)

	for _, positions := range antenas {
		for i := 0; i < len(positions)-1; i++ {
			for j := i + 1; j < len(positions); j++ {
				antiX := 2*positions[i].x - positions[j].x
				antiY := 2*positions[i].y - positions[j].y

				if antiX >= 0 && antiX < len(grid) && antiY >= 0 && antiY < len(grid[0]) {
					if !visitedPos[pos{antiX, antiY}] {
						visitedPos[pos{antiX, antiY}] = true
					}
				}

				antiX = 2*positions[j].x - positions[i].x
				antiY = 2*positions[j].y - positions[i].y

				if antiX >= 0 && antiX < len(grid) && antiY >= 0 && antiY < len(grid[0]) {
					if !visitedPos[pos{antiX, antiY}] {
						visitedPos[pos{antiX, antiY}] = true
					}
				}
			}
		}
	}

	fmt.Println(len(visitedPos))
}

func Part2Day08(input string) {
	fmt.Println("=== Day 8, Part 2 ===")
	grid := utils.Read2dRuneArray(input)

	antenas := make(map[rune][]pos)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != '.' {
				antena := grid[i][j]
				antenas[antena] = append(antenas[antena], pos{i, j})
			}
		}
	}

	visitedPos := make(map[pos]bool)

	for _, positions := range antenas {
		for i := 0; i < len(positions)-1; i++ {
			for j := i + 1; j < len(positions); j++ {
				x1 := positions[i].x
				y1 := positions[i].y
				x2 := positions[j].x
				y2 := positions[j].y

				pos1 := pos{x1, y1}
				pos2 := pos{x2, y2}
				visitedPos[pos1] = true
				visitedPos[pos2] = true

				antiX := 2*x2 - x1
				antiY := 2*y2 - y1

				for antiX >= 0 && antiX < len(grid) && antiY >= 0 && antiY < len(grid[0]) {
					visitedPos[pos{antiX, antiY}] = true
					auxX := antiX
					auxY := antiY
					antiX = 2*antiX - x2
					antiY = 2*antiY - y2
					x2 = auxX
					y2 = auxY
				}

				x1 = pos1.x
				y1 = pos1.y
				x2 = pos2.x
				y2 = pos2.y

				antiX = 2*x1 - x2
				antiY = 2*y1 - y2

				for antiX >= 0 && antiX < len(grid) && antiY >= 0 && antiY < len(grid[0]) {
					visitedPos[pos{antiX, antiY}] = true
					auxX := antiX
					auxY := antiY
					antiX = 2*antiX - x1
					antiY = 2*antiY - y1
					x1 = auxX
					y1 = auxY
				}
			}
		}
	}

	fmt.Println(len(visitedPos))
}
