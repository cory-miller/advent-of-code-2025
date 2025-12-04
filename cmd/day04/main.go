package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	dirs = [][]int{
		// Above
		{-1, -1}, {-1, 0}, {-1, 1},
		// Sides
		{0, -1}, {0, 1},
		// Below
		{1, -1}, {1, 0}, {1, 1},
	}
)

// @ -> 64 as ASCII
// . -> 46 as ASCII

func parseInput(path string) [][]rune {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	grid := make([][]rune, 0)

	for scanner.Scan() {
		rowText := scanner.Text()

		row := []rune(rowText)

		grid = append(grid, row)
	}

	return grid
}

func partOne(grid [][]rune) int {
	removableRolls := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			shelf := grid[row][col]

			if shelf != '@' {
				continue
			}

			neighbors := 0
			for _, dir := range dirs {
				nRow := row + dir[0]
				nCol := col + dir[1]

				if nRow >= 0 && nRow < len(grid) && nCol >= 0 && nCol < len(grid[0]) {
					if grid[nRow][nCol] == '@' {
						neighbors++
					}
				}
			}
			if neighbors < 4 {
				removableRolls++
			}
		}
	}

	return removableRolls
}

func partTwo(grid [][]rune, removedSoFar int) int {
	removableRolls := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			shelf := grid[row][col]

			if shelf != '@' {
				continue
			}

			neighbors := 0
			for _, dir := range dirs {
				nRow := row + dir[0]
				nCol := col + dir[1]

				if nRow >= 0 && nRow < len(grid) && nCol >= 0 && nCol < len(grid[0]) {
					if grid[nRow][nCol] == '@' {
						neighbors++
					}
				}
			}
			if neighbors < 4 {
				removableRolls++
				grid[row][col] = 'x'
			}
		}
	}

	if removableRolls == 0 {
		return removedSoFar
	}

	return partTwo(grid, removedSoFar+removableRolls)
}

func main() {
	d, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := fmt.Sprintf("%s/data/04.txt", d)

	grid := parseInput(path)

	fmt.Printf("Part 1: %d\n", partOne(grid))
	fmt.Printf("Part 2: %d\n", partTwo(grid, 0))
}
