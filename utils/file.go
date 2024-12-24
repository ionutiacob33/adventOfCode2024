package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadInput(filePath string) string {
	data, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filePath, err)
		return ""
	}

	return string(data)
}

func Read2dRuneArray(input string) [][]rune {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	return grid
}

func printGrid(grid [][]rune) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '.' {
				fmt.Printf(" . ")
			} else {
				fmt.Printf("%d ", grid[i][j])
			}
		}
		fmt.Println()
	}
}
