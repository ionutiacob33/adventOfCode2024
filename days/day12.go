package days

import (
	"aoc2024/utils"
	"fmt"
)

func Day12(part int) {
	input := utils.ReadInput("inputs/day12.txt")

	switch part {
	case 1:
		Part1Day12(input)
	case 2:
		Part2Day12(input)
	default:
		fmt.Println("Invalid part specified (use 1 or 2).")
	}
}

func flodFill(grid [][]rune, visited [][]bool, x, y int, plantType rune) (int, int) {
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	stack := [][2]int{{x, y}}
	visited[x][y] = true
	area, perimeter := 0, 0

	for len(stack) > 0 {
		cx, cy := stack[len(stack)-1][0], stack[len(stack)-1][1]
		stack = stack[:len(stack)-1]
		area++

		for _, d := range directions {
			nx, ny := cx+d[0], cy+d[1]

			if nx < 0 || nx >= len(grid) || ny < 0 || ny >= len(grid[0]) {
				perimeter++
			} else if grid[nx][ny] != plantType {
				perimeter++
			} else if !visited[nx][ny] {
				visited[nx][ny] = true
				stack = append(stack, [2]int{nx, ny})
			}
		}
	}

	return area, perimeter
}

func flodFillSides(grid [][]rune, visited [][]bool, x, y int, plantType rune) (int, int) {
	directions := [][3]int{{0, 1, 0}, {1, 0, 1}, {0, -1, 0}, {-1, 0, 1}}
	stack := [][2]int{{x, y}}
	visited[x][y] = true
	area, perimeter := 0, 0

	for len(stack) > 0 {
		cx, cy := stack[len(stack)-1][0], stack[len(stack)-1][1]
		stack = stack[:len(stack)-1]
		area++

		for _, d := range directions {
			nx, ny := cx+d[0], cy+d[1]

			if nx < 0 || nx >= len(grid) || ny < 0 || ny >= len(grid[0]) {
				perimeter++
			} else if grid[nx][ny] != plantType {
				perimeter++
			} else if !visited[nx][ny] {
				visited[nx][ny] = true
				stack = append(stack, [2]int{nx, ny})
			}
		}
	}

	return area, perimeter
}

func Part1Day12(input string) {
	fmt.Println("=== Day 12, Part 1 ===")

	grid := utils.Read2dRuneArray(input)
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	totalPrice := 0
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if !visited[x][y] {
				plantType := grid[x][y]
				area, perimeter := flodFill(grid, visited, x, y, plantType)
				totalPrice += area * perimeter
			}
		}
	}

	fmt.Println(totalPrice)
}

func Part2Day12(input string) {
	grid := utils.Read2dRuneArray(input)
	output := 0

	directions := [][2]int{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	}

	outerCorners := []int{
		0, 0, 0, 1, 0, 0, 1,
		2, 0, 1, 0, 2,
		1, 2, 2, 4,
	}

	checkInnerCorners := [][][]int{
		{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}},
		{{1, -1}, {1, 1}},
		{{-1, 1}, {1, 1}},
		{{1, 1}},
		{{-1, -1}, {-1, 1}},
		{},
		{{-1, 1}},

		{},
		{{-1, -1}, {1, -1}},
		{{1, -1}},
		{},
		{},

		{{-1, -1}},
		{},
		{},
		{},
	}

	visited := make(map[[2]int]bool)

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if visited[[2]int{row, col}] {
				continue
			}
			plant := grid[row][col]
			area := 0
			corners := 0

			queue := [][2]int{{row, col}}
			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]

				currentRow, currentCol := current[0], current[1]
				if visited[[2]int{currentRow, currentCol}] {
					continue
				}
				visited[[2]int{currentRow, currentCol}] = true

				area++
				cornerType := 0

				for i, dir := range directions {
					newRow := currentRow + dir[0]
					newCol := currentCol + dir[1]

					if newRow < 0 || newCol < 0 || newRow >= len(grid) || newCol >= len(grid[row]) {
						cornerType += (1 << i)
					} else if grid[newRow][newCol] != plant {
						cornerType += (1 << i)
					} else if !visited[[2]int{newRow, newCol}] {
						queue = append(queue, [2]int{newRow, newCol})
					}
				}

				outerCornerCount := outerCorners[cornerType]
				innerCornerCount := 0
				for _, corner := range checkInnerCorners[cornerType] {
					newRow := currentRow + corner[0]
					newCol := currentCol + corner[1]

					if newRow < 0 || newCol < 0 || newRow >= len(grid) || newCol >= len(grid[row]) {
						continue
					} else if grid[newRow][newCol] != plant {
						innerCornerCount++
					}
				}
				corners += outerCornerCount + innerCornerCount
			}

			price := area * corners
			output += price
		}
	}

	fmt.Println("Output Day 12 Part 2", output)
}
