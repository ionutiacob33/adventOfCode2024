package days

import (
	"aoc2024/utils"
	"fmt"
)

func Day09(part int) {
	input := utils.ReadInput("inputs/day09.txt")

	switch part {
	case 1:
		Part1Day09(input)
	case 2:
		Part2Day09(input)
	default:
		fmt.Println("Invalid part specified (use 1 or 2).")
	}
}

func Part1Day09(input string) {
	fmt.Println("=== Day 9, Part 1 ===")

	diskMap := []rune(input)
	disk := []int{}
	currentId := 0

	for i := 0; i < len(diskMap); i += 2 {
		for k := 0; k < int(diskMap[i]-'0'); k++ {
			disk = append(disk, currentId)
		}
		if i < len(diskMap)-1 {
			for k := 0; k < int(diskMap[i+1]-'0'); k++ {
				disk = append(disk, -1)
			}
		}
		currentId++
	}

	latestIndex := len(disk) - 1
	for i := 0; i < latestIndex; i++ {
		for disk[latestIndex] == -1 {
			latestIndex--
		}
		if disk[i] == -1 {
			aux := disk[latestIndex]
			disk[latestIndex] = disk[i]
			disk[i] = aux
			latestIndex--
		}
	}

	checksum := 0
	for i := 0; i <= latestIndex; i++ {
		checksum += i * disk[i]
	}

	fmt.Println(checksum)
}

func Part2Day09(input string) {
	fmt.Println("=== Day 9, Part 2 ===")

	diskMap := []rune(input)
	disk := []int{}
	currentId := 0

	for i := 0; i < len(diskMap); i += 2 {
		for k := 0; k < int(diskMap[i]-'0'); k++ {
			disk = append(disk, currentId)
		}
		if i < len(diskMap)-1 {
			for k := 0; k < int(diskMap[i+1]-'0'); k++ {
				disk = append(disk, -1)
			}
		}
		currentId++
	}

	for i := len(disk) - 1; i >= 0; i-- {
		for disk[i] == -1 {
			i--
		}

		currentFileId := disk[i]
		currentFileSize := 0
		for i >= 0 && currentFileId == disk[i] {
			currentFileSize++
			i--
		}

		i++
		for j := 0; j < len(disk); j++ {
			for j < len(disk) && disk[j] != -1 {
				j++
			}

			currentSpaceSize := 0
			for j < len(disk) && disk[j] == -1 {
				currentSpaceSize++
				j++
			}

			if currentSpaceSize >= currentFileSize && i > j-currentSpaceSize {
				disk = move(disk, i, currentFileSize, j-currentSpaceSize)
				break
			}
		}
	}

	checksum := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] != -1 {
			checksum += i * disk[i]
		}
	}

	fmt.Println(checksum)
}

func move(disk []int, fileStart int, currentFileSize int, spaceStart int) []int {
	for i := 0; i < currentFileSize; i++ {
		aux := disk[fileStart+i]
		disk[fileStart+i] = disk[spaceStart+i]
		disk[spaceStart+i] = aux
	}

	return disk
}
