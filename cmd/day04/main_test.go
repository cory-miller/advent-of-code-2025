package main

import (
	"strings"
	"testing"
)

func createGridFromString(textGrid string) [][]rune {
	grid := make([][]rune, 0)

	lines := strings.Split(textGrid, "\n")

	for _, line := range lines {
		row := []rune(line)

		grid = append(grid, row)
	}

	return grid
}

func TestPartOne(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	grid := createGridFromString(input)

	if partOne(grid) != 13 {
		t.Fatal("expected 13")
	}
}

func TestPartTwo(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	grid := createGridFromString(input)

	if partTwo(grid, 0) != 43 {
		t.Fatal("expected 43")
	}
}
